FROM ubuntu:latest


# Variables
VERSION=`git describe --tags --always`
BUILD=`date +%FT%T%z`
MODULE_PATH="github.com/GrolimundSolutions/aws_scheduler/cmd/scheduler/schedulermain"

# Setup the -ldflags option for go build here
LDFLAGS=-ldflags "-w -s -X ${MODULE_PATH}.Version=${VERSION} -X ${MODULE_PATH}.Build=${BUILD}"

# Build Command
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./${BINARY} ./cmd/scheduler/

ENTRYPOINT ["top", "-b"]