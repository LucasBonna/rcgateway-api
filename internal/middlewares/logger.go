package middlewares

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"web/gin/internal/database"
	"web/gin/internal/database/models"

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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Erro ao ler o corpo da requisição")
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		requestParams := formatParams(c.Params)

		requestBody := string(bodyBytes)
		requestHeaders := formatHeaders(c.Request.Header)
		requestQuery := c.Request.URL.RawQuery

		logger := models.Log{
			ID:             uuid.New(),
			Ip:             c.ClientIP(),
			Method:         c.Request.Method,
			Endpoint:       c.Request.URL.Path,
			RequestBody:    &requestBody,
			RequestHeaders: &requestHeaders,
			RequestQuery:   &requestQuery,
			RequestParams:  &requestParams,
		}

		if err := database.Db.Create(&logger).Error; err != nil {
			log.Println("Erro ao criar Log")
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		latency := time.Since(t)
		logger.ResponseTime = latency.String()
		responseBody := blw.body.String()
		responseHeaders := formatHeaders(c.Writer.Header())
		statusCode := c.Writer.Status()

		logger.ResponseBody = &responseBody
		logger.ResponseHeaders = &responseHeaders
		logger.StatusCode = statusCode

		if err := database.Db.Model(&logger).Updates(logger).Error; err != nil {
			log.Println("Erro ao atualizar Log")
		}
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
