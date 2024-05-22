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
	heapify(heap.datos, heap.cmp)

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
	upHeap(heap.cant - 1,heap.datos, heap.cmp)
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
	downHeap(0,heap.datos, heap.cmp)
	heap.cant--
	return dato
}

func swap[T any](dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}

func downHeap[T any](i int, arr []T, func_cmp func(T, T) int) {
	padre := arr[i]
	hijoIzq, hijoDer, tieneHijos := obtenerHijos(i, arr)

	iMayor := obtenerMayor(arr, i, func_cmp)
	mayor := arr[iMayor]

	if !tieneHijos || func_cmp(mayor, padre) == 0 {
		return
	}

	swap(&mayor, &padre)

	if func_cmp(mayor, *hijoIzq) == 0 {

		downHeap((2 * i) + 1, arr, func_cmp)
	
	} else if func_cmp(mayor, *hijoDer) == 0  {
		downHeap((2 * i) + 2, arr, func_cmp)
	}

}

func upHeap[T any](i int, arr []T, func_cmp func(T, T) int) {

	padre, iPadre, tienePadre := obtenerPadre(i,arr)

	if !tienePadre || func_cmp(arr[i], *padre) < 0 {
		return
	}

	swap(&arr[i], padre)
	upHeap(iPadre,arr,func_cmp)
}

func heapify[T any](arr []T, func_cmp func(T, T) int) {

	for i := len(arr) - 1; i > 0; i-- {
		downHeap(i,arr,func_cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	if len(elementos) == 0 || len(elementos) == 1{
		return
	}

	heapify(elementos,funcion_cmp)
	prim := elementos[0]
	ult := elementos[len(elementos) - 1]
	swap(&prim, &ult)
	downHeap(0, elementos, funcion_cmp)

	
	elementosActualizado := elementos[:len(elementos)-2]
	HeapSort(elementosActualizado,funcion_cmp)
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

func obtenerMayor[T any](arr []T, iPadre int, func_cmp func(T,T) int) int {

	max := iPadre
	iIzq := (2*iPadre) + 1
	iDer := (2*iPadre) + 2

	if func_cmp(arr[iPadre], arr[iDer]) < 0 && func_cmp(arr[iDer], arr[iIzq]) > 0{
		max = iDer
	} else if func_cmp(arr[iPadre], arr[iIzq]) < 0 && func_cmp(arr[iIzq], arr[iDer]) > 0 {
		max = iIzq
	}

	return max
}

func (heap *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}
