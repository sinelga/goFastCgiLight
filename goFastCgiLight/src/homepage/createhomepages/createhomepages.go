package createhomepages

import (
	"bytes"
	"domains"
//	"fmt"
	"html/template"
	"io/ioutil"
	"log/syslog"
	"templ_funcmap"
)

func CreatePages(golog syslog.Writer, sitesmap map[string]domains.Sitetohomepage) {

	var base string
	var page string
//	var mediablock string

	base = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/_base.html"
	page = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/homepageindex.html"
//	mediablock = "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/mediablock.html"

	funcMap := template.FuncMap{
//		"FirstWord":                  templ_funcmap.FirstWord,
//		"FirstWordFromSenteces":      templ_funcmap.FirstWordFromSenteces,
//		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
		"SplitPathOnWords": templ_funcmap.SplitPathOnWords,
		"SplitDomainName": templ_funcmap.SplitDomainName,
		"SomeSentences": templ_funcmap.SomeSentences,
	}

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
//		mediablock,
	)

	for _, siteinfo := range sitesmap {

		
		indexpagefullpath :="/home/juno/git/goFastCgiLight/goFastCgiLight/www/"+siteinfo.Locale+"/"+siteinfo.Themes+"/"+siteinfo.Site+"/index.html"
				
		webpage := bytes.NewBuffer(nil)

		if err := index.Execute(webpage, siteinfo); err != nil {
			golog.Err(err.Error())
		}

		webpagebytes := make([]byte, webpage.Len())
		webpagebytes = webpage.Bytes()

		if err := ioutil.WriteFile(indexpagefullpath , webpagebytes, 0666); err != nil {
			golog.Crit(err.Error())
		}

	}

}
