// gocmd
// For the full copyright and license information, please view the LICENSE.txt file.

// Package template provides functions for handling templates
package template

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Options represents the options that can be set when creating a new template
type Options struct {
	// Name holds the template name
	Name string
	// FilePath holds the file path
	FilePath string
	// Content holds the template content
	Content string
	// Data holds the template data
	Data interface{}
}

// New returns a new template by the given options
func New(o Options) (*Template, error) {
	// Init vars
	t := Template{
		name:     o.Name,
		filePath: o.FilePath,
		content:  o.Content,
		data:     o.Data,
	}
	if t.name == "" {
		t.name = fmt.Sprintf("%p", &t) // use pointer
	}

	// If the file path is not empty then
	if t.filePath != "" {
		// Read the file and set the template content
		b, err := os.ReadFile(t.filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create template due to %s", err.Error())
		}
		t.content = string(b)
	}

	// If the content is not empty then
	if t.content != "" {
		var err error
		t.template, err = template.New(t.name).Funcs(template.FuncMap{
			"env":  tplFuncEnv,
			"time": tplFuncTime,
			"exec": tplFuncExec,
		}).Parse(t.content)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template due to %s", err.Error())
		}
	}

	return &t, nil
}

// Template represent a template
type Template struct {
	name     string
	filePath string
	content  string
	data     interface{}
	template *template.Template
}

// Execute executes the template by the template data
func (t *Template) Execute(data interface{}) (string, error) {
	// If the template is nil then
	if t.template == nil {
		return "", nil
	}

	// If the given data is nil and the template data is not then
	if data == nil && t.data != nil {
		data = t.data // use the template data
	}

	// Execute the template and return the content
	b := bytes.Buffer{}
	if err := t.template.Execute(&b, data); err != nil {
		return "", fmt.Errorf("failed to execute template due to %s", err.Error())
	}
	return b.String(), nil
}
