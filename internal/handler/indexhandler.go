package handler

import (
	"github.com/gythialy/jb-helper/internal/response"
	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx) {
	response.Html(ctx, "index.html")
}
