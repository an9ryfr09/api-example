package main

import (
	configure "a6-api/pkg/loader"
	router "a6-api/routers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine
var srv http.Server

func init() {
	initServer()
}

func initServer() {
	gin.SetMode(configure.AppConf.RunMode)
	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	// app.Use(jwt.JWT())
	router.InitRouter(app)
	//setting server options
	srv := &http.Server{
		Addr:           configure.ServerConf.Addr,
		Handler:        app,
		ReadTimeout:    configure.ServerConf.ReadTimeout,
		WriteTimeout:   configure.ServerConf.WriteTimeout,
		IdleTimeout:    configure.ServerConf.IdleTimeout,
		MaxHeaderBytes: configure.ServerConf.MaxHeaderBytes,
	}

	//if this item value is true, then enabled https protocol, otherwise use http protocol only
	if configure.ServerConf.EnableTLS == true {
		go func() {
			srv.ListenAndServeTLS(configure.ServerConf.SSLCertfilePath, configure.ServerConf.SSLKeyfilePath)
		}()
		// app.RunTLS(configure.ServerConf.Addr, configure.ServerConf.SSLCertfilePath, configure.ServerConf.SSLKeyfilePath)
	} else {
		go func() {
			srv.ListenAndServe()
		}()
		// app.Run(configure.ServerConf.Addr)
	}
}

func main() {
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
