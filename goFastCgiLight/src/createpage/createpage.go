package createpage

import (
	"bytes"
	"domains"
	"findfreeparagraph"
	"html/template"
	"log/syslog"
//	"strings"
	"templ_funcmap"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string, bot string, startparameters []string, blocksite bool) []byte {

	var paragrapharr []domains.Paragraph

	var base string
	var page string
	var mediablock string

	base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
	page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/index.html"
	mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/mediablock.html"

	funcMap := template.FuncMap{
		"FirstWord": templ_funcmap.FirstWord,
		"FirstWordFromSenteces": templ_funcmap.FirstWordFromSenteces,
		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
	}


	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
		mediablock,
	)

	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, bot, startparameters)

	if blocksite {

		paragraph.Plocallink = ""
	}

	paragrapharr = append(paragrapharr, paragraph)

	webpage := bytes.NewBuffer(nil)

	htmlpage := domains.Htmlpage{
		Paragraphs: paragrapharr,
	}

	if err := index.Execute(webpage, htmlpage); err != nil {
		golog.Err(err.Error())
	}

	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	return webpagebytes

}



