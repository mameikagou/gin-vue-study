package system

import (
	"github.com/mameikagou/gin-vue-study/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
