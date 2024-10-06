package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/mameikagou/gin-vue-study/server/router"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
