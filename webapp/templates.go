package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

const header = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Anton's app</title>
</head>
<body>`

const footer = `</body>
</html>`

var unparsedTemplates = map[string]string{
	"index": `
<p>
Welcome to Anton's photo app! Show a random photo here, or if no photos are available, a link to upload a photo, and a link to <a href="/photos">all photos</a>.
</p>`,
	"listOfPhotos": `
{{ range . }}
<a href="/photos/{{ .Id }}"><img src="/thumbnails/{{ .Id }}"/></a>
{{ end }}
`,
	"singlePhoto": `
<!-- <h1>{{ .Title }}</h1> -->
<img src="/thumbnails/{{ .Id }}"/>
<!-- <p>{{ .Content }}</p> -->
<p>Taken: {{ .DateTaken }}</p>
<p>Uploaded: {{ .DateUploaded }}</p>

<a href="/photos">All photos</a>
`,
	"editPhoto": `
<h1>{{ .Title }}</h1>
<img src="/thumbnails/{{ .Id }}"/>
<form>
<p>Title <input id="title" type="text" value="{{ .Title }}"></p>
<p>Content <input id="content" type="text" value="{{ .Content }}"></p>
</form>
<p>Taken: {{ .DateTaken }}</p>
<p>Uploaded: {{ .DateUploaded }}</p>

<a href="/photos">All photos</a>
`,
}

var templates map[string]*template.Template

func renderTemplate(templateName string, data interface{}) (*string, error) {
	var doc bytes.Buffer
	err := templates[templateName].Execute(&doc, data)
	if err != nil {
		return nil, err
	} else {
		str := doc.String()
		return &str, nil
	}
}

func initializeTemplates() {
	templates = make(map[string]*template.Template)
	for name, tmpl := range unparsedTemplates {
		var err error
		templates[name], err = template.New(name).Parse(header + tmpl + footer)
		if err != nil {
			log.Fatal(fmt.Sprintf("An error has occurred: %v", err))
		}
	}
}
