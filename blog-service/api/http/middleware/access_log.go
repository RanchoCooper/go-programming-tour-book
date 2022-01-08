package middleware

import (
    "bytes"
    "time"

    "github.com/gin-gonic/gin"

    "blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2022/1/8
 */

type AccessLogWrite struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (alw AccessLogWrite) Write(p []byte) (int, error) {
    if n, err := alw.body.Write(p); err != nil {
        return n, err
    }

    return alw.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
    return func(c *gin.Context) {
        bodyWriter := &AccessLogWrite{
            ResponseWriter: c.Writer,
            body:           bytes.NewBufferString(""),
        }
        c.Writer = bodyWriter

        beginTime := time.Now().Unix()
        c.Next()
        endTime := time.Now().Unix()

        fields := logger.Fields{
            "request":  c.Request.PostForm.Encode(),
            "response": bodyWriter.body.String(),
        }
        s := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"
        logger.Log.WithFields(fields).Infof(c, s, c.Request.Method, bodyWriter.Status(), beginTime, endTime)
    }
}
