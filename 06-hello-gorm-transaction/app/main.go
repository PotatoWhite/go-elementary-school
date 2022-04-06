package main

import (
	errors "errors"
	"gorm.io/gorm"
	"log"
	"potato/hello-gorm-transaction/config"
	"potato/hello-gorm-transaction/entities/persist"
	"time"
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
	if err := db.AutoMigrate(&persist.Crop{}); err != nil {
		log.Fatalln(err)
		return
	}

	// init Test
	deleteAllCrops()
	log.Printf("============Test Classic Transaction============")

	// Insert And Rollback
	classicTransactionRollback()
	log.Printf("============Create And Rollback 4 objects")
	if expectRowCount(0) {
		log.Fatalf("Not expeted records count")
		return
	}

	classicTransactionCommit()
	log.Printf("============Create And Commit 4 objects")
	if expectRowCount(4) {
		log.Fatalf("Not expeted records count")
		return
	}

	deleteAllCrops()
	log.Printf("============Test Callback Transaction============")
	log.Printf("============InsertAndRollback")
	if err := db.Transaction(insertAndRollBack); err != nil {
		log.Printf(err.Error())
	}

	if expectRowCount(0) {
		log.Fatalf("Not expeted records count")
		return
	}

	log.Printf("============InsertAndCommit")
	if err := db.Transaction(insertAndCommit); err != nil {
		log.Printf(err.Error())
	}

	if expectRowCount(4) {
		log.Fatalf("Not expeted records count")
		return
	}

	deleteAllCrops()
	log.Printf("============Test Nested Transaction============")
	db.Transaction(insert2AndSubRollback)

	if expectRowCount(2) {
		log.Fatalf("Not expeted records count")
		return
	}

	deleteAllCrops()
	log.Printf("============Test SavePoint============")
	db.Transaction(insert2AndRollBackTo)

	if expectRowCount(3) {
		log.Fatalf("Not expeted records count")
		return
	}
}

func insert2AndRollBackTo(tx *gorm.DB) error {
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})

	tx.SavePoint("save01")
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	tx.SavePoint("save02")
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})

	tx.RollbackTo("save02")

	return nil
}

func insert2AndSubRollback(tx *gorm.DB) error {
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})

	if err := tx.Transaction(insertSubAndRollback); err != nil {
		log.Printf(err.Error())
	}

	return nil
}

func insertSubAndRollback(tx *gorm.DB) error {
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})

	return errors.New("some error occurred at sub insert")
}

func deleteAllCrops() *gorm.DB {
	log.Printf("============Wiping Records")
	return db.Exec("delete from crops")
}

func expectRowCount(expect int) bool {
	var found []persist.Crop
	db.Where("1=1").Find(&found)
	log.Printf("Expected %v, %v found", expect, len(found))
	return len(found) != expect
}

func classicTransactionCommit() {
	tx := db.Begin()

	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	tx.Commit()
}

func classicTransactionRollback() {
	tx := db.Begin()

	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	tx.Rollback()
}

func disableDefaultTransaction() *gorm.DB {
	session := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	log.Println("Default transaction has disabled")
	return session
}

func enableDefaultTransaction() *gorm.DB {
	session := db.Session(&gorm.Session{SkipDefaultTransaction: false})
	log.Println("Default transaction has enabled")
	return session
}

func insertAndRollBack(tx *gorm.DB) error {
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	return errors.New("some Error occurred")
}

func insertAndCommit(tx *gorm.DB) error {
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	tx.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	return nil
}
