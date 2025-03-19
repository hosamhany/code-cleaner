package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const (
	startMarker = "> Start clean up at "
	endMarker   = "> End clean up"
)

var dateTokenizers = []string{"at", "on"}

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
			isExpired, err := IsExpiredCode(line)
			if err != nil {
				return err
			}
			if isExpired {
				inCleanupBlock = true
				modified = true
			}
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

func IsExpiredCode(line string) (bool, error) {
	tokenizer := whichTokenizer(line)
	if tokenizer == "" {
		return false, errors.New("please use \"at\" or \"on\" to specify the time")
	}

	tokenizedLine := strings.Split(line, tokenizer)
	date := strings.TrimSpace(tokenizedLine[len(tokenizedLine)-1])
	//parsed time in yyyy-mm-dd format
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, err
	}
	//is now before parsedTime
	if time.Now().Compare(t) == -1 {
		return false, nil
	}
	return true, nil
}

func whichTokenizer(line string) string {
	tokenizedLine := strings.Split(line, " ")
	for _, v := range dateTokenizers {
		if slices.Contains(tokenizedLine, v) {
			return v
		}
	}
	return ""
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
