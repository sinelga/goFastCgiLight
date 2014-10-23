package updatehtmlpage

import (
	"bytes"
	"domains"
	"html/template"
	"log/syslog"
	//	"strings"
	"templ_funcmap"
	"time"
)

func UpdatePage(golog syslog.Writer, htmlfile string, paragraphsarr []domains.Paragraph, blocksite bool, locale string, themes string, variant string, created string, updated string) []byte {

	var base string
	var page string
	var mediablock string

	base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/" + locale + "/" + themes + "/" + variant + "/base.html"
	page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/" + locale + "/" + themes + "/" + variant + "/page.html"
	mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/" + locale + "/" + themes + "/" + variant + "/mediablock.html"

	funcMap := template.FuncMap{

		"FirstWord":                  templ_funcmap.FirstWord,
		"FirstWordFromSenteces":      templ_funcmap.FirstWordFromSenteces,
		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
	}

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
		mediablock,
	)
	currenttime := time.Now().Local()

	webpage := bytes.NewBuffer(nil)

	htmlpage := domains.Htmlpage{
		Locale:     locale,
		Themes:     themes,
		Variant:    variant,
		Paragraphs: paragraphsarr,
		Created:    created,
		Updated:    currenttime.Format("2006-01-02 15:04:05"),
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
