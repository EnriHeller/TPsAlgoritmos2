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

}

func (iterador *iterListaEnlazada[T]) Insertar(dato T){

}

func (iterador *iterListaEnlazada[T]) Borrar() T{
	var dato T
	return dato
}





