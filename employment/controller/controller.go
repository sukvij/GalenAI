package controller

import (
	"strconv"
	"strings"
	"sukvij/employment/model"
	"sukvij/employment/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func EmployeeController(app *gin.Engine, db *gorm.DB) {
	controller := &Controller{Db: db}
	router := app.Group("/v1/employees")
	router.GET("", controller.getEmployees)
	router.POST("", controller.createEmployee)
	router.GET("/:id", controller.getEmployeeById)
	router.DELETE("/:id", controller.deleteEmployee)
}

func (controller *Controller) getEmployees(ctx *gin.Context) {
	service := &service.Service{Db: controller.Db, Employee: nil}
	res, err := service.GetEmployee()
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}

func (controller *Controller) deleteEmployee(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	service := &service.Service{Db: controller.Db, Employee: nil}
	err = service.DeleteEmployee(uint(id))
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, err)
}

func (controller *Controller) createEmployee(ctx *gin.Context) {
	var employee model.Employee

	err1 := ctx.ShouldBindJSON(&employee)
	if err1 != nil {
		ctx.JSON(400, err1)
		return
	}
	employee.Country = strings.ToLower(strings.TrimSpace(employee.Country))
	service := &service.Service{Db: controller.Db, Employee: &employee}
	res, err := service.CreateEmployee()
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}

func (controller *Controller) getEmployeeById(ctx *gin.Context) {

	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	service := &service.Service{Db: controller.Db, Employee: nil}
	res, err := service.GetEmployeeById(uint(id))
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}
