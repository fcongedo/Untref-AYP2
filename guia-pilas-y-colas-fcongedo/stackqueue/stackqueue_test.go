package stackqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackQueue(t *testing.T) {
	q := NewStackQueue[int]()

	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
}

func TestStackQueueEnqueue(t *testing.T) {
	q := NewStackQueue[int]()

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestStackQueueDequeue(t *testing.T) {
	q := NewStackQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)

	v, _ := q.Dequeue()
	assert.Equal(t, 1, v)

	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)

	assert.True(t, q.IsEmpty())
}

func TestStackQueueDequeueOnEmptyQueue(t *testing.T) {
	q := NewStackQueue[int]()

	_, err := q.Dequeue()
	assert.EqualError(t, err, "cola vacía")
}

func TestStackQueueFrontOnEmptyQueue(t *testing.T) {
	q := NewStackQueue[int]()

	_, err := q.Front()
	assert.EqualError(t, err, "cola vacía")
}

func TestStackQueueFront(t *testing.T) {
	q := NewStackQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Front()
	assert.Equal(t, 1, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 2, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 3, v)
}
