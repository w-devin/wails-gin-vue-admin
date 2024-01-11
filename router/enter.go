package router

import (
	"wails-gin-vue-admin/router/example"
	"wails-gin-vue-admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
