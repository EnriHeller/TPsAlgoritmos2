package lista

type Lista[T any] interface {

	//Indica a partir de un booleano, si la lista se encuentra vacia.
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int
	Iterar(visitar func(T) bool)
	Iterador() IteradorLista[T]
}