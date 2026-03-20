@echo off
cd /d "%~dp0"

echo Compiling to HOSTStoADGUARDDNS.exe...

go build -o HOSTStoADGUARDDNS.exe main.go

if %errorlevel% equ 0 (
    echo.
    echo SUCCESS! HOSTStoADGUARDDNS.exe created.
) else (
    echo.
    echo ERROR: Compilation failed.
)

echo.
pause
