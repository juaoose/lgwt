package counter

import "sync"

// run `go vet` to chck for suspicious constructs
type Counter struct {
	// You could also embed the mutex here, like:
	// sync.Mutex, which would allow to
	// do c.Lock() and defer c.Unlock() in Inc()
	// But it would also make it part of the public API of Counter
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
