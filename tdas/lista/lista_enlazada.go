package lista


type nodoLista[T any] struct{
	dato T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct{
	primero *nodoLista[T]
	ultimo *nodoLista[T]
	largo int
}

func CrearListaEnlazada[T any]() Lista[T] {
	nuevaLista := new(listaEnlazada[T])

	return nuevaLista
}

func crearNodo[T any](dato T) nodoLista[T]{

	return nodoLista[T]{dato: dato}
}

func (lista listaEnlazada[T]) EstaVacia() bool{
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia(){
		lista.primero = &nuevoNodo
		lista.ultimo = &nuevoNodo
	}else{
		nuevoNodo.siguiente = lista.primero
		lista.primero = &nuevoNodo
	}
	lista.largo ++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia(){
		lista.primero = &nuevoNodo
		lista.ultimo = &nuevoNodo
	}else{
		lista.ultimo.siguiente = &nuevoNodo
		lista.ultimo = &nuevoNodo
	}
	lista.largo ++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}else{
		dato := lista.primero.dato
		lista.primero = lista.primero.siguiente
		return dato
	}

}

func (lista *listaEnlazada[T]) VerPrimero() T {

	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}else{
		return lista.primero.dato
	}
	
}

func (lista *listaEnlazada[T]) VerUltimo() T {

	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}else{
		return lista.ultimo.dato
	}
}


func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}


func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool){

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T]{
	iterador := CrearIterador[T]()
	return iterador
}

