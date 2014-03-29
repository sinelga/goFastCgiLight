package updatesite

import (
	"createfirstgz"
	"domains"
	"findfreeparagraph"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"strconv"
	"strings"
	"time"
	"updatehtmlpage"
)

func Update(golog syslog.Writer, col *db.Col, siteid map[uint64]struct{}) {

	var site domains.Site

	nowUnix := time.Now().Unix()
	var nowUnixInt int
	nowUnixInt = int(nowUnix)
	var htmlfilesplit []string
	var locale string
	var themes string
	var hostsite string
	for id := range siteid {

		col.Read(id, &site)
		golog.Info("updatesite:Update Hits " + strconv.Itoa(site.Hits))

		golog.Info("updatesite:Update Paragraphs " + strconv.Itoa(len(site.Paragraphs)))

		for _, sent := range site.Paragraphs {

			golog.Info(sent.Ptitle)
		}

		htmlfilesplit = strings.Split(site.Pathinfo, "/")

		locale = htmlfilesplit[1]
		themes = htmlfilesplit[2]
		hostsite = htmlfilesplit[3]

		freeparagraph := findfreeparagraph.FindFromQ(golog, locale, themes)
		
		updatedparagraphs := append(site.Paragraphs, freeparagraph)

		err := col.Update(id, map[string]interface{}{
			"Pathinfo":   site.Pathinfo,
			"Created":    site.Created,
			"Updated":    nowUnixInt,
			"Hits":       site.Hits + 1,
//			"Paragraphs": append(site.Paragraphs, freeparagraph)})
			"Paragraphs": updatedparagraphs,
			})
		if err != nil {
			golog.Crit(err.Error())
		} else {

//			webpagebytes := updatehtmlpage.UpdatePage(golog, site.Pathinfo, site.Paragraphs)
			webpagebytes := updatehtmlpage.UpdatePage(golog, site.Pathinfo, updatedparagraphs)

			var pathinfo string
			for _, path := range htmlfilesplit[4:len(htmlfilesplit)] {

				pathinfo = pathinfo + "/" + path

			}

			golog.Info("pathinfo-->" + pathinfo)

			createfirstgz.Creategzhtml(golog, locale, themes, hostsite, pathinfo, webpagebytes)

		}

	}

}
