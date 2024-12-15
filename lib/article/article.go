package article

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/tools/present"
)


var articleTmpl *template.Template

func init() {
	var err error
	articleTmpl, err = present.Template().ParseGlob("templates/article/*")
	if err != nil {
		logs.Error("Failed to parse article render template", err)
	}
	return
}

func Render(doc *present.Doc) (string, error) {
	var buf bytes.Buffer
	err := doc.Render(&buf, articleTmpl)

	return buf.String(), err
}

func Parse(raw string, name string) (*present.Doc, error) {
	reader := strings.NewReader(raw)

	return present.Parse(reader, name, 0)
}

