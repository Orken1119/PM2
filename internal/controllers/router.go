package controller

import (
	"github.com/Orken1119/PM2/internal/controllers/middleware"
	pkg "github.com/Orken1119/PM2/pkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/Orken1119/PM2/docs"


	"github.com/Orken1119/PM2/internal/controllers/auth"
	"github.com/Orken1119/PM2/internal/controllers/user"
	repository "github.com/Orken1119/PM2/internal/repositories"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

	router.Use(middleware.JWTAuth(`access-secret-key`))

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
	}

	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/profile", userController.GetProfile)
	}
}