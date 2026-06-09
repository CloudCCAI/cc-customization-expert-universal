@echo off
setlocal

set "SCRIPT_DIR=%~dp0"
for %%I in ("%SCRIPT_DIR%\..\..") do set "ROOT_DIR=%%~fI"

set "ARCH=%PROCESSOR_ARCHITECTURE%"
if /I "%ARCH%"=="AMD64" set "ARCH=amd64"
if /I "%ARCH%"=="ARM64" set "ARCH=arm64"
if /I "%ARCH%"=="x86" set "ARCH=386"

set "BIN=%ROOT_DIR%\tools\bin-windows-%ARCH%\cloudcc.exe"
if exist "%BIN%" (
  "%BIN%" %*
  exit /b %ERRORLEVEL%
)

echo Bundled Go cloudcc binary is missing: %BIN% 1>&2
echo Use a universal release package that includes tools\bin-windows-%ARCH%\cloudcc.exe. 1>&2
exit /b 127
