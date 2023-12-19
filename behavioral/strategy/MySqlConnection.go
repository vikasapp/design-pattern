package stratgy

import "fmt"

type MySqlConnection struct {
	ConnectionString string
}

func (con *MySqlConnection) Connect() {
	fmt.Println("MySql : " + con.ConnectionString)
}
