package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("postgres", "host=app-66eb4999-9643-4f35-9641-7d6b85ce00fe-do-user-8435387-0.b.db.ondigitalocean.com port=25060 user=mplp password=oa1ddrukmvesu7gi dbname=mplp sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
}
