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
