package response

import "github.com/mameikagou/gin-vue-study/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
