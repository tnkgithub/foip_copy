package main

import (
	"foip/chat/pkg/chat"
	"foip/chat/pkg/log"
	"os"

	handler "foip/chat/pkg/application/handlers"

	"foip/chat/config"

	"github.com/gin-gonic/gin"
)

func main() {

	service := chat.New()
	go service.Run()

	app := gin.Default()

	//create new handler
	ctrl := handler.New(service)
	app.GET("/api/v1/chat", ctrl.Handler)

	app.SetTrustedProxies(config.TrustedProxies)

	log.Call.Infof("Listening and serving HTTP on :%s", os.Getenv("PORT"))
	if err := app.Run(); err != nil {
		log.Call.Fatal(err)
	}
}
