package main

import (
	"log"
	"net/http"

	api "web-service/src/api/controllers"
	config "web-service/src/config"
	s3support "web-service/src/s3support"
	containers "web-service/src/storage_container"
	utils "web-service/src/utils"
)

const (
	configPath = "config/main_config.json"
)

var (
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func configure() {
	if err := config.ReadConfig(configPath); err != nil {
		log.Fatalln("Unable to read program config from", configPath, err)
	}

	if err := utils.InitializeLogger(); err != nil {
		log.Fatalln("Unable to initialize logging at", config.Internal.LoggingDir, err)
	}

	errorLogger = utils.ErrorLogger
	debugLogger = utils.DebugLogger

	db := containers.NewDB()
	if err := db.Initialize(config.Internal.DbPath); err != nil {
		errorLogger.Fatalln("Unable to open db at", config.Internal.DbPath, err)
	}

	api.InitializeControllers(db)
	s3support.InitializeS3Support()
}

func setupRoutes() {
	http.HandleFunc("/upload_files", api.UploadFilesHandler)
	http.HandleFunc("/view/", api.ViewRoomHandler)
	http.HandleFunc("/link", api.GetFileLinkById)

}

func main() {
	configure()

	setupRoutes()

	debugLogger.Println("Starting fair online judge service on", config.Server.Port)
	errorLogger.Fatalln(http.ListenAndServe(config.Server.Port, nil))
}
