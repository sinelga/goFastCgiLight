package main

import (
	"domains"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/mitchellh/mapstructure"
	"log"
	"log/syslog"
	"time"
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
	dir := "tiedotDB"

	tdDB, err := db.OpenDB(dir)
	defer tdDB.Close()
	if err != nil {
		panic(err)
	}

	col := tdDB.Use("Sites")
	queryStr := `["all"]`

	var allsitesarr []domains.Site

	var query interface{}
	var readBack interface{}
	//		queryStr := `{"has": ["ColCreated"],"limit":1000}`
	//
	json.Unmarshal([]byte(queryStr), &query)
	queryResult := make(map[uint64]struct{})
	if err := db.EvalQuery(query, col, &queryResult); err != nil {

		golog.Crit(err.Error())

	}

	for id := range queryResult {

		var siteobj domains.Site
		col.Read(id, &readBack)

		vals := readBack.(map[string]interface{})
		err := mapstructure.Decode(vals, &siteobj)
		if err != nil {
			panic(err)
		}

		allsitesarr = append(allsitesarr, siteobj)

	}
	
	for _,site := range allsitesarr{

		createddate := time.Unix(site.Created, 0)
		updateddate := time.Unix(site.Updated, 0)	
		fmt.Println(site.Pathinfo,createddate,updateddate,site.Hits)
	
	}	

}
