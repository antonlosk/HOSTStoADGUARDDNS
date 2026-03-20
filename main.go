package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputFileName := "hosts.txt"
	outputFileName := "hosts_ADGUARDDNS.txt"

	// Открываем исходный файл для чтения
	inFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла %s: %v", inputFileName, err)
	}
	defer inFile.Close()

	// Создаем выходной файл для записи
	outFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Ошибка при создании файла %s: %v", outputFileName, err)
	}
	defer outFile.Close()

	// Используем сканер для построчного чтения
	scanner := bufio.NewScanner(inFile)
	
	// Используем буферизированную запись для ускорения ввода-вывода
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()

		// Удаляем пробельные символы с конца строки (аналог rstrip в Python)
		line = strings.TrimRightFunc(line, unicode.IsSpace)

		// сохранить пустые строки
		if line == "" {
			writer.WriteString("\n")
			continue
		}

		// сохранить комментарии
		if strings.HasPrefix(line, "#") {
			writer.WriteString(line + "\n")
			continue
		}

		// разбиваем строку по любым пробельным символам (аналог split())
		parts := strings.Fields(line)

		// обработка строк вида "ip domain"
		if len(parts) >= 2 {
			ip := parts[0]
			domain := parts[1]

			// Формируем новую строку
			newLine := fmt.Sprintf("|%s^$dnsrewrite=%s\n", domain, ip)
			writer.WriteString(newLine)
		}
	}

	// Проверяем, не возникло ли ошибок во время чтения файла
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Примечание: в вашем Python-скрипте выводилось hosts_new.txt, 
	// здесь я подставил реальную переменную outputFileName для точности.
	fmt.Printf("Готово: создан файл %s\n", outputFileName)
}
