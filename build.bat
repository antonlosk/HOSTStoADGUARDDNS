@echo off
REM Build script for HOSTStoADGUARDDNS
REM Optimization flags used:
REM -s : Disable symbol table
REM -w : Disable DWARF generation
REM -trimpath : Remove all file system paths from the resulting executable

echo Building optimized executable...
go build -ldflags="-s -w" -trimpath -o HOSTStoADGUARDDNS.exe main.go

if %errorlevel% equ 0 (
    echo Build successful: HOSTStoADGUARDDNS.exe has been created.
) else (
    echo Build failed!
)
pause
