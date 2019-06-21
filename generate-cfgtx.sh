#!/bin/sh

CHANNEL_NAME="tritchannel"
PROJPATH=$(pwd)
CLIPATH=$PROJPATH/cli/peers

echo
echo "##########################################################"
echo "#########  Generating Orderer Genesis block ##############"
echo "##########################################################"
$PROJPATH/configtxgen -profile FourOrgsGenesis -outputBlock $CLIPATH/genesis.block -channelID trit-sys-channel

echo
echo "#################################################################"
echo "### Generating channel configuration transaction 'channel.tx' ###"
echo "#################################################################"
$PROJPATH/configtxgen -profile FourOrgsChannel -outputCreateChannelTx $CLIPATH/channel.tx -channelID $CHANNEL_NAME
cp $CLIPATH/channel.tx $PROJPATH/web
echo
echo "#################################################################"
echo "####### Generating anchor peer update for RealEstateOrg ##########"
echo "#################################################################"
$PROJPATH/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate $CLIPATH/RealEstateOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg RealEstateOrgMSP

echo
echo "##################################################################"
echo "####### Generating anchor peer update for TraderOrg ##########"
echo "##################################################################"
$PROJPATH/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate $CLIPATH/TraderOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg TraderOrgMSP

echo
echo "##################################################################"
echo "#######   Generating anchor peer update for RegulatorOrg   ##########"
echo "##################################################################"
$PROJPATH/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate $CLIPATH/RegulatorOrgMSPAnchors.tx -channelID $CHANNEL_NAME -asOrg RegulatorOrgMSP
