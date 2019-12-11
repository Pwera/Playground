package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TestObjectImplementation struct {
	mock.Mock
}

func printArgumentFunction(str string) {
	fmt.Print(str)
}
func printArgumentLenFunction(str string) {
	fmt.Print(len(str))
}

func Test_Nothing(t *testing.T) {
	object := &Object{}
	Function(object, printArgumentFunction)
}

func Test_Object_F(t *testing.T) {
	var mockedService = new(TestObjectImplementation)
	c := mockedService.On("Method_A").Return("arg")
	assert.Equal(t, []*mock.Call{c}, mockedService.ExpectedCalls)
	assert.Equal(t, "Method_A", c.Method)
	assert.Equal(t, "arg", c.ReturnArguments.Get(0))
}
func Test_Object_F_TestData(t *testing.T) {
	var mockedService = new(TestObjectImplementation)
	mockedService.
		On("Method_B", 0).Return("0", nil)
	expectedCalls := []*mock.Call{
		{
			Parent:          &mockedService.Mock,
			Method:          "Method_B",
			Arguments:       []interface{}{0},
			ReturnArguments: []interface{}{"0", nil},
		},
	}
	assert.Equal(t, "Method_B", expectedCalls[0].Method)
	assert.Equal(t, "0", expectedCalls[0].ReturnArguments[0])
	assert.Equal(t, nil, expectedCalls[0].ReturnArguments[1])
	assert.Equal(t, 0, expectedCalls[0].Arguments[0])
}

func Test_Object_AnythingOfType(t *testing.T) {
	var mockedService = new(TestObjectImplementation)
	c := mockedService.On("Method_C", mock.AnythingOfType("func(int)error)")).Return(0, nil)

	assert.Equal(t, []*mock.Call{c}, mockedService.ExpectedCalls)
	assert.Equal(t, 0, c.ReturnArguments[0])
	assert.Equal(t, nil, c.ReturnArguments[1])
	assert.Equal(t, "Method_C", c.Method)
}
