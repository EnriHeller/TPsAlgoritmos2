package cola_prioridad

const (
	CAPACIDAD_INICIAL int = 10
	CANT_INICIAL      int = 0
	COEF_REDIMENSION  int = 4
	VALOR_REDIMENSION int = 2
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

	heap.datos[heap.cant] = dato
	heap.cant++
	heapity(heap.datos, heap.cmp)
	cap := len(heap.datos)
	if heap.cant == cap {
		heap.redimensionar(heap.cant * VALOR_REDIMENSION)
	}
	upHeap(...)
}

func (heap *colaConPrioridad[T]) Desencolar() T {

	if heap.cant == 0 {
		panic("La cola esta vacia")
	}
	dato := heap.datos[0]
	cap := len(heap.datos)
	if heap.cant * COEF_REDIMENSION <= cap && cap > CAPACIDAD_INICIAL {
		heap.redimensionar(cap / VALOR_REDIMENSION)
	}
	heap.cant--
	downHeap()

	return dato
}


func swap(dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato, *dato1
}


func downHeap[T any]() {

}

func upHeap[T any](){

}


func obtenerMayor(arr []T, padre, der, izq T, func_cmp func(T,T)int) int {

	max := padre

	if func_cmp(arr[padre], arr[der]) < 0 && arr[der] > arr[izq]{
		max = der
	} else if func_cmp(arr[padre], arr[izq]) < 0 && arr[izq] > arr[der] [
		max = izq
	]
	return max
}

func (heap *colaConPrioridad[T])redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nueva, heap.datos)
	heap.datos = nueva
}