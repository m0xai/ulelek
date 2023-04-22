package main

import (
	"backend/internal/entry"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func init() {
}

func main() {
	psqlconn := "postgresql://m0xai:1299@localhost/ulelek?sslmode=disable"

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := entry.NewEntryRouter(db)

	http.ListenAndServe(":8080", router)
}

type Entry struct {
	Id         int
	Content    string
	Is_deleted bool
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
