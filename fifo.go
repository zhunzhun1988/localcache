package localcache

const (
	cacherName = "fifo"
)

func init() {
	RegistCacher(cacherName, NewFIFOCache)
}

type FIFOCache struct {
	data       []interface{}
	startIndex int
	endIndex   int
}

var _ = LocalCache(&FIFOCache{})

func NewFIFOCache(arg LCNewFuncArg) LocalCache {

}

func (fifo *FIFOCache) Get(key LCKey) (value LCValue, ok bool) {

}

func (fifo *FIFOCache) Set(key LCKey, value LCValue) error {

}

func (fifo *FIFOCache) Set(key LCKey, value LCValue, expire time.Duration) error {

}

func (fifo *FIFOCache) Delete(key LCKey) error {

}

func (fifo *FIFOCache) Exists(key LCKey) bool {

}

func (fifo *FIFOCache) Mget(keys []LCKey) (values map[LCKey]LCValue, ok bool) {

}

func (fifo *FIFOCache) Capacity() LCSize {

}

func (fifo *FIFOCache) Size() LCSize {

}
