package main

import (
	"crypto/x509"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
)

func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, string, error) {
	var mspID string
	var err error
	var cert *x509.Certificate

	mspID, err = cid.GetMSPID(stub)
	if err != nil {
		return "", "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)

	if err != nil {
		return "", "", "", err
	}

	return mspID, cert.Issuer.CommonName, cert.Subject.CommonName, nil
}

// For now, just hardcode an ACL
// We will support attribute checks in an upgrade

func authenticateBdsOrg(mspID string, certCN string) bool {
	return (mspID == "BdsOrgMSP") && (certCN == "ca.bds-org")
}

func authenticateCcqOrg(mspID string, certCN string) bool {
	return (mspID == "CcqOrgMSP") && (certCN == "ca.ccq-org")
}

func authenticateTraderOrg(mspID string, certCN string) bool {
	return (mspID == "TraderOrgMSP") && (certCN == "ca.trader-org")
}

func authenticateShareholderOrg(mspID string, certCN string) bool {
	return authenticateBdsOrg(mspID, certCN) || authenticateTraderOrg(mspID, certCN)
}
