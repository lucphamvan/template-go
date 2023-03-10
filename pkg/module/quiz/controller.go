package quiz

import (
	"net/http"
	"strconv"

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
// @Param quiz body model.QuizSetting true "request body"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes [post]
func (h *Handler) CreateQuiz(c *gin.Context) {
	var quizSetting model.QuizSetting
	// validate request body
	err := common.ValidateBodyData(c, &quizSetting)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	var createQuizInput model.CreateQuizInput
	// insert owner id
	createQuizInput.OwnerId = c.Request.Header.Get(common.USER_ID_HEADER)
	createQuizInput.Setting = quizSetting
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
// @Success 200 {object} model.CreateAndInsertQuestionToQuizResponse
// @Failure 400
// @Router /quizzes/{id}/insert-question [patch]
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
	data, err := h.service.CreateAndInsertQuestionToQuiz(quizId, createQuestionInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
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
// @Router /quizzes/{id}/remove-question/{questionId} [delete]
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

// GetQuiz godoc
// @Summary get quizzes
// @Schemes
// @Description get quizzes
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.GetListQuizzesResponse
// @Failure 400
// @Router /quizzes [get]
func (h *Handler) GetQuizzes(c *gin.Context) {
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	// get quizzes
	data, err := h.service.GetQuizzes(ownerId, int64(limit), int64(offset))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetQuestions godoc
// @Summary get question of quiz
// @Schemes
// @Description get question of quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Success 200 {object} []model.Question
// @Failure 400
// @Router /quizzes/{id}/questions [get]
func (h *Handler) GetQuestions(c *gin.Context) {
	// ownder id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	quizId := c.Param("id")

	// get questions of quiz
	questions, err := h.service.GetQuestions(quizId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

// PublishQuiz godoc
// @Summary publish quiz
// @Schemes
// @Description publish quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id}/publish [patch]
func (h *Handler) PublishQuiz(c *gin.Context) {
	// quiz id
	quizId := c.Param("id")
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	// publish quiz
	quiz, err := h.service.PublishQuiz(quizId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

// GetQuiz godoc
// @Summary get quiz
// @Schemes
// @Description get quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id} [get]
func (h *Handler) GetQuiz(c *gin.Context) {
	// quiz id
	quizId := c.Param("id")
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	// get quiz
	quiz, err := h.service.GetQuiz(quizId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

// UpdateQuizSetting godoc
// @Summary update quiz setting
// @Schemes
// @Description update quiz setting
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Param quizSetting body model.QuizSetting true "quiz setting"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id}/update-setting [patch]
func (h *Handler) UpdateQuizSetting(c *gin.Context) {
	// quiz id
	quizId := c.Param("id")
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	// update quiz setting
	var quizSetting model.QuizSetting
	err := c.ShouldBindJSON(&quizSetting)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	quiz, err := h.service.UpdateQuizSetting(quizId, ownerId, quizSetting)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

// DeleteQuiz godoc
// @Summary delete quiz
// @Schemes
// @Description delete quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "quiz id"
// @Success 200 {object} model.Quiz
// @Failure 400
// @Router /quizzes/{id} [delete]
func (h *Handler) DeleteQuiz(c *gin.Context) {
	// quiz id
	quizId := c.Param("id")
	// owner id
	ownerId := c.Request.Header.Get(common.USER_ID_HEADER)
	// delete quiz
	quiz, err := h.service.DeleteQuiz(quizId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}
