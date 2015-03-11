#!/bin/bash

# Exit on error and show the commands as they execute.
set -ex

# Ensure we get the full path of this script's directory.
PROGDIR=`dirname $0`
PROGDIR=`cd $PROGDIR && pwd`

source $PROGDIR/setup_env_linux.txt

cd $GPU_BUILD_ROOT

go build -i -o $GPU_BUILD_ROOT/bin/embed $GPU_RELATIVE_SOURCE_PATH/tools/embed
go generate -x $GPU_RELATIVE_SOURCE_PATH/tools/copyright
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

src/$GPU_RELATIVE_SOURCE_PATH/cc/gradlew -b src/$GPU_RELATIVE_SOURCE_PATH/cc/build.gradle

go generate -x $GPU_RELATIVE_SOURCE_PATH/builder
go build -i -o $GPU_BUILD_ROOT/bin/gazer $GPU_RELATIVE_SOURCE_PATH/server/cmd

# Kill any existing replay daemon before running integration tests.
killall replayd || true

go test $GPU_RELATIVE_SOURCE_PATH/...

# Kill the integration tests' replay daemon.
killall replayd || true
