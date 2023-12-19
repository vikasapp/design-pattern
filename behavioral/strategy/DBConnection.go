package stratgy

type DBConnection struct {
	DB IDBConnection
}

func (conn *DBConnection) DBConnect() {
	conn.DB.Connect()
}
