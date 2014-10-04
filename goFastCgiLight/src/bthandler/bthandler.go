package bthandler

import (
	"clean_pathinfo"
	"createfirstgz"
	"createpage"
	"log/syslog"
	"net/http"
	"strings"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, bot string, startparameters []string, blocksite bool) {

	pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)

	if blocksite {

		golog.Info("BTrequestHandler:will block-> " + site+pathinfo)

	}

	var bytepage []byte
	if strings.HasSuffix(pathinfoclean, ".html")  {

		bytepage = createpage.CreateHtmlPage(golog, locale, themes, bot, startparameters,blocksite)

		resp.Write(bytepage)

	} else {

		resp.WriteHeader(404)
	}

	if strings.HasSuffix(pathinfoclean, ".html") {

		go createfirstgz.Creategzhtml(golog, locale, themes, site, pathinfoclean, bytepage)

	}

}
