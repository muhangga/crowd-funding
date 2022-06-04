package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhangga/controller"
	"github.com/muhangga/helper"
	"github.com/muhangga/service/auth"
	"github.com/muhangga/service/user"

	campaignRepository "github.com/muhangga/repository/campaign"
	userRepository "github.com/muhangga/repository/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Repository
	userRepository := userRepository.NewUserRepository(db)
	campaignsRepository := campaignRepository.NewCampaignRepository(db)

	campaigns, err := campaignsRepository.FindByUserID(24)

	fmt.Println(campaigns)
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
	}

	// Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// Controller
	userController := controller.NewUserController(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/login", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvaible)
	api.POST("/avatars", authMiddleware(authService, userService), userController.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.AuthService, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
