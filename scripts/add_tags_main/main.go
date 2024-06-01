package main

import (
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	targetDir := "scripts/tmp"
	err := filepath.Walk(targetDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			err := processFile(path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
		return
	}
}

func processFile(path string) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}

	// Regular expression to match struct tags with db and json tags
	re := regexp.MustCompile(`db:"([^"]+)"\sjson:"([^"]+)"`)
	updatedContent := re.ReplaceAllFunc(input, func(match []byte) []byte {
		// Extract db tag value
		matches := re.FindSubmatch(match)
		if len(matches) == 3 {
			dbTag := matches[1]
			jsonTag := matches[2]
			newTag := fmt.Sprintf(`db:"%s" json:"%s" mapstructure:"%s"`, dbTag, jsonTag, dbTag)
			return []byte(newTag)
		}
		return match
	})

	// Format the updated content using gofmt
	formattedContent, err := format.Source(updatedContent)
	if err != nil {
		return fmt.Errorf("error formatting file %s: %v", path, err)
	}

	// Write the updated and formatted content back to the file
	err = os.WriteFile(fmt.Sprintf("%s.tmp", path), formattedContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %v", path, err)
	}

	return nil
}
