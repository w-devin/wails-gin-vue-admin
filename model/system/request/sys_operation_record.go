package request

import (
	"wails-gin-vue-admin/model/common/request"
	"wails-gin-vue-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
