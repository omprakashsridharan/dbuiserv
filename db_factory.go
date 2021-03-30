package main

const (
	MYSQL = "mysql"
)

func CreateDbConnection(co ConnectionOptions) *DbConnection {
	var dbConnection DbConnection
	switch co.DbType {
	case MYSQL:
		dbConnection = &MysqlDbConnection{
			connectionOptions: co,
		}
	}
	return &dbConnection
}
