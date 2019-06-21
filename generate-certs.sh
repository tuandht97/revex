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

rm -rf $PROJPATH/orderer/crypto
rm -rf $PROJPATH/bdsPeer/crypto
rm -rf $PROJPATH/ccqPeer/crypto
rm -rf $PROJPATH/traderPeer/crypto

mkdir $PROJPATH/orderer/crypto
mkdir $PROJPATH/bdsPeer/crypto
mkdir $PROJPATH/ccqPeer/crypto
mkdir $PROJPATH/traderPeer/crypto

cp -r $ORDERERS/orderer-org/orderers/orderer0/msp $PROJPATH/orderer/crypto
cp -r $PEERS/bds-org/peers/bds-peer/msp $PROJPATH/bdsPeer/crypto
cp -r $PEERS/ccq-org/peers/ccq-peer/msp $PROJPATH/ccqPeer/crypto
cp -r $PEERS/trader-org/peers/trader-peer/msp $PROJPATH/traderPeer/crypto

cp -r $ORDERERS/orderer-org/orderers/orderer0/tls $PROJPATH/orderer/crypto
cp -r $PEERS/bds-org/peers/bds-peer/tls $PROJPATH/bdsPeer/crypto
cp -r $PEERS/ccq-org/peers/ccq-peer/tls $PROJPATH/ccqPeer/crypto
cp -r $PEERS/trader-org/peers/trader-peer/tls $PROJPATH/traderPeer/crypto

cp $CLIPATH/genesis.block $PROJPATH/orderer/crypto/

bdsCAPATH=$PROJPATH/bdsCA
ccqCAPATH=$PROJPATH/ccqCA
TRADERCAPATH=$PROJPATH/traderCA


rm -rf $bdsCAPATH/ca
mkdir -p $bdsCAPATH/ca
rm -rf $ccqCAPATH/ca
mkdir -p $ccqCAPATH/ca
rm -rf $TRADERCAPATH/ca
mkdir -p $TRADERCAPATH/ca

rm -rf $bdsCAPATH/tls
mkdir -p $bdsCAPATH/tls
rm -rf $ccqCAPATH/tls
mkdir -p $ccqCAPATH/tls
rm -rf $TRADERCAPATH/tls
mkdir -p $TRADERCAPATH/tls


cp $PEERS/bds-org/ca/* $bdsCAPATH/ca
cp $PEERS/bds-org/tlsca/* $bdsCAPATH/tls
mv $bdsCAPATH/ca/*_sk $bdsCAPATH/ca/key.pem
mv $bdsCAPATH/ca/*-cert.pem $bdsCAPATH/ca/cert.pem
mv $bdsCAPATH/tls/*_sk $bdsCAPATH/tls/key.pem
mv $bdsCAPATH/tls/*-cert.pem $bdsCAPATH/tls/cert.pem

cp $PEERS/ccq-org/ca/* $ccqCAPATH/ca
cp $PEERS/ccq-org/tlsca/* $ccqCAPATH/tls
mv $ccqCAPATH/ca/*_sk $ccqCAPATH/ca/key.pem
mv $ccqCAPATH/ca/*-cert.pem $ccqCAPATH/ca/cert.pem
mv $ccqCAPATH/tls/*_sk $ccqCAPATH/tls/key.pem
mv $ccqCAPATH/tls/*-cert.pem $ccqCAPATH/tls/cert.pem

cp $PEERS/trader-org/ca/* $TRADERCAPATH/ca
cp $PEERS/trader-org/tlsca/* $TRADERCAPATH/tls
mv $TRADERCAPATH/ca/*_sk $TRADERCAPATH/ca/key.pem
mv $TRADERCAPATH/ca/*-cert.pem $TRADERCAPATH/ca/cert.pem
mv $TRADERCAPATH/tls/*_sk $TRADERCAPATH/tls/key.pem
mv $TRADERCAPATH/tls/*-cert.pem $TRADERCAPATH/tls/cert.pem


CERTS=$PROJPATH/certs
rm -rf $CERTS
mkdir -p $CERTS
cp $PROJPATH/orderer/crypto/tls/ca.crt $CERTS/ordererOrg.pem
cp $PROJPATH/bdsPeer/crypto/tls/ca.crt $CERTS/bdsOrg.pem
cp $PROJPATH/ccqPeer/crypto/tls/ca.crt $CERTS/ccqOrg.pem
cp $PROJPATH/traderPeer/crypto/tls/ca.crt $CERTS/traderOrg.pem

cp $PEERS/bds-org/users/Admin@bds-org/msp/keystore/* $CERTS/Admin@bds-org-key.pem
cp $PEERS/bds-org/users/Admin@bds-org/msp/signcerts/* $CERTS/
cp $PEERS/ccq-org/users/Admin@ccq-org/msp/keystore/* $CERTS/Admin@ccq-org-key.pem
cp $PEERS/ccq-org/users/Admin@ccq-org/msp/signcerts/* $CERTS/
cp $PEERS/trader-org/users/Admin@trader-org/msp/keystore/* $CERTS/Admin@trader-org-key.pem
cp $PEERS/trader-org/users/Admin@trader-org/msp/signcerts/* $CERTS/
