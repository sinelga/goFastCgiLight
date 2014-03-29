package main

import (
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"math/rand"
	//	"os"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	dir := "tiedotDB"

	myDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}
	if err := myDB.Create("Sites", 2); err != nil {
		panic(err)
	}

	for name := range myDB.StrCol {
		fmt.Printf("I have a collection called %s\n", name)
	}
	//	myDB.Scrub("Sites")

	sites := myDB.Use("Sites")

	if err := sites.Index([]string{"Pathinfo"}); err != nil {
		panic(err)
	}
	if err := sites.Index([]string{"Created"}); err != nil {
		panic(err)
	}
	if err := sites.Index([]string{"Updated"}); err != nil {
		panic(err)
	}

	if err := sites.Index([]string{"Hits"}); err != nil {
		panic(err)
	}

	for path := range sites.SecIndexes {
		fmt.Printf("I have an index on path %s\n", path)
	}

	myDB.Flush()
	myDB.Close()

}
