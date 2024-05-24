#!/bin/bash
if [ $# -lt 1 ]
then
  echo "Usage: import-data.sh [network]"
  exit 1
fi

NETWORK=$1
echo "NETWORK: ${NETWORK}"
rm -rf ~/genesis_export/
mkdir ~/genesis_export/
echo "Download Latest State Export"
LATEST_EXPORT_URL=$(curl https://snapshots.zetachain.com/latest-state-export | jq -r ."${NETWORK}")
echo "LATEST EXPORT URL: ${LATEST_EXPORT_URL}"
wget -q ${LATEST_EXPORT_URL} -O ~/genesis_export/exported-genesis.json