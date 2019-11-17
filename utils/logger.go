package utils

import (
	configure "a6-api/utils/loader"

	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func init() {
	//disabled console colors, write log not use color
	gin.DisableConsoleColor()

	//create logs file
	//format values must be 20060102, I don't know why
	logFile, err := os.Create(fmt.Sprintf("%s/%s.log", configure.CoreConf.LogPath, time.Now().Format("20060102")))

	gin.DefaultWriter = io.MultiWriter(logFile)

	Logger := log.New(logFile, "core: ", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		Logger.Printf("%s is not exists or non-write permission", configure.CoreConf.LogPath)
	}
}
