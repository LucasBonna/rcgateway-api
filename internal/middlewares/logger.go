package middlewares

import (
	"bytes"
	"io/ioutil"
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

        if strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
            fileUploadMsg := "Multipart form data (file upload)"
            logger.RequestBody = &fileUploadMsg
        } else {
            bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
            c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
            requestBody := string(bodyBytes)
            logger.RequestBody = &requestBody
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
