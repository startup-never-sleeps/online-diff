package main

import (
	"log"

	config "online-diff/src/config"
	s3support "online-diff/src/s3support"
	http_server "online-diff/src/server"
	stale_daemon "online-diff/src/stale_daemon"
	containers "online-diff/src/storage_container"
	nlp "online-diff/src/text_similarity"
	utils "online-diff/src/utils"
)

const (
	configPath = "config/main_config.json"
)

var (
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func configure() {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalln("Unable to read program config from", configPath, err)
	}

	if err := utils.InitializeEmptyDir(conf.Internal.TempFilesDir); err != nil {
		log.Fatalln("Unable to initialize temporary directory at", conf.Internal.TempFilesDir, err)
	}

	if err := utils.InitializeLogger(conf.Internal.LoggingDir); err != nil {
		log.Fatalln("Unable to initialize logging at", conf.Internal.LoggingDir, err)
	}

	db, err := containers.NewDbClientService(conf.Internal.DbPath)
	if err != nil {
		errorLogger.Fatalln("Unable to open db at", conf.Internal.DbPath, err)
	}

	nlpCore := nlp.NewPyhonNLP(
		conf.Internal.PythonDifferenceScriptPath,
		conf.Internal.PythonSimilarityScriptPath,
	)

	s3Client, err := s3support.NewMinioService(&conf.Minio)
	if err != nil {
		errorLogger.Fatalln("Unable to initialize s3Client", err)
	}

	daemon := stale_daemon.NewDaemon(db, s3Client, conf.Internal.RefreshStaleDataPeriod)
	daemon.StartAsync()

	server := http_server.NewServer(conf, db, s3Client, nlpCore)
	server.Run()
}

func main() {
	configure()
}
