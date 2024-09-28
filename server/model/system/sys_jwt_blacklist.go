package system

import (
	"github.com/mameikagou/gin-vue-study/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
