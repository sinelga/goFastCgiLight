package templ_funcmap

import (
"strings"
"domains"
"comutils"
"fmt"			
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


func SplitPathOnWords(s string) string {
	
	var link_title string
	words :=strings.Split(strings.TrimSuffix(s,".html"),"/")
	
	for i,word :=range words {
		
		if i == 1 {
			
			link_title = comutils.UpcaseInitial(word)
			fmt.Println(word)
			
			
		} else {
			
			link_title = link_title +" "+word
			fmt.Println(word)
			
		}
			
		
	}
	
	return link_title
	
}

func SplitDomainName(s string) string {
	
	var title string
	words :=strings.Split(s,".")
	
	wordslen := len(words)
	
	title =comutils.UpcaseInitial(words[wordslen -2]) 
	
	return title
	
}

