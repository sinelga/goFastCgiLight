package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"flag"
	"github.com/garyburd/redigo/redis"
		"fmt"
	"log"
	"log/syslog"
	"os"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

func main() {
	flag.Parse() // Scan the arguments list

	var locale string
	var themes string

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}
	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {

		golog.Err("PushInQueue: " + err.Error())
		os.Exit(1)
	} else {

		locale = *localeFlag
		themes = *themesFlag

		var extdomainsarr []domains.Domaincsv
		var extdomainsarrtoinsert []domains.Domaincsv
		var domainstopush []string

		db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/singo.db")
		if err != nil {

			golog.Err(err.Error())
		}

		var sqlstr string

		if locale == "" && themes == "" {

			sqlstr = "select Locale,Themes,Domain,Ip from extdomains where Block=0"

		} else {

			sqlstr = "select Locale,Themes,Domain,Ip from extdomains where Block=0 and Locale='" + locale + "' and Themes='" + themes + "' order by random()"
		}

		golog.Info("syncpushdomain: Start " + sqlstr)

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {

			var extdomain domains.Domaincsv
			rows.Scan(&extdomain.Locale, &extdomain.Themes, &extdomain.Domain, &extdomain.Ip)
			extdomainsarr = append(extdomainsarr, extdomain)
			fmt.Println("extdomain.Domain ",extdomain.Domain)

			if r, err := c.Do("ZADD", "pushdomains 0", extdomain.Domain); err != nil {
				golog.Err("syncpushdomains: " + err.Error())

			} else {
			
//				golog.Info(r)
				fmt.Println("r ",r,)
			}

		}
		rows.Close()
		db.Close()
		
		c.Flush()
		c.Close()

		db, err = sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/pushdomains.db")
		if err != nil {

			golog.Err(err.Error())
		}
		sqlstr = "select Domain from pushdomains"
		rows, err = db.Query(sqlstr)
		if err != nil {
			golog.Err(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var domaintopush string
			rows.Scan(&domaintopush)
			domainstopush = append(domainstopush, domaintopush)
		}

		rows.Close()

		var newpushdomain bool

		for _, extdomain := range extdomainsarr {

			newpushdomain = false

			for _, pushdomain := range domainstopush {

				if extdomain.Domain == pushdomain {

					newpushdomain = true
				}

			}

			if !newpushdomain {

				golog.Info("Add " + extdomain.Domain)
				extdomainsarrtoinsert = append(extdomainsarrtoinsert, extdomain)

			}

		}

		if len(extdomainsarrtoinsert) > 0 {

			sqlstr = "insert into pushdomains(Locale,Themes,Domain,Ip) values(?,?,?,?)"

			tx, err := db.Begin()
			if err != nil {
				golog.Err(err.Error())
			}
			stmt, err := tx.Prepare(sqlstr)
			if err != nil {
				log.Fatal(err)
				golog.Err(err.Error())
			}
			defer stmt.Close()

			for _, domain := range extdomainsarrtoinsert {

				if _, err = stmt.Exec(domain.Locale, domain.Themes, domain.Domain, domain.Ip); err != nil {

					golog.Err(err.Error())

				}

			}
			stmt.Close()
			tx.Commit()

		}

		db.Close()

	}

}
