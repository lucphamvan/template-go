package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

const (
	InvalidTokenMsg = "Invalid Token"
)

func Authen(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	arr := strings.Split(token, "Bearer ")

	if len(arr) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: InvalidTokenMsg})
		return
	}
	claims, err := common.VerifyAccToken(arr[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Error{Error: err.Error()})
		return
	}

	c.Request.Header.Set(common.USER_ID_HEADER, claims.UID)
	c.Next()
}
