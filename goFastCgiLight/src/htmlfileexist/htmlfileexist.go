package htmlfileexist

import (
	//	"addfreeparagraph"
	//	"checkdbexist"
	//	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	//	"createpage"
	//	"database/sql"
	//	"domains"
	//	"getalldbparagraphs"
	"log/syslog"
	//	"checkpathinfo"
	"os"
	//	"sentencesforpr"
	//	"strconv"
	"checksiteexist"
	"github.com/HouzuoGuo/tiedot/db"
//	"math/rand"
	"newsite"
//	"time"
	"updatesite"
)

func StartCheck(golog syslog.Writer, col *db.Col,locale string, themes string, site string, pathinfo string) {

	htmlfile := string("www/" + locale + "/" + themes + "/" + site + pathinfo)

	golog.Info("htmlfileexist:StartCheck: htmlfile " + htmlfile)

//	rand.Seed(time.Now().UTC().UnixNano())
//
//	dir := "tiedotDB"
//
//	tDB, err := db.OpenDB(dir)
//	if err != nil {
//		panic(err)
//	}
//	defer tDB.Flush()
//	defer tDB.Close()

	idsitesarr := checksiteexist.CheckDB(golog, col, htmlfile)

	if len(idsitesarr) == 1 {

		golog.Info("sitesmaker:Update " + htmlfile)
		updatesite.Update(golog, col, idsitesarr)

	} else if len(idsitesarr) == 0 {

		golog.Info("sitesmaker:Createnew site " + htmlfile)

		newsite.CreateSite(golog, col, htmlfile)

		//		fmt.Println(len(sitesarr))
	} else {

		golog.Err("sitesmaker: DB someting wrong " + htmlfile)

	}

	//	var paragraphsarr []domains.Paragraph
	//
	//	thispathinfo := checkpathinfo.Check(pathinfo)
	//
	//	db, err := sql.Open("sqlite3", "file:gofast.db?cache=shared&mode=rwc")
	//	if err != nil {
	//		golog.Err(err.Error())
	//	}
	//
	//	webcontents := checkdbexist.Checkdb(golog, db, host, thispathinfo)
	//
	//	if webcontents.Rowid > 0 {
	//
	//		deltamin := int(time.Since(webcontents.Updated).Minutes())
	//
	//		if webcontents.Hits < 10 && deltamin >= 0 {
	//
	//			paragraphsarr = getalldbparagraphs.GetAllPr(golog, db, webcontents.Rowid, host)
	//
	//			for i, paragraph := range paragraphsarr {
	//
	//				paragraphsarr[i].Sentences = sentencesforpr.GetSents(golog, db, paragraph.Rowid)
	//			}
	//
	//			addfreeparagraph.AddPr(golog, db, webcontents.Rowid, webcontents.Locale, webcontents.Themes)
	//
	//			webcontents.Paragraphs = paragraphsarr
	//			createpage.CreatePg(golog, htmlfile, webcontents)
	//
	//		} else {
	//
	//			golog.Info("Dont Update page hits > (set 10) --> " + strconv.Itoa(webcontents.Hits) + " or delatamin (set 0) to shot " + strconv.Itoa(deltamin) + " " + htmlfile)
	//
	//		}
	//
	//	} else {
	//
	////		golog.Alert("!!!htmlfileexist:StartCheck: no record for " + htmlfile + " let's delete file")
	//
	//		DeleteHtmlFile(golog,htmlfile)
	//
	//	}
	//
	//	if db.Close(); err != nil {
	//
	//		golog.Err(err.Error())
	//	}

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
