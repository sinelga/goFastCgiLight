package createparagraphs

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"comutils"
	//	"cleansentence"
	"domains"
//	"log"
	"log/syslog"
	//	ml "marklib"
	"math/rand"
	"p_create_locallink"
	"prtitlegen"
	"queue/freeparagraphs"
	//	"selectmarkfile"
	//	"time"
	"bytes"
	"fmt"
	"strings"
	"time"
)

//var markfile string
var paragraphs []domains.Paragraph

type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Write(b []byte) {
	br := bytes.NewReader(b)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int, keyword string) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		if i == 0 {
			words = append(words, comutils.UpcaseInitial(keyword))
		} else {

			words = append(words, next)
		}
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func CreatePr(golog syslog.Writer, locale string, themes string, keywords []string, phrases []string, hosts []string, quant int) {

	golog.Info("createparagraphs:CreatePr: Start")

	numWords := 750
	prefixLen := 1

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	c := NewChain(prefixLen)
	var buffer bytes.Buffer

	var sentence string
	for _, phrase := range phrases {

		sentence = sentence + phrase + " "

		if len(sentence) > 80 {

			sentence = sentence + phrase + ". "
			buffer.WriteString(comutils.UpcaseInitial(sentence))
			sentence = ""

		} else {

			sentence = sentence + phrase + " "
		}

	}
	c.Write(buffer.Bytes())

	for i := 0; i < quant; i++ {

		//		var clsentensesarr []string
		//
		prtitle := prtitlegen.Generate(keywords)
		prphrase := comutils.UpcaseInitial(phrases[rand.Intn(len(phrases))]) + "."
		host := hosts[rand.Intn(len(hosts))]
		locallink := p_create_locallink.CreateLink(keywords)
		fmt.Println("-------------")
		fmt.Println(comutils.UpcaseInitial(prtitle[0]) + " " + prtitle[1])
		fmt.Println(prphrase)
		fmt.Println(host)
		fmt.Println(locallink)
		fmt.Println("-------------")

		var finalsentencesarr []string
		text := c.Generate(numWords, prtitle[0]) // Generate text.

		clsentensesarr := strings.Split(text, ". ")

		for _, sentence := range clsentensesarr {

			finalsentencesarr = append(finalsentencesarr, sentence+".")

		}

		text = c.Generate(numWords, prtitle[1]) // Generate text.

		clsentensesarr = strings.Split(text, ". ")

		for _, sentence := range clsentensesarr {

			finalsentencesarr = append(finalsentencesarr, sentence+".")

		}

		for _, fsentence := range finalsentencesarr {

			fmt.Println(fsentence)

		}

		paragraph := domains.Paragraph{

			Ptitle:     comutils.UpcaseInitial(prtitle[0]) + " " + prtitle[1],
			Pphrase:    prphrase,
			Plocallink: locallink,
			Phost:      host,
			Sentences:  finalsentencesarr,
		}
		paragraphs = append(paragraphs, paragraph)

	}

	freeparagraphs.CreateParagraphs(golog, locale, themes, paragraphs)

}
