@echo off
setlocal
REM Requires MinGW gcc on PATH (or set MINGW below).
if defined MINGW (
  set "PATH=%MINGW%;%PATH%"
  set "CC=%MINGW%\gcc.exe"
)
if not exist dist mkdir dist
set CGO_ENABLED=1
echo Building dist\grok-manager.dll ...
go build -buildvcs=false -buildmode=c-shared -trimpath -ldflags="-s -w" -o dist\grok-manager.dll .
if errorlevel 1 (
  echo BUILD FAILED
  exit /b 1
)
copy /y dist\grok-manager.dll dist\grok-manager-v1.0.1.dll >nul
echo OK
dir dist\grok-manager.dll
endlocal
