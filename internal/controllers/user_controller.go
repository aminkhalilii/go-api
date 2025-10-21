package controllers

import (
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Users retrieved successfully", users)
}
func (uc *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())

		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}
	newUser, err := uc.userService.CreateUser(&user)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, "User created successfully", newUser)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	newUser, err := uc.userService.UpdateUser(id, &user)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, "User updated successfully", newUser)
}
func (uc *UserController) DeleteUser(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	err = uc.userService.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	utils.Success(c, http.StatusOK, "User updated successfully", nil)

}
