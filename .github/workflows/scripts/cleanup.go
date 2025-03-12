package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	startMarker = "> Start clean up"
	endMarker   = "> End clean up"
)

// processFile removes code between markers in a given file
func processFile(filePath string) error {
	fmt.Println(filePath)
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var outputLines []string
	inCleanupBlock := false
	modified := false

	//
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, startMarker) {
			inCleanupBlock = true
			modified = true
			continue
		}

		if strings.Contains(line, endMarker) {
			inCleanupBlock = false
			continue
		}

		if !inCleanupBlock {
			outputLines = append(outputLines, line)
		}
	}
	//

	if err := scanner.Err(); err != nil {
		return err
	}

	// Only rewrite the file if it was modified
	if modified {
		err = os.WriteFile(filePath, []byte(strings.Join(outputLines, "\n")+"\n"), 0644)
		if err != nil {
			return err
		}
		fmt.Println("✅ Cleaned:", filePath)
	}
	return nil
}

// scanAndClean searches for files and processes them
func ScanAndClean(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and .git folder
		if info.IsDir() || strings.Contains(path, ".git") {
			return nil
		}

		// Process only .go, .md, or config files
		if strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".md") {
			return processFile(path)
		}

		return nil
	})
}

func main() {
	rootDir := "../.." // Scan current directory
	if err := ScanAndClean(rootDir); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	fmt.Println("🚀 Cleanup complete!")
}
