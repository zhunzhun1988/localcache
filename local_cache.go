package localcache

import (
	"fmt"
	"sync"
	"time"
)

type LCNewFuncArg struct {
	capacity LCSize
}

type LCKey string
type LCValue interface{}
type LCSize int64
type LCNewFunc func(arg LCNewFuncArg) LocalCache

// LocalCache is an interface for a cache to store key value.
type LocalCache interface {
	Get(key LCKey) (value LCValue, ok bool)
	Set(key LCKey, value LCValue) error
	Set(key LCKey, value LCValue, expire time.Duration) error
	Delete(key LCKey) error
	Exists(key LCKey) bool
	Mget(keys []LCKey) (values map[LCKey]LCValue, ok bool)
	Capacity() LCSize
	Size() LCSize
}

type LocalCacheFactory struct {
	cacher map[string]LocalCache
	mu     sync.Mutex
}

func (lcf *LocalCacheFactory) Regist(name string, newFunc LCNewFunc) {
	lcf.mu.Lock()
	defer lcf.mu.Unlock()
	lcf.cacher[name] = newFunc
}

func (lcf *LocalCacheFactory) ExistCacher(name string) bool {
	lcf.mu.Lock()
	defer lcf.mu.Unlock()
	if _, exist := lcf.cacher[name]; exist {
		return true
	}
	return false
}

func (lcf *LocalCacheFactory) GetCacherNewer(name string) (LCNewFunc, error) {
	lcf.mu.Lock()
	defer lcf.mu.Unlock()
	if newFunc, exist := lcf.cacher[name]; exist {
		return newFunc, nil
	} else {
		return nil, NewLocalCacheError(NotFoundError)
	}
}

var cacheFactory *LocalCacheFactory = LocalCacheFactory{
	cacher: make(map[string]LocalCache),
}

func RegistCacher(name string, newFunc LCNewFunc) {
	cacheFactory.Regist(name, newFunc)
}

func ExistCacher(name string) bool {
	return cacheFactory.ExistCacher(name)
}

func GetCacherNewer(name string) (LCNewFunc, error) {
	return cacheFactory.GetCacherNewer(name)
}
