package main

import (
	"flag"
	"fmt"
	"inotify"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch("/home/juno/git/goFastCgiLight/goFastCgiLight/www")
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk("www", scan)
	if err != nil {

		golog.Err(err.Error())
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

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if fileInfo.IsDir() {
	
		fmt.Println("dir ",fileInfo.Name())
	
	
	}

	return
}
