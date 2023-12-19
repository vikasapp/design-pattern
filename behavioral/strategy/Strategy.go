package stratgy

func TestStrategyPattern() {
	mysqlConnection := &MySqlConnection{ConnectionString: "Connecting..."}
	mysqlCon := DBConnection{DB: mysqlConnection}
	mysqlCon.DBConnect()
	postgresConnection := &PostgresConnection{ConnectionString: "Connecting..."}
	postgresCon := DBConnection{DB: postgresConnection}
	postgresCon.DBConnect()
	mongodbConnection := &MongoDBConnection{ConnectionString: "Connecting..."}
	mongodbCon := DBConnection{DB: mongodbConnection}
	mongodbCon.DBConnect()
}
