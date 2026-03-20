@echo off
cd /d "%~dp0"

echo Compiling and optimizing to HOSTStoADGUARDDNS.exe...

go build -ldflags "-s -w" -trimpath -o HOSTStoADGUARDDNS.exe main.go

if %errorlevel% equ 0 (
    echo.
    echo SUCCESS! Optimized HOSTStoADGUARDDNS.exe created.
) else (
    echo.
    echo ERROR: Compilation failed.
)

echo.
pause
