package updatesite

import (
	"domains"
	"findfreeparagraph"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"strconv"
	"time"
	"updatehtmlpage"
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
		golog.Info("updatesite:Update Updated " + strconv.Itoa(site.Updated))

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
		
			updatehtmlpage.UpdatePage(golog,site.Pathinfo,site.Paragraphs)
		}

	}

}
