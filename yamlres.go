package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed tpl/*
var templates embed.FS

func main() {
	in := flag.String("i", "cv.yaml", "Input yaml file")
	out := flag.String("o", "cv.html", "Output html file")
	flag.Parse()

	bytes, err := os.ReadFile(*in)
	if err != nil {
		fmt.Printf("Failed to read %s: %s\n", *in, err)
		os.Exit(1)
	}
	m := map[interface{}]interface{}{}
	if err := yaml.Unmarshal(bytes, &m); err != nil {
		fmt.Printf("Failed to parse yaml: %s\n", err)
		os.Exit(1)
	}
	tpl, _ := template.ParseFS(templates, "*/*")
	f, err := os.Create(*out)
	if err != nil {
		fmt.Printf("Failed to create %s: %s\n", *out, err)
		os.Exit(1)
	}
	defer f.Close()
	if err := tpl.ExecuteTemplate(f, "cv", m); err != nil {
		fmt.Printf("Failed to execute template: %s\n", err)
	}
	fmt.Printf("Created %s\n", *out)
}
