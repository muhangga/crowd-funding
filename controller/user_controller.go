package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhangga/helper"
	"github.com/muhangga/model/response"

	model "github.com/muhangga/model/request"
	user "github.com/muhangga/service/user"
)

type userController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var userRequest model.RegisterRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.RegisterUser(userRequest)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken(user)

	userResponse := response.ResponseUser(user, "aweuaweu")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", userResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var loginRequest model.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedin, err := h.userService.Login(loginRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginResponse := response.ResponseUser(loggedin, "aweuaweu")
	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", loginResponse)
	c.JSON(http.StatusOK, response)
}

func(h *userController) CheckEmailAvaible(c *gin.Context) {
	var emailRequest model.CheckEmailRequest

	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email not found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(emailRequest)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMesasge := "Email has been registered"
	if IsEmailAvailable {
		metaMesasge = "Email is available"
	}

	response := helper.APIResponse(metaMesasge, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}