package templ_funcmap

import (
"strings"
"domains"	
)

func FirstWord(s string) string {

	words := strings.Fields(s)

	return words[0]

}


func FirstWordFromSenteces(s []string) string {
	
	words :=strings.Fields(s[0])
	
	return words[0]
	
}


func FirstWordFromAllParagraphs(s []domains.Paragraph) string {
	
	words :=strings.Fields(s[0].Ptitle)
	
	return words[0]
}