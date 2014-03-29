package htmlfileexist

import (
	"checksiteexist"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"os"
	"newsite"
	"updatesite"
)

func StartCheck(golog syslog.Writer, col *db.Col, locale string, themes string, site string, pathinfo string) {

	htmlfile := string("www/" + locale + "/" + themes + "/" + site + pathinfo)

	golog.Info("htmlfileexist:StartCheck: htmlfile " + htmlfile)

	idsitesarr := checksiteexist.CheckDB(golog, col, htmlfile)

	if len(idsitesarr) == 1 {

		golog.Info("sitesmaker:Update " + htmlfile)
		updatesite.Update(golog, col, idsitesarr)

	} else if len(idsitesarr) == 0 {

		newsite.CreateSite(golog, col, htmlfile)

	} else {

		golog.Err("sitesmaker: DB someting wrong " + htmlfile)

	}

}

func DeleteHtmlFile(golog syslog.Writer, htmlfile string) {

	if finfo, err := os.Stat(htmlfile); err != nil {

		if os.IsNotExist(err) {

			golog.Alert("!!!htmlfileexist:DeleteHtmlFile: file does not exist??? Cant be?? -->" + htmlfile)

		}

	} else {

		if !finfo.IsDir() {

			if err := os.Remove(htmlfile); err != nil {

				golog.Err("!!!htmlfileexist:DeleteHtmlFile: " + err.Error())

			} else {

				golog.Info("htmlfileexist:DeleteHtmlFile: " + htmlfile + " removed OK")

			}

		}

	}

}
