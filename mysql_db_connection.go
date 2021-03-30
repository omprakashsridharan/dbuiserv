package main

import (
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDbConnection struct {
	connectionOptions ConnectionOptions
	db                *gorm.DB
}

func (mdb *MysqlDbConnection) init() error {
	log.Println("Initiating mysql connection")
	if mdb.db == nil {
		username := mdb.connectionOptions.Username
		password := mdb.connectionOptions.Password
		hostname := mdb.connectionOptions.Hostname
		port := mdb.connectionOptions.Port
		dsn := username + ":" + password + "@tcp(" + hostname + ":" + strconv.FormatInt(int64(port), 10) + ")/"
		mysqlConnection := mysql.Open(dsn)
		db, err := gorm.Open(mysqlConnection, &gorm.Config{})
		if err != nil {
			return err
		}
		mdb.db = db
		log.Println("Created connection successfully")
		return nil
	} else {
		log.Println("Reusing existing connection")
		return nil
	}
}

func (mdb *MysqlDbConnection) GetDatabases() []string {
	rows, err := mdb.db.Raw("show databases").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var databases []string

	for rows.Next() {
		var database string
		rows.Scan(&database)
		databases = append(databases, database)
	}
	return databases
}

func (mdb *MysqlDbConnection) GetTables(database string) []string {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = " + "'" + database + "'"
	rows, err := mdb.db.Raw(query).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tables []string

	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
	}
	return tables
}
