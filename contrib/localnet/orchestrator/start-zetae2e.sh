#!/bin/bash

# The script run the zetae2e CLI to run local end-to-end tests
# First argument is the command to run the local e2e
# A second optional argument can be passed and can have the following value:
# upgrade: run the local e2e once, then restart zetaclientd at upgrade height and run the local e2e again

get_zetacored_version() {
  retries=10
  node_info=""
  for ((attempt=1; attempt<=$retries; attempt++)); do
    node_info=$(curl -s -f zetacore0:1317/cosmos/base/tendermint/v1beta1/node_info)
    if [[ $? == 0 ]]; then
      version=$(echo "$node_info" | jq -r '.application_version.version')
      # only return versions containing dots to avoid empty strings and "null"
      if [[ "$version" == *.* ]]; then
        echo "$version"
        return
      fi
    fi
    sleep 1
  done
  echo "Unable to get zetacored version after ${retries} retries"
  exit 1
}

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

echo "waiting for geth RPC to start..."
sleep 2

### Create the accounts and fund them with Ether on local Ethereum network

# unlock the default account account
address=$(yq -r '.default_account.evm_address' config.yml)
echo "funding deployer address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock erc20 tester accounts
address=$(yq -r '.additional_accounts.user_erc20.evm_address' config.yml)
echo "funding erc20 address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock zeta tester accounts
address=$(yq -r '.additional_accounts.user_zeta_test.evm_address' config.yml)
echo "funding zeta tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock zevm message passing tester accounts
address=$(yq -r '.additional_accounts.user_zevm_mp_test.evm_address' config.yml)
echo "funding zevm mp tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock bitcoin tester accounts
address=$(yq -r '.additional_accounts.user_bitcoin.evm_address' config.yml)
echo "funding bitcoin tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock solana tester accounts
address=$(yq -r '.additional_accounts.user_solana.evm_address' config.yml)
echo "funding solana tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock ethers tester accounts
address=$(yq -r '.additional_accounts.user_ether.evm_address' config.yml)
echo "funding ether tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock miscellaneous tests accounts
address=$(yq -r '.additional_accounts.user_misc.evm_address' config.yml)
echo "funding misc tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock admin erc20 tests accounts
address=$(yq -r '.additional_accounts.user_admin.evm_address' config.yml)
echo "funding admin tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock migration tests accounts
address=$(yq -r '.additional_accounts.user_migration.evm_address' config.yml)
echo "funding migration tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock precompile tests accounts
address=$(yq -r '.additional_accounts.user_precompile.evm_address' config.yml)
echo "funding migration tester address ${address} with 10000 Ether"
geth --exec "eth.sendTransaction({from: eth.coinbase, to: '${address}', value: web3.toWei(10000,'ether')})" attach http://eth:8545 > /dev/null

# unlock local solana relayer accounts
if host solana > /dev/null; then
  solana_url=$(yq -r '.rpcs.solana' config.yml)
  solana config set --url "$solana_url" > /dev/null

  relayer=$(yq -r '.observer_relayer_accounts.relayer_accounts[0].solana_address' config.yml)
  echo "funding solana relayer address ${relayer} with 100 SOL"
  solana airdrop 100 "$relayer" > /dev/null

  relayer=$(yq -r '.observer_relayer_accounts.relayer_accounts[1].solana_address' config.yml)
  echo "funding solana relayer address ${relayer} with 100 SOL"
  solana airdrop 100 "$relayer" > /dev/null
fi

### Run zetae2e command depending on the option passed

# Mode migrate is used to run the e2e tests before and after the TSS migration
# It runs the e2e tests with the migrate flag which triggers a TSS migration at the end of the tests. Once the migrationis done the first e2e test is complete
# The second e2e test is run after the migration to ensure the network is still working as expected with the new tss address
if [ "$LOCALNET_MODE" == "tss-migrate" ]; then
  if [[ ! -f deployed.yml ]]; then
    zetae2e local $E2E_ARGS --setup-only --config config.yml --config-out deployed.yml --skip-header-proof
    if [ $? -ne 0 ]; then
      echo "e2e setup failed"
      exit 1
    fi
  else
    echo "skipping e2e setup because it has already been completed"
  fi

  echo "running e2e test before migrating TSS"
  zetae2e local $E2E_ARGS --skip-setup --config deployed.yml --skip-header-proof
  if [ $? -ne 0 ]; then
    echo "first e2e failed"
    exit 1
  fi

  echo "waiting 10 seconds for node to restart"
    sleep 10

  zetae2e local --skip-setup --config deployed.yml --skip-bitcoin-setup --light --skip-header-proof
  ZETAE2E_EXIT_CODE=$?
  if [ $ZETAE2E_EXIT_CODE -eq 0 ]; then
    echo "E2E passed after migration"
    exit 0
  else
    echo "E2E failed after migration"
    exit 1
  fi
