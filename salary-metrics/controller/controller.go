package controller

import (
	"strings"
	servicego "sukvij/salary-metrics/service.go"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func SalaryMetricsController(app *gin.Engine, db *gorm.DB) {
	controller := &Controller{Db: db}
	router := app.Group("/v1/salary-metrics")
	router.GET("country/:country", controller.salaryMetricsCountryWise)
	router.GET("job_title/:title", controller.salaryMetricsJobTitleWise)
}

func (controller *Controller) salaryMetricsCountryWise(ctx *gin.Context) {

	country := ctx.Param("country")
	country = strings.Trim(country, " ")

	service := &servicego.Service{Db: controller.Db}
	res, err := service.SalaryMetricsCountryWise(country)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}

func (controller *Controller) salaryMetricsJobTitleWise(ctx *gin.Context) {

	title := ctx.Param("title")
	title = strings.Trim(title, " ")

	service := &servicego.Service{Db: controller.Db}
	res, err := service.SalaryMetricsJobTitleWise(title)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, res)
}
