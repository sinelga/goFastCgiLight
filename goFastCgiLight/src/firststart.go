package main

import (
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"math/rand"
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
	// Create and open database

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
	myDB.Flush()
	myDB.Close()
	
}
