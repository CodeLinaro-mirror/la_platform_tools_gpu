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
go generate -x $GPU_RELATIVE_SOURCE_PATH/builder
go build -i -o $GPU_BUILD_ROOT/bin/gazer $GPU_RELATIVE_SOURCE_PATH/server/cmd

src/$GPU_RELATIVE_SOURCE_PATH/cc/gradlew -b src/$GPU_RELATIVE_SOURCE_PATH/cc/build.gradle

# Kill any existing replay daemon before running tests.
killall replayd || true

# Try starting a headless X server on display :42 for integration tests.
if [ -x "$(which Xvfb-randr)" ]; then
  Xvfb-randr :42  +extension GLX -screen 0 1280x1024x24 -noreset &
  export XVFB_PID=$!
  export DISPLAY=:42
fi

go test $GPU_RELATIVE_SOURCE_PATH/api/...
go test $GPU_RELATIVE_SOURCE_PATH/atexit/...
go test $GPU_RELATIVE_SOURCE_PATH/atom/...
go test $GPU_RELATIVE_SOURCE_PATH/binary/...
go test $GPU_RELATIVE_SOURCE_PATH/builder/...
go test $GPU_RELATIVE_SOURCE_PATH/database/...
go test $GPU_RELATIVE_SOURCE_PATH/gfxapi/...
go test $GPU_RELATIVE_SOURCE_PATH/image/...
go test $GPU_RELATIVE_SOURCE_PATH/interval/...
go test $GPU_RELATIVE_SOURCE_PATH/log/...
go test $GPU_RELATIVE_SOURCE_PATH/memory/...
go test $GPU_RELATIVE_SOURCE_PATH/multiplexer/...
go test $GPU_RELATIVE_SOURCE_PATH/parse/...
go test $GPU_RELATIVE_SOURCE_PATH/replay/...
go test $GPU_RELATIVE_SOURCE_PATH/ringbuffer/...
go test $GPU_RELATIVE_SOURCE_PATH/rpc/...
go test $GPU_RELATIVE_SOURCE_PATH/server/...
go test $GPU_RELATIVE_SOURCE_PATH/service/...
go test $GPU_RELATIVE_SOURCE_PATH/tools/...

if [ ! -z $XVFB_PID ]; then
  # Only run the integration tests if we were able to start the virtual X display.
  go test $GPU_RELATIVE_SOURCE_PATH/integration...
  kill $XVFB_PID
fi

killall replayd || true
