#!/bin/sh
ls
echo "Clean Older builds"
rm -rf ./bin ./vendor

echo "Build executables"
env GOOS=linux go build -ldflags="-s -w" -o bin/server src/main.go

serverless deploy --aws-profile servicemind-prod --region us-east-2