package handler

import (
	"hacktiv8-msib-final-project-3/database"
	"hacktiv8-msib-final-project-3/handler/httphandler"
	"hacktiv8-msib-final-project-3/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-3/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	db := database.GetPostgresInstance()

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	userRepo := userpg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := httphandler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users/update-account", authService.Authentication(), userHandler.UpdateUser)

	log.Fatalln(r.Run(":" + PORT))
}
