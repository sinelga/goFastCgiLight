package main

import (
	"code.google.com/p/go.net/html"
	//	"code.google.com/p/go.net/html/atom"
	"fmt"
	"log"
	"log/syslog"
	"os"
	//	"strings"
	"domains"
//	"io/ioutil"
	"stringbylineparser"
	"compress/gzip"
//	"updatehtmlpage"
//	"findfreeparagraph"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	var paragrapharr []domains.Paragraph
	var ptitle string
	var pphrase string
	var sentensesarr []string
	var plocallink string

	sf, _ := os.Open("youporn.html.gz")
	s,_ := gzip.NewReader(sf)

	doc, err := html.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {

			if n.Data == "h2" {

				fmt.Println(n.FirstChild.Data)
				ptitle = n.FirstChild.Data

			}

			if n.Data == "h3" {

				fmt.Println(n.FirstChild.Data)
				pphrase = n.FirstChild.Data
			}

			for _, a := range n.Attr {

				if a.Val == "well well-lg" {
					fmt.Println(n.FirstChild.Data)
					pphrase = n.FirstChild.Data

				}
				if a.Val == "well well" {
					fmt.Println(n.FirstChild.Data)
					pphrase = n.FirstChild.Data

				}
				if a.Val == "well well-sm" {
					//					fmt.Println(n.FirstChild.Data)
					sentensesarr = stringbylineparser.Parse(n.FirstChild.Data)
					fmt.Println("sentensescarr len", len(sentensesarr))
					//					paragraph.Sentences = sentensescarr

				}

				if n.Data == "a" {

					if a.Val != "#" {

						fmt.Println(a.Val)
						plocallink = a.Val

					}
				}

			}

			if ptitle != "" && pphrase != "" && plocallink != "" && len(sentensesarr) > 0 {
				//			if ptitle != "" && pphrase != "" && plocallink !="" {

								fmt.Println("Create paragraph ")
				
//				freeparagraph := findfreeparagraph.FindFromQ(*golog, "fi_FI", "porno")				
				
				paragraph := domains.Paragraph{
					Ptitle:     ptitle,
					Pphrase:    pphrase,
					Plocallink: plocallink,
					Phost:      "test.com",
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

	//	fmt.Println("paragrapharr len", len(paragrapharr))
	
//	freeparagraph := findfreeparagraph.FindFromQ(*golog, "it_IT", "finance")
//	paragrapharr =append(paragrapharr,freeparagraph)
//
//	for _, par := range paragrapharr {
//
//		fmt.Println("par.Ptitle", par.Ptitle)
//		fmt.Println("par.Pphrase", par.Pphrase)
//
//	}
//
//	fmt.Println("paragrapharr[0].Sentences  len", len(paragrapharr[0].Sentences))
//
//	webpagebytes := updatehtmlpage.UpdatePage(*golog, "test.com", paragrapharr)
//
//	if err := ioutil.WriteFile("htmlfile.html", webpagebytes, 0666); err != nil {
//		golog.Crit(err.Error())
//	}

	//	w, _ := os.Create("dest.html")
	//	defer w.Close()
	//	html.Render(w, doc)

}
