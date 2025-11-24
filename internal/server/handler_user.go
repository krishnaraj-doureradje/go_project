package server

import (
	"go_project/internal/logger"
	"go_project/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsersHandler godoc
// @Summary List all users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} user.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (s *Server) GetUsersHandler(c *gin.Context) {
	users, err := s.UserService.GetAll()
	if err != nil {
		logger.Log.Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: "Undefined", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// AddUserHandler godoc
// @Summary Add a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.CreateUserInput true "User info"
// @Success 201 {object} user.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (s *Server) AddUserHandler(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		logger.Log.Error().Msg(err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: "Undefined", Message: err.Error()})
		return
	}

	result, err := s.UserService.Create(u)
	if err != nil {
		logger.Log.Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: "Undefined", Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
