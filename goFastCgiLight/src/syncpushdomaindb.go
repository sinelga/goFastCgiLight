package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"os"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var domainFlag = flag.String("domain", "", "optional like naista.fi ...")

func main() {
	flag.Parse() // Scan the arguments list

	var locale string
	var themes string
	var domain string

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}
	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {

		golog.Err("sysncpushdomainsdb: " + err.Error())
		os.Exit(1)
	} else {

		locale = *localeFlag
		themes = *themesFlag
		domain = *domainFlag

		if locale != "" && themes != "" {

			var extdomainsarr []domains.Domaincsv
//			var extdomainsarrtoinsert []domains.Domaincsv
//			var domainstopush []string
			var hostsarr []string

			db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/singolight.db")
			if err != nil {

				golog.Err(err.Error())
			}

			var sqlstr string

			//			sqlstr = "select host from hosts where Locale='" + locale + "' and Themes='" + themes + "' order by random()"
			sqlstr = "select host from hosts where Locale='" + locale + "' and Themes='" + themes + "'"
			rows, err := db.Query(sqlstr)
			if err != nil {

				golog.Err(err.Error())
			}
			defer rows.Close()

			for rows.Next() {

				var host string
				rows.Scan(&host)

				hostsarr = append(hostsarr, host)

			}

			if domain != "" {

				for _, host := range hostsarr {

					if r, err := c.Do("ZADD", "pushdomains", "1", host+"."+domain); err != nil {
						golog.Err("syncpushdomains: " + err.Error())

					} else {

						//				golog.Info(r)
						fmt.Println("r ", r)

					}
				}

			} else {

				// probably I don't need it

				//			sqlstr = "select Locale,Themes,Domain,Ip from extdomains where Block=0 and Locale='" + locale + "' and Themes='" + themes + "' order by random()"

				sqlstr = "select Locale,Themes,Domain,Ip from extdomains where Block=0 and Locale='" + locale + "' and Themes='" + themes + "'"
				golog.Info("syncpushdomaindb: Start " + sqlstr)

				rows, err = db.Query(sqlstr)
				if err != nil {

					golog.Err(err.Error())
				}
				defer rows.Close()

				for rows.Next() {

					var extdomain domains.Domaincsv
					rows.Scan(&extdomain.Locale, &extdomain.Themes, &extdomain.Domain, &extdomain.Ip)
					extdomainsarr = append(extdomainsarr, extdomain)
					fmt.Println("extdomain.Domain ", extdomain.Domain)

					for _, host := range hostsarr {

						if r, err := c.Do("ZADD", "pushdomains", "1", host+"."+extdomain.Domain); err != nil {
							golog.Err("syncpushdomains: " + err.Error())

						} else {

							fmt.Println("r ", r)

						}
					}

				}
				rows.Close()
				//				db.Close()

//				c.Flush()
//				c.Close()

				//				db, err = sql.Open("sqlite3", "/home/juno/git/goFastCgiLight/goFastCgiLight/pushdomains.db")
				//				if err != nil {
				//
				//					golog.Err(err.Error())
				//				}
				//				sqlstr = "select Domain from pushdomains"
				//				rows, err = db.Query(sqlstr)
				//				if err != nil {
				//					golog.Err(err.Error())
				//				}
				//				defer rows.Close()
				//
				//				for rows.Next() {
				//					var domaintopush string
				//					rows.Scan(&domaintopush)
				//					domainstopush = append(domainstopush, domaintopush)
				//				}
				//
				//				rows.Close()
				//
				//				var newpushdomain bool
				//
				//				for _, extdomain := range extdomainsarr {
				//
				//					newpushdomain = false
				//
				//					for _, pushdomain := range domainstopush {
				//
				//						if extdomain.Domain == pushdomain {
				//
				//							newpushdomain = true
				//						}
				//
				//					}
				//
				//					if !newpushdomain {
				//
				//						golog.Info("Add " + extdomain.Domain)
				//						extdomainsarrtoinsert = append(extdomainsarrtoinsert, extdomain)
				//
				//					}
				//
				//				}
				//
				//				if len(extdomainsarrtoinsert) > 0 {
				//
				//					sqlstr = "insert into pushdomains(Locale,Themes,Domain,Ip) values(?,?,?,?)"
				//
				//					tx, err := db.Begin()
				//					if err != nil {
				//						golog.Err(err.Error())
				//					}
				//					stmt, err := tx.Prepare(sqlstr)
				//					if err != nil {
				//						log.Fatal(err)
				//						golog.Err(err.Error())
				//					}
				//					defer stmt.Close()
				//
				//					for _, domain := range extdomainsarrtoinsert {
				//
				//						if _, err = stmt.Exec(domain.Locale, domain.Themes, domain.Domain, domain.Ip); err != nil {
				//
				//							golog.Err(err.Error())
				//
				//						}
				//
				//					}
				//					stmt.Close()
				//					tx.Commit()
				//
				//				}
				//
			}
			//
			db.Close()
			c.Flush()
			c.Close()

		} else {

			print("Check syncpushdomaindb -h")
		}
	}

}
