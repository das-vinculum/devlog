package storage

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"os"
	"time"
)

type Logentry struct {
	Entry string
	Date  time.Time
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
		"date":  logEntry.Date.Format(time.UnixDate)})

	readBack, err := devlog.Read(docID)
	fmt.Printf("Written to log: %v\n", readBack)
	// Gracefully close database
	if err := myDB.Close(); err != nil {
		panic(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
