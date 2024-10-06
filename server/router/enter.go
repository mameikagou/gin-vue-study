package router

import (
	"github.com/mameikagou/gin-vue-study/server/router/example"
	"github.com/mameikagou/gin-vue-study/server/router/study"
	"github.com/mameikagou/gin-vue-study/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Study   study.RouterGroup
}
