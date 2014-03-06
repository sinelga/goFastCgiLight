package updatesite

import (
	"domains"
	"findfreeparagraph"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"strconv"
	"time"
	"updatehtmlpage"
	"strings"
	"createfirstgz"
)

func Update(golog syslog.Writer, tDB *db.DB, siteid map[uint64]struct{}) {

	sites := tDB.Use("Sites")
	var site domains.Site

	nowUnix := time.Now().Unix()
	var nowUnixInt int

	nowUnixInt = int(nowUnix)

	for id := range siteid {

		sites.Read(id, &site)
		golog.Info("updatesite:Update Hits " + strconv.Itoa(site.Hits))
//		golog.Info("updatesite:Update Updated " + strconv.Itoa(site.Updated))

		golog.Info("updatesite:Update Paragraphs " + strconv.Itoa(len(site.Paragraphs)))

		for _, sent := range site.Paragraphs {

			golog.Info(sent.Ptitle)
		}

		freeparagraph := findfreeparagraph.FindFromQ(golog, "fi_FI", "porno")

		err := sites.Update(id, map[string]interface{}{
			"Pathinfo":   site.Pathinfo,
			"Created":    site.Created,
			"Updated":    nowUnixInt,
			"Hits":       site.Hits + 1,
			"Paragraphs": append(site.Paragraphs, freeparagraph)})
		if err != nil {
			golog.Crit(err.Error())
		} else {
		
			webpagebytes :=updatehtmlpage.UpdatePage(golog,site.Pathinfo,site.Paragraphs)
			
			htmlfilesplit :=strings.Split(site.Pathinfo,"/")
			
			locale := htmlfilesplit[1]
			themes :=htmlfilesplit[2]
			site :=htmlfilesplit[3]
			
			var pathinfo string
			for _,path :=range htmlfilesplit[4:len(htmlfilesplit)] {
			
				pathinfo = pathinfo +"/"+path
				
			}
			
			golog.Info("pathinfo-->"+ pathinfo)
			
			createfirstgz.Creategzhtml(golog,locale,themes,site,pathinfo,webpagebytes)
			
			
		}

	}

}
