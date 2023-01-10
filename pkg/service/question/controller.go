package question

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

type Handler struct {
	service IService
}

func NewHanlder(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

// GetQuestions godoc
// @Summary get list questions
// @Schemes
// @Description get list questions with pagination
// @Tags questions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query number false "max number of questions per page"
// @Param offset query number false "page offset"
// @Success 200 {object} model.ListQuestionResponse
// @Failure 400
// @Router /questions [get]
func (h *Handler) GetQuestions(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	data, err := h.service.GetList(int64(limit), int64(offset))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Create godoc
// @Summary create question
// @Schemes
// @Description create question
// @Tags questions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param question body model.CreateQuestionInput true "request body"
// @Success 201 {object} model.Question
// @Failure 400
// @Router /questions [post]
func (h *Handler) Create(c *gin.Context) {
	// request body
	var createQuestionInput model.CreateQuestionInput
	err := common.ValidateBodyData(c, &createQuestionInput)
	// error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	// create question
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	createQuestionInput.OwnerId = ownerId
	question, err := h.service.Create(createQuestionInput)

	// create failed
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, question)
}
