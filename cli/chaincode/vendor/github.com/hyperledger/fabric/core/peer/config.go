/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// The 'viper' package for configuration handling is very flexible, but has
// been found to have extremely poor performance when configuration values are
// accessed repeatedly. The function CacheConfiguration() defined here caches
// all configuration values that are accessed frequently.  These parameters
// are now presented as function calls that access local configuration
// variables.  This seems to be the most robust way to represent these
// parameters in the face of the numerous ways that configuration files are
// loaded and used (e.g, normal usage vs. test cases).

// The CacheConfiguration() function is allowed to be called globally to
// ensure that the correct values are always cached; See for example how
// certain parameters are forced in 'ChaincodeDevMode' in main.go.

package peer

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric/core/comm"
	"github.com/hyperledger/fabric/core/config"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	//Peer config
	LocalMspID                            string
	ListenAddress                         string
	AuthenticationTimeWindow              time.Duration
	PeerTLSEnabled                        bool
	PeerID                                string
	PeerAddress                           string
	PeerEndpoint                          *pb.PeerEndpoint
	NetworkID                             string
	LimitsConcurrencyQSCC                 int
	DiscoveryEnabled                      bool
	ProfileEnabled                        bool
	ProfileListenAddress                  string
	DiscoveryOrgMembersAllowed            bool
	DiscoveryAuthCacheEnabled             bool
	DiscoveryAuthCacheMaxSize             int
	DiscoveryAuthCachePurgeRetentionRatio float64
	ChaincodeListenAddr                   string
	ChaincodeAddr                         string
	AdminListenAddr                       string

	//VM config
	VMEndpoint           string
	VMDockerTLSEnabled   bool
	VMDockerAttachStdout bool

	//Chaincode config
	ChaincodePull bool
}

