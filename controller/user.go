package controller

import (
	"github.com/labstack/echo/v4"
	"golang-rest-api-validation-example/model"
	"golang-rest-api-validation-example/repository"
	"golang-rest-api-validation-example/util"
	"net/http"
)

type UserController struct {
	userRepository *repository.UserRepository
}

func NewUserController() *UserController {
	return &UserController{userRepository: repository.NewUserRepository()}
}

// GetAllUser godoc
// @Summary Get all users
// @Description Get all user items
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Success 200 {array} model.User
// @Failure 500 {object} handler.APIError
// @Router /users [get]
func (userController *UserController) GetAllUser(c echo.Context) error {
	users := userController.userRepository.GetAllUser()
	return util.Negotiate(c, http.StatusOK, users)
}

// SaveUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param user body model.User true "New User"
// @Success 201 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users [post]
func (userController *UserController) SaveUser(c echo.Context) error {
	user := new(model.User)
	if err := util.BindAndValidate(c, user); err != nil {
		return err
	}

	createdUser := userController.userRepository.SaveUser(user)

	return util.Negotiate(c, http.StatusCreated, createdUser)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [get]
func (userController *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := userController.userRepository.GetUser(id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Param user body model.User true "User Info"
// @Success 200 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	user := new(model.User)

	if err := util.BindAndValidate(c, user); err != nil {
		return err
	}

	user, err := userController.userRepository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 204 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [delete]
func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userController.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
