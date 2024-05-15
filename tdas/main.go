package main

import (
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
)

func compararCadenas(cad1, cad2 string) int {
	if cad1 < cad2 {
		fmt.Println(cad1, "es menor que", cad2)
		return -1
	} else if cad1 > cad2 {
		fmt.Println(cad1, "es mayor que", cad2)
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

func TestVolumenIteradorCorteABB(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararEnteros)

	for _, i := range rand.Perm(TAMS_VOLUMEN_ABB[0]) {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {

		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

*/

func main() {

	dic := TDADiccionario.CrearABB[int, int](compararEnteros)

	for _, i := range rand.Perm(21) {
		dic.Guardar(i,i)
	}
	
	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		fmt.Println("Entro la clave", c)

		if c%100 == 0 {
			fmt.Println("entro en la condicion", c)

			seguirEjecutando = false
			return false
		}

		if !seguirEjecutando {
			fmt.Println("Fallo la clave", c)

			siguioEjecutandoCuandoNoDebia = true
			return false
		}

		return true
	})
	
	fmt.Println("se ejecuto cuando no debia", siguioEjecutandoCuandoNoDebia)
}
