package utils

type Database struct {
	host     string
	port     int
	user     string
	password string
	Dsn      string
}

type Conf struct {
	Database
}
