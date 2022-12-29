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

// CreateQuiz godoc
// @Summary create quiz
// @Schemes
// @Description create quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param quiz body model.CreateQuizInput true "request body"
// @Success 200
// @Failure 400
// @Router /quizzes [post]
func (h *Handler) CreateQuiz(c *gin.Context) {
	var createQuizInput model.CreateQuizInput
	err := common.ValidateBodyData(c, &createQuizInput)
	// user id to createQuizInput
	createQuizInput.OwnerId = c.Request.Header.Get(common.USER_ID_HEADER)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	err = h.Service.CreateQuiz(createQuizInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": true})
}
