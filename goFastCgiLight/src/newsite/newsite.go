package newsite

import (
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
	"time"
	"domains"
	"findfreeparagraph"
	"strings"
)

func CreateSite(golog syslog.Writer, col *db.Col, pathinfo string) {

	golog.Info("CreateSite: "+pathinfo)
	var htmlfilesplit []string
	var locale string
	var themes string

	htmlfilesplit = strings.Split(pathinfo, "/")
	locale = htmlfilesplit[1]
	themes = htmlfilesplit[2]
		
	nowUnix :=time.Now().Unix()
	var nowUnixInt int
	
	nowUnixInt = int(nowUnix)
	
	var paragraphs []domains.Paragraph

	paragraph := findfreeparagraph.FindFromQ(golog,locale,themes)
	
	paragraphs = append(paragraphs,paragraph)
	
	
	_, err :=col.Insert(map[string]interface{}{
	
		"Pathinfo": pathinfo,
		"Created": nowUnixInt,
		"Updated": nowUnixInt,
		"Hits": 0,
		"Paragraphs": paragraphs,	
	
	} )
	if err != nil {
		
		golog.Crit(err.Error())
	} 
		
	
}
