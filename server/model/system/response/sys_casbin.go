package response

import (
	"github.com/mameikagou/gin-vue-study/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
