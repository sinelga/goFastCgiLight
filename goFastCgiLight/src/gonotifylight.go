package main

import (
	"flag"
	"fmt"
	"inotify"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"pipelingpush"
	"strings"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var watcher *inotify.Watcher
var pushtoQueueArr [][]string

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

	var INMODIFY bool=false
	for {
		select {
		case ev := <-watcher.Event:
//						log.Println("event:", ev.String(), ev.Mask)
						
						
			if 	 ev.Mask == inotify.IN_OPEN {
				
				INMODIFY = false
			
			}		

			if ev.Mask == inotify.IN_CREATE|inotify.IN_ISDIR {
				//				fmt.Println("Create dir ", ev.Name)
//				watcher.AddWatch(ev.Name, inotify.IN_CREATE|inotify.IN_ISDIR|inotify.IN_CLOSE_NOWRITE|inotify.IN_DELETE)
				watcher.AddWatch(ev.Name, inotify.IN_CREATE|inotify.IN_ISDIR|inotify.IN_CLOSE_NOWRITE|inotify.IN_DELETE|inotify.IN_OPEN|inotify.IN_CLOSE_WRITE)
				err = filepath.Walk(ev.Name, scan)
				if err != nil {
					golog.Err(err.Error())
				}

			}
//			if ev.Mask == inotify.IN_MODIFY {
//				fmt.Println("IN_MODIFY ", ev.Name)
//				INMODIFY = true
//				
//			}
			if ev.Mask == inotify.IN_CLOSE_WRITE {
//				fmt.Println("IN_CLOSE_WRITE ", ev.Name)
				INMODIFY = true
			}
			
			if ev.Mask == inotify.IN_CLOSE_NOWRITE {
//			if ev.Mask == inotify.IN_CLOSE {
																
				if !INMODIFY {
					fmt.Println("Hit on ", ev.Name,INMODIFY )				
					pushHit(*golog, ev.Name)
				}
			}
			if ev.Mask == inotify.IN_DELETE|inotify.IN_ISDIR {
				//				fmt.Println("delete dir ", ev.Name)
				watcher.RemoveWatch(ev.Name)
			}

		case err := <-watcher.Error:
//			log.Println("gonotifylight !!! error:", err)
			golog.Err("gonotifylight: "+err.Error())
		}
	}
}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if fileInfo.IsDir() {

		watcher.AddWatch(path, inotify.IN_CREATE|inotify.IN_ISDIR|inotify.IN_CLOSE_NOWRITE|inotify.IN_DELETE|inotify.IN_OPEN|inotify.IN_CLOSE_WRITE)
//		watcher.AddWatch(path, inotify.IN_ALL_EVENTS)

	}

	return
}

func pushHit(golog syslog.Writer, path string) {

	htmlfilesp := strings.Split(path, "/")
	if len(htmlfilesp) > 2 {
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

		arr := []string{locale, themes, host, pathinfo}

		pushtoQueueArr = append(pushtoQueueArr, arr)

	} else {

		golog.Err("gonotifylight check !!! " + path)
	}

	if len(pushtoQueueArr) > 2 {

		golog.Info("gonotifylight:pipelingpush.PushInQueue 200 hits") 
		pipelingpush.PushInQueue(golog, "redis", pushtoQueueArr)
		pushtoQueueArr = pushtoQueueArr[:0]

		//		fmt.Println("pushtoQueueArr len aftert purge ", len(pushtoQueueArr))

	}

}
