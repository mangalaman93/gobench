package gobench

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	Value int64
	Lock  sync.RWMutex
}

func (c *Counter) Incr() {
	c.Value = c.Value + 1
}

func (c *Counter) IncrAtomic() {
	atomic.AddInt64(&c.Value, 1)
}

func (c *Counter) IncrLock() {
	c.Lock.Lock()
	c.Value = c.Value + 1
	c.Lock.Unlock()
}

func (c *Counter) Read() int64 {
	return c.Value
}

func (c *Counter) ReadAtomic() int64 {
	return atomic.LoadInt64(&c.Value)
}

func (c *Counter) ReadLock() int64 {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.Value
}
