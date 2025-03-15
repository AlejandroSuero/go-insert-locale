package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type LocaleData map[string]map[string]string

func main() {
	inputFile := flag.String("input", "", "Input JSON file containing locale data")
	outputDir := flag.String("output", "", "Output directory for locale files")
	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("No input file provided\n\tUse -input to provide one\n")
		return
	}
	if *outputDir == "" {
		fmt.Printf("No output directory provided\n\tUse -output to provide one\n")
		return
	}

	// Load input translations
	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var inputData LocaleData
	if err := json.Unmarshal(data, &inputData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Process each locale from input data
	for msgKey, translations := range inputData {
		for locale, text := range translations {
			filePath := filepath.Join(*outputDir, locale+".json")

			// 1. Read existing content
			var existing map[string]interface{}
			if fileContent, err := os.ReadFile(filePath); err == nil {
				if err := json.Unmarshal(fileContent, &existing); err != nil {
					fmt.Printf("Error parsing existing JSON for %s: %v\n", locale, err)
					continue
				}
			} else {
				existing = make(map[string]interface{})
			}

			// 2. Preserve existing key order and add new key at the end
			if _, exists := existing[msgKey]; !exists {
				existing[msgKey] = text
			}

			// 3. Write back with original formatting
			if jsonData, err := json.MarshalIndent(existing, "", "  "); err == nil {
				if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
					fmt.Printf("Error writing file %s: %v\n", filePath, err)
				}
			} else {
				fmt.Printf("Error marshaling JSON for %s: %v\n", locale, err)
			}
		}
	}

	fmt.Println("Locale files updated successfully!")
}
