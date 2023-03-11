package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func MockGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(w)
	context.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return context
}

func MockRequestBody(c *gin.Context, content interface{}) {
	c.Request.Method = http.MethodPost
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}
