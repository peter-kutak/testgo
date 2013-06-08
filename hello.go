package main

import ( 
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=require password=network")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error", err)
	}

	rows, err := db.Query("select datid, datname, procpid, usesysid, usename from pg_stat_activity")
	if err != nil {
		log.Fatal("Error query pg_stat", err)
	}
	defer rows.Close()
	fmt.Println("datid","datname")
	for rows.Next() {
		var datid int64
		var datname string
		var procpid int64
		var usesysid int64
		var usename string
		err = rows.Scan(&datid, &datname, &procpid, &usesysid, &usename)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println( datid, datname, procpid, usesysid, usename )
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows", err)
	}
}

