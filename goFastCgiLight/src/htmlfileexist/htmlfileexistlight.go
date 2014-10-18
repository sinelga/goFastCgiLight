package htmlfileexist

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"compress/gzip"
	"domains"
	"findfreeparagraph"
	//	"fmt"
	"io/ioutil"
	"log/syslog"
	"os"
	"stringbylineparser"
	"strings"
	"updatehtmlpage"
)

func StartCheckNoDB(golog syslog.Writer, locale string, themes string, site string, pathinfo string, startparameters []string, blocksite bool) {

	htmlfile := string("/home/juno/git/goFastCgiLight/goFastCgiLight/www/" + locale + "/" + themes + "/" + site + pathinfo)
	finfo, err := os.Stat(htmlfile)
	if err != nil {
		golog.Err("StartCheckNoDB: " + err.Error())
		return
	}

	if !finfo.IsDir() {

		var index bool = false

		if strings.HasSuffix(htmlfile, "/index.html") {

			index = true
		}

		if pathinfo != "/index.html" {

			var paragrapharr []domains.Paragraph
			var ptitle string
			var pphrase string
			var sentensesarr []string
			var plocallink string

			var doc *html.Node

			if !index {
				sf, err := os.Open(htmlfile)
				if err != nil {
					golog.Err("StartCheckNoDB: " + htmlfile + err.Error())
				}
				defer sf.Close()
				s, err := gzip.NewReader(sf)
				if err != nil {
					golog.Err("StartCheckNoDB: " + htmlfile + err.Error())
				}
				defer s.Close()

				doc, err = html.Parse(s)
				if err != nil {
					golog.Err("StartCheckNoDB: " + htmlfile + err.Error())
				}

			} else {

				s, err := os.Open(htmlfile)
				if err != nil {
					golog.Err("StartCheckNoDB: " + htmlfile + err.Error())
				}

				defer s.Close()
				doc, err = html.Parse(s)
				if err != nil {
					golog.Err("StartCheckNoDB: " + htmlfile + err.Error())
				}
			}

			var f func(*html.Node)

			f = func(n *html.Node) {
				if n.Type == html.ElementNode {

					for _, a := range n.Attr {

						if a.Val == "Ptitle" {
							if n.FirstChild != nil {
								ptitle = n.FirstChild.Data
							}

						}

						if a.Val == "Pphrase" {
							if n.FirstChild != nil {
								pphrase = n.FirstChild.Data
							}

						}
						if a.Val == "Sentences" {
							if n.FirstChild != nil {
								sentensesarr = stringbylineparser.Parse(n.FirstChild.Data)
							}
						}

					}

					if n.Data == "a" {

						for _, a := range n.Attr {

							if a.Key == "href" && a.Val != "#" {
								plocallink = a.Val
							}

						}

					}

					if ptitle != "" && pphrase != "" && len(sentensesarr) > 0 {

						if blocksite {

							plocallink = ""

						}

						paragraph := domains.Paragraph{
							Ptitle:     ptitle,
							Pphrase:    pphrase,
							Plocallink: plocallink,
							Phost:      site,
							Sentences:  sentensesarr,
						}

						paragrapharr = append(paragrapharr, paragraph)
						ptitle = ""
						pphrase = ""
						plocallink = ""
						sentensesarr = sentensesarr[:0]
					}

				}

				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
			f(doc)

			freeparagraph := findfreeparagraph.FindFromQ(golog, locale, themes, "google", startparameters)

			if blocksite {
				freeparagraph.Plocallink = ""

			}

			paragrapharr = append(paragrapharr, freeparagraph)
			webpagebytes := updatehtmlpage.UpdatePage(golog, site, paragrapharr, blocksite)

			if !index {

				var b bytes.Buffer

				w := gzip.NewWriter(&b)
				w.Write(webpagebytes)
				w.Close()

				if err := ioutil.WriteFile(htmlfile, b.Bytes(), 0666); err != nil {
					golog.Crit(err.Error())
				}

			} else {

				if err := ioutil.WriteFile(htmlfile, webpagebytes, 0666); err != nil {
					golog.Crit(err.Error())
				}

			}

		}

	} else {

		golog.Info("htmlfileexist:StartCheckNoDB homepage DO nothing!!! " + htmlfile)

	}

}
