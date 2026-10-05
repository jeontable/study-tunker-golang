package main

import (
	"html/template"
)

func main() {
	template.New("foo").Parse(`{{define "T"}}Hello`)
}
