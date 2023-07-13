package main

import (
	"fmt"
	"nodejs_project/db"
	"nodejs_project/routers"
)

func main() {
	db.OpenToMysql()
	router := routers.InitRouter()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
