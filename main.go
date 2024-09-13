package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/lossurdo/golang-starter-project/cmd"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.UnixDate,
	})
}

func main() {
	if err := godotenv.Load("example_dot_env"); err != nil {
		log.Fatalln("Error loading '.env' file")
	}
	cmd.Execute()
}
