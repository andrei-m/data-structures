package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LinkedList_PushPop(t *testing.T) {
	t.Run("push then pop for ints", func(t *testing.T) {
		ll := LinkedList[int]{}
		ll.Push(1)
		ll.Push(2)
		ll.Push(3)
		val, err := ll.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 3, val)
		val, err = ll.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 2, val)
		val, err = ll.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)
	})

	t.Run("push then pop for strings", func(t *testing.T) {
		ll := LinkedList[string]{}
		ll.Push("foo")
		ll.Push("bar")
		val, err := ll.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "bar", val)
		val, err = ll.Pop()
		assert.NoError(t, err)
		assert.Equal(t, "foo", val)
	})

	t.Run("pop from empty list", func(t *testing.T) {
		ll := LinkedList[int]{}
		_, err := ll.Pop()
		assert.Equal(t, errPopFromEmptyList, err)
	})
}

func Test_LinkedList_Size(t *testing.T) {
	ll := LinkedList[int]{}
	assert.Zero(t, ll.Size())
	ll.Push(0)
	assert.Equal(t, 1, ll.Size())
	ll.Push(1)
	assert.Equal(t, 2, ll.Size())
}
