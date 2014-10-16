package checkrootdir

import (
//	"fmt"
	"log/syslog"
	"strings"
)

func Check(golog syslog.Writer, path string) []string {

	var pagepath string
	
	var siteinfo []string

	patharr := strings.Split(path, "/")

	if len(patharr) > 10 {
//		fmt.Println(patharr[7])
		siteinfo =append(siteinfo,patharr[7])
//		fmt.Println(patharr[8])
		siteinfo =append(siteinfo,patharr[8])
//		fmt.Println(patharr[9])
		siteinfo =append(siteinfo,patharr[9])
		//		fmt.Println(patharr[10])

		for i := 10; i < len(patharr); i++ {

			pagepath = pagepath + "/" + patharr[i]

		}
		
		pagepath= strings.TrimSuffix(pagepath,".gz")
		siteinfo =append(siteinfo,pagepath)

//		fmt.Println(pagepath)

	}
	
	
	return siteinfo 

}
