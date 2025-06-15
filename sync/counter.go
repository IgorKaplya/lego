package counter

import "sync"

type Counter struct {
	value int
	mutex sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}
