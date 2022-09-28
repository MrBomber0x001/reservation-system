package config

import (
	"html/template"
	"log"
)

//*AppConfig holds the application configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
