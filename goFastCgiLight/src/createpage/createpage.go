package createpage

import (
	"bytes"
	"findfreeparagraph"
	"html/template"
	"log/syslog"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string) []byte {

	var index = template.Must(template.ParseFiles(
		"/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_firstbase.html",
		"/home/juno/git/goFastCgiLight/goFastCgiLight/templ/firstindex.html",
	))
	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes)
	
//	golog.Info("paragraph.Pushsite "+paragraph.Pushsite)

	webpage := bytes.NewBuffer(nil)

	if err := index.Execute(webpage, paragraph); err != nil {
		golog.Err(err.Error())
	}

	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	return webpagebytes

}
