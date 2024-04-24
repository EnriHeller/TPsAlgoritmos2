package lista

type Lista[T any] interface {

	//Si la lista se encuentra vacia, devuelve verdadero. En caso contrario, devuelve falso
	EstaVacia() bool

	//Inserta un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	//Inserta un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	//Borra el primer elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia"
	BorrarPrimero() T

	//Obtiene el valor del primer elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia".
	VerPrimero() T

	//Obtiene el valor del último elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia".
	VerUltimo() T

	//Devuelve la cantidad de elementos alojados en la lista.
	Largo() int

	//Recorre la lista, aplicandole a cada elemento la función pasada por parámetro hasta encontrar una condición de corte. Si la función devuelve false, significa que se encontro una condición de corte. 
	Iterar(visitar func(T) bool)

	//Permite acceder a un iterador externo, asociado a la lista actual.
	Iterador() IteradorLista[T]
}