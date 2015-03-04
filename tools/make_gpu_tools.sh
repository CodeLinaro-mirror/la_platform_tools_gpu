#!/bin/bash

# Exit on error and show the commands as they execute.
set -ex

# Ensure we get the full path of this script's directory.
PROGDIR=`dirname $0`
PROGDIR=`cd $PROGDIR && pwd`

source $PROGDIR/setup_env_linux.txt

cd $GPU_BUILD_ROOT

go build -i -o $GPU_BUILD_ROOT/bin/embed $GPU_RELATIVE_SOURCE_PATH/embed
go generate -x $GPU_RELATIVE_SOURCE_PATH/binary/generate
go generate -x $GPU_RELATIVE_SOURCE_PATH/rpc/generate
go build -i -o $GPU_BUILD_ROOT/bin/codergen $GPU_RELATIVE_SOURCE_PATH/binary/codergen
go generate -x $GPU_RELATIVE_SOURCE_PATH/replay/protocol
go generate -x $GPU_RELATIVE_SOURCE_PATH/rpc
go build -i -o $GPU_BUILD_ROOT/bin/rpcapi $GPU_RELATIVE_SOURCE_PATH/rpc/rpcapi
go generate -x $GPU_RELATIVE_SOURCE_PATH/rpc/test
go generate -x $GPU_RELATIVE_SOURCE_PATH/service
go build -i -o $GPU_BUILD_ROOT/bin/apic $GPU_RELATIVE_SOURCE_PATH/api/apic
go generate -x $GPU_RELATIVE_SOURCE_PATH/gfxapi/test
go generate -x $GPU_RELATIVE_SOURCE_PATH/gfxapi/gles
go build -i -o $GPU_BUILD_ROOT/bin/protoc-gen-go github.com/golang/protobuf/protoc-gen-go

# Skip the gradle build for now.
# TODO: Enable this once the checked in gcc tool chain is working.
#src/$GPU_RELATIVE_SOURCE_PATH/cc/gradlew -b src/$GPU_RELATIVE_SOURCE_PATH/cc/build.gradle

# The following two go actions have a mutual dependency because we are using
# an outdated version of golang.org/x/tools/go/loader.
# As a temporary work around, run the actions twice and ignore the errors
# during the first build action.
# TODO: Update golang.org/x/tools/go/loader and remove this work around.
go generate -x $GPU_RELATIVE_SOURCE_PATH/builder
set +e
go build -i -o $GPU_BUILD_ROOT/bin/gazer $GPU_RELATIVE_SOURCE_PATH/server/cmd
set -e

# Now repeat the actions.
go generate -x $GPU_RELATIVE_SOURCE_PATH/builder
go build -i -o $GPU_BUILD_ROOT/bin/gazer $GPU_RELATIVE_SOURCE_PATH/server/cmd

go test $GPU_RELATIVE_SOURCE_PATH/...
