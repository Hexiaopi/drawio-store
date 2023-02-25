package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/drawio-store/internal/entity"
	"github.com/hexiaopi/drawio-store/internal/service"
)

type ImageResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func GetDraw(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	image, err := service.GetDrawImage(name)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(err.Error()))
		return
	}
	data, _ := json.Marshal(ImageResponse{
		Status: 200,
		Data:   "data:image/png;base64," + base64.StdEncoding.EncodeToString(image),
	})
	writer.Header().Set("content-type", "application/json")
	writer.Write(data)
}

func PostDraw(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}
	defer request.Body.Close()
	var image entity.DrawImage
	if err := json.Unmarshal(data, &image); err != nil {
		fmt.Fprintln(writer, err)
		return
	}
	if err := service.UpdateDraw(image); err != nil {
		log.Errorf("update draw err:%v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
