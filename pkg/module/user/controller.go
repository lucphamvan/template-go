package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) Handler {
	return Handler{
		service: service,
	}
}

// CreateUser godoc
// @Summary create user
// @Schemes
// @Description create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserInput true "request body"
// @Success 200 {object} model.User
// @Failure 400
// @Router /users [post]
func (h *Handler) Create(c *gin.Context) {
	// bind request body
	var createUserInput model.CreateUserInput
	err := common.ValidateBodyData(c, &createUserInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	// create user
	user, err := h.service.Create(createUserInput)
	if err != nil {
		log.Default().Println("Failed to create user : ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// CheckExitedUser godoc
// @Summary check user exist or not
// @Schemes
// @Description check user exist or not
// @Tags users
// @Accept json
// @Produce json
// @Param email query string true "user email"
// @Success 200 {object} boolean
// @Failure 400
// @Router /users/check [get]
func (h *Handler) CheckExitedUser(c *gin.Context) {
	email := c.Query("email")
	result := h.service.IsUserExist(email)
	c.JSON(http.StatusOK, result)
}

// GetList godoc
// @Summary get list user
// @Schemes
// @Description get list user with pagination
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query number false "max number of user per page"
// @Param offset query number false "page offset"
// @Success 200 {object} model.ListUserResponse
// @Failure 400
// @Router /users [get]
func (h *Handler) GetList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := h.service.GetList(int64(limit), int64(offset))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Get godoc
// @Summary get user by id
// @Schemes
// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "user id"
// @Success 200 {object} model.User
// @Failure 400
// @Router /users/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Get godoc
// @Summary update user
// @Schemes
// @Description update user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "user id"
// @Param user body model.UpdateUserInput true "body data"
// @Success 200 {object} model.User
// @Failure 400
// @Router /users/{id} [patch]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var updateUserInput model.UpdateUserInput
	err := common.ValidateBodyData(c, updateUserInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	user, err := h.service.Update(id, updateUserInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Login godoc
// @Summary Login to system
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Security basicAuth
// @Success 200 {object} model.AuthenResponse
// @Failure 400
// @Failure 401
// @Router /users/login [get]
func (h *Handler) Login(c *gin.Context) {
	email, password, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: "Required email/password in basic authorization"})
		return
	}

	data, err := h.service.Login(email, password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)

}

// GetAccessInfo godoc
// @Summary Get access's information
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.User
// @Failure 401
// @Router /users/me [get]
func (h *Handler) GetAccessInfo(c *gin.Context) {
	id := c.Request.Header.Get(common.USER_ID_HEADER)
	user, err := h.service.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// RefreshToken godoc
// @Summary refresh token
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param token body model.RefreshTokenRequest true "token"
// @Success 200 {object} string
// @Failure 401
// @Router /token/refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	var data model.RefreshTokenRequest
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}

	claim, err := common.VerifyRefToken(data.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}
	token, err := common.GenerateAccToken(claim.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
