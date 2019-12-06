package main

import (
	"github.com/seborama/fuego"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Function(e fuego.Entry) fuego.Entry {
	return e
}
func getFunction() fuego.Function {
	return Function
}
func AllowPredicate(t fuego.Entry) bool {
	return true
}
func BlockPredicate(t fuego.Entry) bool {
	return false
}

func Test_Map_Collect(t *testing.T) {
	c := make(chan fuego.Entry)

	go func() {
		defer close(c)
		c <- fuego.EntryInt(1)
		c <- fuego.EntryInt(2)
		c <- fuego.EntryInt(36)
	}()
	stream := fuego.NewStream(c)
	tests := []struct {
		name   string
		arg    fuego.Stream
		want   fuego.Entry
		mapper fuego.Function
	}{
		{
			name:   "Should return EntrySlice containing 3 element",
			arg:    stream,
			want:   fuego.EntrySlice{fuego.EntryInt(1), fuego.EntryInt(2), fuego.EntryInt(36)},
			mapper: Function,
		},
		{
			name:   "Should return empty EntrySlice",
			arg:    stream,
			want:   fuego.EntrySlice{},
			mapper: Function,
		},
		{
			name:   "Should return empty EntrySlice",
			arg:    fuego.NewStream(c),
			want:   fuego.EntrySlice{},
			mapper: Function,
		},
		{
			name:   "Should return empty EntrySlice",
			arg:    fuego.NewStream(c),
			want:   fuego.EntrySlice{},
			mapper: Function,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.arg.
				Map(tt.mapper).
				Collect(fuego.ToEntrySlice())
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func Test_Map_Filter_Collect(t *testing.T) {
	data := []fuego.Entry{
		fuego.EntryInt(1),
	}

	c := make(chan fuego.Entry)

	go func() {
		defer close(c)
		c <- fuego.EntryInt(1)
		c <- fuego.EntryInt(2)
		c <- fuego.EntryInt(36)
	}()
	stream := fuego.NewStream(c)
	tests := []struct {
		name      string
		arg       fuego.Stream
		want      fuego.Entry
		mapper    fuego.Function
		predicate fuego.Predicate
	}{
		{
			name:      "Should return EntrySlice containing 3 element",
			arg:       stream,
			want:      fuego.EntrySlice{},
			mapper:    Function,
			predicate: BlockPredicate,
		},
		{
			name:      "Should return empty EntrySlice",
			arg:       stream,
			want:      fuego.EntrySlice{},
			mapper:    Function,
			predicate: AllowPredicate,
		},
		{
			name:      "Should return empty EntrySlice",
			arg:       fuego.NewStream(c),
			want:      fuego.EntrySlice{},
			mapper:    getFunction(),
			predicate: AllowPredicate,
		},
		{
			name:      "Should return empty EntrySlice",
			arg:       fuego.NewStream(c),
			want:      fuego.EntrySlice{},
			mapper:    getFunction(),
			predicate: AllowPredicate,
		},
		{
			name:      "Should return one element EntrySlice",
			arg:       fuego.NewStreamFromSlice(data, 0),
			want:      fuego.EntrySlice{fuego.EntryInt(1)},
			mapper:    getFunction(),
			predicate: AllowPredicate,
		},
		{
			name:      "Should return empty EntrySlice",
			arg:       fuego.NewStreamFromSlice(data, 0),
			want:      fuego.EntrySlice{},
			mapper:    getFunction(),
			predicate: BlockPredicate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.arg.
				Map(tt.mapper).
				Filter(tt.predicate).
				Collect(fuego.ToEntrySlice())
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func concatenateStringsBiFunc(i, j fuego.Entry) fuego.Entry {
	iStr := i.(fuego.EntryString)
	jStr := j.(fuego.EntryString)
	return iStr + "-" + jStr
}

func Test_Reduce(t *testing.T) {
	data := []fuego.Entry{
		fuego.EntryString("5"),
		fuego.EntryString("2"),
		fuego.EntryString("1"),
		fuego.EntryString("4"),
		fuego.EntryString("3")}

	res := fuego.NewStreamFromSlice(data, 0).
		Reduce(concatenateStringsBiFunc)
	assert.EqualValues(t, "5-2-1-4-3", res)
}

func Test_AllMatch(t *testing.T) {
	data := []fuego.Entry{
		fuego.EntryString("5*@"),
		fuego.EntryString("2@"),
		fuego.EntryString("@"),
		fuego.EntryString(" @")}

	res := fuego.NewStreamFromSlice(data, 0).
		AllMatch(func(t fuego.Entry) bool {
			return strings.Contains(string(t.(fuego.EntryString)), "@")
		})
	assert.True(t, res)
}
func Test_FlatMap(t *testing.T) {
	a := fuego.EntrySlice{fuego.EntryInt(1), fuego.EntryInt(2), fuego.EntryInt(3)}
	b := fuego.EntrySlice{fuego.EntryInt(4), fuego.EntryInt(5)}
	c := fuego.EntrySlice{fuego.EntryInt(6), fuego.EntryInt(7), fuego.EntryInt(8)}

	sliceOfEntrySlicesOfEntryInts := fuego.EntrySlice{a, b, c}

	sliceOfEntryInts := fuego.NewStreamFromSlice(sliceOfEntrySlicesOfEntryInts, 0).
		FlatMap(fuego.FlattenEntrySliceToEntry(0)).
		Collect(fuego.ToEntrySlice())

	assert.EqualValues(t, fuego.EntrySlice{fuego.EntryInt(1),fuego.EntryInt(2),fuego.EntryInt(3),fuego.EntryInt(4),fuego.EntryInt(5),fuego.EntryInt(6),fuego.EntryInt(7),fuego.EntryInt(8)}, sliceOfEntryInts)
}
