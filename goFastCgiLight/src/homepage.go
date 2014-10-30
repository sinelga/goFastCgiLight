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
	"homepage/checkmetadata"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")
var sitesmap map[string]domains.Sitetohomepage
var startparameters []string
var paragraph domains.Paragraph

var indexhtml_not_exist map[string]struct{}

func main() {
	flag.Parse() // Scan the arguments list
	
	indexhtml_not_exist = make(map[string]struct{})

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
	
	for key := range(indexhtml_not_exist) {
		
		fmt.Println("bin/startsite -locale=fi_FI -themes=porno -site="+key+" -variant=2")
				
	}
	

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

				currenttime := time.Now().Local()

				_, ok := sitesmap[mapkey]

				if !ok {

					metadata := checkmetadata.Check(*golog, locale, themes, site)

					if metadata != nil {

						variantint, err := strconv.Atoi(metadata[0])
						if err != nil {

							golog.Err(err.Error())

						}

						created := metadata[1]

						paragraph = findfreeparagraph.FindFromQ(*golog, locale, themes, "google", startparameters)
						sitesmap[mapkey] = domains.Sitetohomepage{

							Locale:    locale,
							Themes:    themes,
							Site:      site,
							Pages:     []string{pathshot, paragraph.Plocallink},
							Paragraph: paragraph,
							Variant:   variantint,
							Created:   created,
							Updated:   currenttime.Format("2006-01-02 15:04:05"),
						}

					} else {

//						fmt.Println("!!! no metadata in " + locale + "/" + themes + "/" + site + "/index.hmtl")
						indexhtml_not_exist[site] = struct{}{}

					}

				} else {

					pages := append(sitesmap[mapkey].Pages, siteinfo[3])
					paragraph = sitesmap[mapkey].Paragraph
					created := sitesmap[mapkey].Created
					updated := sitesmap[mapkey].Updated
					variant := sitesmap[mapkey].Variant

					sitesmap[mapkey] = domains.Sitetohomepage{
						Locale:    locale,
						Themes:    themes,
						Site:      site,
						Pages:     pages,
						Paragraph: paragraph,
						Variant:   variant,
						Created:   created,
						Updated:   updated,
					}

				}

			} else {

				golog.Err(" nothing to do !!! for " + mapkey + " Path " + path)

			}

		} else {
			golog.Err(" nothing to do !!! for too SHOT?? " + path)
			

		}

	}

	return

}
