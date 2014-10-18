package templ_funcmap

import (
	"comutils"
	"domains"
	"strings"
	//"fmt"
	"math/rand"
	"time"
)

func FirstWord(s string) string {

	words := strings.Fields(s)

	return words[0]

}

func FirstWordFromSenteces(s []string) string {

	words := strings.Fields(s[0])

	return words[0]

}

func FirstWordFromAllParagraphs(s []domains.Paragraph) string {

	words := strings.Fields(s[0].Ptitle)

	return words[0]
}

func SplitPathOnWords(s string) string {

	var link_title string
	words := strings.Split(strings.TrimSuffix(s, ".html"), "/")

	for i, word := range words {

		if i == 1 {

			link_title = comutils.UpcaseInitial(word)

		} else {

			link_title = link_title + " " + word

		}

	}

	return link_title

}

func SplitDomainName(s string) string {

	var title string
	words := strings.Split(s, ".")

	wordslen := len(words)

	title = comutils.UpcaseInitial(words[wordslen-2])

	return title

}

func SomeSentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 4; i++ {

		someSentences = append(someSentences, s.Paragraph.Sentences[i])

	}

	return someSentences

}

func RandomAndLimitPages(s []string) []string {

	var pages []string

	if len(s) > 10 {

		rand.Seed(time.Now().UTC().UnixNano())
		dest := make([]string, len(s))
		perm := rand.Perm(len(s))

		for i, v := range perm {
			dest[v] = s[i]
		}

		for i := 0; i < 10; i++ {

			pages = append(pages, dest[i])

		}
	} else {

		pages = s
	}

	return pages
}

func LatSentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph.Sentences)

	if sentNum > 15 {

		for i := sentNum - 15; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph.Sentences[i])

		}
	}

	return someSentences

}
