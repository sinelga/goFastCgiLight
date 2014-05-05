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

func StartCheckNoDB(golog syslog.Writer, locale string, themes string, site string, pathinfo string) {

	htmlfile := string("www/" + locale + "/" + themes + "/" + site + pathinfo)
	finfo, err := os.Stat(htmlfile)
	if err != nil {
		golog.Err("StartCheckNoDB: " + err.Error())
		return
	}

	if !finfo.IsDir() {
//		golog.Info("htmlfileexist:StartCheckNoDB: htmlfile " + htmlfile + " OK exist")

		var index bool = false

		if strings.HasSuffix(htmlfile, "/index.html") {

			index = true
		}

		var paragrapharr []domains.Paragraph
		var ptitle string
		var pphrase string
		var sentensesarr []string
		var plocallink string

		var doc *html.Node

		if !index {
			sf, _ := os.Open(htmlfile)
			s, _ := gzip.NewReader(sf)
			if err != nil {
				golog.Err("StartCheckNoDB: " + err.Error())
			}
			doc, err = html.Parse(s)
			if err != nil {
				golog.Err("StartCheckNoDB: " + err.Error())
			}

		} else {

			s, _ := os.Open(htmlfile)
			doc, err = html.Parse(s)
			if err != nil {
				golog.Err("StartCheckNoDB: " + err.Error())
			}
		}

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode {

				if n.Data == "h2" {
					if n.FirstChild != nil {
						ptitle = n.FirstChild.Data
					}
				}

				if n.Data == "h3" {
					if n.FirstChild != nil {
						pphrase = n.FirstChild.Data
					}
				}

				for _, a := range n.Attr {

					if a.Val == "well well-lg" {
						if n.FirstChild != nil {
							ptitle = n.FirstChild.Data
						}

					}
					if a.Val == "well well" {
						if n.FirstChild != nil {
							pphrase = n.FirstChild.Data
						}
					}
					if a.Val == "well well-sm" {
						if n.FirstChild != nil {
							sentensesarr = stringbylineparser.Parse(n.FirstChild.Data)
						}
					}

					if n.Data == "a" {

						if a.Val != "#" {
							plocallink = a.Val

						}
					}

				}

				if ptitle != "" && pphrase != "" && plocallink != "" && len(sentensesarr) > 0 {

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

		freeparagraph := findfreeparagraph.FindFromQ(golog, locale, themes)
		paragrapharr = append(paragrapharr, freeparagraph)
		webpagebytes := updatehtmlpage.UpdatePage(golog, site, paragrapharr)

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

}
