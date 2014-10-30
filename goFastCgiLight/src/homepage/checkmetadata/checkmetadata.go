package checkmetadata

import (
	"code.google.com/p/go.net/html"
//	"fmt"
	"log/syslog"
	"os"
)

func Check(golog syslog.Writer, locale string, themes string, site string) []string {

	var retarr []string
	indexpagefullpath := "/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + locale + "/" + themes + "/" + site + "/index.html"

	if _, err := os.Stat(indexpagefullpath); err == nil {

		retarr = Getmetadata(golog, indexpagefullpath)
	}

	return retarr
}

func Getmetadata(golog syslog.Writer, indexpagefullpath string) []string {

	var doc *html.Node

	s, err := os.Open(indexpagefullpath)
	if err != nil {
		golog.Err("checkmetadata:Getmetadata " + indexpagefullpath + err.Error())
	}

	defer s.Close()
	doc, err = html.Parse(s)
	if err != nil {
		golog.Err("checkmetadata:Getmetadata " + indexpagefullpath + err.Error())
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
