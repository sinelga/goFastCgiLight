package bthandler

import (
	"log/syslog"
	"net/http"
	"clean_pathinfo"
	"strings"
	"createpage"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string) {

pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)

//htmlfile := string("www/" + locale + "/" + themes + "/" + site + pathinfoclean)

if strings.HasSuffix(pathinfoclean, ".html") || strings.HasSuffix(pathinfoclean, ".php") || strings.HasSuffix(pathinfoclean, ".jsp") {

	createpage.CreateHtmlPage(golog,locale,themes)

}


}
