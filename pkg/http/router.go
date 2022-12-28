package http

import (
	status "net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/middleware"
	"tchh.lucpham/pkg/service/user"
)

func healcheck(c *gin.Context) {
	c.JSON(status.StatusOK, gin.H{"message": "Ping !"})
}

func SetupRouter() *gin.Engine {
	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))

	userHandler := user.NewHandler(user.ServiceInstance)

	router.GET("/ping", healcheck)

	authRouter := router.Group("/")
	authRouter.Use(middleware.Authen)
	// user
	router.POST("/users", userHandler.Create)
	router.GET("/users/login", userHandler.Login)
	router.GET("/users/check", userHandler.CheckExitedUser)
	router.POST("token/refresh", userHandler.RefreshToken)

	authRouter.GET("/users/me", userHandler.GetAccessInfo)
	authRouter.GET("/users/:id", userHandler.Get)
	authRouter.GET("/users", middleware.ValidateLimitOffset, userHandler.GetList)
	authRouter.PATCH("/users/:id", userHandler.Update)

	return router
}