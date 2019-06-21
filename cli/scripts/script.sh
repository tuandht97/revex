#!/bin/bash
set -e


echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "THANH-TUYET-TUAN TRIT NETWORK"
echo


CHANNEL_NAME="tritchannel"
peer channel create -o orderer0:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "Join regulator peer to channel"
peer channel join -b tritchannel.block

echo "join trader peer to channel"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer channel join -b tritchannel.block

echo "join realestate peer to channel"
CORE_PEER_ADDRESS=realestate-peer:7051
CORE_PEER_LOCALMSPID=RealEstateOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/users/Admin@realestate-org/msp
peer channel join -b tritchannel.block

echo "update the anchor peer for real estate peer"
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./channel-artifacts/RealEstateOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "update the anchor peer for trader peer"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./channel-artifacts/TraderOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "update the anchor peers for regulator peer"
CORE_PEER_ADDRESS=regulator-peer:7051
CORE_PEER_LOCALMSPID=RegulatorOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/regulator-org/peers/regulator-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/regulator-org/peers/regulator-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/regulator-org/peers/regulator-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/regulator-org/users/Admin@regulator-org/msp
peer channel update -o orderer0:7050 -c $CHANNEL_NAME -f ./channel-artifacts/RegulatorOrgMSPAnchors.tx --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem

echo "install chaincode in real regulator peer"
peer chaincode install -n trit_chaincode -v 7.0 -p github.com/chaincode

echo "install chaincode in trader peer"
CORE_PEER_ADDRESS=trader-peer:7051
CORE_PEER_LOCALMSPID=TraderOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/trader-org/peers/trader-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/trader-org/users/Admin@trader-org/msp
peer chaincode install -n trit_chaincode -v 7.0 -p github.com/chaincode

echo "install chaincode in real estate peer"
CORE_PEER_ADDRESS=realestate-peer:7051
CORE_PEER_LOCALMSPID=RealEstateOrgMSP
CORE_PEER_TLS_CERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/peer/crypto/peerOrganizations/realestate-org/users/Admin@realestate-org/msp
peer chaincode install -n trit_chaincode -v 7.0 -p github.com/chaincode

echo "instantiate the chaincode on the channel..."
peer chaincode instantiate -o orderer0:7050 --tls --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem -C $CHANNEL_NAME  -n trit_chaincode -v 7.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('RegulatorOrgMSP.member','TraderOrgMSP.member', 'RealEstateOrgMSP.member')"

echo "TEST CHAINCODE INVOKE"
sleep 3
echo "invoke create publisher, sending invoke transaction on regulator peer..."
peer chaincode invoke -o orderer0:7050 --tls true --peerAddresses realestate-peer:7051   --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem -C tritchannel  -n trit_chaincode  --tlsRootCertFiles /opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/ca.crt  -c '{"Args":["create_and_publish_real_estate","{\"uuid\":\"abc\", \"id\":\"hanh\",\"price\":10,\"square_meter\":10,\"address\":\"123343424234\",\"amount\":10,\"description\":\"123343424234\",\"owner_id\":\"truongthianh\" }"]}'

echo "expected null.."
sleep 1
peer chaincode query -o orderer0:7050 --tls true --cafile /opt/gopath/peer/crypto/ordererOrganizations/orderer-org/orderers/orderer0/msp/tlscacerts/tlsca.orderer-org-cert.pem -C tritchannel  -n trit_chaincode --peerAddresses realestate-peer:7051 --tlsRootCertFiles /opt/gopath/peer/crypto/peerOrganizations/realestate-org/peers/realestate-peer/tls/ca.crt  -c '{"Args":["list_real_estate",""]}'


echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo
