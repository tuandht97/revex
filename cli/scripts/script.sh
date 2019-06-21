#!/bin/bash
set -e


echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "REVEX NETWORK"
echo


CHANNEL_NAME="revex"
peer channel create -o orderer0:7050 -c $CHANNEL_NAME -f ./peer/crypto/channel.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "Join ccq peer to channel"
peer channel join -b revex.block

echo "join trader peer to channel"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer channel join -b revex.block

echo "join bds peer to channel"
CORE_PEER_ADDRESS=bds-peer:7051
CORE_PEER_LOCALMSPID=BdsOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/bds-org/users/Admin@bds-org/msp
peer channel join -b revex.block

echo "update the anchor peer for real estate peer"
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./peer/crypto/BdsOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "update the anchor peer for trader peer"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./peer/crypto/TraderOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "update the anchor peers for ccq peer"
CORE_PEER_ADDRESS=ccq-peer:7051
CORE_PEER_LOCALMSPID=CcqOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/ccq-org/peers/ccq-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/ccq-org/peers/ccq-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/ccq-org/peers/ccq-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/ccq-org/users/Admin@ccq-org/msp
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./peer/crypto/CcqOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "install chaincode in real ccq peer"
peer chaincode install -n revex -v 7.0 -p github.com/chaincode

echo "install chaincode in trader peer"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer chaincode install -n revex -v 7.0 -p github.com/chaincode

echo "install chaincode in bds peer"
CORE_PEER_ADDRESS=bds-peer:7051
CORE_PEER_LOCALMSPID=BdsOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/bds-org/peers/bds-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/bds-org/users/Admin@bds-org/msp
peer chaincode install -n revex -v 7.0 -p github.com/chaincode

echo "instantiate the chaincode on the channel..."
peer chaincode instantiate -o orderer0:7050 --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem -C $CHANNEL_NAME  -n revex -v 7.0 -c '{"Args":["init"]}' -P "AND ('CcqOrgMSP.member','TraderOrgMSP.member', 'BdsOrgMSP.member')"

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo
