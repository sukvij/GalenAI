package user

import (
	response "sukvij/galenfers/Response"
	"sukvij/galenfers/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func UserController(app *gin.Engine, db *gorm.DB) {
	router := app.Group("/v1/users")

	controller := &Controller{Db: db}
	router.POST("/register", controller.userRegistration)
	router.POST("/login", controller.login)
}

func (controller *Controller) userRegistration(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	user.Role = "user" // ensure only normal user can register
	err = CreateUser(&user, controller.Db)
	response.SendResponse(ctx, nil, err)
}

func (controller *Controller) login(ctx *gin.Context) {
	var credential Credentials
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	res, err := GetUserByUserName(&credential, controller.Db)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}
	// create token for this
	middleware.Login(ctx, res.UserName, res.Role)
}
