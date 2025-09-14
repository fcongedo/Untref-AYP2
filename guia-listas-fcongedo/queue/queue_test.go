package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	require.NotNil(t, q)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.False(t, q.IsEmpty())

	v, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	v, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	v, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, v)

	_, err = q.Dequeue()
	assert.Error(t, err)
}

func TestEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	require.NotNil(t, q)

	assert.True(t, q.IsEmpty())

	_, err := q.Dequeue()
	assert.Error(t, err)

	_, err = q.Front()
	assert.Error(t, err)
}

func TestFront(t *testing.T) {
	q := NewQueue[int]()
	require.NotNil(t, q)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, err := q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	q.Dequeue()

	v, err = q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	q.Dequeue()

	v, err = q.Front()
	assert.NoError(t, err)
	assert.Equal(t, 3, v)
}
