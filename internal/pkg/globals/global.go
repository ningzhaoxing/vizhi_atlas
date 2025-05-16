package globals

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	E   *gin.Engine
	Db  *gorm.DB
	C   *Config
	Log *zap.Logger
)
