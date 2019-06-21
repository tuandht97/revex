#!/bin/sh

CHANNEL_NAME="revex"
PROJPATH=$(pwd)
CLIPATH=$PROJPATH/cli/peers

echo
echo "##########################################################"
echo "#########  Generating Orderer Genesis block ##############"
echo "##########################################################"
$PROJPATH/configtxgen -profile ThreeOrgsGenesis -outputBlock $CLIPATH/genesis.block -channelID revex-sys-channel

echo
echo "#################################################################"
echo "### Generating channel configuration transaction 'channel.tx' ###"
echo "#################################################################"
$PROJPATH/configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx $CLIPATH/channel.tx -channelID $CHANNEL_NAME

echo
echo "#################################################################"
echo "####### Generating anchor peer update for BdsOrg ##########"
echo "#################################################################"
$PROJPATH/configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate $CLIPATH/BdsOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg BdsOrgMSP

echo
echo "##################################################################"
echo "####### Generating anchor peer update for TraderOrg ##########"
echo "##################################################################"
$PROJPATH/configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate $CLIPATH/TraderOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg TraderOrgMSP

echo
echo "##################################################################"
echo "#######   Generating anchor peer update for CcqOrg   ##########"
echo "##################################################################"
$PROJPATH/configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate $CLIPATH/CcqOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg CcqOrgMSP
