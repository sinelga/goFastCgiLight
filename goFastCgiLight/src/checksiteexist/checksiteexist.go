package checksiteexist

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"log/syslog"
)

func CheckDB(golog syslog.Writer, col *db.Col, pathinfo string) map[uint64]struct{} {

	queryStr := `[{"eq": "` + pathinfo + `", "in": ["Pathinfo"]}]`

	var query interface{}
	json.Unmarshal([]byte(queryStr), &query)

	queryResult := make(map[uint64]struct{})
	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}

	return queryResult
}
