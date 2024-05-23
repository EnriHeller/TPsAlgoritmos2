package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararCadenas(cad1, cad2 string) int {
	if cad1 < cad2 {
		return -1
	} else if cad1 > cad2 {
		return 1
	}
	return 0
}

func compararEnteros(a, b int) int {
	return a - b
}

func TestColaPrioridadVacia(t *testing.T) {

	t.Log("Prueba que al crear una cola vacia, se comporte como corresponde")
	heap := TDAHeap.CrearHeap(compararEnteros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.Desencolar() })
	require.Equal(t, 0, heap.Cantidad())

}

func TestEncolar(t *testing.T) {

	heap := TDAHeap.CrearHeap(compararEnteros)

	arreglo := []int{3, 5, 0, 8, 11, 23, 54, 17}
	for _, i := range arreglo {
		heap.Encolar(i)
	}
	require.False(t, heap.EstaVacia())
	require.Equal(t, 8, heap.Cantidad())
	require.Equal(t, 54, heap.VerMax())
}

func TestDesencolar(t *testing.T) {

	heap := TDAHeap.CrearHeap(compararEnteros)

	arreglo := []int{3, 5, 0, 8, 11, 23, 54, 17}
	for _, i := range arreglo {
		heap.Encolar(i)
	}

	require.Equal(t, 8, heap.Cantidad())
	require.Equal(t, 54, heap.Desencolar())
	require.Equal(t, 23, heap.Desencolar())
	require.Equal(t, 17, heap.Desencolar())
	require.Equal(t, 11, heap.Desencolar())
	require.Equal(t, 8, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 0, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
}
