package handler

import (
	"log"

	"geteway-service/internal/client"
	"geteway-service/internal/service"
)

type Dependencies struct {
	AuthHandler        *AuthHandler
	AdminHandler       *AdminHandler
	UserHandler        *UserHandler
	TestHandler        *TestHandler
	JobHandler         *JobHandler
	FilterHandler      *FilterHandler
	TestSectionHandler *TestSectionHandler
	ExamHandler        *ExamHandler
}

func InitDependencies() *Dependencies {
	var authHandler *AuthHandler
	if authClient, err := client.NewAuthClient(); err != nil {
		log.Printf("[Dependencies] AuthClient ulanishda xatolik: %v", err)
	} else {
		authService := service.NewAuthService(authClient)
		authHandler = NewAuthHandler(authService)
	}

	var adminHandler *AdminHandler
	if adminService, err := service.NewAdminService(); err != nil {
		log.Printf("[Dependencies] AdminService ulanishda xatolik: %v", err)
	} else {
		adminHandler = NewAdminHandler(adminService)
	}

	var userHandler *UserHandler
	if userService, err := service.NewUserService(); err != nil {
		log.Printf("[Dependencies] UserService ulanishda xatolik: %v", err)
	} else {
		userHandler = NewUserHandler(userService)
	}

	var testHandler *TestHandler
	if generateClient, err := client.NewGenerateClient(); err != nil {
		log.Printf("[Dependencies] GenerateClient ulanishda xatolik: %v", err)
	} else {
		generateService := service.NewGenerateService(generateClient)
		testHandler = NewGenerateHandler(generateService)
	}

	var jobHandler *JobHandler
	if jobClient, err := client.NewJobClient(); err != nil {
		log.Printf("[Dependencies] JobClient ulanishda xatolik: %v", err)
	} else {
		jobService := service.NewJobService(jobClient)
		jobHandler = NewJobHandler(jobService)
	}

	var filterHandler *FilterHandler
	if filterClient, err := client.NewFilterClient(); err != nil {
		log.Printf("[Dependencies] FilterClient ulanishda xatolik: %v", err)
	} else {
		filterService := service.NewFilterService(filterClient)
		filterHandler = NewFilterHandler(filterService)
	}

	var testSectionHandler *TestSectionHandler
	if testSectionClient, err := client.NewTestClient(); err != nil {
		log.Printf("[Dependencies] TestSectionClient ulanishda xatolik: %v", err)
	} else {
		testSectionService := service.NewTestService(testSectionClient)
		testSectionHandler = NewTestSectionHandler(testSectionService)
	}

	// ✅ ExamHandler
var examHandler *ExamHandler
if examService, err := service.NewExamService(); err != nil {
    log.Printf("[Dependencies] ExamService ulanishda xatolik: %v", err)
} else {
    examHandler = NewExamHandler(examService)
}
	return &Dependencies{
		AuthHandler:        authHandler,
		AdminHandler:       adminHandler,
		UserHandler:        userHandler,
		TestHandler:        testHandler,
		JobHandler:         jobHandler,
		FilterHandler:      filterHandler,
		TestSectionHandler: testSectionHandler,
		ExamHandler:        examHandler,
	}
}
