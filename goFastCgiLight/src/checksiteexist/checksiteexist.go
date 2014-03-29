package checksiteexist

import (

	"encoding/json"
//	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
)

func CheckDB(golog syslog.Writer, col *db.Col, pathinfo string) map[uint64]struct{}{

//	sites := tDB.Use("Sites")

	queryStr := `[{"eq": "` + pathinfo + `", "in": ["Pathinfo"]}]`
//	queryStr := `[{"eq": "fi_FI/porno/www.test.com/index.html", "in": ["Pathinfo"]}]`

	var query interface{}
	json.Unmarshal([]byte(queryStr), &query)

	queryResult := make(map[uint64]struct{})
	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}

	return queryResult
//	sites.Close()
}
