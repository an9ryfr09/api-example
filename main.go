package main

import (
	"a6-api/routes"
	"a6-api/utils/loader"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
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
	runtime.GOMAXPROCS(int(loader.Load().Core.CpuCoreNum))

	//run mode [debug|test|release]
	gin.SetMode(loader.Load().Core.RunMode)

	app = gin.New()

	//set middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	// app.Use(middleware.JWT())
	// app.Use(middleware.Validator())

	//initialization routers
	router.InitRouter(app)

	//set server options
	srv := &http.Server{
		Addr:           loader.Load().Server.Addr,
		Handler:        app,
		ReadTimeout:    loader.Load().Server.ReadTimeout * time.Second,
		WriteTimeout:   loader.Load().Server.WriteTimeout * time.Second,
		IdleTimeout:    loader.Load().Server.IdleTimeout * time.Second,
		MaxHeaderBytes: int(loader.Load().Server.MaxHeaderBytes),
	}

	//if this item value is true, then enabled the https protocol, otherwise use the http protocol only
	if loader.Load().Server.EnableTLS == true {
		go func() {
			srv.ListenAndServeTLS(loader.Load().Server.SSLCertfilePath, loader.Load().Server.SSLKeyfilePath)
		}()
	} else {
		go func() {
			srv.ListenAndServe()
		}()
	}
}

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			loader.Reload()
			log.Println("Reloaded config")
		}
	}()

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
