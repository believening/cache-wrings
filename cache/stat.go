package cache

// Stat cache 状态
type Stat struct {
	Cnt       int
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(k string, v []byte) {
	s.Cnt++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Cnt--
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
