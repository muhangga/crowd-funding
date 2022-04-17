package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/controller"

	// model "github.com/muhangga/model/request"
	repository "github.com/muhangga/repository/user"
	user "github.com/muhangga/service/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// User
	userRepository := repository.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/login", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvaible)
	api.POST("/avatars", userController.UploadAvatar)

	router.Run()
}