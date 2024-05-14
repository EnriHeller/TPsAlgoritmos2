package diccionario

import (
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

func (arbol *abb[K, V]) buscar(padre **nodoAbb[K, V], clave K) **nodoAbb[K, V] {

	if *padre == nil || (*padre).clave == clave {
		return padre
	} else {
		claveActual := (*padre).clave

		if arbol.cmp(clave, claveActual) < 0 {
			return arbol.buscar(&(*padre).izquierdo, clave)
		} else if arbol.cmp(clave, claveActual) > 0 {
			return arbol.buscar(&(*padre).derecho, clave)
		} else {
			return padre
		}
	}

}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {
	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		*busqueda = crearNodo(clave, valor)
		arbol.cantidad++
	} else {
		(*busqueda).dato = valor
	}

}

func (arbol *abb[K, V]) Pertenece(clave K) bool {

	busqueda := arbol.buscar(&arbol.raiz, clave)
	return *busqueda != nil
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Borrar(clave K) V {

	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	refeActual := *busqueda

	if refeActual.izquierdo == nil && refeActual.derecho == nil {
		*busqueda = nil
	}

	if refeActual.izquierdo == nil && refeActual.derecho != nil {
		*busqueda = refeActual.derecho
	}

	if refeActual.izquierdo != nil && refeActual.derecho == nil {
		*busqueda = refeActual.izquierdo
	}
	if refeActual.izquierdo != nil && refeActual.derecho != nil {
		reemplazante := buscarMasDerecho[K, V](refeActual.izquierdo)
		refeActual.clave, refeActual.dato = reemplazante.clave, reemplazante.dato
		reemplazante = nil
	}

	arbol.cantidad--
	return refeActual.dato
}

func (arbol *abb[K, V]) Obtener(clave K) V {

	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	return (*busqueda).dato
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	arbol.iteradorInterno(arbol.raiz, visitar)
}

func (arbol *abb[K, V]) iteradorInterno(nodoActual *nodoAbb[K, V], visitar func(clave K, valor V) bool) {

	if nodoActual == nil {
		return
	}

	arbol.iteradorInterno(nodoActual.izquierdo, visitar)
	if !visitar(nodoActual.clave, nodoActual.dato) {
		return
	}
	arbol.iteradorInterno(nodoActual.derecho, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterABB[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.arbol = arbol
	if arbol.raiz != nil {
		minimo := buscarMasIzquierdo[K, V](arbol.raiz)
		maximo := buscarMasDerecho[K, V](arbol.raiz)
		iter.desde = &minimo.clave
		iter.hasta = &maximo.clave

		iter.apilarSiguientes(arbol.raiz, iter.pila)
	}
	return iter
}

func (iter *iterABB[K, V]) apilarSiguientes(primerNodo *nodoAbb[K, V], pila TDAPila.Pila[*nodoAbb[K, V]]) {

	if primerNodo == nil {
		return
	}

	if iter.arbol.cmp(primerNodo.clave, *iter.desde) >= 0 && iter.arbol.cmp(primerNodo.clave, *iter.hasta) <= 0 {
		iter.pila.Apilar(primerNodo)
		iter.apilarSiguientes(primerNodo.izquierdo, pila)

	} else if iter.arbol.cmp(primerNodo.clave, *iter.desde) < 0 {
		iter.apilarSiguientes(primerNodo.derecho, pila)

	} else if iter.arbol.cmp(primerNodo.clave, *iter.hasta) > 0 {
		iter.apilarSiguientes(primerNodo.izquierdo, pila)

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
	iter.apilarSiguientes(nodoActual.derecho, iter.pila)
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
		if !visitar(nodoActual.clave, nodoActual.dato) {
			return
		}
	}

	if arbol.cmp(nodoActual.clave, *hasta) <= 0 {
		arbol._iterarRango(nodoActual.derecho, desde, hasta, visitar)
	}
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterABB[K, V])
	iter.desde = desde
	iter.hasta = hasta
	iter.apilarSiguientes(arbol.raiz, iter.pila)
	return iter
}

func buscarMasDerecho[K comparable, V any](padre *nodoAbb[K, V]) *nodoAbb[K, V] {

	if padre.derecho == nil {
		return padre
	}
	return buscarMasDerecho(padre.derecho)
}

func buscarMasIzquierdo[K comparable, V any](padre *nodoAbb[K, V]) *nodoAbb[K, V] {
	if padre.izquierdo == nil {
		return padre
	}
	return buscarMasIzquierdo(padre.izquierdo)
}
