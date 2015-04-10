#!/bin/bash

# Exit on error and show the commands as they execute.
set -ex

# Ensure we get the full path of this script's directory.
PROGDIR=`dirname $0`
PROGDIR=`cd $PROGDIR && pwd`

source $PROGDIR/setup_env_common.txt
source $PROGDIR/setup_toolchain_linux64.txt

crosscompile_windows=1
run_integration_tests=0
use_xvfb=0

function show_help {
  # Turn off command echoing so the help message is readable.
  set +x
  echo "USAGE: "`basename $0`" [-h] [-i] [-w] [-x]"
  echo "  -h    Show this message."
  echo "  -i    Run integration tests."
  echo "  -w    Do NOT cross-compile the server for Windows."
  echo "  -x    Start an Xvfb-randr server for running integration tests"
  echo "        without an X server."
}

while getopts "h?iwx" opt; do
    case "$opt" in
    h|\?)  show_help
        exit 0
        ;;
    i)  run_integration_tests=1
        ;;
    w)  crosscompile_windows=0
        ;;
    x)  use_xvfb=1
        ;;
    esac
done

cd $GPU_BUILD_ROOT

if [ $use_xvfb -eq 1 ]; then
  # Start a headless X server on display :42 for integration tests.
  # Xvfb is not able to run the integration tests (the cause is not entirely
  # understood, but it doesn't appear to expose GLX to the X client).
  # Instead, run Xvfb-randr, which appears to be installed as part of
  # Chrome remote desktop.
  Xvfb-randr :42 +extension GLX -screen 0 1280x1024x24 -noreset &
  export XVFB_PID=$!
  export DISPLAY=:42
fi

export GO_BUILD_FLAGS="-i -v -x -o"
export GO_GENERATE_FLAGS="-v -x"
export GO_TEST_FLAGS="-v -x"

go build    $GO_BUILD_FLAGS    $GPU_BUILD_ROOT/bin/embed    $GPU_RELATIVE_SOURCE_PATH/tools/embed
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/tools/copyright
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/binary/generate
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/rpc/generate
go build    $GO_BUILD_FLAGS    $GPU_BUILD_ROOT/bin/codergen $GPU_RELATIVE_SOURCE_PATH/binary/codergen
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/replay/protocol
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/rpc
go build    $GO_BUILD_FLAGS    $GPU_BUILD_ROOT/bin/rpcapi   $GPU_RELATIVE_SOURCE_PATH/rpc/rpcapi
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/rpc/test
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/service
go build    $GO_BUILD_FLAGS    $GPU_BUILD_ROOT/bin/apic     $GPU_RELATIVE_SOURCE_PATH/api/apic
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/gfxapi/test
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/gfxapi/gles
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/atom
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/builder
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/database
go generate $GO_GENERATE_FLAGS                              $GPU_RELATIVE_SOURCE_PATH/memory
go build    $GO_BUILD_FLAGS    $GPU_BUILD_ROOT/bin/gazer    $GPU_RELATIVE_SOURCE_PATH/server/cmd

go run src/$GPU_RELATIVE_SOURCE_PATH/cc/build.go --v --runtests

# Kill any existing replay daemon before running tests.
killall replayd || true

# Run non-integration tests.
go list android.googlesource.com/platform/tools/gpu/... | egrep -v '/integration/?' | xargs go test $GO_TEST_FLAGS

if [ $run_integration_tests -eq 1 ]; then
  # Run the integration tests.
  go list android.googlesource.com/platform/tools/gpu/... | egrep '/integration/?' | xargs go test $GO_TEST_FLAGS
fi

if [ ! -z $XVFB_PID ]; then
  kill $XVFB_PID
fi

killall replayd || true

if [ $crosscompile_windows -eq 1 ]; then
  go run src/$GPU_RELATIVE_SOURCE_PATH/cc/build.go --v --target=windows
  source $PROGDIR/setup_toolchain_linux_xc_win64.txt
  go build $GO_BUILD_FLAGS $GPU_BUILD_ROOT/bin/windows_amd64/gazer.exe -ldflags="-extld=$CC" $GPU_RELATIVE_SOURCE_PATH/server/cmd
fi
