package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"rc/gateway/internal/database"
	"rc/gateway/internal/database/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func isJSONContentType(contentType string) bool {
	return strings.HasPrefix(contentType, "application/json")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/swagger") || c.Request.URL.Path == "/docs" {
			c.Next()
			return
		}

		t := time.Now()

		requestHeaders := formatHeaders(c.Request.Header)
		requestQuery := c.Request.URL.RawQuery
		requestParams := formatParams(c.Params)

		logger := models.Log{
			ID:             uuid.New(),
			Ip:             c.ClientIP(),
			Method:         c.Request.Method,
			Endpoint:       c.Request.URL.Path,
			RequestHeaders: &requestHeaders,
			RequestQuery:   &requestQuery,
			RequestParams:  &requestParams,
		}

		contentType := c.GetHeader("Content-Type")
		if strings.HasPrefix(contentType, "multipart/form-data") {
			fileUploadMsg := "Multipart form data (file upload)"
			logger.RequestBody = &fileUploadMsg
			nonJSONMsg := fmt.Sprintf("Non-JSON content type: %s", contentType)
			logger.RequestBody = &nonJSONMsg
			logger.RequestHeaders = &nonJSONMsg
			logger.ResponseHeaders = &nonJSONMsg
			logger.ResponseBody = &nonJSONMsg
		} else if !isJSONContentType(contentType) {
			nonJSONMsg := fmt.Sprintf("Non-JSON content type: %s", contentType)
			logger.RequestBody = &nonJSONMsg
			logger.RequestHeaders = &nonJSONMsg
			logger.ResponseHeaders = &nonJSONMsg
			logger.ResponseBody = &nonJSONMsg
		} else {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		database.Db.Create(&logger)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		responseTime := time.Since(t).String()
		responseBody := blw.body.String()
		responseHeaders := formatHeaders(c.Writer.Header())

		logger.ResponseTime = responseTime
		logger.ResponseBody = &responseBody
		logger.ResponseHeaders = &responseHeaders
		logger.StatusCode = c.Writer.Status()

		database.Db.Model(&logger).Updates(logger)
	}
}

func formatParams(params gin.Params) string {
	var buffer bytes.Buffer
	for _, p := range params {
		buffer.WriteString(p.Key)
		buffer.WriteString("=")
		buffer.WriteString(p.Value)
		buffer.WriteString("&")
	}
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 1)
	}
	return buffer.String()
}

func formatHeaders(headers map[string][]string) string {
	var buffer bytes.Buffer
	for key, values := range headers {
		buffer.WriteString(key)
		buffer.WriteString(": ")
		buffer.WriteString(strings.Join(values, ", "))
		buffer.WriteString("\n")
	}
	return buffer.String()
}
