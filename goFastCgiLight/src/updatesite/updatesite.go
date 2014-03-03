package updatesite

import (
	"domains"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"strconv"
	"time"
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

		err := sites.Update(id, map[string]interface{}{
			"Pathinfo": site.Pathinfo,
			"Created":  site.Created,
			"Updated":  nowUnixInt,
			"Hits":     site.Hits + 1})
		if err != nil {

			golog.Crit(err.Error())
		}

	}

	//	err = sites.Update(docID, map[string]interface{}{
	//		"name": "Go is very popular",
	//		"url":  "google.com"})
	//	if err != nil {
	//		panic(err)
	//	}

}
