package service

import (
	"wails-gin-vue-admin/service/example"
	"wails-gin-vue-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
