package updatehtmlpage

import (
	"bytes"
	"domains"
	"html/template"
	"log/syslog"
)

func UpdatePage(golog syslog.Writer, htmlfile string, paragraphsarr []domains.Paragraph) []byte {


	var index = template.Must(template.ParseFiles(
		"templ/_base.html",
		"templ/index.html",
	))
	webpage := bytes.NewBuffer(nil)

	htmlpage := domains.Htmlpage{
		Paragraphs: paragraphsarr,
	}

	if err := index.Execute(webpage, htmlpage); err != nil {
		golog.Err(err.Error())
	}

	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	return webpagebytes

}
