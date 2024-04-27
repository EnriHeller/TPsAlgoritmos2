package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodo[T any](dato T) nodoLista[T] {

	return nodoLista[T]{dato: dato}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = &nuevoNodo
		lista.ultimo = &nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = &nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = &nuevoNodo
	} else {
		lista.ultimo.siguiente = &nuevoNodo

	}
	lista.ultimo = &nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	} else {
		return lista.primero.dato
	}

}

func (lista *listaEnlazada[T]) VerUltimo() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	} else {
		return lista.ultimo.dato
	}
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

	actual := lista.primero
	iterarHastaFinal := true

	for actual != nil || !iterarHastaFinal {
		iterarHastaFinal = visitar(actual.dato)
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {

	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	//La lista esta vacia
	if iterador.anterior == nil && iterador.actual == nil {
		iterador.actual = &nuevoNodo
		iterador.lista.primero = &nuevoNodo
		iterador.lista.ultimo = &nuevoNodo
	}
	//Si tengo un nodo antes, el siguiente de este nodo es el nuevo nodo.
	if iterador.anterior != nil {
		iterador.anterior.siguiente = &nuevoNodo
	}
	//Actualizo el que estaba antes de insertar como actual al siguiente del que se va a añadir
	if iterador.actual != nil {
		nuevoNodo.siguiente = iterador.actual
	}
	//Si estoy parado en el ultimo nodo, tengo que actualizar la referencia al ultimo al momento de añadir el nuevo.
	if iterador.lista.ultimo == iterador.actual {
		iterador.lista.ultimo.siguiente = &nuevoNodo
		iterador.lista.ultimo = &nuevoNodo
	}
	//Apunto al primero con una lista que no está vacia, por lo que actualizo referencia al primer nodo
	if iterador.anterior == nil && iterador.HaySiguiente() {
		nuevoNodo.siguiente = iterador.lista.primero
		iterador.lista.primero = &nuevoNodo
	}
	iterador.actual = &nuevoNodo
	iterador.lista.largo++

}

func (iterador *iterListaEnlazada[T]) Borrar() T {

    if !iterador.HaySiguiente() {
        panic("El iterador termino de iterar")
    }
    datoBorrado := iterador.actual.dato

	//Estoy parado en el primer elemento
    if iterador.anterior == nil {
        if iterador.actual.siguiente == nil {
            iterador.lista.primero = nil
            iterador.lista.ultimo = nil
            iterador.actual = nil
        } else {
            iterador.actual = iterador.lista.primero.siguiente
			iterador.lista.primero = iterador.lista.primero.siguiente
        }
    }else{
        iterador.anterior.siguiente = iterador.actual.siguiente
		// Si el siguiente es nil (el iterador apunta al ultimo elemento), hacemos referencia al anterior
        if iterador.actual.siguiente == nil {
            iterador.actual = iterador.anterior 
            iterador.lista.ultimo = iterador.anterior
        } else {
            iterador.actual = iterador.actual.siguiente
        }
    }
    iterador.lista.largo--
    return datoBorrado
}
