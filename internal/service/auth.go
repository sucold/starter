package service

import (
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hinego/authentic"
	"github.com/hinego/systemd/internal/cache"
	"github.com/hinego/systemd/internal/consts"
	"github.com/hinego/systemd/internal/dao"
	"github.com/hinego/systemd/internal/model"
	"gorm.io/gen"
	"time"
)

var (
	Auth *authentic.Authentic
)

func StartAuth() error {
	Auth = &authentic.Authentic{
		AddCode:  addCode,
		SetCode:  setCode,
		DelCode:  delCode,
		LoadCode: loadCode,
		Cache:    cache.DefaultCache,
		Key:      consts.UserKey,
	}
	return Auth.Init()
}
func loadCode(c *authentic.Context) error {
	var result = make([]*model.Token, 0)
	t := dao.Token
	return t.Where(t.Expire.Gt(time.Now().Unix()), t.LogoutAt.Eq(0)).FindInBatches(&result, 1000, func(tx gen.Dao, batch int) error {
		for _, res := range result {
			data := map[string]any{}
			if err := json.Unmarshal([]byte(res.Data), &data); err == nil {
				if err := c.Cache.Set(c.Context, res.Code, data, time.Unix(res.Expire, 0).Sub(time.Now())); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
func addCode(c *authentic.Context) error {
	request := g.RequestFromCtx(c.Context)
	c.Data["expire"] = c.Expire.Unix()
	c.Data["device_id"] = 0

	token := &model.Token{
		Code:      c.Token.Raw,
		Expire:    c.Expire.Unix(),
		UserID:    c.Data[consts.UserKey].(int64),
		IP:        request.GetClientIp(),
		UserAgent: request.UserAgent(),
		Data:      gjson.MustEncodeString(c.Data),
		Role:      "login",
	}
	if role, ok := c.Data["token_type"]; ok {
		token.Role = gconv.String(role)
	}
	err := dao.Token.Create(token)
	c.Data["token_id"] = token.ID
	return err
}
func setCode(c *authentic.Context) error {
	_, err := dao.Token.Where(dao.Token.Code.Like(c.Token.Raw)).UpdateSimple(dao.Token.Expire.Value(c.Expire.Unix()))
	return err
}
func delCode(c *authentic.Context) error {
	_, err := dao.Token.Where(dao.Token.Code.Like(c.Token.Raw)).UpdateSimple(dao.Token.LogoutAt.Value(time.Now().Unix()))
	return err
}
