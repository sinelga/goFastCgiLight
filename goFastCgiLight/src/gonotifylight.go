package main

import (
	"flag"
	"fmt"
	"inotify"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"pushinqueue"
	"strings"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var watcher *inotify.Watcher

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

	watcher, err = inotify.NewWatcher()
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
//			log.Println("event:", ev.String(), ev.Mask)

			if ev.Mask == 1073742080 {
				fmt.Println("Create dir ", ev.Name)
				//				watcher.AddWatch(ev.Name,inotify.IN_CREATE|inotify.IN_ISDIR|inotify.IN_CLOSE_NOWRITE|inotify.IN_DELETE)
				watcher.AddWatch(ev.Name, inotify.IN_ALL_EVENTS)
				err = filepath.Walk(ev.Name, scan)
				if err != nil {

					golog.Err(err.Error())
				}

			}
			if ev.Mask == 16 {
				fmt.Println("Hit on ", ev.Name)
				pushHit(*golog, ev.Name)
			}
			if ev.Mask == 1073742336 {
				fmt.Println("delete dir ", ev.Name)
				watcher.RemoveWatch(ev.Name)
			}

		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if fileInfo.IsDir() {

		//			watcher.AddWatch(path,inotify.IN_CREATE| inotify.IN_ISDIR|inotify.IN_CLOSE_NOWRITE|inotify.IN_DELETE)
		watcher.AddWatch(path, inotify.IN_ALL_EVENTS)


	}

	return
}

func pushHit(golog syslog.Writer, path string) {

	htmlfilesp := strings.Split(path, "/")
	locale := htmlfilesp[1]
	themes := htmlfilesp[2]
	host := htmlfilesp[3]

	var pathinfo string = "/"

	for i := 4; i < len(htmlfilesp); i++ {

		pathinfo = pathinfo + htmlfilesp[i]

		if i < (len(htmlfilesp) - 1) {
			pathinfo = pathinfo + "/"
		}
	}
	fmt.Println(locale, themes, host, pathinfo)
	pushinqueue.PushInQueue(golog, "redis", locale, themes, host, pathinfo)

}
