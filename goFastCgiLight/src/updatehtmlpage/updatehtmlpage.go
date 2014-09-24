package updatehtmlpage

import (
	"bytes"
	"domains"
	"html/template"
	"log/syslog"
)

func UpdatePage(golog syslog.Writer, htmlfile string, paragraphsarr []domains.Paragraph, blocksite bool) []byte {

	var base string
	var page string

	if blocksite {

		base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
		page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/indexnolinks.html"

	} else {

		base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
		page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/index.html"

	}

	var index = template.Must(template.ParseFiles(
		base,
		page,
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
