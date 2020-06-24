package md2txt

import (
	"strings"

	"github.com/gomarkdown/markdown"
)

func ToTXT(md string) string {
	html := markdown.ToHTML([]byte(md), nil, nil)
	content := strings.Replace(string(html), "<li><p>", "<li><p>• ", -1)

	return trimHtml(content)
}
