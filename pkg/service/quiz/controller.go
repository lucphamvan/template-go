package quiz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

type Handler struct {
	service Service
}

func NewHanlder(service Service) *Handler {
	return &Handler{
		service: service,
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
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes [post]
func (h *Handler) CreateQuiz(c *gin.Context) {
	var createQuizInput model.CreateQuizInput
	// validate request body
	err := common.ValidateBodyData(c, &createQuizInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	// insert owner id
	createQuizInput.OwnerId = c.Request.Header.Get(common.USER_ID_HEADER)
	// create quiz
	quiz, err := h.service.CreateQuiz(createQuizInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// InsertQuestion godoc
// @Summary create and insert question to quiz
// @Schemes
// @Description create and insert question to quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Param question body model.CreateQuestionInput true "request body"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id}/add-question [patch]
func (h *Handler) InsertQuestion(c *gin.Context) {
	// validate request body
	var createQuestionInput model.CreateQuestionInput
	err := common.ValidateBodyData(c, &createQuestionInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	// insert owner id
	createQuestionInput.OwnerId = c.Request.Header.Get(common.USER_ID_HEADER)
	createQuestionInput.Deleted = false

	// create and insert question to quiz
	quizId := c.Param("id")
	quiz, err := h.service.CreateAndInsertQuestionToQuiz(quizId, createQuestionInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// RemoveQuestion godoc
// @Summary remove question from quiz
// @Schemes
// @Description remove question from quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Param questionId path string true "question id"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id}/remove-question/{questionId} [patch]
func (h *Handler) RemoveQuestion(c *gin.Context) {
	// quiz id
	quizId := c.Param("id")
	// question id
	questionId := c.Param("questionId")
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	// remove question from quiz
	quiz, err := h.service.RemoveQuestionFromQuiz(quizId, questionId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}
