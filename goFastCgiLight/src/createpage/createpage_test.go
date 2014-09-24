package createpage

import (
//	"io/ioutil"
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

	startparameters := []string{"tcp", ":6379", "2"}
	
	
	locale :="fi_FI"
	themes :="porno"
	site :="test.com"
	pathinfo := "/test.html"
	blocksite := false

	testbitepage := CreateHtmlPage(*golog, locale, themes, "google", startparameters, blocksite)
	
	createfirstgz.Creategzhtml(*golog, locale, themes, site, pathinfo,testbitepage)
	
}
