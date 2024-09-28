package study

import (
	"github.com/gin-gonic/gin"
)

type SecondRouter struct{}

func (r *FirstRouter) InitSecondRouter(Router *gin.RouterGroup) {
	secondRouter := Router.Group("study")
	{
		secondRouter.POST("second")
	}
}
