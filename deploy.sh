#!/bin/bash
set -e

# variables
FUNC_NAME=test

# choosing aws profile
export AWS_PROFILE=xich-cn

# build binary for amazonlinux os (the lambda runtime OS)
mkdir -p tmp/gocache
docker build . -f docker/Dockerfile.builder -t amazonlinux-go-builder
docker run -v $(pwd)/src:/app/src -v $(pwd)/tmp/gocache:/gocache -e "GOCACHE=/gocache" amazonlinux-go-builder

# upload binary to lambda
zip -j tmp/function.zip src/out/handler
aws lambda update-function-code --function-name $FUNC_NAME --zip-file fileb://tmp/function.zip > /dev/null
