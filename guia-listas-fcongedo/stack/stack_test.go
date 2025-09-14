package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPushPop(t *testing.T) {
	s := NewStack[int]()
	require.NotNil(t, s)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.False(t, s.IsEmpty())

	v, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, v)

	v, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	v, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	_, err = s.Pop()
	assert.Error(t, err)
}

func TestEmptyStack(t *testing.T) {
	s := NewStack[int]()
	require.NotNil(t, s)

	assert.True(t, s.IsEmpty())

	_, err := s.Pop()
	assert.Error(t, err)

	_, err = s.Top()
	assert.Error(t, err)
}

func TestTop(t *testing.T) {
	s := NewStack[int]()
	require.NotNil(t, s)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Top()
	assert.NoError(t, err)
	assert.Equal(t, 3, v)

	s.Pop()

	v, err = s.Top()
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	s.Pop()

	v, err = s.Top()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)
}
