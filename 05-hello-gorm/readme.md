# Hello-Gorm

## GORM Spec
- Full-Featured ORM
- Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
- Hooks (Before/After Create/Save/Update/Delete/Find)
- Eager loading with Preload, Joins
- Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- Context, Prepared Statement Mode, DryRun Mode
- Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
- Composite Primary Key, Indexes, Constraints
- Auto Migrations
- Logger
- Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
- Every feature comes with tests
- Developer Friendly


## 0. 개발환경 설정 : install DBMS (postgresql) by docker
- docker를 이용한 image pull

```shell
❯ docker pull postgres
```
- postgres 설치

```shell
❯ docker run -d -p 5432:5432 -e POSTGRES_PASSWORD='potato' --name local_postgres postgres
```

- 테스트용 계정생성
```shell
❯ docker exec -it local_postgres /bin/bash

root@9d50d28d720e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.

postgres=# create user potato password 'test1234';
CREATE ROLE
postgres=# create database hellogorm owner potato;
CREATE DATABASE
```

## 1. 프로젝트 생성 및 gorm 설치

- go module 생성
```shell
❯ go mod init potato/hello-gorm
```
- install gorm module, postgres Driver

```shell
❯ go get -u gorm.io/gorm
go: downloading gorm.io/gorm v1.23.3
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading github.com/jinzhu/now v1.1.4
go: downloading github.com/jinzhu/now v1.1.5
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added gorm.io/gorm v1.23.3

❯ go get -u gorm.io/driver/postgres
go: downloading gorm.io/driver/postgres v1.3.1
go: downloading github.com/jackc/pgx/v4 v4.14.1
go: downloading github.com/jackc/pgx/v4 v4.15.0
go: downloading github.com/jackc/pgx v3.6.2+incompatible
go: downloading github.com/jackc/pgtype v1.9.1
go: downloading github.com/jackc/pgproto3/v2 v2.2.0
go: downloading github.com/jackc/pgconn v1.10.1
go: downloading github.com/jackc/pgio v1.0.0
go: downloading github.com/jackc/pgproto3 v1.1.0
go: downloading github.com/jackc/pgconn v1.11.0
go: downloading github.com/jackc/chunkreader/v2 v2.0.1
go: downloading github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b
go: downloading github.com/jackc/pgpassfile v1.0.0
go: downloading golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
go: downloading golang.org/x/text v0.3.7
go: downloading github.com/jackc/chunkreader v1.0.0
go: downloading github.com/jackc/pgtype v1.10.0
go: downloading golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
go: added github.com/jackc/chunkreader/v2 v2.0.1
go: added github.com/jackc/pgconn v1.11.0
go: added github.com/jackc/pgio v1.0.0
go: added github.com/jackc/pgpassfile v1.0.0
go: added github.com/jackc/pgproto3/v2 v2.2.0
go: added github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b
go: added github.com/jackc/pgtype v1.10.0
go: added github.com/jackc/pgx/v4 v4.15.0
go: added golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
go: added golang.org/x/text v0.3.7
go: added gorm.io/driver/postgres v1.3.1
```

## 3. ORM을 위한 Object 생성

- /entities/persist/vo.go

```go
type Crop struct {
	gorm.Model
	Name      string
	Quantity  int64
	HarvestAt time.Time
}

type CropStatistics struct {
	Name  string
	Total int64
}
```
- Crop은 ORM의 Object로 정의 하였고, CropStatistics는 Raw Query(Native Query) 예시에서 사용할 목적이다.

- /config/database.go
```go
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
```
- database 설정을 관리할 DatabaseConfig struct를 만들었다.  
외부에서 string 형태러 전달될 dsn을 전달 받아 postgres sql의 드라이버를 통해 실제 Database와 연결될 connection을 만든다.
- SetConnectionPool 를 통해 기본적인 Connection Pool을 설정 할 수 있게 한다.
    - maxIdleConns
    - maxOpenConns

## 4. CRUD 확인코드

```go
func testCrud() {
	// remove all
	var cropModel persist.Crop
	db.Where("1=1").Delete(&persist.Crop{})       // soft delete
	db.Unscoped().Where("1=1").Delete(&cropModel) // delete permanently

	// create a crop
	db.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	db.Create(&persist.Crop{Name: "carrot", Quantity: 10, HarvestAt: time.Now()})
	db.Create(&persist.Crop{Name: "potato", Quantity: 10, HarvestAt: time.Now()})
	db.Create(&persist.Crop{Name: "tomato", Quantity: 5, HarvestAt: time.Now()})

	// find all
	log.Println("=========== First Find ")
	var all []persist.Crop
	if result := db.Find(&all); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
			db.Save(element)
		}
	}
	log.Println("=========== Updated ")
	log.Println("done")

	log.Println("=========== Select raw Query ")
	// Native SQL
	if rows, err := db.Raw("select name, sum(quantity) as total from Crops group by name").Rows(); err != nil {
		log.Fatalln(err)
		return
	} else {
		defer rows.Close()
		log.Println(toCropStatistics(rows))
	}

	log.Println("=========== Find After update ")
	// find all again
	var allAgain []persist.Crop
	if result := db.Find(&allAgain); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(allAgain) > 0 {
		for _, element := range allAgain {
			log.Printf("%v:%v at %vs after created", element.Name, element.Quantity, math.Round(element.UpdatedAt.Sub(element.CreatedAt).Seconds()))
		}
	}

	log.Println("=========== Remove all potatoes(soft_deleted)")

	// remove potato
	db.Where("Name='potato'").Delete(&persist.Crop{})

	log.Println("done")

	log.Println("=========== Find after remove potatoes")

	// find all again
	var afterDeletePotato []persist.Crop
	if result := db.Find(&afterDeletePotato); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(afterDeletePotato) > 0 {
		for _, element := range afterDeletePotato {
			log.Printf("%v:%v before", element.Name, element.Quantity)
		}
	}

	log.Println("=========== Find removed potatoes")
	// find deleted
	var findDeleted []persist.Crop
	if result := db.Unscoped().Where("deleted_at < now()").Find(&findDeleted); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(findDeleted) > 0 {
		for _, element := range findDeleted {
			log.Printf("%v:%v deletedAt %v", element.Name, element.Quantity, element.DeletedAt.Time.Format(time.RFC3339))
		}
	}

	log.Println("=========== Delete potatoes")
	// rollback from soft_delete
	db.Unscoped().Where("deleted_at is not null").Delete(&persist.Crop{})
	log.Println("done")

	log.Println("=========== Find removed potatoes again : will not find")
	// find deleted
	var findDeletedAgain []persist.Crop
	if result := db.Unscoped().Where("deleted_at is not null").Find(&findDeletedAgain); result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalln(result.Error)
	} else if len(findDeletedAgain) > 0 {
		for _, element := range findDeletedAgain {
			log.Printf("%v:%v deletedAt %v", element.Name, element.Quantity, element.DeletedAt.Time.Format(time.RFC3339))
		}
	}
}
```