package diccionario

const (
	VACIO int = 0
	OCUPADO int = 1
	BORRADO int = 2

	TAMANO int = 6
) 

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K,V]
	cantidad int
	tam      int
	borrados int
}

type iterDiccionario[K comparable, V any] struct{
	hash hashCerrado[K, V]
	pos int
}

func crearCeldaHash[K comparable, V any]() celdaHash[K, V]{
	nuevaCelda := new(celdaHash[K, V])
	nuevaCelda.estado = 0
	return *nuevaCelda
}

func CrearHash[K comparable, V any]() Diccionario[K, V]{
	
	nuevoHash := new(hashCerrado[K, V])
	nuevoHash.tam = TAMANO
	nuevoHash.cantidad = 0
	nuevoHash.borrados = 0

	nuevaTabla := make([]celdaHash[K,V], nuevoHash.tam) 

	for i := 0 ; i < nuevoHash.tam ; i ++{
		nuevaCelda := crearCeldaHash[K, V]()
		nuevaTabla = append(nuevaTabla, nuevaCelda)
	}

	nuevoHash.tabla = nuevaTabla

	return nuevoHash
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V){

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool{
	return true
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V{
	var dato V
	return dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V{
	var dato V
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int{
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool){

}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V]{

	return &iterDiccionario[K,V]{}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool{
	return true
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V){
	var clave K
	var valor V
	return clave,valor
}

func (iter *iterDiccionario[K, V]) Siguiente(){
	
}


