package page

import (
	"net/http"

	"github.com/hexiaopi/drawio-store/internal/service"
	log "github.com/sirupsen/logrus"
)

func GetDraw(writer http.ResponseWriter, _ *http.Request) {
	generateHTML(writer, nil, "layout", "navbar", "draw")
}

func QueryDraw(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Errorf("parse form error:%v", err)
		errorMessage(writer, request, err)
		return
	}
	images, err := service.ListDrawImages()
	if err != nil {
		errorMessage(writer, request, err)
		return
	}
	generateHTML(writer, images, "layout", "navbar", "draw")
}
