package v1

import (
	"wails-gin-vue-admin/api/v1/example"
	"wails-gin-vue-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
