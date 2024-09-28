package study

import (
	api "github.com/mameikagou/gin-vue-study/server/api/v1"
)

type RouterGroup struct {
	FirstRouter
	SecondRouter
}

var (
	FirstApi  = api.ApiGroupApp.StudyApiGroup
	SecondApi = api.ApiGroupApp.StudyApiGroup
)
