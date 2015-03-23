set PROGDIR=%~dp0

:: Top-level directory from which to run go commands.  This is currently the
:: same as %GOPATH%.
set GPU_BUILD_ROOT=%PROGDIR%..\..\..\..\..\..\

:: The root of the Android Studio repo.
set REPO_ROOT=%PROGDIR%..\..\..\..\..\..\..\..\

:: The path to the gpu source code (relative to %GPU_BUILD_ROOT%\src).
set GPU_RELATIVE_SOURCE_PATH=android.googlesource.com\platform\tools\gpu

set GOROOT=%REPO_ROOT%\prebuilts\go\windows-x86
set GOPATH=%GPU_BUILD_ROOT%

:: Add the pre-built go tools to the path.
set PATH=%GOROOT%\bin;%PATH%

:: Add the tools we build to the path.
set PATH=%GPU_BUILD_ROOT%\bin;%PATH%
