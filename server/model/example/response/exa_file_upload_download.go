package response

import "github.com/mameikagou/gin-vue-study/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
