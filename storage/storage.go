package storage

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"os"
	"time"
)

type Logentry struct {
	Entry string
	Date  *time.Time
}

func (logEntry Logentry) Store() {
	homeDir := os.Getenv("HOME")
	myDBDir := homeDir + "/MyDatabase"
	myDB, err := db.OpenDB(myDBDir)
	check(err)
	devlog := myDB.Use("Devlog")
	if devlog == nil {
		if err := myDB.Create("Devlog"); err != nil {
			panic(err)
		}
	}
	docID, err := devlog.Insert(map[string]interface{}{
		"entry": logEntry.Entry,
		"date":  &logEntry.Date,
	})

	readBack, err := devlog.Read(docID)
	fmt.Printf("Written to log: %v\n", readBack)
	// Gracefully close database
	if err := myDB.Close(); err != nil {
		panic(err)
	}
}

func LoadAllEntries() map[int]*Logentry {
	homeDir := os.Getenv("HOME")
	myDBDir := homeDir + "/MyDatabase"
	myDB, err := db.OpenDB(myDBDir)
	check(err)
	devlog := myDB.Use("Devlog")

	var query interface{}
	json.Unmarshal([]byte(`["all"]`), &query)
	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, devlog, &queryResult); err != nil {
		panic(err)
	}
	entries := make(map[int]*Logentry)
	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := devlog.Read(id)
		if err != nil {
			panic(err)
		}
		entryDate, err := time.Parse(time.RFC3339Nano, readBack["date"].(string))
		check(err)
		entries[id] = &Logentry{
			Entry: readBack["entry"].(string),
			Date:  &entryDate,
		}
	}
	// Gracefully close database
	if err := myDB.Close(); err != nil {
		panic(err)
	}
	return entries
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
