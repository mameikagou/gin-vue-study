package example

import (
	api "github.com/mameikagou/gin-vue-study/server/api/v1"
)

type RouterGroup struct {
	CustomerRouter
	FileUploadAndDownloadRouter
}

var (
	exaCustomerApi              = api.ApiGroupApp.ExampleApiGroup.CustomerApi
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
)
