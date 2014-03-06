package newsite

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
//	"encoding/json"
	"time"
	"domains"
	"findfreeparagraph"
)

func CreateSite(golog syslog.Writer, tDB *db.DB, pathinfo string) {
		
	sites := tDB.Use("Sites")
	
	nowUnix :=time.Now().Unix()
	var nowUnixInt int
	
	nowUnixInt = int(nowUnix)
	
	var paragraphs []domains.Paragraph

	paragraph := findfreeparagraph.FindFromQ(golog,"fi_FI","porno")
	
	paragraphs = append(paragraphs,paragraph)
	
	
	docID, err :=sites.Insert(map[string]interface{}{
	
		"Pathinfo": pathinfo,
		"Created": nowUnixInt,
		"Updated": nowUnixInt,
		"Hits": 0,
		"Paragraphs": paragraphs,	
	
	} )
	if err != nil {
		panic(err)
	} else {
	
		fmt.Println(docID)
	
	}
	
//	sites.Close()
	
	
//	queryStr := `[{"eq": "fi_FI/porno/www.test.com/index.html", "in": ["Pathinfo"]}]`
//	queryStr := `[{"int-from": 1393734423,"int-to": 1393735206,"limit": 3,"in": ["Created"]}]`
//	
//	var query interface{}
//	json.Unmarshal([]byte(queryStr), &query)
//	
//	queryResult := make(map[uint64]struct{})
//	if err := db.EvalQuery(query, sites, &queryResult); err != nil {
//		panic(err)
//	}
//	for id := range queryResult {
//		fmt.Printf("Query returned document ID %d\n", id)
//	}
//	var readBack domains.Site
//	for id := range queryResult {
//		sites.Read(id, &readBack)
//		fmt.Printf("Query returned document %v\n", readBack.Created)
//	}
	
	
	
}
