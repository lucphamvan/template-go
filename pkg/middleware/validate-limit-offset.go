package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

func ValidateLimitOffset(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	// return error if one of limit/offset is empty and another is not
	if (limitStr != "" && offsetStr == "") || (limitStr == "" && offsetStr != "") {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: common.ERROR_REQUIRE_LIMIT_OFFSET})
		return
	}

	// return error if limit/offset is not empty and not number
	if limitStr != "" && offsetStr != "" {
		// return error if limit is 0
		if limitStr == "0" {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: common.ERROR_QUERY_LIMIT_ZERO})
			return
		}
		_, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: common.ERROR_QUERY_LIMIT})
			return
		}

		_, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.Error{Error: common.ERROR_QUERY_OFFSET})
			return
		}
	}

	c.Next()
}
