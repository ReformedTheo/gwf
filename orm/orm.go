package orm

import (
	"database/sql"
	"fmt"
)

type Config struct {
	//Db       string
	User     string
	Password string
	DBName   string
	// Add other configuration fields as needed
}

var connectionString string

func CreateConnection() {

	// new instance of config
	var cfg Config
	//fmt.Printf("Db (mysql, postgrees, MS server):")
	//fmt.Scanln(&cfg.Db

	fmt.Printf("DB User: ")
	fmt.Scanln(&cfg.User)

	fmt.Printf("\n")

	fmt.Printf("DB Password: ")
	fmt.Scanln(&cfg.Password)

	fmt.Printf("\n")

	fmt.Printf("DB Name: ")
	fmt.Scanln(&cfg.DBName)

	setConnectionString(cfg)

}

func setConnectionString(cfg Config) {
	connectionString = cfg.User + ":" + cfg.Password + "@/" + cfg.DBName
}

func getConnectionString() string {
	return connectionString
}

func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
