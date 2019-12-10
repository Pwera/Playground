package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_f(t *testing.T){
	t.Log("Invoking " + t.Name())
	assert.Equal(t, 1,1)
}
