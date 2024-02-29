package database

var connection string

func init() {
	connection = "Mysql"
}

func GetDbConnect() string {
	return connection
}
