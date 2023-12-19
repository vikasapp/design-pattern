package stratgy

import "fmt"

type MongoDBConnection struct {
	ConnectionString string
}

func (con *MongoDBConnection) Connect() {
	fmt.Println("MongoDBs : " + con.ConnectionString)
}
