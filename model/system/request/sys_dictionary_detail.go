package request

import (
	"wails-gin-vue-admin/model/common/request"
	"wails-gin-vue-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
