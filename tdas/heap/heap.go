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

	heap := new(colaConPrioridad[T])
	heap.datos = arreglo
	heap.cant = len(arreglo)
	heap.cmp = funcion_cmp
	heap.heapify(heap.datos)

	return heap
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
	heap.upHeap(heap.cant - 1)
	cap := len(heap.datos)
	if heap.cant == cap {
		heap.redimensionar(heap.cant * VALOR_REDIMENSION)
	}

}

func (heap *colaConPrioridad[T]) Desencolar() T {

	if heap.cant == 0 {
		panic("La cola esta vacia")
	}
	dato := heap.datos[0]
	cap := len(heap.datos)
	if heap.cant*COEF_REDIMENSION <= cap && cap > CAPACIDAD_INICIAL {
		heap.redimensionar(cap / VALOR_REDIMENSION)
	}
	heap.datos[0] = heap.datos[heap.cant-1]
	heap.downHeap(0)
	heap.cant--
	return dato
}

func (heap *colaConPrioridad[T]) swap(dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}

func (heap *colaConPrioridad[T]) downHeap(i int) {
	padre := &heap.datos[i]
	hijoIzq, hijoDer, tieneHijos := obtenerHijos(i, heap.datos)

	mayor := obtenerMayor(heap.datos, heap.cmp, padre, hijoIzq, hijoDer)

	if !tieneHijos || mayor == padre {
		return
	}

	heap.swap(mayor, padre)

	if mayor == hijoIzq {
		heap.downHeap((2 * i) + 1)
	} else if mayor == hijoDer {
		heap.downHeap((2 * i) + 2)
	}

}

func (heap *colaConPrioridad[T]) upHeap(i int) {

	padre, iPadre, tienePadre := obtenerPadre(i,heap.datos)

	if !tienePadre || heap.cmp(heap.datos[i], *padre) < 0 {
		return
	}

	heap.swap(&heap.datos[i], padre)
	heap.upHeap(iPadre)
}

func (heap *colaConPrioridad[T]) heapify(arr []T) {

	for i := len(arr) - 1; i > 0; i-- {
		heap.downHeap(i)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	heap := new(colaConPrioridad[T])
	heap.datos = elementos
	heap.cant = len(elementos)
	heap.cmp = funcion_cmp
	heap.heapify(heap.datos)

	prim := heap.datos[0]
	ult := heap.datos[heap.cant-1]
	heap.swap(&prim, &ult)
	
}

func obtenerHijos[T any](i int, arr []T) (*T, *T, bool) {
	hijoIzq := (2 * i) + 1
	hijoDer := (2 * i) + 2

	if hijoIzq > len(arr) || hijoDer > len(arr) {
		return nil, nil, false
	}

	return &arr[hijoIzq], &arr[hijoDer], true
}

func obtenerPadre[T any](i int, arr []T) (*T, int, bool) {

	padre := (i - 1) / 2

	if padre < 0 {
		return &arr[0], 0, false
	}

	return &arr[padre], padre, true
}

func obtenerMayor[T any](arr []T, func_cmp func(T, T) int , padre, der, izq *T) *T {

    max := padre

    if func_cmp(*padre, *der) < 0 && func_cmp(*der, *izq) > 0{
        max = der
    } else if func_cmp(*padre, *izq) < 0 && func_cmp(*izq, *der) > 0{
        max = izq
	}
    return max
}

func (heap *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}
