setlocal

echo on

:: Ensure we get the full path of this script's directory.
set PROGDIR=%~dp0

call %PROGDIR%setup_env_windows.bat

cd %GPU_BUILD_ROOT%

set GO_BUILD_FLAGS=-i -v -x -o
set GO_GENERATE_FLAGS=-v -x
set GO_TEST_FLAGS=-v -x

go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\embed.exe    %GPU_RELATIVE_SOURCE_PATH%\tools\embed
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\tools\copyright
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\binary\generate
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\rpc\generate
go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\codergen.exe %GPU_RELATIVE_SOURCE_PATH%\binary\codergen
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\replay\protocol
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\rpc
go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\rpcapi.exe   %GPU_RELATIVE_SOURCE_PATH%\rpc\rpcapi
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\rpc\test
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\service
go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\apic.exe     %GPU_RELATIVE_SOURCE_PATH%\api\apic
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\gfxapi\test
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\gfxapi\gles
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\atom
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\builder
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\database
go generate %GO_GENERATE_FLAGS%                                   %GPU_RELATIVE_SOURCE_PATH%\memory
go build    %GO_BUILD_FLAGS%    %GPU_BUILD_ROOT%\bin\gazer.exe    %GPU_RELATIVE_SOURCE_PATH%\server\cmd

:: TODO: enable gradle build for .cpp files.
::src\%GPU_RELATIVE_SOURCE_PATH%\cc\gradlew.bat -b src\%GPU_RELATIVE_SOURCE_PATH%\cc\build.gradle --info
