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

// serve for test
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

	tdDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}
	defer tdDB.Flush()
	defer tdDB.Close()

	for name := range tdDB.StrCol {
		fmt.Printf("I have a collection called %s\n", name)
	}
	col := tdDB.Use("Sites")
	
	pathinfo := "fi_FI/porno/www.test10.com/index.html"

	idsitesarr := checksiteexist.CheckDB(*golog, col, pathinfo)

	if len(idsitesarr) == 1 {

		golog.Info("sitesmaker:Update " + pathinfo)
		updatesite.Update(*golog, col,idsitesarr)

	} else if len(idsitesarr) == 0 {

		golog.Info("sitesmaker:Createnew site "+pathinfo)

		newsite.CreateSite(*golog, col, pathinfo)

		//		fmt.Println(len(sitesarr))
	} else {

		golog.Err("sitesmaker: DB someting wrong " + pathinfo)

	}


	//	tDB.Flush()
	//	tDB.Close()
}
