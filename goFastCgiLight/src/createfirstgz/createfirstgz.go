package createfirstgz

import (
	"bytes"
	//	"checkpathinfo"
	"compress/gzip"
	//	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	//	"log"
	"log/syslog"
	"os"
	"path"
	"strings"
	"checkpathinfo"
)

func Creategzhtml(golog syslog.Writer, locale string, themes string, site string, pathinfoclean string, webpagebytes []byte) {

	_htmlfile := checkpathinfo.Check("www/" + locale + "/" + themes + "/" + site + pathinfoclean)

	var index bool = false

	if strings.HasSuffix(_htmlfile, "/index.html") {

		index = true
	}

	pathstr := path.Dir(_htmlfile)
	
	golog.Info("pathstr "+pathstr) 

	if _, err := os.Stat(pathstr); err != nil {

		if os.IsNotExist(err) {

			if err := os.MkdirAll(pathstr, 0700);err != nil {
			
				golog.Crit(err.Error())
			
			} 			

		}

	}

	if !index {

		var b bytes.Buffer

		w := gzip.NewWriter(&b)
		w.Write(webpagebytes)
		w.Close()

		if err := ioutil.WriteFile(_htmlfile, b.Bytes(), 0666); err != nil {
			golog.Crit(err.Error())
		}

	} else {
		if err := ioutil.WriteFile(_htmlfile, webpagebytes, 0666); err != nil {
			golog.Crit(err.Error())
		}

	}
}
