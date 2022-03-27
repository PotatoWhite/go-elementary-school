package main

import (
	"log"
	"potato/hello-gorm/config"
	"potato/hello-gorm/entities/persist"
	"time"
)

var dbConfig *config.DatabaseConfig

func main() {
	dbConfig = &config.DatabaseConfig{}

	// initialize database
	if db, err := dbConfig.Init("host=localhost user=potato password=test1234 dbname=hellogorm port=5432 sslmode=disable TimeZone=Asia/Seoul"); err != nil {
		log.Fatalln(err)
		return
	} else {
		dbConfig = db
	}

	// migrate tables
	if err := dbConfig.DB.AutoMigrate(&persist.Crop{}); err != nil {
		log.Fatalln(err)
		return
	}

	// create a crop
	dbConfig.DB.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	dbConfig.DB.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	dbConfig.DB.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})

	var records []persist.CropStatistics
	if rows, err := dbConfig.DB.Raw("select name, sum(quantity) as total from Crops group by name").Rows(); err != nil {
		log.Fatalln(err)
		return
	} else {
		defer rows.Close()

		for rows.Next() {
			var record persist.CropStatistics
			rows.Scan(
				&record.Name,
				&record.Total)

			records = append(records, record)
		}
	}

	log.Println(records)

}
