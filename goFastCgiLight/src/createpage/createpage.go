package createpage

import (
	"bytes"
	"findfreeparagraph"
	"html/template"
	"log/syslog"
	"domains"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string, bot string, startparameters []string, blocksite bool) []byte {
	
	
	var paragrapharr []domains.Paragraph
	
	var firstbase string
	var firstindex string
	var mediablock string
	
	

//	if blocksite {
//
//		firstbase = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_firstbase.html"
//		firstindex = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/firstindexnolinks.html"
//
//	} else {


		firstbase = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
		firstindex = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/index.html"
		mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/mediablock.html"

		
//	}

	var index = template.Must(template.ParseFiles(
		firstbase,
		firstindex,
		mediablock,
	))
	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, bot, startparameters)
	
	if blocksite {
		
		paragraph.Plocallink=""
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
