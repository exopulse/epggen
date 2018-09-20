package generator

import (
	"bytes"
	"fmt"
	"html/template"
)

// TemplateLoader is implemented by template loaders.
type TemplateLoader interface {
	LoadTemplate(name string) (string, error)
}

// Generator is main generator interface.
type Generator struct {
	tl TemplateLoader
}

// NewGenerator creates new generator.
func NewGenerator(tl TemplateLoader) *Generator {
	return &Generator{tl: tl}
}

// Generate generates file contents from template and params.
func (g *Generator) Generate(templateName string, params map[string]string) (string, error) {
	tmplContent, err := g.tl.LoadTemplate(fmt.Sprintf("tmpl/%s.tmpl", templateName))

	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Parse(tmplContent))
	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, params); err != nil {
		return "", err
	}

	return buf.String(), nil
}
