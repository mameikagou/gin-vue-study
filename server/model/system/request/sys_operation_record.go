package request

import (
	"github.com/mameikagou/gin-vue-study/server/model/common/request"
	"github.com/mameikagou/gin-vue-study/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
