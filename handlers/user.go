package handlers

import (
	"go_project/db"
	"go_project/logger"
	"go_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary List all users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := db.Connection().Find(&users).Error; err != nil {
		logger.Log.Error().Err(err).Msg("Failed to fetch users")
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    "Undefined",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// AddUser godoc
// @Summary Add a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserInput true "User info"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func AddUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Log.Error().Err(err).Msg("Invalid input")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    "Undefined",
			Message: err.Error(),
		})
		return
	}

	// Map input to User model
	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := db.Connection().Create(&user).Error; err != nil {
		logger.Log.Error().Err(err).Msg("Failed to add user")
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:    "Undefined",
			Message: err.Error(),
		})
		return
	}

	// Return full user including ID
	c.JSON(http.StatusCreated, user)
}
