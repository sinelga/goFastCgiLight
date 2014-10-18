package templ_funcmap

import (
	"comutils"
	"domains"
	"strings"
	//"fmt"
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

		for i := 0; i < 10; i++ {

			pages = append(pages, s[i])

		}
	} else {

		pages = s
	}

	return pages

}
