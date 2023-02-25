package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)
	return
}

// Logger 日志记录
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		//记录请求包
		buf, _ := ioutil.ReadAll(request.Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
		request.Body = rdr //rewrite

		log.WithFields(log.Fields{
			"path":   request.URL.Path,
			"param":  request.URL.RawQuery,
			"method": request.Method,
		}).Infof("receive request body:%s ", string(buf))

		//记录返回包
		wc := &ResponseWithRecorder{
			ResponseWriter: writer,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		handler.ServeHTTP(writer, request)

		defer func() { //日志记录扫尾工作
			log.WithFields(log.Fields{
				"path":   request.URL.Path,
				"param":  request.URL.RawQuery,
				"method": request.Method,
			}).Infof("done use time:%s with status:%v", time.Since(start).String(), wc.statusCode)
		}()
	})
}
