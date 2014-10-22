package createhomepages

import (
	"bytes"
	"domains"
	//	"fmt"
	"html/template"
	"io/ioutil"
	"log/syslog"
	"os"
	"strconv"
	"templ_funcmap"
)

func CreatePages(golog syslog.Writer, sitesmap map[string]domains.Sitetohomepage) {


	funcMap := template.FuncMap{
		"FirstWord":           templ_funcmap.FirstWord,
		"SplitPathOnWords":    templ_funcmap.SplitPathOnWords,
		"SplitDomainName":     templ_funcmap.SplitDomainName,
		"SomeSentences":       templ_funcmap.SomeSentences,
		"RandomAndLimitPages": templ_funcmap.RandomAndLimitPages,
		"LatSentences":        templ_funcmap.LatSentences,
	}


	for _, siteinfo := range sitesmap {

		variant := siteinfo.Variant
		variantstr := strconv.Itoa(variant)

		base := "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/"+variantstr+"/base.html"
		page := "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/"+variantstr+"/homepageindex.html"

		index, _ := template.New("base").Funcs(funcMap).ParseFiles(
			base,
			page,
			//		mediablock,
		)

		indexdirectory := "/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + siteinfo.Locale + "/" + siteinfo.Themes + "/" + siteinfo.Site

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
