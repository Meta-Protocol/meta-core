#!/bin/bash

run_ssm_cmds_validators () {
    COMMAND_ID=$(aws ssm send-command \
    --targets Key=tag:ZetaChainValidator,Values=true \
    --document-name "AWS-RunShellScript" \
    --parameters "commands=[$1]" | jq .Command.CommandId -r || exit 1)
    echo "$COMMAND_ID"
}

run_ssm_cmds_indexer () {
    COMMAND_ID=$(aws ssm send-command \
    --targets Key=tag:Name,Values=zetachain-indexer \
    --document-name "AWS-RunShellScript" \
    --parameters "commands=[$1]" | jq .Command.CommandId -r || exit 1)
    echo "$COMMAND_ID"
}

check_cmd_status () {
    COMMAND_ID=$1
    echo "COMMAND_ID: $COMMAND_ID"
    COMMAND_STATUS=$(aws ssm list-commands --command-id "$COMMAND_ID" | jq '.Commands[0].Status' -r)
    until [[ "$COMMAND_STATUS" == "Success" || "$COMMAND_STATUS" == "Failed" ]]; do 
        echo "Waiting for Command to complete. ID: $COMMAND_ID | Status: $COMMAND_STATUS"
        sleep 2
        COMMAND_STATUS=$(aws ssm list-commands --command-id "$COMMAND_ID" | jq '.Commands[0].Status' -r)
    done
    echo "Complete. ID: $COMMAND_ID | Status: $COMMAND_STATUS"
    if [ "$COMMAND_STATUS" == "Failed" ]; then
        echo "Command ID $COMMAND_ID Failed" && exit 1
    fi
}



