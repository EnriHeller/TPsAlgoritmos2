package diccionario

import (
	//"fmt"
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterABB[K comparable, V any] struct {
	desde *K
	hasta *K
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	arbol *abb[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {

	nuevoArbol := new(abb[K, V])

	nuevoArbol.cantidad = 0
	nuevoArbol.cmp = funcion_cmp

	return nuevoArbol
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {

	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      valor,
		izquierdo: nil,
		derecho:   nil,
	}
}

func (arbol *abb[K, V]) buscar(padre *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if padre == nil {
		return padre
	}

	if arbol.cmp(padre.clave, clave) > 0 {
		// Va a la izquierda
		return arbol.buscar(padre.izquierdo, clave)
	} 
	
	if arbol.cmp(padre.clave, clave) < 0 { 
		// Va a la derecha
		return arbol.buscar(padre.derecho, clave)
	}

	return padre
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {

	if arbol.raiz == nil {
		arbol.raiz = crearNodo(clave, valor)
		arbol.cantidad++
		return
	}

	nuevoNodo := crearNodo(clave, valor)
	busqueda := arbol.buscar(arbol.raiz, clave)

	//No existe nodo con dicha clave
	if busqueda == nil {
		busqueda = nuevoNodo
		arbol.cantidad++
	}

	busqueda.dato = nuevoNodo.dato
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {

	busqueda := arbol.buscar(arbol.raiz, clave)
	return busqueda != nil
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Borrar(clave K) V {

	busqueda := arbol.buscar(arbol.raiz, clave)

	if busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	if busqueda.izquierdo == nil && busqueda.derecho == nil {
		busqueda = nil
	}

	if busqueda.izquierdo == nil && busqueda.derecho != nil {
		busqueda = busqueda.derecho
	}

	if busqueda.izquierdo != nil && busqueda.derecho == nil {
		busqueda = busqueda.izquierdo
	}
	if busqueda.izquierdo != nil && busqueda.derecho != nil {
		padre, reemplazante := buscarMasDerecho[K, V](busqueda.izquierdo)
		busqueda.clave, busqueda.dato = reemplazante.clave, reemplazante.dato
		padre.derecho = nil
	}

	arbol.cantidad--
	return busqueda.dato
}

func (arbol *abb[K, V]) Obtener(clave K) V {

	busqueda := arbol.buscar(arbol.raiz, clave)

	if busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	return busqueda.dato
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	arbol.iteradorInterno(arbol.raiz, visitar)
}

func (arbol *abb[K, V]) iteradorInterno(nodoActual *nodoAbb[K, V], visitar func(clave K, valor V) bool) {

	if nodoActual == nil {
		return
	}

	arbol.iteradorInterno(nodoActual.izquierdo, visitar)
	visitar(nodoActual.clave, nodoActual.dato)
	arbol.iteradorInterno(nodoActual.derecho, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {

	iter := new(iterABB[K, V])
	_, minimo := buscarMasIzquierdo[K, V](arbol.raiz)
	_, maximo := buscarMasDerecho[K, V](arbol.raiz)
	iter.desde = &minimo.clave
	iter.hasta = &maximo.clave
	iter.apilarSiguientes(arbol.raiz)

	return iter
}

func (iter *iterABB[K, V]) apilarSiguientes(primerNodo *nodoAbb[K, V]) {

	if primerNodo == nil {
		return
	}
	if iter.arbol.cmp(primerNodo.clave, *iter.desde) >= 0 && iter.arbol.cmp(primerNodo.clave, *iter.hasta) <= 0 {
		iter.pila.Apilar(primerNodo)
		iter.apilarSiguientes(primerNodo.izquierdo)
	} else if iter.arbol.cmp(primerNodo.clave, *iter.desde) < 0 {
		iter.apilarSiguientes(primerNodo.derecho)
	} else if iter.arbol.cmp(primerNodo.clave, *iter.hasta) > 0 {
		iter.apilarSiguientes(primerNodo.izquierdo)
	}
}

func (iter *iterABB[K, V]) HaySiguiente() bool {

	return !iter.pila.EstaVacia()
}

func (iter *iterABB[K, V]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodoActual := iter.pila.Desapilar()
	iter.apilarSiguientes(nodoActual.derecho)
}

func (iter *iterABB[K, V]) VerActual() (K, V) {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	clave := iter.pila.VerTope().clave
	valor := iter.pila.VerTope().dato

	return clave, valor
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	arbol._iterarRango(arbol.raiz, desde, hasta, visitar)
}

func (arbol *abb[K, V]) _iterarRango(nodoActual *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	if nodoActual == nil {
		return
	}

	if arbol.cmp(nodoActual.clave, *desde) >= 0 {
		arbol._iterarRango(nodoActual.izquierdo, desde, hasta, visitar)
	}

	if arbol.cmp(nodoActual.clave, *desde) >= 0 && arbol.cmp(nodoActual.clave, *hasta) <= 0 {
		visitar(nodoActual.clave, nodoActual.dato)
	}

	if arbol.cmp(nodoActual.clave, *hasta) <= 0 {
		arbol._iterarRango(nodoActual.derecho, desde, hasta, visitar)
	}
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := new(iterABB[K, V])
	iterador.desde = desde
	iterador.hasta = hasta
	iterador.apilarSiguientes(arbol.raiz)
	return iterador
}

func buscarMasDerecho[K comparable, V any](padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {

	if padre.derecho != nil && padre.derecho.derecho == nil {
		return padre, padre.derecho
	}

	return buscarMasDerecho(padre.derecho)
}

func buscarMasIzquierdo[K comparable, V any](padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {

	if padre.izquierdo != nil && padre.izquierdo.izquierdo == nil {
		return padre, padre.izquierdo
	}
	return buscarMasIzquierdo(padre.izquierdo)
}
