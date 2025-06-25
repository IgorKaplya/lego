package blogrender

import (
	"embed"
	"html/template"
	"io"

	"github.com/russross/blackfriday/v2"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

//go:embed "templates/*"
var postTemplatesFS embed.FS
var templates, templatesParsingError = template.ParseFS(postTemplatesFS, "templates/*.gohtml")

func RenderPost(writer io.Writer, post Post) error {
	if templatesParsingError != nil {
		return templatesParsingError
	}

	return templates.ExecuteTemplate(writer, "blog.gohtml", post)
}

func (p Post) BodyAsHtml() template.HTML {
	var input = []byte(p.Body)
	var output = blackfriday.Run(input)
	return template.HTML(output)
}

func RenderIndex(writer io.Writer, posts []Post) error {
	if templatesParsingError != nil {
		return templatesParsingError
	}

	return templates.ExecuteTemplate(writer, "index.gohtml", posts)
}
