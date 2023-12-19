package stratgy

import "fmt"

type PostgresConnection struct {
	ConnectionString string
}

func (con *PostgresConnection) Connect() {
	fmt.Println("Postgress : " + con.ConnectionString)
}
