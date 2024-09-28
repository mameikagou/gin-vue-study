package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SecondApi struct{}

// @Router    /study/first [post]
func (e *SecondApi) CreateSecondApi(c *gin.Context) {
	fmt.Println("This is a test output for second api.")
	c.JSON(200, gin.H{"message": "Success", "data": []int{1, 2, 3, 4}})
}
