package createpage

import (
	"bytes"
	"findfreeparagraph"
	"html/template"
	"log/syslog"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string, bot string, startparameters []string, blocksite bool) []byte {
	
	var firstbase string
	var firstindex string

	if blocksite {

		firstbase = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_firstbase.html"
		firstindex = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/firstindexnolinks.html"

	} else {

		firstbase = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_firstbase.html"
		firstindex = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/firstindex.html"
	}

	var index = template.Must(template.ParseFiles(
		firstbase,
		firstindex,
	))
	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, bot, startparameters)

	webpage := bytes.NewBuffer(nil)

	if err := index.Execute(webpage, paragraph); err != nil {
		golog.Err(err.Error())
	}

	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	return webpagebytes

}
