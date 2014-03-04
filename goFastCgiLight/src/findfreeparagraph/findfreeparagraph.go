package findfreeparagraph

import (
	"domains"
		"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func GetRecqueParagraph(locale string, themes string) domains.Paragraph {

	sentences := []string{"sentence1", "sentence2", "sentence3", "sentence4", "sentence5", "sentence6", "sentence7", "sentence8"}

	paragraph := domains.Paragraph{
		Ptitle:     "ptitle",
		Pphrase:    "pphrase",
		Plocallink: "plocallink",
		Phost:      "phost",
		Sentences:  sentences,
	}

	return paragraph
}

func FindFromQ(golog syslog.Writer,locale string, themes string) domains.Paragraph  {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
	
		golog.Crit(err.Error())
		
	}

	queuename := locale + ":" + themes
//	newdomainsqueuename := "newdomains:" + queuename
//	newdomainsqueuename := "newdomains"
	
//	var newdomain string
//	
//	if quan_newdomains,err := redis.Int(c.Do("SCARD", newdomainsqueuename)); err != nil {
//		log.Fatal(err)
//	} else {
//		
//		if quan_newdomains > 0{
//		
//			 newdomain,_ =redis.String(c.Do("SRANDMEMBER", newdomainsqueuename))
//		
//		}
//	
//	}
			
	var unmarPar domains.Paragraph

	if quan_prs, err := redis.Int(c.Do("LLEN", queuename)); err != nil {

		golog.Crit(err.Error())

	} else {

		if quan_prs > 1 {

			bparagraph, _ := redis.Bytes(c.Do("LPOP", queuename))

			
			err := json.Unmarshal(bparagraph, &unmarPar)
			if err != nil {
	
				golog.Crit(err.Error())

			}

		} else {

	
			golog.Crit("need more free paragraphs!!!!")

		}

	}
	return unmarPar
}
