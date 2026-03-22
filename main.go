package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"

	_ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
