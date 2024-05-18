package cola_prioridad

const (
	CAPACIDAD_INICIAL int = 10
	CANT_INICIAL      int = 0
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {

	nuevo := make([]T, CAPACIDAD_INICIAL)
	return &colaConPrioridad[T]{datos: nuevo, cant: CANT_INICIAL, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == CANT_INICIAL
}

func (heap *colaConPrioridad[T]) VerMax() T {

	if heap.cant == CANT_INICIAL {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Encolar(dato T) {

}

func (heap *colaConPrioridad[T]) Desencolar() T {

	return heap.datos[0]
}
