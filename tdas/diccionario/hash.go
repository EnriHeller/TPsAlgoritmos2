package diccionario

import (
	"fmt"
)

const (
	COEFICIENTE_REDIMENSION float64 = 0.7
	FACTOR_REDIMENSION      int     = 2
	VACIO                   int     = 0
	OCUPADO                 int     = 1
	BORRADO                 int     = 2

	TAMANO int = 6
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
	hash hashCerrado[K, V]
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

	nuevaTabla := make([]celdaHash[K, V], nuevoHash.tam)

	for i := 0; i < nuevoHash.tam; i++ {
		nuevaCelda := crearCeldaHash[K, V]()
		nuevaTabla = append(nuevaTabla, nuevaCelda)
	}

	nuevoHash.tabla = nuevaTabla

	return nuevoHash
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], capacidad)

}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	factorCarga := float64((hash.cantidad + hash.borrados)) / float64(hash.tam)
	if factorCarga >= COEFICIENTE_REDIMENSION {
		hash.redimensionar(FACTOR_REDIMENSION * hash.tam)
	}

	posicion := hash.buscar(clave)
	if hash.Pertenece(clave) {
		hash.tabla[posicion].dato = dato
	} else {
		hash.tabla[posicion].clave = clave
		hash.tabla[posicion].dato = dato
		hash.tabla[posicion].estado = OCUPADO
	}
	hash.cantidad++

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	if hash.buscar(clave) == -1 {
		return false
	}
	return true
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos := hash.buscar(clave)
	return hash.tabla[pos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	var dato V
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {

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

func (hash *hashCerrado[K, V]) buscar(clave K) int {
	posicion := hash.hashear(clave)
	primeraPorcion := hash.tabla[posicion:hash.tam]
	porcionAuxiliar := hash.tabla[:posicion]

	for i := posicion; i < hash.tam; i++ {
		celdaActual := primeraPorcion[i]
		if celdaActual.estado == OCUPADO && celdaActual.clave == clave {
			return posicion
		} else if celdaActual == nil {
			return -1
		} else {
			continue
		}
	}

	for _, celdaActual := range porcionAuxiliar {
		if celdaActual.estado == OCUPADO && celdaActual.clave == clave {
			return posicion
		} else if celdaActual.estado == VACIO {
			return -1
		} else {
			continue
		}
	}
	return -1
}

func (hash *hashCerrado[K, V]) hashear(clave K) int {
	claveByte := convertirABytes(clave)
	hashing := sdbmHash(claveByte)
	return int(hashing) % hash.tam
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

	return &iterDiccionario[K, V]{}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return true
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	var clave K
	var valor V
	return clave, valor
}

func (iter *iterDiccionario[K, V]) Siguiente() {

}

func (hash *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {

	nuevaTabla := make([]celdaHash[K, V], nuevaCapacidad)
	tablaAnterior := hash.tabla
	hash.tabla = nuevaTabla
	hash.tam = nuevaCapacidad

	for i := 0; i < nuevaCapacidad; i++ {
		hash.tabla[i].estado = VACIO
	}

	for _, elem := range tablaAnterior {
		if elem.estado == OCUPADO {
			K, V := elem.clave, elem.dato
			hash.Guardar(K, V)
		}
	}

}
