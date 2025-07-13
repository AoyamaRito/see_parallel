package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExpandFiles(patterns []string) ([]string, error) {
	var expandedFiles []string
	seen := make(map[string]bool)

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("error expanding pattern %s: %v", pattern, err)
		}

		if len(matches) == 0 {
			if _, err := os.Stat(pattern); err == nil {
				matches = []string{pattern}
			}
		}

		for _, match := range matches {
			if !seen[match] {
				seen[match] = true
				expandedFiles = append(expandedFiles, match)
			}
		}
	}

	return expandedFiles, nil
}

func ReadAndCombineFiles(files []string) (string, error) {
	var combined strings.Builder

	for i, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("error reading file %s: %v", file, err)
		}

		if i > 0 {
			combined.WriteString("\n\n")
		}

		combined.WriteString(fmt.Sprintf("=======↓ここから %s ==========\n", file))
		combined.Write(content)
	}

	return combined.String(), nil
}