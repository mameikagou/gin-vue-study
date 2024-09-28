package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type FirstRouter struct{}

func (r *FirstRouter) InitFirstRouter(Router *gin.RouterGroup) {
	firstRouter := Router.Group("study")
	{
		firstRouter.POST("/first", func(context *gin.Context) {
			fmt.Println("This is a test output.")
			context.JSON(200, gin.H{"message": "Success", "data": []int{1, 2, 3, 4}})
		})
	}
}
