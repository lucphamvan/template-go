package common

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	ACC_SECRET_KEY = "0BY9n_Qs272miuuR7wum8E"
	REF_SECRET_KEY = "1CZ9n_Qs272hiuuR3w5m4O"
	ACC_DURATION   = 15 * time.Minute
	REF_DURATION   = 30 * 24 * time.Hour
)

// claim struct ..
type Claims struct {
	UID string `json:"uid,omitempty"`
	jwt.RegisteredClaims
}

func GeneratePassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(pass, hashpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	return err == nil
}

func ValidateBodyData(c *gin.Context, data interface{}) error {
	err := c.BindJSON(data)
	if err != nil {
		return errors.New(ERROR_BIND_JSON)
	}

	err = validator.New().Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func generateToken(id string, duration time.Duration, secret string) (string, error) {
	expiredTime := time.Now().Add(duration)
	claims := &Claims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtClaims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(token string, secret string) (*Claims, error) {
	claims := &Claims{}
	jwtClaim, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtClaim.Valid {
		return nil, err
	}

	return claims, nil
}

func GenerateAccToken(id string) (string, error) {
	return generateToken(id, ACC_DURATION, ACC_SECRET_KEY)
}

func GenerateRefToken(id string) (string, error) {
	return generateToken(id, REF_DURATION, REF_SECRET_KEY)
}

func VerifyAccToken(token string) (*Claims, error) {
	return verifyToken(token, ACC_SECRET_KEY)
}

func VerifyRefToken(token string) (*Claims, error) {
	return verifyToken(token, REF_SECRET_KEY)
}
