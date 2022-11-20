package base

import (
	"encoding/json"
	"github.com/gogf/gf/v2/os/gfile"
)

type ProLayout struct {
	NavTheme     string `json:"navTheme"`
	ColorPrimary string `json:"colorPrimary"`
	Logo         string `json:"logo"`
	Title        string `json:"title"`
}
type Setting struct {
	Layout   ProLayout `json:"layout"`
	Register bool      `json:"register" dc:"开放注册"`
	Forget   bool      `json:"forget" dc:"开放找回"`
	Verify   bool      `json:"verify" dc:"注册验证"`
}

var DefaultSetting = Setting{
	Layout: ProLayout{
		NavTheme:     "light",
		ColorPrimary: "#13C2C2",
		Logo:         "/logo.svg",
		Title:        "默认名称",
	},
	Register: false,
	Forget:   false,
	Verify:   false,
}

const SettingPath = "setting.json"

func init() {
	if gfile.Exists(SettingPath) {
		data := gfile.GetBytes(SettingPath)
		_ = json.Unmarshal(data, &DefaultSetting)
	}
}
