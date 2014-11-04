package main

import (
	"domains"
	"findfreeparagraph"
	"flag"
	"fmt"
	"homepage/createhomepages"
	"io/ioutil"
	"log/syslog"
	"strings"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var startparameters []string

var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")
var siteFlag = flag.String("site", "", "any valid domain ")
var variantFlag = flag.Int("variant", 0, " 0 1 2 3 ...")
//var paragraph domains.Paragraph
var paragraphs []domains.Paragraph

//var sitesmap map[string]domains.Sitetohomepage
var pages []string

var sitetohomepage domains.Sitetohomepage

func main() {

	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	locale := *localeFlag
	themes := *themesFlag
	site := *siteFlag
	variant := *variantFlag

	content, err := ioutil.ReadFile("/home/juno/git/goFastCgiLight/goFastCgiLight/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	startparameters = []string{strings.TrimSpace(parameters[0]), strings.TrimSpace(parameters[1]), strings.TrimSpace(parameters[2])}

	//	sitesmap = make(map[string]domains.Sitetohomepage)
	//	paragraphs = make([]domains.Paragraph,10)

	for i := 0; i < 10; i++ {

		paragraph := findfreeparagraph.FindFromQ(*golog, locale, themes, "google", startparameters)
		paragraphs = append(paragraphs, paragraph)
		pages = append(pages, paragraph.Plocallink)
	}

	currenttime := time.Now().Local()

	sitetohomepage = domains.Sitetohomepage{

		Locale:    locale,
		Themes:    themes,
		Site:      site,
		Pages:     pages,
		Paragraph: paragraphs,
		Variant:   variant,
		Created:   currenttime.Format("2006-01-02 15:04:05"),
		Updated:   currenttime.Format("2006-01-02 15:04:05"),
	}


	createhomepages.CreatePages(*golog, sitetohomepage)

}
