package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrUnsupportedFormat = errors.New("unsupported file format")
	ErrFileNotFound      = errors.New("file not found")
	ErrInvalidContent    = errors.New("invalid file content")
)

// FileType represents supported file formats
type FileType int

const (
	Unknown FileType = iota
	JSON
	YAML
)

type Parser struct {
	filePath string
	fileType FileType
}

// Creates a new parser instance
func NewParser(filepath string) (*Parser, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w : %s", ErrFileNotFound, filepath)
	}

	fileType, err := detectFileType(filepath)
	if err != nil {
		return nil, err
	}

	return &Parser{
		filePath: filepath,
		fileType: fileType,
	}, nil
}

// Reads the file and returns its content as a map
func (p *Parser) Parse() (map[string]interface{}, error) {
	content, err := os.ReadFile(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: %w", err)
	}

	var result map[string]interface{}

	//Parse based on file type
	switch p.fileType {
	case JSON:
		if err := json.Unmarshal(content, &result); err != nil {
			return nil, fmt.Errorf("%w: invalid JSON: %s", ErrInvalidContent, err)
		}
	case YAML:
		if err := yaml.Unmarshal(content, &result); err != nil {
			return nil, fmt.Errorf("%w: invalid YAML: %s", ErrInvalidContent, err)
		}
	}

	return result, nil
}

// Detects the file format from its extension
func detectFileType(filepath string) (FileType, error) {
	ext := strings.ToLower(path.Ext(filepath))
	switch ext {
	case ".json":
		return JSON, nil
	case ".yaml", ".yml":
		return YAML, nil
	default:
		return Unknown, fmt.Errorf("%w: %s", ErrUnsupportedFormat, ext)
	}
}
