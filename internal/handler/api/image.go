package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/drawio-store/internal/service"
)

func GetImage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	image, err := service.GetDrawImage(name)
	if err != nil {
		log.Errorf("get image:%s err:%v", name, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("content-type", "image/png")
	writer.Write(image)
}

func CreateImage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	if err := service.CreateDrawImage(name); err != nil {
		log.Errorf("create image:%s err:%v", name, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func DeleteImage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	if err := service.DeleteDrawImage(name); err != nil {
		log.Errorf("delete image:%s err:%v", name, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
}
