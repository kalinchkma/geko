package lib

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

// This function load the html template with the help of embed package
// Return html template string and nil error if success
// return "", and error if error occures
func LoadHtmlTemplateToString(FS embed.FS, path string, data any) (string, error) {
	// Parse html template based of provided embed.FS
	temp, err := template.ParseFS(FS, path)
	// Handle template parsign error
	if err != nil {
		return "", err
	}
	// Make loaded templated into byte buffer
	var tempBuffer bytes.Buffer

	// bind data and template byte buffer
	// Handle error if error occures
	if err := temp.Execute(&tempBuffer, data); err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	// Return template string data and nil error
	return tempBuffer.String(), nil
}
