package main

type DbConnection interface {
	init() error
	GetDatabases() []string
	GetTables(database string) []string
}
