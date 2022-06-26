package main

import (
	httpx "cloud-storage-api/http"
	"cloud-storage-api/postgres"
	"cloud-storage-api/storage"
	"fmt"
	"log"
	"os"
	"time"
)

type app struct {
	httpServer *httpx.Server
	db         *postgres.DB
}

func main() {
	v := new(app)
	v.run()
}

func (a *app) run() {
	pg := postgres.NewDB(a.dbConnStr())
	a.db = pg
	err := a.db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer a.db.Close()

	var currentTime time.Time
	if err := a.db.DB.QueryRow("SELECT now()").Scan(&currentTime); err != nil {
		log.Fatal(err)
	}

	server := httpx.NewServer()
	server.Addr = "localhost:8080"
	a.httpServer = server

	st := storage.NewStorageService()
	a.httpServer.UploadService = st

	log.Fatal(a.httpServer.ListenAndServe())
}

func (a *app) dbConnStr() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	return connStr
}
