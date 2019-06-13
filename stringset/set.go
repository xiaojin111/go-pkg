package stringset

// Set 用于检查ip,mac,zone,country等是否存在
type Set map[string]bool

// 往Set里面添加元素
func (s Set) Put(key string) {
	s[key] = true
}

// 判断某个元素key是否存在于Set
func (s Set) Exist(key string) bool {
	_, ok := s[key]
	return ok
}
