package lista_test

import (
	//"fmt"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

//Lista está vacia

// Insertar un elemento mediante el iterador en la posición en la que se crea, efectivamente lo añade primero y es equivalente a InsertarPrimero.

func InsertarAlInicio(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[string]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	iteradorL2 := lista2.Iterador()

	elem := "prueba"

	lista1.InsertarPrimero(elem)
	iteradorL2.Insertar(elem)

	require.Equal(t, lista1.VerPrimero(), lista2.VerPrimero())
}

// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.

func InsertarAlFinal(t *testing.T) {
	lista1 := TDALista.CrearListaEnlazada[string]()
	lista2 := TDALista.CrearListaEnlazada[string]()
	iteradorL2 := lista2.Iterador()

	elementosBase := [5]string{"hola","como","estas","todo","bien"}
	elementoNuevo := "Messi"

	for i := range(elementosBase){
		lista1.InsertarUltimo(elementosBase[i])
		lista2.InsertarUltimo(elementosBase[i])
	}

	for iteradorL2.HaySiguiente() {
		iteradorL2.Siguiente()
	}

	iteradorL2.Insertar(elementoNuevo)
	lista1.InsertarUltimo(elementoNuevo)

	require.Equal(t, lista1.VerUltimo(), lista2.VerUltimo())
}

// Insertar un elemento en el medio se hace en la posición correcta. Por definición, el elemento que se encontraba en el medio antes debe ser el siguiente del nuevo que se inserta. 

func InsertarEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iteradorL := lista.Iterador()
	elementosBase := [10]int{0,1,2,3,4,5,6,7,8,9}
	elementoNuevo := 1000

	for i := range elementosBase{
		lista.InsertarUltimo(elementosBase[i])
	}

	medio := (lista.Largo()/2)

	for i:= 0 ; i < medio ; i++{
		iteradorL.Siguiente()
	}

	iteradorL.Insertar(elementoNuevo)
	iteradorL.Siguiente()

	require.Equal(t, elementosBase[medio-1], iteradorL.VerActual())
}

// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func BorrarUltimoElemento(t *testing.T) {

}

// Remover el último elemento con el iterador cambia el último de la lista.
func CambiaUltimoElementoAlRemover(t *testing.T) {

}

// Verifica que al remover un elemento del medio, este no esté.
func NoExisteMedioAlRemoverlo(t *testing.T) {

}
