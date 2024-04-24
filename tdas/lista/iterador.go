package lista

type iterListaEnlazada[T any] struct {
    actual *nodoLista[T]
	anterior *nodoLista[T]
	siguiente *nodoLista[T]
}

func CrearIterador[T any]() IteradorLista[T]{
	return new(iterListaEnlazada[T])
}

func (iterador *iterListaEnlazada[T]) VerActual() T{
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool{
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente(){
	iterador.anterior = iterador.actual
	iterador.actual = iterador.siguiente
	iterador.siguiente = iterador.siguiente.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T){
	nuevoNodo := crearNodo(dato)

	//la lista esta vacia
	if iterador.actual == nil{
		iterador.actual = &nuevoNodo
	}

	iterador.anterior.siguiente = &nuevoNodo
	nuevoNodo.siguiente = iterador.actual
	iterador.actual = &nuevoNodo
}

func (iterador *iterListaEnlazada[T]) Borrar() T{

	datoBorrado := iterador.actual.dato

	iterador.anterior.siguiente = iterador.siguiente
	iterador.actual = iterador.siguiente
	iterador.siguiente = iterador.siguiente.siguiente

	return datoBorrado
}





