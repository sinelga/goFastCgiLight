package main

import (
	"checksiteexist"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"log"
	"log/syslog"
	"math/rand"
	"newsite"
	"time"
	//	"domains"
	"updatesite"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	} else {

		golog.Info("Start sitesmaker")

	}

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	dir := "tiedotDB"

	tDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}
	defer tDB.Flush()
	defer tDB.Close()

	for name := range tDB.StrCol {
		fmt.Printf("I have a collection called %s\n", name)
	}

	pathinfo := "fi_FI/porno/www.test5.com/index.html"

	idsitesarr := checksiteexist.CheckDB(*golog, tDB, pathinfo)

	if len(idsitesarr) == 1 {

		golog.Info("sitesmaker:Update " + pathinfo)
		updatesite.Update(*golog, tDB,idsitesarr)

	} else if len(idsitesarr) == 0 {

		golog.Info("sitesmaker:Createnew site "+pathinfo)

		newsite.CreateSite(*golog, tDB, pathinfo)

		//		fmt.Println(len(sitesarr))
	} else {

		golog.Err("sitesmaker: DB someting wrong " + pathinfo)

	}


	//	tDB.Flush()
	//	tDB.Close()
}
