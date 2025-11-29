package handler

import (
	"geteway-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	deps := InitDependencies()

	// 🔹 Auth routes (JWTsiz)
	auth := r.Group("/auth")
	{
		auth.POST("/login", deps.AuthHandler.Login)
		auth.POST("/refresh", deps.AuthHandler.Refresh)
	}

	// 🔹 Admin
	admin := r.Group("/admin")
	admin.POST("/create", deps.AdminHandler.Create)
	admin.Use(middleware.JWTMiddleware())
	{
		admin.GET("/list", deps.AdminHandler.GetAll)
		admin.GET("/:id", deps.AdminHandler.GetByID)
		admin.PUT("/update/:id", deps.AdminHandler.Update)
		admin.DELETE("/delete/:id", deps.AdminHandler.Delete)
	}

	// 🔹 User
	user := r.Group("/user")
	user.Use(middleware.JWTMiddleware())
	{
		user.GET("/list", deps.UserHandler.GetAll)
		user.GET("/:id", deps.UserHandler.GetByID)
		user.POST("/create", deps.UserHandler.Create)
		user.PUT("/update/:id", deps.UserHandler.Update)
		user.DELETE("/delete/:id", deps.UserHandler.Delete)
	}

	// 🔹 Test (ochiq va JWT aralash)
	test := r.Group("/test")
	{
		test.POST("/generate", middleware.JWTMiddleware(), deps.TestHandler.GenerateTest)
		test.POST("/get", deps.TestHandler.GetTest)
	}

	// 🔹 Job
	job := r.Group("/job")
	{
		job.POST("/submit", deps.JobHandler.SubmitTest)
	}

	// 🔹 Filter
	filter := r.Group("/filter")
	filter.Use(middleware.JWTMiddleware())
	{
		filter.POST("/users", deps.FilterHandler.GetUsers)
	}

	// 🔹 SECTION CRUD
	section := r.Group("/section")
	section.Use(middleware.JWTMiddleware())
	{
		section.GET("/list", deps.TestSectionHandler.GetAllSections)
		section.GET("/:id", deps.TestSectionHandler.GetSectionById)
		section.POST("/create", deps.TestSectionHandler.CreateSection)
		section.PUT("/update/:id", deps.TestSectionHandler.UpdateSection)
		section.DELETE("/delete/:id", deps.TestSectionHandler.DeleteSection)
	}

	// 🔹 QUESTION CRUD
	question := r.Group("/question")
	question.Use(middleware.JWTMiddleware())
	{
		question.GET("/list", deps.TestSectionHandler.GetAllQuestions)
		question.GET("/:id", deps.TestSectionHandler.GetQuestionById)
		question.POST("/create", deps.TestSectionHandler.CreateQuestion)
		question.PUT("/update/:id", deps.TestSectionHandler.UpdateQuestion)
		question.DELETE("/delete/:id", deps.TestSectionHandler.DeleteQuestion)
	}

	// 🔹 OPTION CRUD
	option := r.Group("/option")
	option.Use(middleware.JWTMiddleware())
	{
		option.GET("/list", deps.TestSectionHandler.GetAllOptions)
		option.GET("/:id", deps.TestSectionHandler.GetOptionById)
		option.POST("/create", deps.TestSectionHandler.CreateOption)
		option.PUT("/update/:id", deps.TestSectionHandler.UpdateOption)
		option.DELETE("/delete/:id", deps.TestSectionHandler.DeleteOption)
	}

	// 🔹 EXAM CRUD
exam := r.Group("/exam")
exam.Use(middleware.JWTMiddleware())
{
	exam.GET("/list", deps.ExamHandler.GetAll)       
	exam.GET("/:id", deps.ExamHandler.GetByID)        
	exam.POST("/create", deps.ExamHandler.Create)     
	exam.PUT("/update/:id", deps.ExamHandler.Update) 
	exam.DELETE("/delete/:id", deps.ExamHandler.Delete) 
}



}
