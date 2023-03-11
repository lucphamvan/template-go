package db

import "os"

const (
	DATABASE                  = "Kun"
	USER_COLLECTION           = "users"
	QUESTION_COLLECTION       = "questions"
	QUIZ_COLLECTION           = "quiz"
	ACCEPTED_EMAIL_COLLECTION = "emails"
)

var MONGO_URL = os.Getenv("MONGO_URL")
