package reflect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type User struct {
	name     string
	birthday time.Time
	age      int
}

func (u User) ToString() string {
	return fmt.Sprintf("User: %v %v %d", u.name, u.birthday, u.age)
}

func TestUserProperties(t *testing.T) {
	tom := User{}
	userType := reflect.TypeOf(tom)

	assert.Equal(t, userType.NumField(), 3)
	assert.True(t, userType.Comparable())
	assert.NotNil(t, userType.Kind())
	assert.NotNil(t, userType.NumMethod())
	userType.MethodByName("ToString")
}
