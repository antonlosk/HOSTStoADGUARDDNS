package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "hosts.txt"
	outputFile := "hosts_ADGUARDDNS.txt"

	// Open input file
	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer in.Close()

	// Create output file
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer out.Close()

	// Use buffered I/O for better performance
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimRight(line, " \t\r\n")

		// Keep empty lines
		if trimmedLine == "" {
			writer.WriteString("\n")
			continue
		}

		// Keep comments
		if strings.HasPrefix(trimmedLine, "#") {
			writer.WriteString(trimmedLine + "\n")
			continue
		}

		// Process "ip domain" lines
		parts := strings.Fields(trimmedLine)
		if len(parts) >= 2 {
			ip := parts[0]
			domain := parts[1]
			
			// Format: |domain^$dnsrewrite=ip
			newLine := fmt.Sprintf("|%s^$dnsrewrite=%s\n", domain, ip)
			writer.WriteString(newLine)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// Flush the buffer to ensure all data is written to the file
	writer.Flush()
	fmt.Println("Done: created file hosts_ADGUARDDNS.txt")
}
