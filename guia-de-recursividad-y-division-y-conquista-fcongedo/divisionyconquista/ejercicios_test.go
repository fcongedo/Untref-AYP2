package divisionyconquista

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusquedaBinaria(t *testing.T) {
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 1))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 2))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 3))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 4))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 6))
}

func TestBusquedaTernariaRecursiva(t *testing.T) {
	assert.Equal(t, 8, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
	assert.Equal(t, 0, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))
	assert.Equal(t, 4, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	assert.Equal(t, -1, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10))
}

func TestPico(t *testing.T) {
	assert.Equal(t, 4, Pico([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 8, Pico([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, 0, Pico([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestSubsecuenciaSumaMaxima(t *testing.T) {
	arreglo1 := []int{-1, 6, -2, 5, -1, 4, 3, -4, 3}
	sumaMaxima1, posInicial1, posFinal1 := SubsecuenciaSumaMaxima(arreglo1)
	assert.Equal(t, 15, sumaMaxima1)
	assert.Equal(t, 1, posInicial1)
	assert.Equal(t, 6, posFinal1)

	arreglo2 := []int{2, -1, 3, 4, -2, 5, -1, 6, -3, 7}
	sumaMaxima2, posInicial2, posFinal2 := SubsecuenciaSumaMaxima(arreglo2)
	assert.Equal(t, 20, sumaMaxima2)
	assert.Equal(t, 0, posInicial2)
	assert.Equal(t, 9, posFinal2)

	arreglo3 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sumaMaxima3, posInicial3, posFinal3 := SubsecuenciaSumaMaxima(arreglo3)
	assert.Equal(t, 6, sumaMaxima3)
	assert.Equal(t, 3, posInicial3)
	assert.Equal(t, 6, posFinal3)
}

func TestEsArreloMagico(t *testing.T) {
	assert.True(t, EsArregloMagico([]int{0}))
	assert.True(t, EsArregloMagico([]int{0, 2, 3, 4}))
	assert.True(t, EsArregloMagico([]int{-1, 0, 2, 5, 6}))
	assert.True(t, EsArregloMagico([]int{-2, -1, 0, 1, 4}))
	assert.True(t, EsArregloMagico([]int{-5, -3, -1, 0, 4}))

	assert.False(t, EsArregloMagico([]int{}))
	assert.False(t, EsArregloMagico([]int{1}))
	assert.False(t, EsArregloMagico([]int{1, 2, 3, 4, 5}))
	assert.False(t, EsArregloMagico([]int{1, 3, 5, 7, 9}))
	assert.False(t, EsArregloMagico([]int{-5, -3, 1, 2, 3}))
	assert.False(t, EsArregloMagico([]int{-10, -5, -2, 0, 1}))
}
