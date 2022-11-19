package base

type Setting struct {
	Name  string `json:"name" dc:"标题"`
	Logo  string `json:"logo" dc:"图标"`
	Color string `json:"color" dc:"主色"`
	Theme string `json:"theme" dc:"主题"`
}

var DefaultSetting = Setting{
	Name:  "默认",
	Logo:  "/logo.svg",
	Color: "#13C2C2",
	Theme: "light",
}
