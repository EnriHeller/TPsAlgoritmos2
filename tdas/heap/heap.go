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
	//Aca habrá que hacer downheap del ultimo al primero con cada elemento del arreglo
	nuevo := make([]T, CAPACIDAD_INICIAL)
	return &colaConPrioridad[T]{datos: nuevo, cant: CANT_INICIAL, cmp: funcion_cmp}
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

	/*heapify(heap.datos, heap.cmp)
	cap := len(heap.datos)
	if heap.cant == cap {
		heap.redimensionar(heap.cant * VALOR_REDIMENSION)
	}*/

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
	heap.cant--
	//downHeap()

	return dato
}

func (heap *colaConPrioridad[T]) swap(dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}


func (heap *colaConPrioridad[T]) downHeap(i int) {
	padre := &heap.datos[i]
	hijoIzq, hijoDer, tieneHijos := heap.obtenerHijos(i)

	mayor := heap.obtenerMayor(i)

	if !tieneHijos || mayor == padre{
		return
	}

	heap.swap(mayor,padre)

	if mayor == hijoIzq{
		heap.downHeap((2*i)+1)
	}else if mayor == hijoDer{
		heap.downHeap((2*i)+2)
	}
	

}

func (heap *colaConPrioridad[T]) upHeap(i int) {

	padre, iPadre, tienePadre := heap.obtenerPadre(i)

	if !tienePadre || heap.cmp(heap.datos[i], *padre) < 0 {
		return
	}

	heap.swap(&heap.datos[i], padre)
	heap.upHeap(iPadre)
}

func (heap *colaConPrioridad[T]) obtenerHijos(i int) (*T, *T, bool) {
	hijoIzq := (2 * i) + 1
	hijoDer := (2 * i) + 2

	if hijoIzq > len(heap.datos) || hijoDer > len(heap.datos) {
		return nil, nil, false
	}

	return &heap.datos[hijoIzq], &heap.datos[hijoDer], true
}

func (heap *colaConPrioridad[T]) obtenerPadre(i int) (*T, int, bool) {

	padre := (i - 1) / 2

	if padre < 0 {
		return &heap.datos[0], 0, false
	}

	return &heap.datos[padre], padre, true
}

func (heap *colaConPrioridad[T]) obtenerMayor(iPadre int) *T {
	padre := heap.datos[iPadre]
	max := padre
	hijoIzq, hijoDer, tieneHijos := heap.obtenerHijos(iPadre)

	if !tieneHijos {
		return &max
	}

	if heap.cmp(padre, *hijoDer) < 0 && heap.cmp(*hijoIzq, *hijoDer) < 0 {
		max = *hijoDer
	} else if heap.cmp(padre, *hijoIzq) < 0 && heap.cmp(*hijoDer, *hijoIzq) < 0 {
		max = *hijoIzq
	}

	return &max
}

func (heap *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}
