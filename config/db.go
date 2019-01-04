package config

const DrivierName = "mysql"

type DbConf struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

var MasterDbConfig DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	Username: "root",
	Password: "root",
	DbName:   "iris_db",
}

var SlaveDbConfig DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	Username: "root",
	Password: "root",
	DbName:   "iris_db",
}