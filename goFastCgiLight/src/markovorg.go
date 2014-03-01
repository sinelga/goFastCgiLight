package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"

	"database/sql"
	//		"bufio"
	//	"flag"
	"fmt"
	//	"io"
	"math/rand"
	//	"os"
	"bytes"
	//	"io/ioutil"
	"log"
	"log/syslog"
	"paragraphs/getphrasesfromdb"
	"strings"
	"time"
)

// Prefix is a Markov chain prefix of one or more words.
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

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
//func (c *Chain) Build(r io.Reader) {
//	br := bufio.NewReader(r)
//	p := make(Prefix, c.prefixLen)
//	for {
//		var s string
//		if _, err := fmt.Fscan(br, &s); err != nil {
//			break
//		}
//		key := p.String()
//		c.chain[key] = append(c.chain[key], s)
//		p.Shift(s)
//	}
//}

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

// Generate returns a string of at most n words generated from Chain.
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
			words = append(words, keyword)
		} else {

			words = append(words, next)
		}
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	numWords := 25
	prefixLen := 1

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	c := NewChain(prefixLen) // Initialize a new Chain.

	db, err := sql.Open("sqlite3", "singo.db")
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer

	for _, phrase := range getphrasesfromdb.GetAll(*golog, db, "fi_FI", "porno") {

		buffer.WriteString(phrase + " ")

	}

	c.Write(buffer.Bytes())               // Build chains from standard input.
	text := c.Generate(numWords, "porno") // Generate text.
	fmt.Println(text)
	text = c.Generate(numWords, "pillu")
	fmt.Println(text)
}
