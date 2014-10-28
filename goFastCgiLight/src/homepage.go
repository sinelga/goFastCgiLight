package main

import (
	"domains"
	"findfreeparagraph"
	"flag"
	"fmt"
	"homepage/checkrootdir"
	//	"homepage/createhomepages"
	"log/syslog"
	"os"
	"path/filepath"
	//	"startones"
	"homepage/checkmetadata"
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

	//	createhomepages.CreatePages(*golog, sitesmap)

}

func scan(path string, fileInfo os.FileInfo, inpErr error) (err error) {

	if !fileInfo.IsDir() {
		//		fmt.Println("path ",path)
		siteinfo := checkrootdir.Check(*golog, path)

		if len(siteinfo) > 2 {

			mapkey := siteinfo[0] + siteinfo[1] + siteinfo[2]

			locale := siteinfo[0]
			themes := siteinfo[1]
			site := siteinfo[2]
			pathshot := siteinfo[3]

			if len(siteinfo) == 4 && len(strings.Split(siteinfo[2], ".")) > 1 && (strings.HasSuffix(pathshot, ".html") || strings.HasSuffix(pathshot, ".gz")) {

				_, ok := sitesmap[mapkey]

				if !ok {

					metadata := checkmetadata.Check(*golog, locale, themes, site)

					if metadata != nil {
						for _, meta := range metadata {

							fmt.Println(meta)
						}

						paragraph = findfreeparagraph.FindFromQ(*golog, locale, themes, "google", startparameters)
						sitesmap[mapkey] = domains.Sitetohomepage{

							Locale:    locale,
							Themes:    themes,
							Site:      site,
							Pages:     []string{pathshot, paragraph.Plocallink},
							Paragraph: paragraph,
						}

					} else {

						fmt.Println("!!! no metadata in " + locale+"/"+themes+"/"+site+"/index.hmtl")

					}

				} else {

					pages := append(sitesmap[mapkey].Pages, siteinfo[3])
					paragraph = sitesmap[mapkey].Paragraph

					sitesmap[mapkey] = domains.Sitetohomepage{
						Locale:    locale,
						Themes:    themes,
						Site:      site,
						Pages:     pages,
						Paragraph: paragraph,
					}

				}

			} else {

				golog.Info(" nothing to do !!! for " + mapkey + " Path " + path)

			}

		} else {
			golog.Info(" nothing to do !!! for too SHOT?? " + path)

		}

	}

	return

}
