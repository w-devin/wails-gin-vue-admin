package utils

import (
	"fmt"
	"strconv"
	"wails-gin-vue-admin/global"
	"wails-gin-vue-admin/model/system"
)

func RegisterApis(apis ...system.SysApi) {
	var count int64
	var apiPaths []string
	for i := range apis {
		apiPaths = append(apiPaths, apis[i].Path)
	}
	global.GVA_DB.Find(&[]system.SysApi{}, "path in (?)", apiPaths).Count(&count)
	if count > 0 {
		fmt.Println("插件已安装或存在同名路由")
		return
	}
	err := global.GVA_DB.Create(&apis).Error
	if err != nil {
		fmt.Println(err)
	}
}

func RegisterMenus(menus ...system.SysBaseMenu) {
	var count int64
	var menuNames []string
	parentMenu := menus[0]
	otherMenus := menus[1:]
	for i := range menus {
		menuNames = append(menuNames, menus[i].Name)
	}
	global.GVA_DB.Find(&[]system.SysBaseMenu{}, "name in (?)", menuNames).Count(&count)
	if count > 0 {
		fmt.Println("插件已安装或存在同名菜单")
		return
	}
	parentMenu.ParentId = "0"
	err := global.GVA_DB.Create(&parentMenu).Error
	if err != nil {
		fmt.Println(err)
	}
	for i := range otherMenus {
		pid := strconv.Itoa(int(parentMenu.ID))
		otherMenus[i].ParentId = pid
	}
	err = global.GVA_DB.Create(&otherMenus).Error
	if err != nil {
		fmt.Println(err)
	}
}
