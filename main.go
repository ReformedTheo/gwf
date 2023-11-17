package main

import (
	// Import ORM
	"GWF/orm"

	// Other needed imports
	"fmt"
	"os"
)

const help string = "\ngo run main.go help to see avaliable commands."

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("No command! %s", help)
		return
	}

	switch args[0] {
	case "initproject":
		if len(args) < 2 {
			fmt.Printf("Project name is required %s", help)
			return
		}
		name := args[1]
		initProject(name)
	case "startserver":
		startServer()
	case "createconnection":
		createORMConnection()
	case "connect":
		connectORM()
	case "help":
		listAll()
	default:
		fmt.Printf("Unknown command! %s", help)
	}
}

func startServer() {
	fmt.Println("Server running on localhost:7830.")
}

func createORMConnection() {
	// calling function from ORM project
	orm.CreateConnection()
}

func connectORM() {
	orm.Connect()
}

func listAll() {
	fmt.Printf("startserver - Run server on localhost:7830 \ncreateconnection - To create a connection with a database \n connect -  Connect and migrate data stored in models to DB")
}

func initProject(name string) {
	orm.InitDb(name)
}
