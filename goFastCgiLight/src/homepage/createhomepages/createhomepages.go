package createhomepages

import (
	"bytes"
	"domains"
	//	"fmt"
	"html/template"
	"io/ioutil"
	"log/syslog"
	"os"
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
		"FirstWord": templ_funcmap.FirstWord,
		//		"FirstWordFromSenteces":      templ_funcmap.FirstWordFromSenteces,
		//		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
		"SplitPathOnWords": templ_funcmap.SplitPathOnWords,
		"SplitDomainName":  templ_funcmap.SplitDomainName,
		"SomeSentences":    templ_funcmap.SomeSentences,
		"RandomAndLimitPages": templ_funcmap.RandomAndLimitPages,
	}

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
		//		mediablock,
	)

	for _, siteinfo := range sitesmap {

		indexdirectory := "/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + siteinfo.Locale + "/" + siteinfo.Themes + "/" + siteinfo.Site

		//		src, err := os.Stat(indexdirectory)
		if _, err := os.Stat(indexdirectory); os.IsNotExist(err) {

			if os.MkdirAll(indexdirectory, 0777) != nil {
				golog.Err(err.Error())

			}

		}


		indexpagefullpath := "/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + siteinfo.Locale + "/" + siteinfo.Themes + "/" + siteinfo.Site + "/index.html"

		webpage := bytes.NewBuffer(nil)

		if err := index.Execute(webpage, siteinfo); err != nil {
			golog.Err(err.Error())
		}

		webpagebytes := make([]byte, webpage.Len())
		webpagebytes = webpage.Bytes()

		if err := ioutil.WriteFile(indexpagefullpath, webpagebytes, 0666); err != nil {
			golog.Crit(err.Error())
		}

	}

}
