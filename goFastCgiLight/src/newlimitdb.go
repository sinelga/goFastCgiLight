package main

import (
	"flag"
	"fmt"
	//	"log"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"os"
	"path/filepath"
	"scanfiles"
	"limitdbbyredis"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

//var c, _ = redis.Dial("tcp", ":6379")
var sitesmap map[string]int

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}

	sitesmap = make(map[string]int)

	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	err = filepath.Walk("/home/juno/git/goFastCgiLight/goFastCgiLight/www", scan)
	if err != nil {

		golog.Err(err.Error())
	}

	
	limitdbbyredis.MakeNewDb(*golog,c,sitesmap)
	

}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {


	if !fileInfo.IsDir() {

		site := scanfiles.Scan(*golog, path)

		i, ok := sitesmap[site]

		if !ok {

			sitesmap[site] = 0

		} else {

			sitesmap[site] = i + 1

		}

	}

	return
}
