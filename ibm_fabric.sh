#!/bin/sh
set -e

echo
echo "#################################################################"
echo "#######        Generating cryptographic material       ##########"
echo "#################################################################"
PROJPATH=$(pwd)
CLIPATH=$PROJPATH/cli/peers
ORDERERS=$CLIPATH/ordererOrganizations
PEERS=$CLIPATH/peerOrganizations

rm -rf $CLIPATH
$PROJPATH/cryptogen generate --config=$PROJPATH/crypto-config.yaml --output=$CLIPATH

sh generate-cfgtx.sh

# rm -rf $PROJPATH/{orderer,realestatePeer,regulatorPeer,traderPeer}/crypto
rm -rf $PROJPATH/orderer/crypto
rm -rf $PROJPATH/realestatePeer/crypto
rm -rf $PROJPATH/regulatorPeer/crypto
rm -rf $PROJPATH/traderPeer/crypto

# mkdir $PROJPATH/{orderer,realestatePeer,regulatorPeer,traderPeer}/crypto
mkdir $PROJPATH/orderer/crypto
mkdir $PROJPATH/realestatePeer/crypto
mkdir $PROJPATH/regulatorPeer/crypto
mkdir $PROJPATH/traderPeer/crypto

# cp -r $ORDERERS/orderer-org/orderers/orderer0/{msp,tls} $PROJPATH/orderer/crypto
cp -r $ORDERERS/orderer-org/orderers/orderer0/msp $PROJPATH/orderer/crypto
cp -r $ORDERERS/orderer-org/orderers/orderer0/tls $PROJPATH/orderer/crypto
# cp -r $PEERS/realestate-org/peers/realestate-peer/{msp,tls} $PROJPATH/realestatePeer/crypto
cp -r $PEERS/realestate-org/peers/realestate-peer/msp $PROJPATH/realestatePeer/crypto
cp -r $PEERS/realestate-org/peers/realestate-peer/tls $PROJPATH/realestatePeer/crypto
# cp -r $PEERS/regulator-org/peers/regulator-peer/{msp,tls} $PROJPATH/regulatorPeer/crypto
cp -r $PEERS/regulator-org/peers/regulator-peer/msp $PROJPATH/regulatorPeer/crypto
cp -r $PEERS/regulator-org/peers/regulator-peer/tls $PROJPATH/regulatorPeer/crypto
# cp -r $PEERS/trader-org/peers/trader-peer/{msp,tls} $PROJPATH/traderPeer/crypto
cp -r $PEERS/trader-org/peers/trader-peer/msp $PROJPATH/traderPeer/crypto
cp -r $PEERS/trader-org/peers/trader-peer/tls $PROJPATH/traderPeer/crypto


cp $CLIPATH/genesis.block $PROJPATH/orderer/crypto/

REALESTATECAPATH=$PROJPATH/realestateCA
REGULATORCAPATH=$PROJPATH/regulatorCA
TRADERCAPATH=$PROJPATH/traderCA


# rm -rf {$REALESTATECAPATH,$REGULATORCAPATH,$TRADERCAPATH,$SHOPCAPATH}/{ca,tls}
rm -rf $REALESTATECAPATH/ca
rm -rf $REGULATORCAPATH/ca
rm -rf $TRADERCAPATH/ca

rm -rf $REALESTATECAPATH/tls
rm -rf $REGULATORCAPATH/tls
rm -rf $TRADERCAPATH/tls

# mkdir -p {$REALESTATECAPATH,$REGULATORCAPATH,$TRADERCAPATH,$SHOPCAPATH}/{ca,tls}
mkdir -p $REALESTATECAPATH/ca
mkdir -p $REGULATORCAPATH/ca
mkdir -p $TRADERCAPATH/ca

mkdir -p $REALESTATECAPATH/tls
mkdir -p $REGULATORCAPATH/tls
mkdir -p $TRADERCAPATH/tls

cp $PEERS/realestate-org/ca/* $REALESTATECAPATH/ca
cp $PEERS/realestate-org/tlsca/* $REALESTATECAPATH/tls
mv $REALESTATECAPATH/ca/*_sk $REALESTATECAPATH/ca/key.pem
mv $REALESTATECAPATH/ca/*-cert.pem $REALESTATECAPATH/ca/cert.pem
mv $REALESTATECAPATH/tls/*_sk $REALESTATECAPATH/tls/key.pem
mv $REALESTATECAPATH/tls/*-cert.pem $REALESTATECAPATH/tls/cert.pem

cp $PEERS/regulator-org/ca/* $REGULATORCAPATH/ca
cp $PEERS/regulator-org/tlsca/* $REGULATORCAPATH/tls
mv $REGULATORCAPATH/ca/*_sk $REGULATORCAPATH/ca/key.pem
mv $REGULATORCAPATH/ca/*-cert.pem $REGULATORCAPATH/ca/cert.pem
mv $REGULATORCAPATH/tls/*_sk $REGULATORCAPATH/tls/key.pem
mv $REGULATORCAPATH/tls/*-cert.pem $REGULATORCAPATH/tls/cert.pem

cp $PEERS/trader-org/ca/* $TRADERCAPATH/ca
cp $PEERS/trader-org/tlsca/* $TRADERCAPATH/tls
mv $TRADERCAPATH/ca/*_sk $TRADERCAPATH/ca/key.pem
mv $TRADERCAPATH/ca/*-cert.pem $TRADERCAPATH/ca/cert.pem
mv $TRADERCAPATH/tls/*_sk $TRADERCAPATH/tls/key.pem
mv $TRADERCAPATH/tls/*-cert.pem $TRADERCAPATH/tls/cert.pem


WEBCERTS=$PROJPATH/web/certs
rm -rf $WEBCERTS
mkdir -p $WEBCERTS
cp $PROJPATH/orderer/crypto/tls/ca.crt $WEBCERTS/ordererOrg.pem
cp $PROJPATH/realestatePeer/crypto/tls/ca.crt $WEBCERTS/realestateOrg.pem
cp $PROJPATH/regulatorPeer/crypto/tls/ca.crt $WEBCERTS/regulatorOrg.pem
cp $PROJPATH/traderPeer/crypto/tls/ca.crt $WEBCERTS/traderOrg.pem
cp $PEERS/realestate-org/users/Admin@realestate-org/msp/keystore/* $WEBCERTS/Admin@realestate-org-key.pem
cp $PEERS/realestate-org/users/Admin@realestate-org/msp/signcerts/* $WEBCERTS/
cp $PEERS/regulator-org/users/Admin@regulator-org/msp/keystore/* $WEBCERTS/Admin@regulator-org-key.pem
cp $PEERS/regulator-org/users/Admin@regulator-org/msp/signcerts/* $WEBCERTS/
cp $PEERS/trader-org/users/Admin@trader-org/msp/keystore/* $WEBCERTS/Admin@trader-org-key.pem
cp $PEERS/trader-org/users/Admin@trader-org/msp/signcerts/* $WEBCERTS/