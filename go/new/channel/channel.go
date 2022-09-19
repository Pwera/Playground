package channel

import (
	"sync"
)

type T int

func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

type ClosableChan struct {
	ch       chan T
	once     sync.Once
	isClosed bool
}

func (cc *ClosableChan) Close() {
	cc.once.Do(func() {
		close(cc.ch)
		cc.isClosed = true
	})
}

func (cc *ClosableChan) IsClosed() bool {
	return cc.isClosed
}

func Ping(ch chan string) {
	ch <- "ping"
}