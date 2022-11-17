package table

type App struct {
	Model
	Name string `json:"name" dc:"软件名称"`
	Desc string `json:"desc" dc:"软件简介"`
	Tags string `json:"tags" dc:"文件MD5"`
	Size int64  `json:"size" dc:"文件大小"`
}

type Node struct {
	Model
	Name   string   `json:"name" dc:"节点名称"`
	Addr   string   `json:"addr" dc:"节点IP"`
	Addrs  []string `json:"addrs" dc:"所有IP"`
	Cpu    string   `json:"cpu"`
	Core   int64    `json:"core"`
	Os     string   `json:"os" dc:"节点系统"`
	Mem    int64    `json:"mem" dc:"节点内存"`
	Disk   int64    `json:"disk"`
	Online bool     `json:"online"`
}

type Service struct {
	Model
	Name   string `json:"name" dc:"服务名称"`
	NodeID int64  `json:"node_id" dc:"绑定节点ID"`
	AppID  int64  `json:"app_id" dc:"软件ID"`
	Memory int64  `json:"memory" dc:"内存占用"`
	Status bool   `json:"status"`
}
