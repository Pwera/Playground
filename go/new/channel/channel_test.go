package channel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestChannel(t *testing.T) {
	var c = make(chan T, 1)
	assert.False(t, IsClosed(c))
	assert.Equal(t, cap(c), 1)
	assert.Equal(t, len(c), 0)
	c <- 67
	assert.True(t, IsClosed(c))
	assert.Equal(t, cap(c), 1)
	assert.Equal(t, len(c), 0)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
