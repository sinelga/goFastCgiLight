package main

import (
	"bthandler"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
	"startones"
	"sync"
)

var startOnce sync.Once
var startparameters []string
var sitestoblock map[string]struct{}

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	themes := req.Header.Get("X-THEMES")
	locale := req.Header.Get("X-LOCALE")
	variant := req.Header.Get("X-VARIANT")
	site := req.Header.Get("X-DOMAIN")
	pathinfo := req.Header.Get("X-PATHINFO")
	bot := req.Header.Get("X-BOT")
	uid_got := req.Header.Get("X-NGINX-BROWSER-ID-GOT")
	uid_set := req.Header.Get("X-NGINX-BROWSER-ID-SET")

	startOnce.Do(func() {
		startparameters, sitestoblock = startones.Start(*golog)

		for sitetoblock := range sitestoblock {

			golog.Info("block-> "+sitetoblock)

		}

	})
	
	bloksite := false
	
	_, ok := sitestoblock[site]
	
	if ok {

		bloksite = true			
		
	}
	
	
	golog.Info("uid_set "+uid_set+"  uid_got "+uid_got)
	

	if uid_got !="" && uid_set =="" {
		
		golog.Info("!!!not first visit "+site+pathinfo)
		
	}


	bthandler.BTrequestHandler(*golog, resp, req, locale, themes, site, pathinfo, bot, startparameters,bloksite,variant)

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
