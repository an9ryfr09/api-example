package main

import (
	middleware "a6-api/middleware/verification"
	"a6-api/router"
	"a6-api/utils/loader"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
)

var app *gin.Engine
var srv http.Server

func init() {
	initServer()
}

//initServer Initialization http or https server configure
func initServer() {
	//set cpu core numbers
	runtime.GOMAXPROCS(int(loader.Load().Core.CpuCoreNum))

	//run mode [debug|test|release]
	gin.SetMode(loader.Load().Core.RunMode)

	app = gin.New()

	//add middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.Validator())
	// app.Use(middleware.Cors())
	// app.Use(middleware.JWT())

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
			//enabled http2
			http2.ConfigureServer(srv, &http2.Server{})
			srv.ListenAndServeTLS(loader.Load().Server.SSLCertfilePath, loader.Load().Server.SSLKeyfilePath)
		}()
	} else {
		go func() {
			srv.ListenAndServe()
		}()
	}
}

//startServer start http or https server on goroutine
func startServer() {
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

func main() {
	startServer()
}
