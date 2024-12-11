#!/bin/bash

# This script is used to start ZetaClient for the localnet
# An optional argument can be passed and can have the following value:
# background: start the ZetaClient in the background, this prevent the image from being stopped when ZetaClient must be restarted

/usr/sbin/sshd

HOSTNAME=$(hostname)
export ZETACLIENTD_SUPERVISOR_ENABLE_AUTO_DOWNLOAD=true

# sepolia is used in chain migration tests, this functions set the sepolia endpoint in the zetaclient_config.json
set_sepolia_endpoint() {
  jq '.EVMChainConfigs."11155111".Endpoint = "http://eth2:8545"' /root/.zetacored/config/zetaclient_config.json > tmp.json && mv tmp.json /root/.zetacored/config/zetaclient_config.json
}

# import a relayer private key (e.g. Solana relayer key)
import_relayer_key() {
    local num="$1"

  # import solana (network=7) relayer private key
  privkey_solana=$(yq -r ".observer_relayer_accounts.relayer_accounts[${num}].solana_private_key" /root/config.yml)
  zetaclientd relayer import-key --network=7 --private-key="$privkey_solana" --password=pass_relayerkey
}

PREPARAMS_PATH="/root/preparams/${HOSTNAME}.json"
if [[ -n "${ZETACLIENTD_GEN_PREPARAMS}" ]]; then
  # generate pre-params as early as possible
  # to reach keygen height on schedule
  if [ ! -f "$PREPARAMS_PATH" ]; then
    zetaclientd gen-pre-params "$PREPARAMS_PATH"
  fi
else
  echo "Using static preparams"
  cp "/root/static-preparams/${HOSTNAME}.json" "${PREPARAMS_PATH}"
fi

# Wait for authorized_keys file to exist (generated by zetacore0)
while [ ! -f ~/.ssh/authorized_keys ]; do
    echo "Waiting for authorized_keys file to exist..."
    sleep 1
done

# need to wait for zetacore0 to be up
while ! curl -s -o /dev/null zetacore0:26657/status ; do
    echo "Waiting for zetacore0 rpc"
    sleep 10
done

# read HOTKEY_BACKEND env var for hotkey keyring backend and set default to test
BACKEND="test"
if [ "$HOTKEY_BACKEND" == "file" ]; then
    BACKEND="file"
fi

num=$(echo $HOSTNAME | tr -dc '0-9')
node="zetacore$num"

while [ ! -f $HOME/.zetacored/os.json ]; do
    echo "Waiting for zetacore to exchange os.json file..."
    sleep 1
done
operator=$(cat $HOME/.zetacored/os.json | jq '.ObserverAddress' )
operatorAddress=$(echo "$operator" | tr -d '"')
echo "operatorAddress: $operatorAddress"

# create the path that holds observer relayer private keys (e.g. Solana relayer key)
RELAYER_KEY_PATH="$HOME/.zetacored/relayer-keys"
mkdir -p "${RELAYER_KEY_PATH}"

mkdir -p "$HOME/.tss/"
zetae2e get-zetaclient-bootstrap > "$HOME/.tss/address_book.seed"

echo "Start zetaclientd"
# skip initialization if the config file already exists (zetaclientd init has already been run)
if [[ $HOSTNAME == "zetaclient0" && ! -f ~/.zetacored/config/zetaclient_config.json ]]
then
    MYIP=$(/sbin/ip -o -4 addr list eth0 | awk '{print $4}' | cut -d/ -f1)
    zetaclientd init --zetacore-url zetacore0 --chain-id athens_101-1 --operator "$operatorAddress" --log-format=text --public-ip "$MYIP" --keyring-backend "$BACKEND" --pre-params "$PREPARAMS_PATH"

    # import relayer private key for zetaclient0
    import_relayer_key 0

    # if eth2 is enabled, set the endpoint in the zetaclient_config.json
    # in this case, the additional evm is represented with the sepolia chain, we set manually the eth2 endpoint to the sepolia chain (11155111 -> http://eth2:8545)
    # in /root/.zetacored/config/zetaclient_config.json
    if host eth2 ; then
     echo "enabling additional evm (eth2)"
     set_sepolia_endpoint
    fi
fi
if [[ $HOSTNAME != "zetaclient0" && ! -f ~/.zetacored/config/zetaclient_config.json ]]
then
  num=$(echo $HOSTNAME | tr -dc '0-9')
  node="zetacore$num"
  zetaclientd init --zetacore-url "$node" --chain-id athens_101-1 --operator "$operatorAddress" --log-format=text --public-ip "$MYIP" --log-level 1 --keyring-backend "$BACKEND" --pre-params "$PREPARAMS_PATH"

  # import relayer private key for zetaclient{$num}
  import_relayer_key "${num}"

  # check if the option is additional-evm
  # in this case, the additional evm is represented with the sepolia chain, we set manually the eth2 endpoint to the sepolia chain (11155111 -> http://eth2:8545)
  # in /root/.zetacored/config/zetaclient_config.json
  if [[ -n $ADDITIONAL_EVM ]]; then
   set_sepolia_endpoint
  fi
fi

# merge zetaclient-config-overlay.json into zetaclient_config.json if specified
if [[ -f /root/zetaclient-config-overlay.json ]]; then
  jq -s '.[0] * .[1]' /root/.zetacored/config/zetaclient_config.json /root/zetaclient-config-overlay.json > /tmp/merged_config.json
  mv /tmp/merged_config.json /root/.zetacored/config/zetaclient_config.json
fi

zetaclientd-supervisor start < /root/password.file