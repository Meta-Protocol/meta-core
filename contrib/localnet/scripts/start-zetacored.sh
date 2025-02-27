#!/bin/bash

# This script is used to start the zetacored nodes
# It initializes the nodes and creates the genesis.json file
# It also starts the nodes
# The number of nodes is passed as an first argument to the script
# The second argument is optional and can have the following value:
#  - import-data: import data into the genesis file

/usr/sbin/sshd

add_emissions_withdraw_authorizations() {

    config_file="/root/config.yml"
    json_file="/root/.zetacored/config/genesis.json"

    # Check if config file exists
    if [[ ! -f "$config_file" ]]; then
        echo "Error: Config file not found at $config_file"
        return 1
    fi
    # Address to add emissions withdraw authorizations
    address=$(yq -r '.additional_accounts.user_emissions_withdraw.bech32_address' "$config_file")

    # Check if genesis file exists
    if [[ ! -f "$json_file" ]]; then
        echo "Error: Genesis file not found at $json_file"
        return 1
    fi

    echo "Adding emissions withdraw authorizations for address: $address"


     # Using jq to parse JSON, create new entries, and append them to the authorization array
     if ! jq --arg address "$address" '
         # Store the nodeAccountList array
         .app_state.observer.nodeAccountList as $list |
         # Iterate over the stored list to construct new objects and append to the authorization array
         .app_state.authz.authorization += [
             $list[] |
             {
                 "granter": .operator,
                 "grantee": $address,
                 "authorization": {
                     "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
                     "msg": "/zetachain.zetacore.emissions.MsgWithdrawEmission"
                 },
                 "expiration": null
             }
         ]
     ' "$json_file" > temp.json; then
         echo "Error: Failed to update genesis file"
         return 1
     fi
     mv temp.json "$json_file"
}



# create keys
CHAINID="athens_101-1"
KEYRING="test"
HOSTNAME=$(hostname)
INDEX=${HOSTNAME:0-1}

# Environment variables used for upgrade testing
export DAEMON_HOME=$HOME/.zetacored
export DAEMON_NAME=zetacored
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true
export CLIENT_DAEMON_NAME=zetaclientd
export CLIENT_DAEMON_ARGS="-enable-chains,GOERLI,-val,operator"
export DAEMON_DATA_BACKUP_DIR=$DAEMON_HOME
export CLIENT_SKIP_UPGRADE=true
export CLIENT_START_PROCESS=false
export UNSAFE_SKIP_BACKUP=true

echo "HOSTNAME: $HOSTNAME"

# init ssh keys
# we generate keys at runtime to ensure that keys are never pushed to
# a docker registry
if [ $HOSTNAME == "zetacore0" ]; then
  if [[ ! -f ~/.ssh/id_rsa ]]; then
    ssh-keygen -t rsa -q -N "" -f ~/.ssh/id_rsa
    cp ~/.ssh/id_rsa.pub ~/.ssh/authorized_keys
    # keep localtest.pem for compatibility
    cp ~/.ssh/id_rsa ~/.ssh/localtest.pem
    chmod 600 ~/.ssh/*
  fi
fi

# Wait for authorized_keys file to exist (zetacore1+)
while [ ! -f ~/.ssh/authorized_keys ]; do
    echo "Waiting for authorized_keys file to exist..."
    sleep 1
done

# Skip init if it has already been completed (marked by presence of ~/.zetacored/init_complete file)
if [[ ! -f ~/.zetacored/init_complete ]]
then
  # Init a new node to generate genesis file .
  # Copy config files from existing folders which get copied via Docker Copy when building images
  mkdir -p ~/.backup/config
  zetacored init Zetanode-Localnet --chain-id=$CHAINID --default-denom azeta
  rm -rf ~/.zetacored/config/app.toml
  rm -rf ~/.zetacored/config/client.toml
  rm -rf ~/.zetacored/config/config.toml
  cp -r ~/zetacored/common/app.toml ~/.zetacored/config/
  cp -r ~/zetacored/common/client.toml ~/.zetacored/config/
  cp -r ~/zetacored/common/config.toml ~/.zetacored/config/
  sed -i -e "/moniker =/s/=.*/= \"$HOSTNAME\"/" "$HOME"/.zetacored/config/config.toml

  # Add two new keys for operator and hotkey and create the required json structure for os_info
  source ~/add-keys.sh
fi

# Pause other nodes so that the primary can node can do the genesis creation
if [ $HOSTNAME != "zetacore0" ]
then
  while [ ! -f ~/.zetacored/config/genesis.json ]; do
    echo "Waiting for genesis.json file to exist..."
    sleep 1
  done
  # need to wait for zetacore0 to be up
  while ! curl -s -o /dev/null zetacore0:26657/status ; do
    echo "Waiting for zetacore0 rpc"
    sleep 1
done
fi

