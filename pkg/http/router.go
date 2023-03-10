package http

import (
	status "net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/middleware"
	"tchh.lucpham/pkg/module/question"
	"tchh.lucpham/pkg/module/quiz"
	"tchh.lucpham/pkg/module/user"
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
	quizHandler := quiz.NewHanlder(*quiz.ServiceInstance)
	questionHandler := question.NewHanlder(question.ServiceInstance)

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

	// quiz
	authRouter.GET("/quizzes/:id", quizHandler.GetQuiz)
	authRouter.GET("/quizzes/:id/questions", quizHandler.GetQuestions)
	authRouter.GET("/quizzes", middleware.ValidateLimitOffset, quizHandler.GetQuizzes)
	authRouter.POST("/quizzes", quizHandler.CreateQuiz)
	authRouter.PATCH("/quizzes/:id/insert-question", quizHandler.InsertQuestion)
	authRouter.PATCH("/quizzes/:id/publish", quizHandler.PublishQuiz)
	authRouter.PATCH("/quizzes/:id/update-setting", quizHandler.UpdateQuizSetting)
	authRouter.DELETE("/quizzes/:id/remove-question/:questionId", quizHandler.RemoveQuestion)
	authRouter.DELETE("/quizzes/:id", quizHandler.DeleteQuiz)

	// questions
	authRouter.GET("/questions", middleware.ValidateLimitOffset, questionHandler.GetQuestions)
	authRouter.POST("/questions", questionHandler.Create)
	authRouter.DELETE("/questions/:id", questionHandler.Delete)
	authRouter.PATCH("/questions/:id", questionHandler.Update)

	return router
}
