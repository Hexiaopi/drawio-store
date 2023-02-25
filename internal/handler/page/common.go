package page

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
)

func Err(writer http.ResponseWriter, request *http.Request) {
	val := request.URL.Query()
	msg := val.Get("msg")
	generateHTML(writer, msg, "layout", "navbar", "error")
}

// 生成前段代码
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	funcMap := template.FuncMap{"fdate": formatDate, "fbyte": formatByte}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		log.Println(err)
		return
	}
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, err error) {
	url := []string{"/err?msg=", err.Error()}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 其他重定向
func redirectPage(writer http.ResponseWriter, request *http.Request, err error, page string) {
	url := []string{"/", page}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化
func formatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func formatByte(size int64) string {
	return humanize.Bytes(uint64(size))
}
