package cache

import (
	"sync"
)

type inMemoryCache struct {
	sync.RWMutex
	m map[string][]byte
	Stat
}

func (imc *inMemoryCache) Set(k string, v []byte) error {
	imc.Lock()
	defer imc.Unlock()
	imc.m[k] = v
	imc.add(k, v)
	return nil
}

func (imc *inMemoryCache) Get(k string) ([]byte, error) {
	// var err error
	imc.RLock()
	defer imc.RUnlock()
	// v, exist := imc.m[k]
	// if !exist {
	// 	err = ErrNoExistKey
	// }
	// return v, err

	// 没有向上返回错误，便于上层判断不存在
	// 错误类型保留在包内使用，小写？
	return imc.m[k], nil
}

func (imc *inMemoryCache) Del(k string) error {
	imc.Lock()
	defer imc.Unlock()
	v, exist := imc.m[k]
	if exist {
		delete(imc.m, k)
		imc.del(k, v)
	}
	return nil
}

func (imc *inMemoryCache) GetStat() Stat {
	return imc.Stat
}

// NewInMemCache 返回内存缓存
func NewInMemCache() Cache {
	return &inMemoryCache{
		m:    make(map[string][]byte),
		Stat: Stat{0, 0, 0},
	}
}
