package base

import (
	"encoding/json"
	"github.com/gogf/gf/v2/os/gfile"
)

type Setting struct {
	Name     string `json:"name" dc:"标题"`
	Logo     string `json:"logo" dc:"图标"`
	Color    string `json:"color" dc:"主色"`
	Theme    string `json:"theme" dc:"主题"`
	Register bool   `json:"register" dc:"开放注册"`
	Forget   bool   `json:"forget" dc:"开放找回"`
}

var DefaultSetting = Setting{
	Name:  "默认",
	Logo:  "/logo.svg",
	Color: "#13C2C2",
	Theme: "light",
}

const SettingPath = "setting.json"

func init() {
	if gfile.Exists(SettingPath) {
		data := gfile.GetBytes(SettingPath)
		_ = json.Unmarshal(data, &DefaultSetting)
	}
}
