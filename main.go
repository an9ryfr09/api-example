package main

import (
	router "a6-api/routers"
	configure "a6-api/utils/loader"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var srv http.Server

func init() {
	initServer()
}

func initServer() {
	//set cpu core numbers
	runtime.GOMAXPROCS(configure.AppConf.CpuCoreNum)

	//run mode [debug|test|release]
	gin.SetMode(configure.AppConf.RunMode)

	app = gin.New()

	//set middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	// app.Use(jwt.JWT())

	//initialization routers
	router.InitRouter(app)

	//set server options
	srv := &http.Server{
		Addr:           configure.ServerConf.Addr,
		Handler:        app,
		ReadTimeout:    configure.ServerConf.ReadTimeout,
		WriteTimeout:   configure.ServerConf.WriteTimeout,
		IdleTimeout:    configure.ServerConf.IdleTimeout,
		MaxHeaderBytes: configure.ServerConf.MaxHeaderBytes,
	}

	//if this item value is true, then enabled the https protocol, otherwise use the http protocol only
	if configure.ServerConf.EnableTLS == true {
		go func() {
			srv.ListenAndServeTLS(configure.ServerConf.SSLCertfilePath, configure.ServerConf.SSLKeyfilePath)
		}()
	} else {
		go func() {
			srv.ListenAndServe()
		}()
	}
}

func main() {
	//start up http server on goroutine, and receive unix signals to shutdown;
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
