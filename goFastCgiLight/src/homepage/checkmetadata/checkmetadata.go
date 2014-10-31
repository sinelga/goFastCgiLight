package checkmetadata

import (
	"code.google.com/p/go.net/html"
	"compress/gzip"
	"fmt"
	"log/syslog"
	"os"
	"strings"
)

//var index bool = false

func Check(golog syslog.Writer, locale string, themes string, site string, path string) []string {

	var retarr []string

	indexpagefullpath := "/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + locale + "/" + themes + "/" + site + "/index.html"

	_,err := os.Stat(indexpagefullpath)
	if  os.IsNotExist(err) {
		
		fmt.Println("not exist",indexpagefullpath)
		retarr = Getmetadata(golog, path)
		
	} else {
		
		retarr = Getmetadata(golog, indexpagefullpath)
		
	}

	return retarr
}

func Getmetadata(golog syslog.Writer, indexpagefullpath string) []string {

	var index bool = false
	
	if strings.HasSuffix(indexpagefullpath, "/index.html") {

		index = true
	}


	var doc *html.Node

	if !index {

		sf, err := os.Open(indexpagefullpath)
		if err != nil {
			golog.Err("StartCheckNoDB: " + indexpagefullpath + err.Error())
		}
		defer sf.Close()
		s, err := gzip.NewReader(sf)
		if err != nil {
			golog.Err("StartCheckNoDB: " + indexpagefullpath + err.Error())
		}
		defer s.Close()
		doc, err = html.Parse(s)
		if err != nil {
			golog.Err("checkmetadata:Getmetadata " + indexpagefullpath + err.Error())
		}

	} else {

		s, err := os.Open(indexpagefullpath)
		if err != nil {
			golog.Err("checkmetadata:Getmetadata " + indexpagefullpath + err.Error())
		}

		defer s.Close()
		doc, err = html.Parse(s)
		if err != nil {
			golog.Err("checkmetadata:Getmetadata " + indexpagefullpath + err.Error())
		}

	}

	var variant string
	var created string
	var updated string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode {

			if n.Data == "meta" {

				for _, meta := range n.Attr {

					if meta.Key == "variant" {

						variant = meta.Val

					}

					if meta.Key == "created" {

						created = meta.Val

					}

					if meta.Key == "updated" {

						updated = meta.Val

					}

				}

			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	if variant != "" && created != "" && updated != "" {

		return []string{variant, created, updated}

	} else {
		return nil
	}

}
