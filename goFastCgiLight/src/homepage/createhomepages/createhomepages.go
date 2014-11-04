package createhomepages

import (
	"bytes"
	"domains"
	//		"fmt"
	"html/template"
	"io/ioutil"
	"log/syslog"
	"os"
	"strconv"
	"templ_funcmap"
)

func CreatePages(golog syslog.Writer, siteinfo domains.Sitetohomepage) {

	funcMap := template.FuncMap{
		"FirstWord":           templ_funcmap.FirstWord,
		"SplitPathOnWords":    templ_funcmap.SplitPathOnWords,
		"SplitDomainName":     templ_funcmap.SplitDomainName,
		"RandomAndLimitPages": templ_funcmap.RandomAndLimitPages,

		"Ptitle_0":           templ_funcmap.Ptitle_0,
		"Pphrase_0":          templ_funcmap.Pphrase_0,
		"Fist5_0_Sentences":  templ_funcmap.Fist5_0_Sentences,
		"Last5_0_Sentences":  templ_funcmap.Last5_0_Sentences,
		"Fist10_0_Sentences": templ_funcmap.Fist10_0_Sentences,
		"Last10_0_Sentences": templ_funcmap.Last10_0_Sentences,

		"Ptitle_1":           templ_funcmap.Ptitle_1,
		"Pphrase_1":          templ_funcmap.Pphrase_1,
		"Fist5_1_Sentences":  templ_funcmap.Fist5_1_Sentences,
		"Last5_1_Sentences":  templ_funcmap.Last5_1_Sentences,
		"Fist10_1_Sentences": templ_funcmap.Fist10_1_Sentences,
		"Last10_1_Sentences": templ_funcmap.Last10_1_Sentences,
		
		"Ptitle_2":           templ_funcmap.Ptitle_2,
		"Pphrase_2":          templ_funcmap.Pphrase_2,
		"Fist5_2_Sentences":  templ_funcmap.Fist5_2_Sentences,
		"Last5_2_Sentences":  templ_funcmap.Last5_2_Sentences,
		"Fist10_2_Sentences": templ_funcmap.Fist10_2_Sentences,
		"Last10_2_Sentences": templ_funcmap.Last10_2_Sentences,
		
		"Ptitle_3":           templ_funcmap.Ptitle_3,
		"Pphrase_3":          templ_funcmap.Pphrase_3,
		"Fist5_3_Sentences":  templ_funcmap.Fist5_3_Sentences,
		"Last5_3_Sentences":  templ_funcmap.Last5_3_Sentences,
		"Fist10_3_Sentences": templ_funcmap.Fist10_3_Sentences,
		"Last10_3_Sentences": templ_funcmap.Last10_3_Sentences,
		
		"Ptitle_4":           templ_funcmap.Ptitle_4,
		"Pphrase_4":          templ_funcmap.Pphrase_4,
		"Fist5_4_Sentences":  templ_funcmap.Fist5_4_Sentences,
		"Last5_4_Sentences":  templ_funcmap.Last5_4_Sentences,
		"Fist10_4_Sentences": templ_funcmap.Fist10_4_Sentences,
		"Last10_4_Sentences": templ_funcmap.Last10_4_Sentences,
		
		"Ptitle_5":           templ_funcmap.Ptitle_5,
		"Pphrase_5":          templ_funcmap.Pphrase_5,
		"Fist5_5_Sentences":  templ_funcmap.Fist5_5_Sentences,
		"Last5_5_Sentences":  templ_funcmap.Last5_5_Sentences,
		"Fist10_5_Sentences": templ_funcmap.Fist10_5_Sentences,
		"Last10_5_Sentences": templ_funcmap.Last10_5_Sentences,		
		
		"Ptitle_6":           templ_funcmap.Ptitle_6,
		"Pphrase_6":          templ_funcmap.Pphrase_6,
		"Fist5_6_Sentences":  templ_funcmap.Fist5_6_Sentences,
		"Last5_6_Sentences":  templ_funcmap.Last5_6_Sentences,
		"Fist10_6_Sentences": templ_funcmap.Fist10_6_Sentences,
		"Last10_6_Sentences": templ_funcmap.Last10_6_Sentences,			
		
		"Ptitle_7":           templ_funcmap.Ptitle_7,
		"Pphrase_7":          templ_funcmap.Pphrase_7,
		"Fist5_7_Sentences":  templ_funcmap.Fist5_7_Sentences,
		"Last5_7_Sentences":  templ_funcmap.Last5_7_Sentences,
		"Fist10_7_Sentences": templ_funcmap.Fist10_7_Sentences,
		"Last10_7_Sentences": templ_funcmap.Last10_7_Sentences,
		
		"Ptitle_8":           templ_funcmap.Ptitle_8,
		"Pphrase_8":          templ_funcmap.Pphrase_8,
		"Fist5_8_Sentences":  templ_funcmap.Fist5_8_Sentences,
		"Last5_8_Sentences":  templ_funcmap.Last5_8_Sentences,
		"Fist10_8_Sentences": templ_funcmap.Fist10_8_Sentences,
		"Last10_8_Sentences": templ_funcmap.Last10_8_Sentences,
		
		"Ptitle_9":           templ_funcmap.Ptitle_9,
		"Pphrase_9":          templ_funcmap.Pphrase_9,
		"Fist5_9_Sentences":  templ_funcmap.Fist5_9_Sentences,
		"Last5_9_Sentences":  templ_funcmap.Last5_9_Sentences,
		"Fist10_9_Sentences": templ_funcmap.Fist10_9_Sentences,
		"Last10_9_Sentences": templ_funcmap.Last10_9_Sentences,				
		
					
								
	}

	locale := siteinfo.Locale
	themes := siteinfo.Themes

	variant := siteinfo.Variant
	variantstr := strconv.Itoa(variant)

	base := "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/" + locale + "/" + themes + "/" + variantstr + "/base.html"

	page := "/home/juno/git/goFastCgiLight/goFastCgiLight/templ/" + locale + "/" + themes + "/" + variantstr + "/homepageindex.html"

	index := template.Must(template.New("base").Funcs(funcMap).ParseFiles(base, page))

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
