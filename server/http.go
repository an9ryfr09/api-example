package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr           string
	Handler        *gin.Engine
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes uint64
}
