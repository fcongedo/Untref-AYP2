package ejercicios

import (
	"testing"
	"untref-ayp2/guia-pilas-colas/queue"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvertirCadena(t *testing.T) {
	tests := []struct {
		want string
		arg  string
	}{
		{
			want: "",
			arg:  "",
		},
		{
			want: "dcba",
			arg:  "abcd",
		},
		{
			want: "edcba",
			arg:  "abcde",
		},
	}

	for _, test := range tests {
		got := InvertirCadena(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestPalindromo(t *testing.T) {
	tests := []struct {
		want bool
		arg  string
	}{
		{
			want: true,
			arg:  "",
		},
		{
			want: true,
			arg:  "1456541",
		},
		{
			want: false,
			arg:  "145654",
		},
		{
			want: true,
			arg:  "145541",
		},
	}

	for _, test := range tests {
		got := Palindromo(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestBalanceada(t *testing.T) {
	tests := []struct {
		want bool
		arg  string
	}{
		{
			want: true,
			arg:  "",
		},
		{
			want: true,
			arg:  "[()]{}{[()()]()}",
		},
		{
			want: false,
			arg:  "[(])",
		},
	}

	for _, test := range tests {
		got := Balanceada(test.arg)
		assert.Equal(t, test.want, got)
	}
}

func TestUnirColas(t *testing.T) {
	q1 := queue.NewQueue[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	q2 := queue.NewQueue[int]()
	q2.Enqueue(5)
	q2.Enqueue(7)

	q12 := UnirColas(q1, q2)
	require.NotNil(t, q12, "La cola no puede ser nula")

	value, _ := q12.Dequeue()
	assert.Equal(t, 1, value)

	value, _ = q12.Dequeue()
	assert.Equal(t, 2, value)

	value, _ = q12.Dequeue()
	assert.Equal(t, 3, value)

	value, _ = q12.Dequeue()
	assert.Equal(t, 5, value)

	value, _ = q12.Dequeue()
	assert.Equal(t, 7, value)

	assert.True(t, q12.IsEmpty())
}

func TestUnirColasVacias(t *testing.T) {
	q1 := queue.NewQueue[int]()
	q2 := queue.NewQueue[int]()

	q12 := UnirColas(q1, q2)
	require.NotNil(t, q12, "La cola no puede ser nula")

	assert.True(t, q12.IsEmpty())
}

func TestUnirColaVacia(t *testing.T) {
	q1 := queue.NewQueue[int]()
	q1.Enqueue(1)

	q2 := queue.NewQueue[int]()

	q12 := UnirColas(q1, q2)
	require.NotNil(t, q12, "La cola no puede ser nula")

	value, _ := q12.Dequeue()
	assert.Equal(t, 1, value)

	assert.True(t, q12.IsEmpty())
}
