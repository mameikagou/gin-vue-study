package initialize

import (
	"github.com/mameikagou/gin-vue-study/server/global"
	"github.com/mameikagou/gin-vue-study/server/model/learn"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(learn.Book{})
	if err != nil {
		return err
	}
	return nil
}
