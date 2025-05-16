package initialize

import (
	"github.com/gin-gonic/gin"
	"vizhi_atlas/internal/pkg/globals"
)

func initEngine() {
	e := gin.Default()
	e.Use(gin.Recovery())
	globals.E = e
}
