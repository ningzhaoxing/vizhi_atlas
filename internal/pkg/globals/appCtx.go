package globals

import (
	"gorm.io/gorm"
)

type AppCtx struct {
	db     *gorm.DB
	config *Config
}

type Options func(a *AppCtx)

func NewAppCtx(op ...Options) *AppCtx {
	appCtx := new(AppCtx)
	for _, options := range op {
		options(appCtx)
	}
	return appCtx
}

func NewDefaultAppCtx() *AppCtx {
	return NewAppCtx(WithOptionDb(Db), WithOptionConfig(C))
}

func WithOptionDb(db *gorm.DB) Options {
	return func(a *AppCtx) {
		a.db = db
	}
}

func WithOptionConfig(c *Config) Options {
	return func(a *AppCtx) {
		a.config = c
	}
}

func (a *AppCtx) SetDb(db *gorm.DB) {
	a.db = db
}
