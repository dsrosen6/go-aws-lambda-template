#!/bin/bash

# Your variables
role="YOUR_ROLE_ARN"
function_name="YOUR_FUNCTION_NAME"

# Main Logic
export AWS_PAGER=""

GOOS=linux GOARCH=arm64 go build -o bootstrap main.go
zip function.zip bootstrap

aws lambda create-function --function-name "$function_name" \
--runtime provided.al2023 --handler bootstrap \
--architectures arm64 \
--role "$role" \
--zip-file fileb://function.zip