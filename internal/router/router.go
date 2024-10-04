package router

import (
	"github.com/fasthttp/router"
	"github.com/gythialy/jb-helper/internal/handler"
	"github.com/gythialy/jb-helper/internal/utils"
	"github.com/valyala/fasthttp"
)

func NewRouter() *router.Router {
	r := router.New()
	r.NotFound = fasthttp.FSHandler(utils.GetStaticPath(), 0)
	r.GET("/", handler.IndexHandler)
	api := r.Group("/api")
	api.POST("/generate", handler.LicenseHandler)
	return r
}
