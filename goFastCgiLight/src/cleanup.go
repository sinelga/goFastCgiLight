package main

import (
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"log"
	"log/syslog"
	"math/rand"
	"orphance"
	"os"
	"path/filepath"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var hoursFlag *int = flag.Int("hours", 0, "Hours was not visited")

func main() {
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	flag.Parse() // Scan the arguments list
	//	int64_0 := int64(0)

	if *hoursFlag > 0 {

		startCleanup(*golog, *hoursFlag)

	}

}

func startCleanup(golog syslog.Writer, hours int) {

	hoursint64 := float64(hours)

	dir := "tiedotDB"
	rand.Seed(time.Now().UTC().UnixNano())

	tdDB, err := db.OpenDB(dir)

	defer tdDB.Close()
//	defer tdDB.Scrub("Sites")
	if err != nil {
		panic(err)
		
	}

	col := tdDB.Use("Sites")

	var numScanned = 0
	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		numScanned++

		if !fileInfo.IsDir() {

//						fmt.Println(path, time.Since( fileInfo.ModTime()).Hours())

			if hoursint64 < time.Since(fileInfo.ModTime()).Hours() {

				orphance.LookUp(golog, col, path)
			}

		}
		return
	}

	err = filepath.Walk("www", scan)
	if err != nil {

		golog.Err(err.Error())
	}

	fmt.Println("Total scanned", numScanned)

}
