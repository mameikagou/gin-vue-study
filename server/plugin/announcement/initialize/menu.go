package initialize

import (
	"context"
	model "github.com/mameikagou/gin-vue-study/server/model/system"
	"github.com/mameikagou/gin-vue-study/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  24,
			Path:      "anInfo",
			Name:      "anInfo",
			Hidden:    false,
			Component: "plugin/announcement/view/info.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "公告管理", Icon: "box"},
		},
	}
	utils.RegisterMenus(entities...)
}
