package set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := New()
	assert.Equal(t, len(set.items), 0)
	entity := struct {
		value int
	}{
		value: 3,
	}
	set.Add(entity)
	fmt.Println(set.Values())

	assert.Equal(t, len(set.items), 1)
}

func TestHashSet_Size(t *testing.T) {
	set := New()
	assert.Equal(t, set.Size(), 0)
	entity := struct {
		value int
	}{
		value: 3,
	}
	entity2 := struct {
		value       int
		secondValue string
	}{
		value:       1,
		secondValue: "",
	}
	set.Add(entity)
	set.Add(entity2)
	fmt.Println(set.Values())

	assert.Equal(t, set.Size(), 2)
}

func TestHashSet_Remove(t *testing.T) {
	set := New()
	assert.Equal(t, set.Size(), 0)
	set.Remove()
	entity := struct {
		value int
	}{
		value: 3,
	}
	set.Add(entity)
}