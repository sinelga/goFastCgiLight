package orphance

import (
	"checksiteexist"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"os"
)

func LookUp(golog syslog.Writer, col *db.Col, path string) {

	idsitesarr := checksiteexist.CheckDB(golog, col, path)

	if len(idsitesarr) == 1 {

//		golog.Info("orphance:LookUp "+path  + " exist in db")

		for id := range idsitesarr {

			golog.Info("Delete record for "+path)
			col.Delete(id)

		}

	} else if len(idsitesarr) > 1 {
		golog.Err("orphance:LookUp !!!Someting wrong double record? "+path)

	}
	
	if err := os.Remove(path); err != nil {
		
		golog.Err(err.Error())
	
	}
	

}
