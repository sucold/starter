package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

var DefaultCache = gcache.New()

func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return DefaultCache.Set(ctx, key, value, duration)
}
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	return DefaultCache.SetMap(ctx, data, duration)
}
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	return DefaultCache.SetIfNotExist(ctx, key, value, duration)
}

func SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (bool, error) {
	return DefaultCache.SetIfNotExistFunc(ctx, key, f, duration)
}

func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (bool, error) {
	return DefaultCache.SetIfNotExistFuncLock(ctx, key, f, duration)
}
func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return DefaultCache.Get(ctx, key)
}

func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {
	return DefaultCache.GetOrSet(ctx, key, value, duration)
}

func GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return DefaultCache.GetOrSetFunc(ctx, key, f, duration)
}

func GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return DefaultCache.GetOrSetFuncLock(ctx, key, f, duration)
}
func Contains(ctx context.Context, key interface{}) (bool, error) {
	return DefaultCache.Contains(ctx, key)
}
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return DefaultCache.GetExpire(ctx, key)
}
func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error) {
	return DefaultCache.Remove(ctx, keys...)
}
func Removes(ctx context.Context, keys []interface{}) error {
	return DefaultCache.Removes(ctx, keys)
}
func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	return DefaultCache.Update(ctx, key, value)
}
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	return DefaultCache.UpdateExpire(ctx, key, duration)
}
func Size(ctx context.Context) (int, error) {
	return DefaultCache.Size(ctx)
}
func Data(ctx context.Context) (map[interface{}]interface{}, error) {
	return DefaultCache.Data(ctx)
}
func Keys(ctx context.Context) ([]interface{}, error) {
	return DefaultCache.Keys(ctx)
}
func KeyStrings(ctx context.Context) ([]string, error) {
	return DefaultCache.KeyStrings(ctx)
}
func Values(ctx context.Context) ([]interface{}, error) {
	return DefaultCache.Values(ctx)
}
func MustGet(ctx context.Context, key interface{}) *gvar.Var {
	return DefaultCache.MustGet(ctx, key)
}
func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {
	return DefaultCache.MustGetOrSet(ctx, key, value, duration)
}
func MustGetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) *gvar.Var {
	return DefaultCache.MustGetOrSet(ctx, key, f, duration)
}
func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) *gvar.Var {
	return DefaultCache.MustGetOrSetFuncLock(ctx, key, f, duration)
}
func MustContains(ctx context.Context, key interface{}) bool {
	return DefaultCache.MustContains(ctx, key)
}
func MustGetExpire(ctx context.Context, key interface{}) time.Duration {
	return DefaultCache.MustGetExpire(ctx, key)
}
func MustSize(ctx context.Context) int {
	return DefaultCache.MustSize(ctx)
}
func MustData(ctx context.Context) map[interface{}]interface{} {
	return DefaultCache.MustData(ctx)
}
func MustKeys(ctx context.Context) []interface{} {
	return DefaultCache.MustKeys(ctx)
}
func MustKeyStrings(ctx context.Context) []string {
	return DefaultCache.MustKeyStrings(ctx)
}
func MustValues(ctx context.Context) []interface{} {
	return DefaultCache.MustValues(ctx)
}
