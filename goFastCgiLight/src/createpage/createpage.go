package createpage

import (
	"bytes"
	"domains"
	"findfreeparagraph"
	"html/template"
	"log/syslog"
	"strings"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string, bot string, startparameters []string, blocksite bool) []byte {

	var paragrapharr []domains.Paragraph

	var firstbase string
	var firstindex string
	var mediablock string

	firstbase = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
	firstindex = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/index.html"
	mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/mediablock.html"

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"FirstWord": firstWord,
	}

//	var index = template.Must(template.ParseFiles(
//		firstbase,
//		firstindex,
//		mediablock,
//	))

//	index.Funcs(funcMap)

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		firstbase,
		firstindex,
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


func firstWord(s string) string {
	
		words := strings.Fields(s)
		
		return words[0]
	
	
}
