package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/common/log"

	"github.com/prometheus/prometheus/storage/remote"
)

type valueConverter struct {
	float float64
	str   string
}

// Implements sql.Scanner
func (vc *valueConverter) Scan(v interface{}) error {
	vc.str = fmt.Sprintf("%v", v)
	switch v.(type) {
	case int64:
		vc.float = float64(v.(int64))
	case float64:
		vc.float = v.(float64)
	case []byte:
		vc.str = fmt.Sprintf("%s", v)
	}
	return nil
}

func runQuery(q *remote.Query, db *sql.DB) []*remote.TimeSeries {
	resp := []*remote.TimeSeries{}
	var query string
	var hasJob bool
	for _, m := range q.Matchers {
		if m.Type == remote.MatchType_EQUAL && m.Name == "query" {
			query = m.Value
		}
		if m.Type == remote.MatchType_EQUAL && m.Name == "job" && m.Value == "sql" {
			hasJob = true
		}

	}
	if query == "" || !hasJob {
		return resp
	}

	log.Infof("Running query %q", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Errorf("Error executing query: %s", err)
		return resp
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Errorf("Error fetching column names: %s", err)
		return resp
	}
	for rows.Next() {
		items := make([]interface{}, len(columns))
		for i := range columns {
			items[i] = &valueConverter{}
		}
		err = rows.Scan(items...)
		if err != nil {
			log.Errorf("Error scanning row: %s", err)
			return resp
		}

		// Make timeseries from row.
		ts := &remote.TimeSeries{}
		value := 1.0
		for i, name := range columns {
			vc := items[i].(*valueConverter)
			// If there's a column called "value" use it as the value,
			// otherwise use the column as a label.
			if name == "value" {
				value = vc.float
			} else {
				ts.Labels = append(ts.Labels, &remote.LabelPair{Name: name, Value: vc.str})
			}
		}

		// Create the same sample every 60s.
		for t := q.StartTimestampMs; t <= q.EndTimestampMs; t += 60000 {
			ts.Samples = append(ts.Samples, &remote.Sample{Value: value, TimestampMs: t})
		}
		resp = append(resp, ts)
	}
	log.Infof("Returned %d time series.", len(resp))

	return resp
}

func main() {
	// Connect to and initialise database.
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS foo(bar TEXT NOT NULL PRIMARY KEY, baz TEXT, value INT)")
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
	}
	_, err = db.Exec(`REPLACE INTO foo(bar, baz, value) VALUES
    ("hello", "world", 42),
    ("prometheus", "started", 2012),
    ("company", "robust perception", 1)
 `)
	if err != nil {
		log.Fatalf("Error filling table: %s", err)
	}
	log.Info("Database opened and setup.")

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		reqBuf, err := ioutil.ReadAll(snappy.NewReader(r.Body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req remote.ReadRequest
		if err := proto.Unmarshal(reqBuf, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(req.Queries) != 1 {
			http.Error(w, "Can only handle one query.", http.StatusBadRequest)
			return
		}

		resp := remote.ReadResponse{
			Responses: []*remote.QueryResponse{
				{Timeseries: runQuery(req.Queries[0], db)},
			},
		}
		data, err := proto.Marshal(&resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/x-protobuf")
		if _, err := snappy.NewWriter(w).Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
