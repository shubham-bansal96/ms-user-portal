package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ms-user-portal/app/config"
	"github.com/ms-user-portal/app/database"
	"github.com/ms-user-portal/app/logging"
	"github.com/ms-user-portal/app/routes"
)

const listenPort = ":4376"

func main() {
	gin := gin.Default()

	config.Initialize()
	log.Println("config initialized successfully")

	logging.Initialize(config.Config)
	lw := logging.LogForFunc()
	lw.Debug("log initialized successfully")

	database.Initialize(config.Config)

	routes.Initialize(gin)
	lw.Debug("routes initialized successfully")

	if err := gin.Run(listenPort); err != nil {
		lw.Fatalf("gin engine failed to run ", err.Error())
	}
}
