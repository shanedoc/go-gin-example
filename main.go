package main

import (
	"fmt"
	"net/http"

	"github.com/shanedoc/go-gin-example/pkg/setting"
	"github.com/shanedoc/go-gin-example/routes"
)

func main() {
	router := routes.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%d", setting.HttPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
