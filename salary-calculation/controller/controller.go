package controller

import (
	"strconv"
	"sukvij/salary-calculation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func SalaryCalculationController(app *gin.Engine, db *gorm.DB) {
	controller := &Controller{Db: db}
	router := app.Group("/v1/salary-calculation")
	router.GET("/:id", controller.salaryCalculation)
}

func (controller *Controller) salaryCalculation(ctx *gin.Context) {

	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	service := &service.Service{Db: controller.Db}
	res, err := service.SalaryCalculation(uint(id))
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}
