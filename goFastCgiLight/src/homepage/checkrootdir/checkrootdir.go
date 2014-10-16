package checkrootdir

import (
	"log/syslog"
	"strings"
)

func Check(golog syslog.Writer, path string) []string {

	var pagepath string
	
	var siteinfo []string

	patharr := strings.Split(path, "/")

	if len(patharr) > 10 {
		siteinfo =append(siteinfo,patharr[7])
		siteinfo =append(siteinfo,patharr[8])
		siteinfo =append(siteinfo,patharr[9])

		for i := 10; i < len(patharr); i++ {

			pagepath = pagepath + "/" + patharr[i]

		}
		
		pagepath= strings.TrimSuffix(pagepath,".gz")
		siteinfo =append(siteinfo,pagepath)

	}
	
	
	return siteinfo 

}
