package diccionario

import (
	"fmt"
	"math"
)

const (
	COEFICIENTE_REDIMENSION float64 = 0.7
	FACTOR_REDIMENSION      int     = 2
	VACIO                   int     = 0
	OCUPADO                 int     = 1
	BORRADO                 int     = 2

	TAMANO int = 7
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int // Solo hace referencia a ocupados
	tam      int
	borrados int
}

type iterDiccionario[K comparable, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

func crearCeldaHash[K comparable, V any]() celdaHash[K, V] {
	nuevaCelda := new(celdaHash[K, V])
	nuevaCelda.estado = 0
	return *nuevaCelda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	nuevoHash := new(hashCerrado[K, V])
	nuevoHash.tam = TAMANO
	nuevoHash.cantidad = 0
	nuevoHash.borrados = 0

	nuevoHash.tabla = crearTabla[K, V](nuevoHash.tam)

	return nuevoHash
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	nuevaTabla := make([]celdaHash[K, V], capacidad)
	for i := 0; i < capacidad; i++ {
		nuevaCelda := crearCeldaHash[K, V]()
		nuevaTabla[i] = nuevaCelda
	}
	return nuevaTabla
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	factorCarga := float64((hash.cantidad + hash.borrados)) / float64(hash.tam)
	if factorCarga >= COEFICIENTE_REDIMENSION {
		hash.redimensionar(FACTOR_REDIMENSION * hash.tam)
	}
	posicion, err := hash.buscar(clave)

	if err != nil { // La llave no esta presente, guarda una llave nueva
		hash.tabla[posicion].clave = clave
		hash.tabla[posicion].estado = OCUPADO
		hash.cantidad++
		// La llave esta borrada y la guardamos nuevamente
	} else if hash.tabla[posicion].estado == BORRADO && hash.tabla[posicion].clave == clave {
		hash.tabla[posicion].estado = OCUPADO
		hash.cantidad++
		hash.borrados--
	}
	hash.tabla[posicion].dato = dato // Guardamos el dato para cualquiera de los casos

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {

	_, err := hash.buscar(clave)
	return err == nil
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos, err := hash.buscar(clave)
	if err != nil {
		panic(err.Error())
	}
	return hash.tabla[pos].dato

}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion, err := hash.buscar(clave)
	if err != nil {
		panic(err.Error())
	}
	elemento := &hash.tabla[posicion]
	elemento.estado = BORRADO
	hash.cantidad--
	hash.borrados++
	return elemento.dato
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	for _, elem := range hash.tabla {

		if elem.estado == VACIO || elem.estado == BORRADO {
			continue
		}
		if !visitar(elem.clave, elem.dato) {
			break
		}

	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

	for ind, elemento := range hash.tabla {
		if elemento.estado == OCUPADO {
			return &iterDiccionario[K, V]{hash: hash, pos: ind}
		}
	}

	return &iterDiccionario[K, V]{hash: hash, pos: hash.tam}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {

	return iter.pos != iter.hash.tam
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.hash.tabla[iter.pos].clave
	dato := iter.hash.tabla[iter.pos].dato
	return clave, dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.pos++

	for iter.pos < iter.hash.tam {
		if iter.hash.tabla[iter.pos].estado == OCUPADO {
			return
		}
		iter.pos++
	}

	iter.pos = iter.hash.tam
}

func (hash *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {

	tablaAnterior := hash.tabla
	hash.tabla = crearTabla[K, V](nuevaCapacidad)
	hash.tam = nuevaCapacidad
	hash.borrados = 0

	for _, elem := range tablaAnterior {
		if elem.estado == OCUPADO {
			K, V := elem.clave, elem.dato
			hash.Guardar(K, V)
		}
	}
}

func convertirABytes[K comparable](clave K) []byte {

	return []byte(fmt.Sprintf("%v", clave))
}

func sdbmHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func (hash *hashCerrado[K, V]) hashear(clave K) int {
	claveByte := convertirABytes(clave)
	hashing := sdbmHash(claveByte)
	resultado := int(math.Mod(float64(hashing), float64(hash.tam)))
	return resultado

}

func (hash hashCerrado[K, V]) buscar(clave K) (int, error) {
	posicion := hash.hashear(clave)

	primeraBusqueda, err := hash.buscarEnPorcion(posicion, hash.tam, clave)

	if err == nil {
		return primeraBusqueda, nil
	}

	segundaBusqueda, err2 := hash.buscarEnPorcion(0, posicion, clave)

	if err2 == nil {
		return segundaBusqueda, nil
	}

	return primeraBusqueda, fmt.Errorf("La clave no pertenece al diccionario")
}

func (hash *hashCerrado[K, V]) buscarEnPorcion(inicio, fin int, clave K) (int, error) {
	for i := inicio; i < fin; i++ {
		celdaActual := hash.tabla[i]
		if celdaActual.estado != VACIO && celdaActual.clave == clave {
			return i, nil
		} else if celdaActual.estado == VACIO {
			return i, fmt.Errorf("La clave no pertenece al diccionario")
		}
	}
	return fin - 1, fmt.Errorf("La clave no pertenece al diccionario")
}
