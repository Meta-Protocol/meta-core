#!/bin/bash

echo $1 $2

NODE_NUMBER=$1
SEED_NODE=$2

echo "Starting ZetaClient Node $NODE_NUMBER"
 
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/root/go/bin
export MYIP=$(hostname -i)
export IDX=$NODE_NUMBER 
export TSSPATH=/root/.tssnew 

if  (( $NODE_NUMBER == 0 )); then
    sleep 5 # Wait for Zetacored to start
    yes | zetaclientd -val val 2>&1 | tee ~/.zetaclient/zetaclient.log
else
    until [ -f SEED_NODE_ID ]
        do
            echo "Waiting for Seed Node Validator ID"
            sleep 10
            curl -s ${SEED_NODE}:8123/p2p -o SEED_NODE_ID

        done
    SEED_NODE_ID=$(cat SEED_NODE_ID)
    echo "SEED_NODE_ID=${SEED_NODE_ID}"
    
    yes | zetaclientd -val val \
        --peer /dns/${SEED_NODE}/tcp/6668/p2p/${SEED_NODE_ID} \
        2>&1 | tee ~/.zetaclient/zetaclient.log

fi



