package main

import (
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
)

func compararCadenas(cad1, cad2 string) int {
	if cad1 < cad2 {
		return -1
	} else if cad1 > cad2 {
		return 1
	}
	return 0
}

func compararEnteros(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0

}

/*
TestGeneracionTodasCombinacionesReverso
    catedra_test.go:459: Esta prueba genera todos los árboles posibles de 7 nodos con valores del 1 al 7, con sus diferentes estructuras. Para cada uno, borra los nodos del 5 al 2 (en ese orden)
    catedra_test.go:478:
                Error Trace:    /tmp/corrector.i30og9ah/skel/diccionario/catedra_test.go:478
                Error:          Should be false
                Test:           TestGeneracionTodasCombinacionesReverso
                Messages:       Fallo al ver si pertenece 4 tras insertar en orden [1 2 3 5 4 6 7] y borrar justamente 4 (habiendo borrado los anteriores)
*/

func main() {

	
	dic := TDADiccionario.CrearABB[int, int](compararEnteros)

	for _,i := range(rand.Perm(7)){
		dic.Guardar(i,i)
	}

	
	dic.Borrar(5)
	dic.Borrar(4)
	dic.Borrar(3)
	dic.Borrar(2)

	for i:= 5; i > 1; i --{
		fmt.Println("El ", i, dic.Pertenece(i) )
		
	}

}
