package main

import (
	"domains"
	"findfreeparagraph"
	"flag"
	"fmt"
	"homepage/checkrootdir"
	"homepage/createhomepages"
	"log/syslog"
	"os"
	"path/filepath"
	//	"startones"
	"io/ioutil"
	"strings"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

var sitesmap map[string]domains.Sitetohomepage

var startparameters []string

var paragraph domains.Paragraph

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	content, err := ioutil.ReadFile("/home/juno/git/goFastCgiLight/goFastCgiLight/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	startparameters = []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

	sitesmap = make(map[string]domains.Sitetohomepage)

	err = filepath.Walk("/home/juno/git/goFastCgiLight/goFastCgiLight/www", scan)
	if err != nil {

		golog.Err(err.Error())
	}

	createhomepages.CreatePages(*golog, sitesmap)

}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if !fileInfo.IsDir() {

		siteinfo := checkrootdir.Check(*golog, path)

		mapkey := siteinfo[0] + siteinfo[1] + siteinfo[2]
		//		fmt.Println(mapkey)

		if len(siteinfo) == 4 && len(strings.Split(siteinfo[2], ".")) > 1 {

			_, ok := sitesmap[mapkey]

			if !ok {

				paragraph = findfreeparagraph.FindFromQ(*golog, "fi_FI", "porno", "google", startparameters)

				//				fmt.Println(paragraph.Pphrase)

				sitesmap[mapkey] = domains.Sitetohomepage{

					Locale:    siteinfo[0],
					Themes:    siteinfo[1],
					Site:      siteinfo[2],
					Pages:     []string{siteinfo[3],paragraph.Plocallink},
					Paragraph: paragraph,
				}

			} else {

				pages := append(sitesmap[mapkey].Pages, siteinfo[3])
//				pages = 
				paragraph = sitesmap[mapkey].Paragraph

				sitesmap[mapkey] = domains.Sitetohomepage{
					Locale:    siteinfo[0],
					Themes:    siteinfo[1],
					Site:      siteinfo[2],
					Pages:     pages,
					Paragraph: paragraph,
				}

			}

		} else {
			
			golog.Info(" nothing to do !!! for "+mapkey)

		}

	}

	return

}
