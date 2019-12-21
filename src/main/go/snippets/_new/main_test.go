package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func Test_f(t *testing.T) {
	t.Log("Invoking " + t.Name())
	assert.Equal(t, 1, 1)
}

func TestIntComparable(t *testing.T) {
	n := 1
	var pn *int = &n
	var up = unsafe.Pointer(&pn)

	assert.True(t, reflect.TypeOf(pn).Comparable())
	assert.True(t, reflect.TypeOf(up).Comparable())
}

func TestFunctionComparable(t *testing.T) {
	type EmptyFunc func()
	type ConsumerFunc func() EmptyFunc
	type SupplierFunc func(EmptyFunc)
	type VariadicFunc func(...EmptyFunc)
	var f1 EmptyFunc = func() {}
	var f2 ConsumerFunc = func() EmptyFunc {
		return f1
	}
	var f3 SupplierFunc = func(f EmptyFunc) {}
	var f4 VariadicFunc = func(f ...EmptyFunc) {}

	assert.False(t, reflect.TypeOf(f1).Comparable())
	assert.False(t, reflect.TypeOf(f2).Comparable())
	assert.False(t, reflect.TypeOf(f3).Comparable())
	assert.False(t, reflect.TypeOf(f4).Comparable())
}

type Helper = interface {
	Help() string
}
type RealHelper struct{}
type NotSoRealHelper struct{}

func (rh RealHelper) Help() string      { return "" }
func (rh NotSoRealHelper) Help() string { return "nil" }

func TestInterface(t *testing.T) {
	//compile-time check
	var a = Helper(RealHelper{})
	var b Helper = RealHelper{}
	var explicit = interface{ Help() string }.Help(a)

	helpers := []Helper{
		a,
		b,
		&NotSoRealHelper{},
	}
	assert.Equal(t, a, b)
	assert.NotEqual(t, a, explicit)
	assert.NotEqual(t, b, explicit)
	assert.NotNil(t, helpers)
}
