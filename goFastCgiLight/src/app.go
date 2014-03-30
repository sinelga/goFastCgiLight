package main

import (
	"bthandler"
	"jswebserv"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	themes := req.Header.Get("X-THEMES")
	locale := req.Header.Get("X-LOCALE")
	site := req.Header.Get("X-DOMAIN")
	pathinfo := req.Header.Get("X-PATHINFO")
	bot := req.Header.Get("X-BOT")
	rootdir := req.Header.Get("X-ROOTDIR")
	//	golog.Info("REQUEST")

	if bot == "1" {

		bthandler.BTrequestHandler(*golog, resp, req, locale, themes, site, pathinfo)
	} else if bot == "0" {

		jswebserv.JsServ(*golog, resp, req, rootdir, site)
	}

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)

}
