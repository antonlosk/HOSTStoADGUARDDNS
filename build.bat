@echo off
chcp 65001 > nul
echo Компиляция программы в HOSTStoADGUARDDNS.exe...

:: Команда для компиляции с указанием выходного имени
go build -o HOSTStoADGUARDDNS.exe main.go

:: Проверка на успешное завершение
if %errorlevel% equ 0 (
    echo.
    echo Успешно! Файл HOSTStoADGUARDDNS.exe создан.
) else (
    echo.
    echo Ошибка при компиляции. Проверьте правильность кода и установлен ли Go.
)

echo.
pause
