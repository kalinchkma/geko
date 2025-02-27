package helper

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

func LoadHtmlTemplateToString(FS embed.FS, path string, data interface{}) (string, error) {

	temp, err := template.ParseFS(FS, path)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	var tempBuffer bytes.Buffer
	if err := temp.Execute(&tempBuffer, data); err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}
	return tempBuffer.String(), nil
}
