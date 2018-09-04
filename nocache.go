package localcache


type NoCache struct {
}

var _ = LocalCache(&NoCache{})

func (nc *NoCache) Get(key LCKey) (value LCValue, ok bool) {
	return LCValue{}, false
}

func (nc *NoCache) Set(key LCKey, value LCValue) (err error) {
	return NewLocalCacheError(NotSupportError)
}

func (nc *NoCache) Set(key LCKey, value LCValue, expire time.Duration) error {
	return NewLocalCacheError(NotSupportError)
}

func (nc *NoCache) Delete(key LCKey) error {
	return NewLocalCacheError(NotSupportError)
}

func (nc *NoCache) Exists(key LCKey) bool {
	return false
}

func (nc *NoCache) Mget(keys []LCKey) (values map[LCKey]LCValue, ok bool) {
	return map[LCKey]LCValue{}, false
}

func (nc *NoCache) Capacity() LCSize {
	return LCSize(0)
}

func (nc *NoCache) Size() LCSize {
	return LCSize(0)
}