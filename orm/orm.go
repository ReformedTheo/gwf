package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//"GWF/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	//Db       string
	User     string
	Password string
	DBName   string
	// Add other configuration fields as needed
}

// GLOBAL VARIABLES
var connectionString string

func check(e error) {
	if e != nil {
		panic(e)
	}

}
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

	// Construct the connection string
	connectionString = fmt.Sprintf("%s:%s@/%s", cfg.User, cfg.Password, cfg.DBName)

	// Create a map of environment variables
	envMap := map[string]string{
		"CONNECTION_STRING": connectionString,
	}

	//Marshall connection string to env format
	envContent, err := godotenv.Marshal(envMap)
	if err != nil {
		fmt.Printf("Error marshalling env content %s", err)
	}

	//create .env file
	f, err := os.Create(".env")
	check(err)
	defer f.Close()

	//write connection string on .env file
	_, err = f.WriteString(envContent)
	if err != nil {
		fmt.Printf("error writing to .env file: %s", err)
	}
}

func getConnectionString() string {

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the connection string from the loaded environment variables
	connectionString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		log.Fatal("CONNECTION_STRING not found in .env file")
	}

	return connectionString
}

func Connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func InitDb(name string) {

	var cfg Config

	//Get db user and password
	fmt.Printf("DB User: ")
	fmt.Scanln(&cfg.User)

	fmt.Printf("\n")

	fmt.Printf("DB Password: ")
	fmt.Scanln(&cfg.Password)

	// Construct the connection string
	connectionStringInit := fmt.Sprintf("%s:%s@/", cfg.User, cfg.Password)

	db, err := sql.Open("mysql", connectionStringInit)
	check(err)

	_, err = db.Exec("CREATE DATABASE " + name)
	check(err)
}

/*func createTable(){
	db, err := sql.Open("mysql", getConnectionString())
	check(err)
	defer db.Close()
	db.Exec("CREATE TABLE USER")
}*/
