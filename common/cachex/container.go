package cachex

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

type Container[V any] struct {
	c *collection.Cache
}

func NewContainer[V any](expire time.Duration, opts ...collection.CacheOption) (*Container[V], error) {
	c, err := collection.NewCache(expire, opts...)
	if err != nil {
		return nil, err
	}
	return &Container[V]{c}, nil
}

func (t Container[V]) Set(key string, v *V) {
	t.c.Set(key, v)
}

func (t Container[V]) Get(key string) (*V, bool) {
	v, ok := t.c.Get(key)

	if ok {
		return v.(*V), ok
	}

	return nil, ok
}
