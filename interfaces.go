package main

type DbConnection interface {
	init() error
	GetDatabases() []string
	GetTables(database string) []string
	GetTableData(database string, table string) []interface{}
}
