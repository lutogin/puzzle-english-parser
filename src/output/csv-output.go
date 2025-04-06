package outputCsv

import (
	"fmt"
	"go-pe-parser/src/config"
	peParser "go-pe-parser/src/parser"
	"os"
	"path/filepath"
)

func SaveToCSV(words []peParser.WordPair, cfg *config.Config) error {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	// Get separator from config or use tab as fallback
	separator := cfg.APP.Config["csvSeparator"]
	if separator == "" {
		separator = ";"
	}

	// Check if file exists
	var file *os.File
	filename := cfg.APP.Config["fileName"]
	fullPath := filepath.Join(workDir, filename)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// Create new file with headers if it doesn't exist
		file, err = os.Create(fullPath)
		if err != nil {
			return err
		}
		// Write CSV header
		file.WriteString(fmt.Sprintf("Word%sTranslation%sPhrase%sPhraseTranslation\n", separator, separator, separator))
	} else {
		// Open file in append mode if it exists
		var err error
		file, err = os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	// Write words
	for _, pair := range words {
		file.WriteString(fmt.Sprintf("%s%s%s%s%s%s%s\n",
			pair.Word, separator,
			pair.Translation, separator,
			pair.Phrase, separator,
			pair.PhraseTranslation))
	}

	return nil
}
