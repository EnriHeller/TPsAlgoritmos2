package diccionario

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterABB[K comparable, V any] struct {
	arbol *abb[K, V]
	desde *K
	hasta *K
	pos   int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {

	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: valor, izquierdo: nil, derecho: nil}
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {

}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	return true
}

func (arbol *abb[K, V]) Cantidad() int {
	return 0
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	var holis V
	return holis
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	var chauchis V
	return chauchis
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

	return
}

func (iter *iterABB[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return iter
}
