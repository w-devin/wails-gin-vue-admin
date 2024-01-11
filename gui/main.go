package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/w-devin/poketto/file"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/zap"
	"os"
	"wails-gin-vue-admin/core"
	"wails-gin-vue-admin/global"
	"wails-gin-vue-admin/initialize"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	CONFIG_FILE_PATH = "config.json"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go core.RunWindowsServer()

	appConfig := &AppConfig{}
	if file.IsFileExists(CONFIG_FILE_PATH) {
		configFile, _ := os.Open(CONFIG_FILE_PATH)
		defer configFile.Close()

		jsonBytes, _ := os.ReadFile(CONFIG_FILE_PATH)

		err := json.Unmarshal(jsonBytes, appConfig)
		if err != nil {
			fmt.Printf("failed to unmarshal config file, %v\n", err)
		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gui",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewAssetHandler(appConfig),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
