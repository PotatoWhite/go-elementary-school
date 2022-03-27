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


## 개발환경 설정 : install DBMS (postgresql) by docker
- docker를 이용한 image pull
```shell
docker pull postgres
```
- postgres 설치
```shell
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD='potato' --name local_postgres postgres
```

- 테스트용 계정생성
```shell
docker exec -it local_postgres /bin/bash
root@9d50d28d720e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.

postgres=# create user potato password 'test1234';
CREATE ROLE
postgres=# create database hellogorm owner potato;
CREATE DATABASE
```

## 프로젝트 생성 
- go module 생성
```shell
go mod init potato/hello-gorm
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