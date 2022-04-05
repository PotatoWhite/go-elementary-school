package main

import (
	"gorm.io/gorm"
	"log"
	"potato/hello-gorm-relation/config"
	"potato/hello-gorm-relation/entities/persist"
)

var db *gorm.DB

func main() {
	var dbms config.DatabaseConfig

	// initialize database
	if _db, err := dbms.Init("host=localhost user=potato password=test1234 dbname=hellogorm port=5432 sslmode=disable TimeZone=Asia/Seoul"); err != nil {
		log.Fatalln(err)
		return
	} else {
		db = _db
	}

	// migrate tables
	if err := db.AutoMigrate(&persist.Farm{}, &persist.Crop{}); err != nil {
		log.Fatalln(err)
		return
	}

	// init Test
	log.Printf("============Test Classic Transaction============")

}
