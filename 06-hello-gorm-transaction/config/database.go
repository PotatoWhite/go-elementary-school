package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DatabaseConfig struct {
	DSN string
	DB  *gorm.DB
}

func (d *DatabaseConfig) Init(dsn string) (*gorm.DB, error) {
	d.DSN = dsn
	if open, err := gorm.Open(postgres.Open(d.DSN)); err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		d.DB = open
	}

	return d.DB, nil
}

func (d *DatabaseConfig) SetConnectionPool(maxFree int, maxOpen int) {
	db, err := d.DB.DB()
	if err != nil {
		return
	}
	db.SetMaxIdleConns(maxFree)
	db.SetMaxOpenConns(maxOpen)
}