# Genesis creation following steps
# 1. Accumulate all the os_info files from other nodes on zetcacore0 and create a genesis.json
# 2. Add the observers , authorizations and required params to the genesis.json
# 3. Copy the genesis.json to all the nodes .And use it to create a gentx for every node
# 4. Collect all the gentx files in zetacore0 and create the final genesis.json
# 5. Copy the final genesis.json to all the nodes and start the nodes
# 6. Update Config in zetacore0 so that it has the correct persistent peer list
# 7. Start the nodes

# Start of genesis creation . This is done only on zetacore0.
# Skip genesis if it has already been completed (marked by presence of ~/.zetacored/init_complete file)
if [[ $HOSTNAME == "zetacore0" && ! -f ~/.zetacored/init_complete ]]
then
  ZETACORED_REPLICAS=2
  if host zetacore3 ; then
    echo "zetacore3 exists, setting ZETACORED_REPLICAS to 4"
    ZETACORED_REPLICAS=4
  fi
  # generate node list
  START=1
  # shellcheck disable=SC2100
  END=$((ZETACORED_REPLICAS - 1))
  NODELIST=()
  for i in $(eval echo "{$START..$END}")
  do
    NODELIST+=("zetacore$i")
  done

  # Misc : Copying the keyring to the client nodes so that they can sign the transactions
  ssh zetaclient0 mkdir -p ~/.zetacored/keyring-test/
  scp ~/.zetacored/keyring-test/* zetaclient0:~/.zetacored/keyring-test/
  ssh zetaclient0 mkdir -p ~/.zetacored/keyring-file/
  scp ~/.zetacored/keyring-file/* zetaclient0:~/.zetacored/keyring-file/

# 1. Accumulate all the os_info files from other nodes on zetcacore0 and create a genesis.json
  for NODE in "${NODELIST[@]}"; do
    INDEX=${NODE:0-1}
    ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/
    while ! scp "$NODE":~/.zetacored/os_info/os.json ~/.zetacored/os_info/os_z"$INDEX".json; do
      echo "Waiting for os_info.json from node $NODE"
      sleep 1
    done
    scp ~/.zetacored/os_info/os_z"$INDEX".json zetaclient"$INDEX":~/.zetacored/os.json
  done

  ssh zetaclient0 mkdir -p ~/.zetacored/
  scp ~/.zetacored/os_info/os.json zetaclient0:/root/.zetacored/os.json

# 2. Add the observers, authorizations, required params and accounts to the genesis.json
  zetacored collect-observer-info
  zetacored add-observer-list --keygen-block 25

  # Add emissions withdraw authorizations
  if ! add_emissions_withdraw_authorizations; then
      echo "Error: Failed to add emissions withdraw authorizations"
      exit 1
  fi

  # Update governance and other chain parameters for localnet
  jq '
    .app_state.gov.params.voting_period="30s" |
    .app_state.gov.params.quorum="0.1" |
    .app_state.gov.params.threshold="0.1" |
    .app_state.gov.params.expedited_voting_period = "10s" |
    .app_state.gov.deposit_params.min_deposit[0].denom = "azeta" |
    .app_state.gov.params.min_deposit[0].denom = "azeta" |
    .app_state.staking.params.bond_denom = "azeta" |
    .app_state.crisis.constant_fee.denom = "azeta" |
    .app_state.mint.params.mint_denom = "azeta" |
    .app_state.evm.params.evm_denom = "azeta" |
    .app_state.feemarket.params.min_gas_price = "10000000000.0000" |
    .consensus.params.block.max_gas = "500000000"
  ' "$HOME/.zetacored/config/genesis.json" > "$HOME/.zetacored/config/tmp_genesis.json" \
    && mv "$HOME/.zetacored/config/tmp_genesis.json" "$HOME/.zetacored/config/genesis.json"

# set admin account
  zetacored add-genesis-account zeta1n0rn6sne54hv7w2uu93fl48ncyqz97d3kty6sh 100000000000000000000000000azeta # Funds the localnet_gov_admin account

  emergency_policy=$(yq -r '.policy_accounts.emergency_policy_account.bech32_address' /root/config.yml)
  admin_policy=$(yq -r '.policy_accounts.admin_policy_account.bech32_address' /root/config.yml)
  operational_policy=$(yq -r '.policy_accounts.operational_policy_account.bech32_address' /root/config.yml)

  zetacored add-genesis-account "$emergency_policy" 100000000000000000000000000azeta
  zetacored add-genesis-account "$admin_policy" 100000000000000000000000000azeta
  zetacored add-genesis-account "$operational_policy" 100000000000000000000000000azeta

  jq --arg emergency "$emergency_policy" \
    --arg operational "$operational_policy" \
    --arg admin "$admin_policy" '
      .app_state.authority.policies.items[0].address = $emergency |
      .app_state.authority.policies.items[1].address = $operational |
      .app_state.authority.policies.items[2].address = $admin
  ' "$HOME/.zetacored/config/genesis.json" > "$HOME/.zetacored/config/tmp_genesis.json" \
    && mv "$HOME/.zetacored/config/tmp_genesis.json" "$HOME/.zetacored/config/genesis.json"

# give balance to runner accounts to deploy contracts directly on zEVM
# default account
  address=$(yq -r '.default_account.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# legacy erc20 tester
  address=$(yq -r '.additional_accounts.user_legacy_erc20.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# legacy zeta tester
  address=$(yq -r '.additional_accounts.user_legacy_zeta.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# legacy ethers tester
  address=$(yq -r '.additional_accounts.user_legacy_ether.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# bitcoin deposit tester
  address=$(yq -r '.additional_accounts.user_bitcoin_deposit.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# bitcoin withdraw tester
  address=$(yq -r '.additional_accounts.user_bitcoin_withdraw.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# solana tester
  address=$(yq -r '.additional_accounts.user_solana.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# sui tester
  address=$(yq -r '.additional_accounts.user_sui.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta# migration tester
  address=$(yq -r '.additional_accounts.user_migration.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# precompile tester
  address=$(yq -r '.additional_accounts.user_precompile.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# ether tester
  address=$(yq -r '.additional_accounts.user_ether.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# erc20 tester
  address=$(yq -r '.additional_accounts.user_erc20.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# ether revert tester
  address=$(yq -r '.additional_accounts.user_ether_revert.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# erc20 revert tester
  address=$(yq -r '.additional_accounts.user_erc20_revert.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta
# emissions withdraw tester
  address=$(yq -r '.additional_accounts.user_emissions_withdraw.bech32_address' /root/config.yml)
  zetacored add-genesis-account "$address" 100000000000000000000000000azeta

# 3. Copy the genesis.json to all the nodes .And use it to create a gentx for every node
  zetacored gentx operator 1000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING --gas-prices 20000000000azeta
  # Copy host gentx to other nodes
  for NODE in "${NODELIST[@]}"; do
    ssh $NODE mkdir -p ~/.zetacored/config/gentx/peer/
    scp ~/.zetacored/config/gentx/* $NODE:~/.zetacored/config/gentx/peer/
  done
  # Create gentx files on other nodes and copy them to host node
  mkdir ~/.zetacored/config/gentx/z2gentx
  for NODE in "${NODELIST[@]}"; do
      ssh $NODE rm -rf ~/.zetacored/genesis.json
      scp ~/.zetacored/config/genesis.json $NODE:~/.zetacored/config/genesis.json
      ssh $NODE zetacored gentx operator 1000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING
      scp $NODE:~/.zetacored/config/gentx/* ~/.zetacored/config/gentx/
      scp $NODE:~/.zetacored/config/gentx/* ~/.zetacored/config/gentx/z2gentx/
  done

#  TODO : USE --modify flag to modify the genesis file when v18 is released
  if [[ -n "$ZETACORED_IMPORT_GENESIS_DATA" ]]; then
    echo "Importing data"
    zetacored parse-genesis-file /root/genesis_data/exported-genesis.json
  fi

# 4. Collect all the gentx files in zetacore0 and create the final genesis.json
  zetacored collect-gentxs
  zetacored validate-genesis
# 5. Copy the final genesis.json to all the nodes
  for NODE in "${NODELIST[@]}"; do
      ssh $NODE rm -rf ~/.zetacored/genesis.json
      scp ~/.zetacored/config/genesis.json $NODE:~/.zetacored/config/genesis.json
  done
# 6. Update Config in zetacore0 so that it has the correct persistent peer list
   pp=$(cat $HOME/.zetacored/config/gentx/z2gentx/*.json | jq '.body.memo' )
   pps=${pp:1:58}
   sed -i -e 's/^persistent_peers =.*/persistent_peers = "'$pps'"/' "$HOME"/.zetacored/config/config.toml
fi
# End of genesis creation steps . The steps below are common to all the nodes

# Update persistent peers
if [[ $HOSTNAME != "zetacore0" && ! -f ~/.zetacored/init_complete ]]
then
  # Misc : Copying the keyring to the client nodes so that they can sign the transactions
  ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/keyring-test/
  scp ~/.zetacored/keyring-test/* "zetaclient$INDEX":~/.zetacored/keyring-test/
  ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/keyring-file/
  scp ~/.zetacored/keyring-file/* "zetaclient$INDEX":~/.zetacored/keyring-file/

  pp=$(cat $HOME/.zetacored/config/gentx/peer/*.json | jq '.body.memo' )
  pps=${pp:1:58}
  sed -i -e "/persistent_peers =/s/=.*/= \"$pps\"/" "$HOME"/.zetacored/config/config.toml
fi

# mark init completed so we skip it if container is restarted
touch ~/.zetacored/init_complete

cosmovisor run start --pruning=nothing --minimum-gas-prices=0.0001azeta --json-rpc.api eth,txpool,personal,net,debug,web3,miner --api.enable --home /root/.zetacored