package response

import "wails-gin-vue-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
