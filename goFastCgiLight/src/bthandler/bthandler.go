package bthandler

import (
	"clean_pathinfo"
	"createpage"
	"log/syslog"
	"net/http"
	"strings"
	"createfirstgz"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string) {

	pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)

	//htmlfile := string("www/" + locale + "/" + themes + "/" + site + pathinfoclean)

	var bytepage []byte
	if strings.HasSuffix(pathinfoclean, ".html") || strings.HasSuffix(pathinfoclean, ".php") || strings.HasSuffix(pathinfoclean, ".jsp") {

		bytepage = createpage.CreateHtmlPage(golog, locale, themes)

		resp.Write(bytepage)

	} else {

		resp.WriteHeader(404)
	}
	
	if strings.HasSuffix(pathinfoclean, ".html") {

		go createfirstgz.Creategzhtml(golog,locale,themes,site,pathinfoclean,bytepage)

	}	

}
