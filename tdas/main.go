package main

import (
	"fmt"
	//"math/rand"
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

dic := TDADiccionario.CrearABB[int, int](compararEnteros)

	arr := rand.Perm(n)

	for _, randIndice := range arr {
		dic.Guardar(randIndice, randIndice)
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for _, randIndice := range arr {
		ok = dic.Pertenece(randIndice)
		if !ok {
			break
		}
		ok = dic.Obtener(randIndice) == randIndice
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	for _, randIndice := range arr {
		require.EqualValues(b, true, dic.Pertenece(randIndice))
		require.EqualValues(b, randIndice, dic.Obtener(randIndice))
		require.EqualValues(b, randIndice, dic.Borrar(randIndice))
	}

*/

func main() {

	dic := TDADiccionario.CrearABB[int, int](compararEnteros)
	arr := [10]int{25, 10, 7, 15, 5, 9, 30, 27, 50, 28}

	for i := 0; i < len(arr); i++ {
		dic.Guardar(arr[i], arr[i])
	}
	desde := 125
	hasta := 300
	iter := dic.IteradorRango(&desde, &hasta)
	fmt.Println(iter)
	fmt.Println(iter.VerActual())

}
