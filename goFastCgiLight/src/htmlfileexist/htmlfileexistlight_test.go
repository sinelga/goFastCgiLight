package htmlfileexist

import (
	"log"
	"log/syslog"
	"testing"
)

func TestStartCheckNoDB(t *testing.T) {
	
	var blocksite bool
	var locale string
	var themes string
	var site string
	var pathinfo string
	

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")		
	}
	startparameters := []string{"tcp", ":6379", "2"}
	
	blocksite = false
	locale = "fi_FI"
	themes = "porno"
	site = "test.com"
	pathinfo = "/test.html.gz"
	StartCheckNoDB(*golog, locale, themes, site, pathinfo, startparameters, blocksite)
	
	blocksite = false
	locale = "fi_FI"
	themes = "porno"
	site = "test.com"
	pathinfo = "/test.html.gz"
	StartCheckNoDB(*golog, locale, themes, site, pathinfo, startparameters, blocksite)
	
	blocksite = false
	locale = "fi_FI"
	themes = "porno"
	site = "test.com"
	pathinfo = "/test.html.gz"
	StartCheckNoDB(*golog, locale, themes, site, pathinfo, startparameters, blocksite)
	
	
	blocksite = false
	locale = "fi_FI"
	themes = "porno"
	site = "test.com"
	pathinfo = "/test.html.gz"
	StartCheckNoDB(*golog, locale, themes, site, pathinfo, startparameters, blocksite)
	
	
//	
//	
//	blocksite = false
//	locale = "fi_FI"
//	themes = "porno"
//	site = "test.com"
//	pathinfo = "/test.html.gz"
//	StartCheckNoDB(*golog, locale, themes, site, pathinfo, startparameters, blocksite)
				

}
