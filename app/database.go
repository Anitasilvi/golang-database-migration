package app

import (
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang-database-migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

//migrate create -ext sql -dir db/migrations create_table_first
//migrate create -ext sql -dir db/migrations create_table_second
//migrate create -ext sql -dir db/migrations create_table_third
//migrate create -ext sql -dir db/migrations create_table_sample_dirty_state
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path db/migrations up
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path db/migrations down
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path db/migrations up/down (tambah versi)
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path db/migrations force (versi sebelumnya) 20231005001137
//migrate -database "mysql://root@tcp(localhost:3306)/golang-database-migration" -path db/migrations version
