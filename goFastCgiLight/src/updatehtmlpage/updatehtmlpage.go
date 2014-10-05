package updatehtmlpage

import (
	"bytes"
	"domains"
	"html/template"
	"log/syslog"
//	"strings"
	"templ_funcmap"
)

func UpdatePage(golog syslog.Writer, htmlfile string, paragraphsarr []domains.Paragraph, blocksite bool) []byte {

	var base string
	var page string
	var mediablock string

	base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
	page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/index.html"
	mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/mediablock.html"

	funcMap := template.FuncMap{

		"FirstWord":             templ_funcmap.FirstWord,
		"FirstWordFromSenteces": templ_funcmap.FirstWordFromSenteces,
		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
	}

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
		mediablock,
	)

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

//func firstWord(s string) string {
//
//	words := strings.Fields(s)
//
//	return words[0]
//
//}
