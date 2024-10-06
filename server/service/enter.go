package service

import (
	"github.com/mameikagou/gin-vue-study/server/service/example"
	"github.com/mameikagou/gin-vue-study/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
