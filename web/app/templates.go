package app

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

func InitTemplates() (*template.Template, error) {
	templatesDir := os.Getenv("TEMPLATES_DIR")
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}
	var allFiles []string
	for _, file := range files {
		allFiles = append(allFiles, path.Join(templatesDir, file.Name()))
	}
	templates, err := template.ParseFiles()
	if err != nil {
		return nil, err
	}
	return templates, nil
}
