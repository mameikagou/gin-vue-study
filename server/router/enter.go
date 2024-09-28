package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/mameikagou/gin-vue-study/server/router/example"
	_ "github.com/mameikagou/gin-vue-study/server/router/study"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Study   example.RouterGroup
}
