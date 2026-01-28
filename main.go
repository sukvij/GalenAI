package main

import (
	"fmt"
	empController "sukvij/employment/controller"
	"sukvij/galenfers/configs"
	"sukvij/galenfers/database"
	"sukvij/salary-calculation/controller"
	salmetricsController "sukvij/salary-metrics/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()
	conn, err := database.Connect(cfg)
	fmt.Println(conn, err)

	app := gin.Default()
	empController.EmployeeController(app, conn)
	controller.SalaryCalculationController(app, conn)
	salmetricsController.SalaryMetricsController(app, conn)
	app.Run(":8080")
}
