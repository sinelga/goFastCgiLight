package main

import (
	"flag"
	"fmt"
	"inotify"
	"log"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch("www")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
