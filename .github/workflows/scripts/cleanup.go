package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	startMarker = "> Start clean up"
	endMarker   = "> End clean up"
)

// removeExpiredCode removes code between markers in a given file
func removeExpiredCode(filePath string) error {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var outputLines []string
	inCleanupBlock := false
	modified := false
	hasBothMarkers := false

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, startMarker) {
			inCleanupBlock = true
			modified = true
			continue
		}

		if strings.Contains(line, endMarker) {
			hasBothMarkers = true
			inCleanupBlock = false
			continue
		}

		if !inCleanupBlock {
			outputLines = append(outputLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Only rewrite the file if it was modified
	if modified && hasBothMarkers {
		err = os.WriteFile(filePath, []byte(strings.Join(outputLines, "\n")+"\n"), 0644)
		if err != nil {
			return err
		}
		fmt.Println("✅ Cleaned:", filePath)
	}
	return nil
}

func ScanFilesWithExt(root string, extensions []string) ([]string, error) {
	var golangFiles []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories and .git folder
		if info.IsDir() || strings.Contains(path, ".git") || strings.Contains(path, ".github") {
			return nil
		}

		// Process only .go, .md, or config files
		if FileWithinExtensions(path, extensions) {
			fmt.Println(golangFiles)
			golangFiles = append(golangFiles, path)
			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return golangFiles, nil
}

func FileWithinExtensions(filePath string, extensions []string) bool {
	for _, v := range extensions {
		if strings.HasSuffix(filePath, v) {
			return true
		}
	}
	return false
}

func main() {
	rootDir := "."                // Scan current directory
	extensions := []string{".go"} // Scan go files only

	// Scan the rootDir for the extensions mentioned and return files matching
	files, err := ScanFilesWithExt(rootDir, extensions)

	if err != nil {
		log.Fatal("❌ Error:", err)
	}

	// Remove expired code for every file matching the extension specified
	for _, file := range files {
		err := removeExpiredCode(file)
		if err != nil {
			log.Fatal("❌ Error:", err)
		}
	}

	fmt.Println("🚀 Cleanup complete!")
}
