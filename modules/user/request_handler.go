package user

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerUser struct {
	ctr ControllerUser
}

func NewUserRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerUser {
	return RequestHandlerUser{
		ctr: controllerUser{
			userUseCase: useCaseUser{
				userRepo: repository.NewUser(dbCrud),
			},
		}}
}

// CreateUser ... Create User
// @Summary Create User based on parameter
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body user.UserParam true "User Data"
// @Success 200 {object} user.UserParam
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /user/ [post]
func (h RequestHandlerUser) CreateUser(c *gin.Context) {

	request := UserParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUserById ... user by id
// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path int true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} user.FindUser
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /user/{id} [get]
func (h RequestHandlerUser) GetUserById(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteUser ... delete User by email
// @Summary Delete User
// @Description Delete User by email
// @Tags Users
// @Param email path string true "User Email"
// @Success 200 {object} dto.ResponseMeta
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /user/{email} [delete]
func (h RequestHandlerUser) DeleteUser(c *gin.Context) {
	email := c.Param("email")
	res, err := h.ctr.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateUser ... update user
// @Summary Create User based on parameter
// @Description Create new user
// @Tags Users
// @Accept json
// @Param id path int true "User ID"
// @Param user body user.UserParam true "User Data"
// @Success 200 {object} user.UserParam
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /user/{id} [put]
func (h RequestHandlerUser) UpdateUser(c *gin.Context) {
	request := UserParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateUser(request, uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllUser get all user
func (h RequestHandlerUser) GetAllUser(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	usernameStr := c.DefaultQuery("name", "")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetAllUser(uint(page), usernameStr)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}
