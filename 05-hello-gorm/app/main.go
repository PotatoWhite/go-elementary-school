package main

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"log"
	"math"
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

	// remove all
	var cropModel persist.Crop
	dbConfig.DB.Where("1=1").Delete(&persist.Crop{})       // soft delete
	dbConfig.DB.Unscoped().Where("1=1").Delete(&cropModel) // delete permanently

	// create a crop
	dbConfig.DB.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	dbConfig.DB.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	dbConfig.DB.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	dbConfig.DB.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	// find all
	log.Println("=========== First Find ")
	var all []persist.Crop
	if result := dbConfig.DB.Find(&all); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(all) > 0 {
		for _, element := range all {
			log.Printf("%v:%v at %v", element.Name, element.Quantity, element.HarvestAt.Format(time.RFC3339))
		}
	}

	// update
	time.Sleep(1 * time.Second)
	for _, element := range all {
		if element.Name == "potato" {
			element.Quantity += 1
			dbConfig.DB.Save(element)
		}
	}
	log.Println("=========== Updated ")
	log.Println("done")

	log.Println("=========== Select raw Query ")
	// Native SQL
	if rows, err := dbConfig.DB.Raw("select name, sum(quantity) as total from Crops group by name").Rows(); err != nil {
		log.Fatalln(err)
		return
	} else {
		defer rows.Close()
		log.Println(toCropStatistics(rows))
	}

	log.Println("=========== Find After update ")
	// find all again
	var allAgain []persist.Crop
	if result := dbConfig.DB.Find(&allAgain); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(allAgain) > 0 {
		for _, element := range allAgain {
			log.Printf("%v:%v at %vs after created", element.Name, element.Quantity, math.Round(element.UpdatedAt.Sub(element.CreatedAt).Seconds()))
		}
	}

	log.Println("=========== Remove all potatoes(soft_deleted)")

	// remove potato
	dbConfig.DB.Where("Name='potato'").Delete(&persist.Crop{})

	log.Println("done")

	log.Println("=========== Find after remove potatoes")

	// find all again
	var afterDeletePotato []persist.Crop
	if result := dbConfig.DB.Find(&afterDeletePotato); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(afterDeletePotato) > 0 {
		for _, element := range afterDeletePotato {
			log.Printf("%v:%v before", element.Name, element.Quantity)
		}
	}

	log.Println("=========== Find removed potatoes")
	// find deleted
	var findDeleted []persist.Crop
	if result := dbConfig.DB.Unscoped().Where("deleted_at < now()").Find(&findDeleted); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(findDeleted) > 0 {
		for _, element := range findDeleted {
			log.Printf("%v:%v deletedAt %v", element.Name, element.Quantity, element.DeletedAt.Time.Format(time.RFC3339))
		}
	}

	log.Println("=========== Delete potatoes")
	// rollback from soft_delete
	dbConfig.DB.Unscoped().Where("deleted_at is not null").Delete(&persist.Crop{})
	log.Println("done")

	log.Println("=========== Find removed potatoes again : will not find")
	// find deleted
	var findDeletedAgain []persist.Crop
	if result := dbConfig.DB.Unscoped().Where("deleted_at is not null").Find(&findDeletedAgain); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(findDeletedAgain) > 0 {
		for _, element := range findDeletedAgain {
			log.Printf("%v:%v deletedAt %v", element.Name, element.Quantity, element.DeletedAt.Time.Format(time.RFC3339))
		}
	}
}

func toCropStatistics(rows *sql.Rows) []persist.CropStatistics {
	var records []persist.CropStatistics

	for rows.Next() {
		var record persist.CropStatistics
		rows.Scan(
			&record.Name,
			&record.Total)

		records = append(records, record)
	}
	return records
}
