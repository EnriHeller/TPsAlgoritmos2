package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TEST_VOLUMEN int = 10000
)
var arregloOrdenadoHeapsort = []int{0,3,5,8,11,17,23,54}

var arregloOrdenado = []int{54, 23, 17, 11, 8, 5, 3, 0}
var arregloDesordenado = []int{3, 5, 0, 8, 11, 23, 54, 17}

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

	for _, i := range arregloDesordenado {
		heap.Encolar(i)
	}

	require.Equal(t, 8, heap.Cantidad())

	for _, elem := range arregloOrdenado {
		require.Equal(t, elem, heap.Desencolar())
	}
}

func TestHeapDeCadenas(t *testing.T) {

	t.Log("Pruebas primitivas de heap de strings")
	heap := TDAHeap.CrearHeap(compararCadenas)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar("Algoritmos")
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "Algoritmos", heap.VerMax())

	heap.Encolar("Estructuras")
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, "Estructuras", heap.VerMax())

	heap.Encolar("Hola")
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, "Hola", heap.VerMax())

	heap.Encolar("Mundo")
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, "Mundo", heap.VerMax())

	heap.Encolar("Datos")
	require.Equal(t, 5, heap.Cantidad())
	require.Equal(t, "Mundo", heap.VerMax())

	require.Equal(t, "Mundo", heap.Desencolar())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, "Hola", heap.VerMax())
	require.Equal(t, "Hola", heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, "Estructuras", heap.VerMax())
	require.Equal(t, "Estructuras", heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, "Datos", heap.VerMax())
	require.Equal(t, "Datos", heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "Algoritmos", heap.VerMax())
	require.Equal(t, "Algoritmos", heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestVolumen(t *testing.T) {

	t.Log("Pruebas de volumen de heap")

	heap := TDAHeap.CrearHeap(compararEnteros)

	for i := 0; i < TEST_VOLUMEN; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i+1, heap.Cantidad())
	}

	for i := TEST_VOLUMEN; i > 0; i-- {
		require.Equal(t, i-1, heap.VerMax())
		require.Equal(t, i-1, heap.Desencolar())
		require.Equal(t, i-1, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
}

func TestCrearHeapArregloVacio(t *testing.T) {
	t.Log("Prueba crea un heap con un arreglo vacio")
	arr := []int{}
	heap := TDAHeap.CrearHeapArr(arr, compararEnteros)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola está vacía", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(9)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 9, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
}

func TestCrearHeapArreglo(t *testing.T) {
	t.Log("Prueba crea un heap con un arreglo de enteros")

	heap := TDAHeap.CrearHeapArr(arregloDesordenado, compararEnteros)

	for _,elem := range arregloOrdenado{
		require.Equal(t, elem, heap.Desencolar())
	}

	require.True(t, !heap.EstaVacia())
}

func TestHeapsort(t *testing.T) {
	test := arregloDesordenado
	TDAHeap.HeapSort(test, compararEnteros)
	require.Equal(t, arregloOrdenadoHeapsort, test)
}
