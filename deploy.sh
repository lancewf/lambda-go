#!/bin/bash

GOOS=linux go build main.go
zip function.zip main
aws lambda update-function-code --function-name GetStartedLambdaGo2 --zip-file fileb://function.zip --region us-west-2
