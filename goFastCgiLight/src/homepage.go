package main

import (
	"flag"
	"fmt"
	"homepage/checkrootdir"
	"log/syslog"
	"os"
	"path/filepath"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

type Sitetohomepage struct {
	Locale string
	Themes string
	Site   string
	Pages  []string
}

var sitesmap map[string]Sitetohomepage

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	sitesmap = make(map[string]Sitetohomepage)

	err := filepath.Walk("/home/juno/git/goFastCgiLight/goFastCgiLight/www", scan)
	if err != nil {

		golog.Err(err.Error())
	}

	fmt.Println("sites ", len(sitesmap))

	for site, siteinfo := range sitesmap {

		fmt.Println(site)
		for _, page := range siteinfo.Pages {

			fmt.Println(page)

		}

	}

}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if !fileInfo.IsDir() {

		siteinfo := checkrootdir.Check(*golog, path)

		if len(siteinfo) == 4 {

			_, ok := sitesmap[siteinfo[2]]

			if !ok {

				//				fmt.Println(siteinfo[2])
				sitesmap[siteinfo[2]] = Sitetohomepage{

					siteinfo[0], siteinfo[1], siteinfo[2], []string{siteinfo[3]}}

			} else {

				pages := append(sitesmap[siteinfo[2]].Pages, siteinfo[3])

				sitesmap[siteinfo[2]] = Sitetohomepage{

					siteinfo[0],
					siteinfo[1],
					siteinfo[2],
					pages,
				}

			}

		}

	}

	return

}
