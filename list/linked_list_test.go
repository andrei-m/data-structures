package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LinkedList_PushPop(t *testing.T) {
	t.Run("push them pop", func(t *testing.T) {
		ll := LinkedList{}
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

	t.Run("pop from empty list", func(t *testing.T) {
		ll := LinkedList{}
		_, err := ll.Pop()
		assert.Equal(t, errPopFromEmptyList, err)
	})
}
