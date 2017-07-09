#!/bin/sh
docker run --rm -e CGO_ENABLED=0 -v "$PWD:/go/src/github.com/pinfake/pes6go" golang go build -ldflags "-linkmode external -extldflags -static" -o /go/src/github.com/pinfake/pes6go/bin/pes6go /go/src/github.com/pinfake/pes6go/main/pes6go.go
