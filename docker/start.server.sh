#!/bin/bash

if [ -z $HOST ]
then
    HOST="0.0.0.0"
fi
if [ -z $PORT ]
then
    PORT="3000"
fi

gomon --exec "make build" &
make build
mkdir -p .data/dynamodb/ && touch .data/dynamodb/shared-local-instance.db && chmod 0777 -R .data
lambda-local start --host $HOST s--port $PORT --volume $VOLUME_APP --network partner-location-api_partners-network --env .local.env