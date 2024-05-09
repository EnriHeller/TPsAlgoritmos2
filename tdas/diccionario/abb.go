package diccionario

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
	arbol *abb[K, V]
	desde *K
	hasta *K
	pos   int
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

	if padre == nil || padre.clave == clave {
		return padre
	}

	if arbol.cmp(padre.clave, clave) > 0 {
		//Va a la izquierda
		return arbol.buscar(padre.izquierdo, clave)

	} else if arbol.cmp(padre.clave, clave) < 0 {
		//Va a la derecha
		return arbol.buscar(padre.derecho, clave)
	}

	return padre
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {
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
		panic("La clave no pertenece al diccionario.")
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

}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterABB[K, V])

	return iter

}

func (iter *iterABB[K, V]) HaySiguiente() bool {
	return true
}

func (iter *iterABB[K, V]) Siguiente() {

}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	var clave K
	var valor V
	return clave, valor
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := new(iterABB[K, V])
	return iterador
}

func buscarMasDerecho[K comparable, V any](padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {

	if padre.derecho != nil && padre.derecho.derecho == nil {
		return padre, padre.derecho
	}
	return buscarMasDerecho(padre.derecho)
}
