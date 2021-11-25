#!/bin/sh
ls
echo "Starting Server Locally"

./node_modules/nodemon/bin/nodemon.js --watch src -e go,js --exec 'env GOOS=linux IS_RUNNING_WO_SERVERLESS=true go run src/main.go' --signal SIGTERM

#sls offline --host=0.0.0.0 --useDocker