package createpage

import (

	"log"
	"log/syslog"
	"testing"
	"createfirstgz"
)

func TestCreateHtmlPage(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	var blocksite bool
	var locale string
	var themes string
	var variant string
	var site string
	var pathinfo string
	var testbitepage []byte


	startparameters := []string{"tcp", ":6379", "2"}
	
	
	locale ="fi_FI"
	themes ="porno"
	variant ="3"
	site ="test.com"
	pathinfo = "/test.html"
	blocksite = true

	testbitepage = CreateHtmlPage(*golog, locale, themes, "google", startparameters, blocksite,variant)
	createfirstgz.Creategzhtml(*golog, locale, themes, site, pathinfo,testbitepage)
	
	locale ="fi_FI"
	themes ="porno"
	variant ="3"
	site ="test.com"
	pathinfo = "/test2.html"
	blocksite = true

	testbitepage = CreateHtmlPage(*golog, locale, themes, "google", startparameters, blocksite,variant)
	createfirstgz.Creategzhtml(*golog, locale, themes, site, pathinfo,testbitepage)
	
	
	
	
	
	
	
}