func GlobalConfig() (*Config, error) {
	c := &Config{}
	if err := c.load(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) load() error {
	preeAddress, err := getLocalAddress()
	if err != nil {
		return err
	}
	c.PeerAddress = preeAddress
	c.PeerID = viper.GetString("peer.id")
	c.PeerEndpoint = &pb.PeerEndpoint{
		Id: &pb.PeerID{
			Name: c.PeerID,
		},
		Address: c.PeerAddress,
	}
	c.LocalMspID = viper.GetString("peer.localMspId")
	c.ListenAddress = viper.GetString("peer.listenAddress")
	c.AuthenticationTimeWindow = viper.GetDuration("peer.authentication.timewindow")
	c.PeerTLSEnabled = viper.GetBool("peer.tls.enabled")
	c.NetworkID = viper.GetString("peer.networkId")
	c.LimitsConcurrencyQSCC = viper.GetInt("peer.limits.concurrency.qscc")
	c.DiscoveryEnabled = viper.GetBool("peer.discovery.enabled")
	c.ProfileEnabled = viper.GetBool("peer.profile.enabled")
	c.ProfileListenAddress = viper.GetString("peer.profile.listenAddress")
	c.DiscoveryOrgMembersAllowed = viper.GetBool("peer.discovery.orgMembersAllowedAccess")
	c.DiscoveryAuthCacheEnabled = viper.GetBool("peer.discovery.authCacheEnabled")
	c.DiscoveryAuthCacheMaxSize = viper.GetInt("peer.discovery.authCacheMaxSize")
	c.DiscoveryAuthCachePurgeRetentionRatio = viper.GetFloat64("peer.discovery.authCachePurgeRetentionRatio")
	c.ChaincodeListenAddr = viper.GetString("peer.chaincodeListenAddress")
	c.ChaincodeAddr = viper.GetString("peer.chaincodeAddress")
	c.AdminListenAddr = viper.GetString("peer.adminService.listenAddress")

	c.VMEndpoint = viper.GetString("vm.endpoint")
	c.VMDockerTLSEnabled = viper.GetBool("vm.docker.tls.enabled")
	c.VMDockerAttachStdout = viper.GetBool("vm.docker.attachStdout")

	c.ChaincodePull = viper.GetBool("chaincode.pull")

	return nil
}

// getLocalAddress returns the address:port the local peer is operating on.  Affected by env:peer.addressAutoDetect
func getLocalAddress() (string, error) {
	peerAddress := viper.GetString("peer.address")
	if peerAddress == "" {
		return "", fmt.Errorf("peer.address isn't set")
	}
	host, port, err := net.SplitHostPort(peerAddress)
	if err != nil {
		return "", errors.Errorf("peer.address isn't in host:port format: %s", peerAddress)
	}

	localIP, err := GetLocalIP()
	if err != nil {
		peerLogger.Errorf("Local ip address not auto-detectable: %s", err)
		return "", err
	}
	autoDetectedIPAndPort := net.JoinHostPort(localIP, port)
	peerLogger.Info("Auto-detected peer address:", autoDetectedIPAndPort)
	// If host is the IPv4 address "0.0.0.0" or the IPv6 address "::",
	// then fallback to auto-detected address
	if ip := net.ParseIP(host); ip != nil && ip.IsUnspecified() {
		peerLogger.Info("Host is", host, ", falling back to auto-detected address:", autoDetectedIPAndPort)
		return autoDetectedIPAndPort, nil
	}

	if viper.GetBool("peer.addressAutoDetect") {
		peerLogger.Info("Auto-detect flag is set, returning", autoDetectedIPAndPort)
		return autoDetectedIPAndPort, nil
	}
	peerLogger.Info("Returning", peerAddress)
	return peerAddress, nil

}

// GetServerConfig returns the gRPC server configuration for the peer
func GetServerConfig() (comm.ServerConfig, error) {
	secureOptions := &comm.SecureOptions{
		UseTLS: viper.GetBool("peer.tls.enabled"),
	}
	serverConfig := comm.ServerConfig{SecOpts: secureOptions}
	if secureOptions.UseTLS {
		// get the certs from the file system
		serverKey, err := ioutil.ReadFile(config.GetPath("peer.tls.key.file"))
		if err != nil {
			return serverConfig, fmt.Errorf("error loading TLS key (%s)", err)
		}
		serverCert, err := ioutil.ReadFile(config.GetPath("peer.tls.cert.file"))
		if err != nil {
			return serverConfig, fmt.Errorf("error loading TLS certificate (%s)", err)
		}
		secureOptions.Certificate = serverCert
		secureOptions.Key = serverKey
		secureOptions.RequireClientCert = viper.GetBool("peer.tls.clientAuthRequired")
		if secureOptions.RequireClientCert {
			var clientRoots [][]byte
			for _, file := range viper.GetStringSlice("peer.tls.clientRootCAs.files") {
				clientRoot, err := ioutil.ReadFile(
					config.TranslatePath(filepath.Dir(viper.ConfigFileUsed()), file))
				if err != nil {
					return serverConfig,
						fmt.Errorf("error loading client root CAs (%s)", err)
				}
				clientRoots = append(clientRoots, clientRoot)
			}
			secureOptions.ClientRootCAs = clientRoots
		}
		// check for root cert
		if config.GetPath("peer.tls.rootcert.file") != "" {
			rootCert, err := ioutil.ReadFile(config.GetPath("peer.tls.rootcert.file"))
			if err != nil {
				return serverConfig, fmt.Errorf("error loading TLS root certificate (%s)", err)
			}
			secureOptions.ServerRootCAs = [][]byte{rootCert}
		}
	}
	// get the default keepalive options
	serverConfig.KaOpts = comm.DefaultKeepaliveOptions
	// check to see if interval is set for the env
	if viper.IsSet("peer.keepalive.interval") {
		serverConfig.KaOpts.ServerInterval = viper.GetDuration("peer.keepalive.interval")
	}
	// check to see if timeout is set for the env
	if viper.IsSet("peer.keepalive.timeout") {
		serverConfig.KaOpts.ServerTimeout = viper.GetDuration("peer.keepalive.timeout")
	}
	// check to see if minInterval is set for the env
	if viper.IsSet("peer.keepalive.minInterval") {
		serverConfig.KaOpts.ServerMinInterval = viper.GetDuration("peer.keepalive.minInterval")
	}
	return serverConfig, nil
}

// GetClientCertificate returns the TLS certificate to use for gRPC client
// connections
func GetClientCertificate() (tls.Certificate, error) {
	cert := tls.Certificate{}

	keyPath := viper.GetString("peer.tls.clientKey.file")
	certPath := viper.GetString("peer.tls.clientCert.file")

	if keyPath != "" || certPath != "" {
		// need both keyPath and certPath to be set
		if keyPath == "" || certPath == "" {
			return cert, errors.New("peer.tls.clientKey.file and " +
				"peer.tls.clientCert.file must both be set or must both be empty")
		}
		keyPath = config.GetPath("peer.tls.clientKey.file")
		certPath = config.GetPath("peer.tls.clientCert.file")

	} else {
		// use the TLS server keypair
		keyPath = viper.GetString("peer.tls.key.file")
		certPath = viper.GetString("peer.tls.cert.file")

		if keyPath != "" || certPath != "" {
			// need both keyPath and certPath to be set
			if keyPath == "" || certPath == "" {
				return cert, errors.New("peer.tls.key.file and " +
					"peer.tls.cert.file must both be set or must both be empty")
			}
			keyPath = config.GetPath("peer.tls.key.file")
			certPath = config.GetPath("peer.tls.cert.file")
		} else {
			return cert, errors.New("must set either " +
				"[peer.tls.key.file and peer.tls.cert.file] or " +
				"[peer.tls.clientKey.file and peer.tls.clientCert.file]" +
				"when peer.tls.clientAuthEnabled is set to true")
		}
	}
	// get the keypair from the file system
	clientKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return cert, errors.WithMessage(err,
			"error loading client TLS key")
	}
	clientCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return cert, errors.WithMessage(err,
			"error loading client TLS certificate")
	}
	cert, err = tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		return cert, errors.WithMessage(err,
			"error parsing client TLS key pair")
	}
	return cert, nil
}
