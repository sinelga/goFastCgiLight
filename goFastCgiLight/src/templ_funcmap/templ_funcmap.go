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

// 0
func Ptitle_0(s domains.Sitetohomepage) string {

	return s.Paragraph[0].Ptitle

}

func Pphrase_0(s domains.Sitetohomepage) string {

	return s.Paragraph[0].Pphrase

}

func Fist5_0_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

	}

	return someSentences

}

func Last5_0_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[0].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

		}
	}

	return someSentences

}

func Fist10_0_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

	}

	return someSentences

}

func Last10_0_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[0].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

		}
	}

	return someSentences

}

//1
func Ptitle_1(s domains.Sitetohomepage) string {

	return s.Paragraph[1].Ptitle

}

func Pphrase_1(s domains.Sitetohomepage) string {

	return s.Paragraph[1].Pphrase

}

func Fist5_1_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[1].Sentences[i])

	}

	return someSentences

}

func Last5_1_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[1].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

		}
	}

	return someSentences

}

func Fist10_1_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[1].Sentences[i])

	}

	return someSentences

}

func Last10_1_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[1].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[0].Sentences[i])

		}
	}

	return someSentences

}

// 2
func Ptitle_2(s domains.Sitetohomepage) string {

	return s.Paragraph[2].Ptitle
}

func Pphrase_2(s domains.Sitetohomepage) string {

	return s.Paragraph[2].Pphrase
}

func Fist5_2_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[2].Sentences[i])

	}

	return someSentences
}

func Last5_2_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[2].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[2].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_2_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[2].Sentences[i])

	}

	return someSentences

}

func Last10_2_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[2].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[2].Sentences[i])

		}
	}

	return someSentences
}

// 3
func Ptitle_3(s domains.Sitetohomepage) string {

	return s.Paragraph[3].Ptitle
}

func Pphrase_3(s domains.Sitetohomepage) string {

	return s.Paragraph[3].Pphrase
}

func Fist5_3_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[3].Sentences[i])

	}

	return someSentences
}

func Last5_3_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[3].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[3].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_3_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[3].Sentences[i])

	}

	return someSentences

}

func Last10_3_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[3].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[3].Sentences[i])

		}
	}

	return someSentences
}

// 4
func Ptitle_4(s domains.Sitetohomepage) string {

	return s.Paragraph[4].Ptitle
}

func Pphrase_4(s domains.Sitetohomepage) string {

	return s.Paragraph[4].Pphrase
}

func Fist5_4_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[4].Sentences[i])

	}
	return someSentences
}

func Last5_4_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[4].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[4].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_4_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[4].Sentences[i])

	}
	return someSentences
}

func Last10_4_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[4].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[4].Sentences[i])

		}
	}

	return someSentences
}


// 5
func Ptitle_5(s domains.Sitetohomepage) string {

	return s.Paragraph[5].Ptitle
}

func Pphrase_5(s domains.Sitetohomepage) string {

	return s.Paragraph[5].Pphrase
}

func Fist5_5_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[5].Sentences[i])

	}
	return someSentences
}

func Last5_5_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[5].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[5].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_5_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[5].Sentences[i])

	}
	return someSentences
}

func Last10_5_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[5].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[5].Sentences[i])

		}
	}

	return someSentences
}

// 6
func Ptitle_6(s domains.Sitetohomepage) string {

	return s.Paragraph[6].Ptitle
}

func Pphrase_6(s domains.Sitetohomepage) string {

	return s.Paragraph[6].Pphrase
}

func Fist5_6_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[6].Sentences[i])

	}
	return someSentences
}

func Last5_6_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[6].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[6].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_6_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[6].Sentences[i])

	}
	return someSentences
}

func Last10_6_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[6].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[6].Sentences[i])

		}
	}

	return someSentences
}

// 7
func Ptitle_7(s domains.Sitetohomepage) string {

	return s.Paragraph[7].Ptitle
}

func Pphrase_7(s domains.Sitetohomepage) string {

	return s.Paragraph[7].Pphrase
}

func Fist5_7_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[7].Sentences[i])

	}
	return someSentences
}

func Last5_7_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[7].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[7].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_7_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[7].Sentences[i])

	}
	return someSentences
}

func Last10_7_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[7].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[7].Sentences[i])

		}
	}

	return someSentences
}

// 8
func Ptitle_8(s domains.Sitetohomepage) string {

	return s.Paragraph[8].Ptitle
}

func Pphrase_8(s domains.Sitetohomepage) string {

	return s.Paragraph[8].Pphrase
}

func Fist5_8_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[8].Sentences[i])

	}
	return someSentences
}

func Last5_8_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[8].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[8].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_8_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[8].Sentences[i])

	}
	return someSentences
}

func Last10_8_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[8].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[8].Sentences[i])

		}
	}

	return someSentences
}

// 9
func Ptitle_9(s domains.Sitetohomepage) string {

	return s.Paragraph[9].Ptitle
}

func Pphrase_9(s domains.Sitetohomepage) string {

	return s.Paragraph[8].Pphrase
}

func Fist5_9_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 5; i++ {

		someSentences = append(someSentences, s.Paragraph[9].Sentences[i])

	}
	return someSentences
}

func Last5_9_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[9].Sentences)

	if sentNum > 5 {

		for i := sentNum - 5; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[9].Sentences[i])

		}
	}

	return someSentences
}

func Fist10_9_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	for i := 0; i < 10; i++ {

		someSentences = append(someSentences, s.Paragraph[9].Sentences[i])

	}
	return someSentences
}

func Last10_9_Sentences(s domains.Sitetohomepage) []string {
	var someSentences []string

	sentNum := len(s.Paragraph[9].Sentences)

	if sentNum > 10 {

		for i := sentNum - 10; i < sentNum; i++ {

			someSentences = append(someSentences, s.Paragraph[9].Sentences[i])

		}
	}

	return someSentences
}
