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

	// Open input file for reading
	inFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Error opening input file %s: %v", inputFileName, err)
	}
	defer inFile.Close()

	// Create output file for writing
	outFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Error creating output file %s: %v", outputFileName, err)
	}
	defer outFile.Close()

	// Use scanner to read line by line
	scanner := bufio.NewScanner(inFile)
	
	// Use buffered writer for better I/O performance
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()

		// Remove trailing whitespace (like rstrip in Python)
		line = strings.TrimRightFunc(line, unicode.IsSpace)

		// Keep empty lines
		if line == "" {
			writer.WriteString("\n")
			continue
		}

		// Keep comments
		if strings.HasPrefix(line, "#") {
			writer.WriteString(line + "\n")
			continue
		}

		// Split the line by any whitespace
		parts := strings.Fields(line)

		// Process lines in the format "ip domain"
		if len(parts) >= 2 {
			ip := parts[0]
			domain := parts[1]

			// Format the new line
			newLine := fmt.Sprintf("|%s^$dnsrewrite=%s\n", domain, ip)
			writer.WriteString(newLine)
		}
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("Done: created file %s\n", outputFileName)
}
