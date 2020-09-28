/**
 * Copyright 2020 gd Author. All rights reserved.
 * Author: Chuck1024
 */

package main

import (
	"github.com/chuck1024/gd"
	"github.com/chuck1024/gd/dlog"
	"github.com/chuck1024/gd/net/dhttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerHttp(c *gin.Context, req interface{}) (code int, message string, err error, ret string) {
	dlog.Debug("httpServerTest req:%v", req)
	ret = "ok!!!"
	return http.StatusOK, "ok", nil, ret
}

func main() {
	d := gd.Default()
	d.HttpServer.SetInit(func(g *gin.Engine) error {
		r := g.Group("")
		r.Use(
			dhttp.GlFilter(),
			dhttp.GroupFilter(),
			dhttp.Logger("quick-start"),
		)

		d.HttpServer.GET(r, "test", HandlerHttp)

		if err := d.HttpServer.CheckHandle(); err != nil {
			return err
		}
		return nil
	})

	d.SetConfig("Server", "httpPort", "10240")

	if err := d.Run(); err != nil {
		dlog.Error("Error occurs, error = %s", err.Error())
		return
	}
}