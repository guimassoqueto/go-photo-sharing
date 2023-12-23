package views

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func Parse(filePath string) (Template, error) {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template failed: %w", err)
	}
    return Template{
        htmlTemplate: tmpl,
    }, nil
}

type Template struct {
	htmlTemplate *template.Template
}

func (t Template) ExecuteTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("executing template failed: %v", err)
		http.Error(w, "There was an error executing the temmplate", http.StatusInternalServerError)
	}
}

func (t Template) ExecuteTemplateMinified(w http.ResponseWriter, data interface{}) {
	m := minify.New()
    m.AddFunc("text/html", html.Minify)

    var buf bytes.Buffer
    err := t.htmlTemplate.Execute(&buf, data)
    if err != nil {
        log.Printf("executing template failed: %v", err)
        http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    minifiedHTML, err := m.String("text/html", buf.String())
    if err != nil {
        log.Printf("minifying template failed: %v", err)
        http.Error(w, "There was an error minifying the template", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(minifiedHTML))
}