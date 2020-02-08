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
lambda-local start --host $HOST s--port $PORT --volume $VOLUME_APP --network partner-location-api_partners-network --env .local.env