fi


# Mode upgrade is used to run the e2e tests before and after the upgrade
# It runs the e2e tests , waits for the upgrade height to be reached, and then runs the e2e tests again once the ungrade is done.
# The second e2e test is run after the upgrade to ensure the network is still working as expected with the new version
if [ "$LOCALNET_MODE" == "upgrade" ]; then

  # Run the e2e tests, then restart zetaclientd at upgrade height and run the e2e tests again

  # set upgrade height to 225 by default
  UPGRADE_HEIGHT=${UPGRADE_HEIGHT:=225}
  OLD_VERSION=$(get_zetacored_version)
  COMMON_ARGS="--skip-header-proof --skip-tracker-check"

  if [[ ! -f deployed.yml ]]; then
    zetae2e local $E2E_ARGS --setup-only --config config.yml --config-out deployed.yml ${COMMON_ARGS}
    if [ $? -ne 0 ]; then
      echo "e2e setup failed"
      exit 1
    fi
  else
    echo "skipping e2e setup because it has already been completed"
  fi

  # Run zetae2e, if the upgrade height is greater than 100 to populate the state
  if [ "$UPGRADE_HEIGHT" -gt 100 ]; then
    echo "running E2E command to setup the networks and populate the state..."

    # Use light flag to ensure tests can complete before the upgrade height
    zetae2e local $E2E_ARGS --skip-setup --config deployed.yml --light --skip-precompiles ${COMMON_ARGS}
    if [ $? -ne 0 ]; then
      echo "first e2e failed"
      exit 1
    fi
  fi

  echo "Waiting for upgrade height..."
  CURRENT_HEIGHT=0
  WAIT_HEIGHT=$(( UPGRADE_HEIGHT - 1 ))
  # wait for upgrade height
  while [[ $CURRENT_HEIGHT -lt $WAIT_HEIGHT ]]
  do
    CURRENT_HEIGHT=$(curl -s zetacore0:26657/status | jq -r '.result.sync_info.latest_block_height')
    echo current height is "$CURRENT_HEIGHT", waiting for "$WAIT_HEIGHT"
    sleep 2
  done

  echo "waiting 10 seconds for node to restart..."
  sleep 10

  NEW_VERSION=$(get_zetacored_version)

  echo "upgrade result: ${OLD_VERSION} -> ${NEW_VERSION}"

  if [[ "$OLD_VERSION" == "$NEW_VERSION" ]]; then
    echo "version did not change after upgrade height, maybe the upgrade did not run?"
    exit 2
  fi

  # wait for zevm endpoint to come up
  sleep 10

  echo "running E2E command to test the network after upgrade..."

  # Run zetae2e again
  # When the upgrade height is greater than 100 for upgrade test, the Bitcoin tests have been run once, therefore the Bitcoin wallet is already set up
  # Use light flag to skip advanced tests
  if [ "$UPGRADE_HEIGHT" -lt 100 ]; then
    zetae2e local $E2E_ARGS --skip-setup --config deployed.yml --light --skip-precompiles ${COMMON_ARGS}
  else
    zetae2e local $E2E_ARGS --skip-setup --config deployed.yml --skip-bitcoin-setup --light --skip-precompiles ${COMMON_ARGS}
  fi

  ZETAE2E_EXIT_CODE=$?
  if [ $ZETAE2E_EXIT_CODE -eq 0 ]; then
    echo "E2E passed after upgrade"
    exit 0
  else
    echo "E2E failed after upgrade"
    exit 1
  fi

else
  # If no mode is passed, run the e2e tests normally
  echo "running e2e setup..."

  if [[ ! -f deployed.yml ]]; then
    zetae2e local $E2E_ARGS --config config.yml --setup-only --config-out deployed.yml
    if [ $? -ne 0 ]; then
      echo "e2e setup failed"
      exit 1
    fi
  else
    echo "skipping e2e setup because it has already been completed"
  fi

  if [ "$LOCALNET_MODE" == "setup-only" ]; then
    exit 0
  fi

  echo "running e2e tests with arguments: $E2E_ARGS"

  zetae2e local $E2E_ARGS --skip-setup --config deployed.yml
  ZETAE2E_EXIT_CODE=$?

  # if e2e passed, exit with 0, otherwise exit with 1
  if [ $ZETAE2E_EXIT_CODE -eq 0 ]; then
    echo "e2e passed"
    exit 0
  else
    echo "e2e failed"
    exit 1
  fi

fi
