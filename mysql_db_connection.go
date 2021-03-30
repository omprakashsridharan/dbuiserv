package main

import (
	"database/sql"
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

func (mdb *MysqlDbConnection) GetTableData(database string, table string) []interface{} {
	query := "SELECT * from `" + database + "`." + table
	rows, err := mdb.db.Raw(query).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}
	count := len(columnTypes)
	finalRows := []interface{}{}
	for rows.Next() {
		scanArgs := make([]interface{}, count)
		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		masterData := map[string]interface{}{}
		for i, v := range columnTypes {
			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}
			masterData[v.Name()] = scanArgs[i]
		}
		finalRows = append(finalRows, masterData)
	}
	return finalRows
}
