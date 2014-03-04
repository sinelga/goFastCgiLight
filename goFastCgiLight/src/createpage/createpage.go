package createpage

import (
//"html/template"
"findfreeparagraph"
"log/syslog"

)

func CreateHtmlPage(golog syslog.Writer,locale string,themes string) {

//	var index = template.Must(template.ParseFiles(
//		"templ/_base.html",
//		"templ/index.html",
//	))
	paragraph := findfreeparagraph.FindFromQ(golog,locale,themes)
	
	golog.Info(paragraph.Ptitle)


}