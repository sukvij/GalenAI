package main

import (
	"fmt"
	"sukvij/employment/controller"
	"sukvij/galenfers/configs"
	"sukvij/galenfers/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()
	conn, err := database.Connect(cfg)
	fmt.Println(conn, err)

	app := gin.Default()
	controller.EmployeeController(app, conn)
	app.Run(":8080")
}
