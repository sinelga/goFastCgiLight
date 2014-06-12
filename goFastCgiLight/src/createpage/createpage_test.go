package createpage

import (
	"log"
	"log/syslog"
	"testing"
		"io/ioutil"
)

func TestCreateHtmlPage(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	testbitepage := CreateHtmlPage(*golog, "fi_FI", "porno")

	if err := ioutil.WriteFile("/tmp/index.html", testbitepage, 0666); err != nil {
		golog.Crit(err.Error())
	}
}
