package library

import (
	"bytes"
	"fmt"
	"html/template"
)

func LoadHtmlTemplateToString(path string, data interface{}) (string, error) {
	temp := template.Must(template.ParseFiles(path))
	var tempBuffer bytes.Buffer
	if err := temp.Execute(&tempBuffer, data); err != nil {
		return "", fmt.Errorf("error parsing tempalte: %w", err)
	}
	return tempBuffer.String(), nil
}
