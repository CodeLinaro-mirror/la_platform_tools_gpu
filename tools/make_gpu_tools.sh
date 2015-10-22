#!/bin/bash

# Exit on error and show the commands as they execute.
set -ex

# Ensure we get the full path of this script's directory.
PROGDIR=`dirname $0`
PROGDIR=`cd $PROGDIR && pwd`

# Use osx-X86_64 instead of darwin-x64 for Mac.
# TODO: Switch the build to use darwin-x64 to be consistent with
# other Android repositories.
HOST_OS=$(uname | tr A-Z a-z | sed -e "s/darwin/osx/g")
HOST_ARCH="x86_64"

source $PROGDIR/setup_env_common.txt
source $PROGDIR/setup_toolchain_$HOST_OS-$HOST_ARCH.txt

if [[ $HOST_OS == "linux" ]]; then
  crosscompile_windows=1
else
  crosscompile_windows=0
fi

run_integration_tests=0
use_xvfb=0
build_pkginfo_apk=1

BUILD_NUMBER="SNAPSHOT-"`date "+%Y-%m-%dT%H:%M:%S%z"`
BUILD_FLAVOR="release"
DIST_DIR=$GPU_BUILD_ROOT/dist

function show_help {
  # Turn off command echoing so the help message is readable.
  set +x
  echo "USAGE: "`basename $0`" [options]"
  echo ""
  echo "  -b <name/number> Sets the build number (used for artifact naming)."
  echo "  -d <absolute path> Sets the distribution directory."
  echo "  -f <release|debug> Sets the build flavor."
  echo "  -h    Show this message."
  echo "  -i    Run integration tests."
  echo "  -p    Do NOT build the pkginfo APK."
  echo "  -w    Do NOT cross-compile the server for Windows."
  echo "  -x    Start an Xvfb-randr server for running integration tests"
  echo "        without an X server."
}

while getopts "b:d:f:h?iwx" opt; do
    case "$opt" in
    b)  BUILD_NUMBER=$OPTARG
        ;;
    d)  DIST_DIR=$OPTARG
        ;;
    f)  BUILD_FLAVOR=$OPTARG
        ;;
    h|\?)  show_help
        exit 0
        ;;
    i)  run_integration_tests=1
        ;;
    p)  build_pkginfo_apk=0
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
export GO_TEST_FLAGS="-v -x"

go build $GO_BUILD_FLAGS $GPU_BUILD_ROOT/bin/$HOST_OS-$HOST_ARCH/$BUILD_FLAVOR/gapis $GPU_RELATIVE_SOURCE_PATH/server/gapis
go build $GO_BUILD_FLAGS $GPU_BUILD_ROOT/bin/$HOST_OS-$HOST_ARCH/$BUILD_FLAVOR/gapit $GPU_RELATIVE_SOURCE_PATH/tools/gapit

# Kill any existing replay daemon before running tests.
killall gapir || true

go run src/$GPU_RELATIVE_SOURCE_PATH/make.go -f -v=1 --disable=code cc

# Kill any existing replay daemon before running tests.
killall gapir || true

# Run non-integration tests.
go list android.googlesource.com/platform/tools/gpu/... | egrep -v '/integration/?' | xargs go test $GO_TEST_FLAGS

if [ $run_integration_tests -eq 1 ]; then
  # Run the integration tests.
  go list android.googlesource.com/platform/tools/gpu/... | egrep '/integration/?' | xargs go test $GO_TEST_FLAGS
fi

if [ ! -z $XVFB_PID ]; then
  kill $XVFB_PID
fi

killall gapir || true

if [ $build_pkginfo_apk -eq 1 ]; then
  go run src/$GPU_RELATIVE_SOURCE_PATH/make.go -f -v=1 --disable=code pkginfo
fi

if [ $crosscompile_windows -eq 1 ]; then
  go run src/$GPU_RELATIVE_SOURCE_PATH/make.go -f -v=1 -targetos=windows --disable=code cc:gapir
  source $PROGDIR/setup_toolchain_linux_xc_win64.txt
  go build $GO_BUILD_FLAGS $GPU_BUILD_ROOT/bin/windows-$HOST_ARCH/$BUILD_FLAVOR/gapis.exe -ldflags="-extld=$CC -s" $GPU_RELATIVE_SOURCE_PATH/server/gapis
  go build $GO_BUILD_FLAGS $GPU_BUILD_ROOT/bin/windows-$HOST_ARCH/$BUILD_FLAVOR/gapit.exe -ldflags="-extld=$CC -s" $GPU_RELATIVE_SOURCE_PATH/tools/gapit
fi

# Create zip files for the build artifacts.
if [[ -n "$DIST_DIR" ]]; then
  mkdir -p $DIST_DIR
  DIST_DIR=$(cd $DIST_DIR && pwd)

  cd $GPU_BUILD_ROOT
  for TARGET_OS in linux windows osx; do
    if [[ $TARGET_OS == $HOST_OS || ( $HOST_OS == linux && $TARGET_OS == windows && $crosscompile_windows == 1 ) ]]; then
      if [[ $TARGET_OS == windows ]]; then
        EXE_EXTENSION=".exe"
      else
        EXE_EXTENSION=""
      fi

      OUT=$DIST_DIR/$TARGET_OS/gapid
      rm -rf $OUT
      mkdir -p $OUT/android/armeabi-v7a
      mkdir -p $OUT/android/arm64-v8a
      mkdir -p $OUT/$TARGET_OS/$HOST_ARCH

      cp bin/pkginfo.apk $OUT/android
      for ARTIFACT in libgapii.so gapir.apk; do
        cp bin/android-arm/$BUILD_FLAVOR/$ARTIFACT $OUT/android/armeabi-v7a
        cp bin/android-arm64/$BUILD_FLAVOR/$ARTIFACT $OUT/android/arm64-v8a
      done

      for ARTIFACT in gapir gapis gapit; do
        BINARY=bin/$TARGET_OS-$HOST_ARCH/$BUILD_FLAVOR/$ARTIFACT$EXE_EXTENSION
        # strip does not work on OS-X!
        if [ "$HOST_OS" != "osx" ]; then
          strip $BINARY
        fi

        cp $BINARY $OUT/$TARGET_OS/$HOST_ARCH
      done

      ZIP="$DIST_DIR/gapid_r01_$TARGET_OS.zip"
      rm -f $ZIP
      pushd $OUT/..
      zip -9r $ZIP gapid
      popd
    fi
  done
fi
