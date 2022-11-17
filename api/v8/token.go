package v8

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/systemd/api"
	"github.com/hinego/systemd/internal/model"
)

type TokenFetchReq struct {
	g.Meta `path:"/token/fetch" method:"get" sm:"密钥列表" tags:"密钥管理"`
	api.PageReq
}
type TokenFetchRes struct {
	*api.PageReq
	*api.PageRes
	Data []*model.Token `json:"data"`
}
type TokenUpdateReq struct {
	g.Meta `path:"/token/update" method:"post" sm:"更新密钥" tags:"密钥管理"`
	*model.Token
}
type TokenUpdateRes struct {
}

type TokenCreateReq struct {
	g.Meta `path:"/token/create" method:"post" sm:"创建密钥" tags:"密钥管理"`
}
type TokenCreateRes struct {
}
type TokenDeleteReq struct {
	g.Meta `path:"/token/delete" method:"post" sm:"删除密钥" tags:"密钥管理"`
	ID     int64 `json:"id"`
}
type TokenDeleteRes model.Token
