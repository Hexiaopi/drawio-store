package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hexiaopi/drawio-store/internal/handler"
	log "github.com/sirupsen/logrus"
)

// 日志打印初始化
func initLogrus(logLevel string) error {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}

func init() {
	if err := initLogrus("debug"); err != nil {
		panic(err)
	}
}

func main() {
	router := handler.RegisterRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		log.Infof("service start at port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start http server fail:%v", err)
		}
	}()

	killerChan := make(chan os.Signal)
	signal.Notify(killerChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-killerChan
	log.Infof("get killer signal. %v", sig)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Errorf("http server shut down server fail %v", err)
	}
}
