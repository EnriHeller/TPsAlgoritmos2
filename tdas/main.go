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

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](compararCadenas)

	claves := make([]string, n)
	valores := make([]int, n)

	for _, i := range rand.Perm(n) {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for _, i := range rand.Perm(n) {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	for _, i := range rand.Perm(n){
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

*/

func main() {
	n := 25000

	dic := TDADiccionario.CrearABB[int, int](compararEnteros)

	ok := true

	for _, i := range rand.Perm(n){
		dic.Guardar(i, i)
		ok = dic.Pertenece(i)

		if !ok  { 
			break
		}
	}

	for _, i := range rand.Perm(n) {

		ok = dic.Pertenece(i)

		if !ok  { 
			fmt.Println("sep")
			break
		}


		ok = dic.Obtener(i) == i
		if !ok { 
			fmt.Println("Se rompe en el obtener")
			break
		}

		ok = dic.Borrar(i) == i
		if !ok { 
			fmt.Println("se rompio en borrar")
			break
		}

		ok = !dic.Pertenece(i)
		if !ok {
			fmt.Println("Se rompe en el segundo pertenece")
			break
		}
	}

	fmt.Println("me dio", ok)

}
