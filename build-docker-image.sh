#!/bin/sh
source ./build-linux-binaries.sh
docker build -t pes6go .
docker rmi $(docker images -f "dangling=true" -q)