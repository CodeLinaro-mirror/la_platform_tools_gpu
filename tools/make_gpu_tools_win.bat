setlocal

echo on

:: Ensure we get the full path of this script's directory.
set PROGDIR=%~dp0

call %PROGDIR%setup_env_windows.bat

cd %GPU_BUILD_ROOT%

set GO_BUILD_FLAGS=-i -v -x -o
set GO_TEST_FLAGS=-v -x

go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\gapis.exe    %GPU_RELATIVE_SOURCE_PATH%\server\gapis

go run src/$GPU_RELATIVE_SOURCE_PATH/cc/build.go --v --f --runtests
