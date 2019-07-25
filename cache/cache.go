package cache

import (
	"errors"
)

var (
	// ErrNoExistKey 查询键不存在
	ErrNoExistKey = errors.New("Cache: key is not exist")
)

// Cache 接口
type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
	GetStat() Stat
}
