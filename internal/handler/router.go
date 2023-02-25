package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hexiaopi/drawio-store/internal/handler/api"
	"github.com/hexiaopi/drawio-store/internal/handler/page"
	"github.com/hexiaopi/drawio-store/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	//静态文件
	assets := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.Logger)
	apiRouter.HandleFunc("/v1/image/{name}", api.GetImage).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/image/{name}", api.CreateImage).Methods(http.MethodPost)
	apiRouter.HandleFunc("/v1/image/{name}", api.DeleteImage).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/v1/draw/{name}", api.GetDraw).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/draw", api.PostDraw).Methods(http.MethodPost)

	router.Use(middleware.Logger)
	router.HandleFunc("/", page.QueryDraw).Methods(http.MethodGet,http.MethodPost)

	return router
}
