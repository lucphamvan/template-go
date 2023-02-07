package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

type Handler struct {
	Service Service
}

func NewHanlder(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// CreateQuizzes godoc
// @Summary create quizzes
// @Schemes
// @Description create quizzes
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param quiz body model.CreateQuizzesInput true "request body"
// @Success 200 {object} model.Quizzes
// @Failure 400
// @Router /quizzes [post]
func (h *Handler) CreateQuizzes(c *gin.Context) {
	var createQuizzesInput model.CreateQuizzesInput
	err := common.ValidateBodyData(c, &createQuizzesInput)
	// user id to createQuizInput
	createQuizzesInput.OwnerId = c.Request.Header.Get(common.USER_ID_HEADER)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	quizzes, err := h.Service.CreateQuizzes(createQuizzesInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}
