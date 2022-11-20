package base

import (
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
)

type ProLayout struct {
	NavTheme     string `json:"navTheme"`
	ColorPrimary string `json:"colorPrimary"`
	Logo         string `json:"logo"`
	Title        string `json:"title"`
}
type Setting struct {
	NavTheme     string `json:"navTheme"`
	ColorPrimary string `json:"colorPrimary"`
	Logo         string `json:"logo"`
	Title        string `json:"title"`
	Icon         string `json:"icon" dc:"ICON图标"`
	Register     bool   `json:"register" dc:"开放注册"`
	Forget       bool   `json:"forget" dc:"开放找回"`
	Verify       bool   `json:"verify" dc:"注册验证"`
	Desc         string `json:"desc" dc:"登录说明"`
}

func (r *Setting) JSON() map[string]any {
	return map[string]any{
		"register": r.Register,
		"verify":   r.Verify,
		"desc":     r.Desc,
		"icon":     r.Icon,
		"forget":   r.Forget,
		"layout": map[string]any{
			"navTheme":     r.NavTheme,
			"colorPrimary": r.ColorPrimary,
			"logo":         r.Logo,
			"title":        r.Title,
		},
	}
}

var DefaultSetting = Setting{
	NavTheme:     "light",
	ColorPrimary: "#13C2C2",
	Logo:         "/logo.svg",
	Title:        "默认名称",
	Register:     true,
	Forget:       true,
	Verify:       true,
	Desc:         "最牛逼的卖牛平台面板",
}

const SettingPath = "setting.json"

func SaveConfig() {
	_ = gfile.PutBytes(SettingPath, gjson.MustEncode(DefaultSetting))
}
func init() {
	if gfile.Exists(SettingPath) {
		data := gfile.GetBytes(SettingPath)
		_ = json.Unmarshal(data, &DefaultSetting)
	}
}
