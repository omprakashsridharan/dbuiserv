package main

func main() {
	r := SetupRoutes()
	r.Run()

	// co := ConnectionOptions{
	// 	DbType:   "mysql",
	// 	Hostname: "localhost",
	// 	Port:     3306,
	// 	Username: "root",
	// 	Password: "12qwmn09",
	// }
	// dbConnection := CreateDbConnection(co)
	// err := (*dbConnection).init()
	// if err != nil {
	// 	log.Println("Error while getting connection")
	// }
	// (*dbConnection).GetTableData("desctop", "user")
}
