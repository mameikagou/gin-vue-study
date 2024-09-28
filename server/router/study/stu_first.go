package study

import (
	"github.com/gin-gonic/gin"
)

type FirstRouter struct{}

func (r *FirstRouter) InitFirstRouter(Router *gin.RouterGroup) {
	firstRouter := Router.Group("study")
	{
		firstRouter.POST("/first", r.F)
	}
}
