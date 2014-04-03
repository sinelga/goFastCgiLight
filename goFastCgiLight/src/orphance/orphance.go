package orphance

import (
	"checksiteexist"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
)

func LookUp(golog syslog.Writer, col *db.Col, path string) {

	idsitesarr := checksiteexist.CheckDB(golog, col, path)

	if len(idsitesarr) == 1 {

		golog.Info(path + " exist in db")

	} else if len(idsitesarr) == 0 {

//		golog.Info(path + " NOT in db")

	}

}
