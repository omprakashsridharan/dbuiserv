package main

import "testing"

func TestShowDatabases(t *testing.T) {
	co := ConnectionOptions{
		DbType:   "mysql",
		Hostname: "localhost",
		Port:     3306,
		Username: "root",
		Password: "12qwmn09",
	}
	dbConnection := CreateDbConnection(co)
	err := (*dbConnection).init()
	if err != nil {
		t.Errorf("Error while getting connection")
	}
}
