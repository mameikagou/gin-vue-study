package v1

import (
	"github.com/mameikagou/gin-vue-study/server/api/v1/example"
	"github.com/mameikagou/gin-vue-study/server/api/v1/learn"
	"github.com/mameikagou/gin-vue-study/server/api/v1/study"
	"github.com/mameikagou/gin-vue-study/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	StudyApiGroup   study.ApiGroup
	LearnApiGroup   learn.ApiGroup
}
