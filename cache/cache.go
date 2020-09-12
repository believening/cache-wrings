package cache

import "log"

// Cache .
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

func New(typ string) Cache {
	switch typ {
	case "inmemory":
		log.Println(typ, "ready to serve")
		return newInMemoryCache()
	default:
		log.Println("not support")
	}
	return nil
}
