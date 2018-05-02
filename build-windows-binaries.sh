#!/bin/sh
docker run --rm -e CGO_ENABLED=0 -e GOOS=windows -e GOARCH=386 -v "$PWD:/go/src/github.com/pinfake/pes6go" golang /bin/bash -c "go get github.com/pinfake/pes6go/... && go build -o /go/src/github.com/pinfake/pes6go/bin/pes6go.exe /go/src/github.com/pinfake/pes6go/main/pes6go.go"
