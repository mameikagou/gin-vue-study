package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type FirstApi struct{}

// @Router    /study/first [post]
func (e *FirstApi) CreateFirstApi(c *gin.Context) {
	fmt.Println("This is a test output.")
	c.JSON(200, gin.H{"message": "Success", "data": []int{1, 2, 3, 4}})
}
