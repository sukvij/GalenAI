package controller

import (
	"strconv"
	"strings"
	"sukvij/employment/model"
	"sukvij/employment/service"
	response "sukvij/galenfers/Response"
	"sukvij/galenfers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func EmployeeController(app *gin.Engine, db *gorm.DB) {
	controller := &Controller{Db: db}
	router := app.Group("/v1/employees")
	router.Use(middleware.JWTAuthMiddleware())
	router.GET("", controller.getEmployees)
	router.POST("", controller.createEmployee)
	router.GET("/:id", controller.getEmployeeById)
	router.DELETE("/:id", controller.deleteEmployee)
}

func (controller *Controller) getEmployees(ctx *gin.Context) {
	service := &service.Service{Db: controller.Db, Employee: nil}
	res, err := service.GetEmployee()
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	response.SendResponse(ctx, res, err)
}

func (controller *Controller) deleteEmployee(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}
	service := &service.Service{Db: controller.Db, Employee: nil}
	err = service.DeleteEmployee(uint(id))
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	response.SendResponse(ctx, map[string]string{"result": "successfully deleted"}, nil)
}

func (controller *Controller) createEmployee(ctx *gin.Context) {
	var employee model.Employee

	err1 := ctx.ShouldBindJSON(&employee)
	if err1 != nil {
		response.SendResponse(ctx, nil, err1)
		return
	}
	employee.Country = strings.ToLower(strings.TrimSpace(employee.Country))
	service := &service.Service{Db: controller.Db, Employee: &employee}
	res, err := service.CreateEmployee()
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}
	response.SendResponse(ctx, res, nil)
}

func (controller *Controller) getEmployeeById(ctx *gin.Context) {

	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}
	service := &service.Service{Db: controller.Db, Employee: nil}
	res, err := service.GetEmployeeById(uint(id))
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	response.SendResponse(ctx, res, nil)
}
