// Code generated by hertz generator.

package main

import (
	"context"
	handler "github.com/1359332949/douyin/cmd/api/biz/handler"
	"github.com/1359332949/douyin/cmd/api/biz/handler/api"
	"github.com/1359332949/douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		api.SendResponse(c, errno.ServiceErr, nil)
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		api.SendResponse(c, errno.ServiceErr, nil)
	})
}