package getphrasesfromdb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	//	"comutils"
	"database/sql"
	//	"domains"
	//	"insertsentences"
	//	"io/ioutil"
	//	"log"
	"log/syslog"
)

func GetAll(golog syslog.Writer, db *sql.DB, locale string, themes string) []string {

	var phrasesarr []string

	sqlstr := "select phrase from phrases where locale='" + locale + "' and themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {

		golog.Crit(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var phrase string
		rows.Scan(&phrase)
		phrasesarr = append(phrasesarr, phrase)

	}
	rows.Close()

	return phrasesarr
}
