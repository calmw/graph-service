package api

import (
	"github.com/gin-gonic/gin"
	"graph-service/api/middlewares"
	"graph-service/api/routes"
	"graph-service/api/validate"
	"graph-service/db"
)

func Start() {

	//init mysql
	db.InitMysql()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":8000")

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
