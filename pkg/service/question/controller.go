package question

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
