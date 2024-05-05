package main

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

func MasDeLaMitad(arr []int) bool {
    dic := TDADiccionario.CrearHash[int, int]()

    for _, clave := range arr {
        if !dic.Pertenece(clave) {
            dic.Guardar(clave, 1)
        } else {
            cantGuardada := dic.Obtener(clave)
            dic.Guardar(clave, cantGuardada+1)
        }
    }

    iter := dic.Iterador()

    for iter.HaySiguiente() {
        _, valor := iter.VerActual()
        if valor >= len(arr)/2 {
            return true
        }
        iter.Siguiente()
    }

    return false
}

func main() {
    arr := []int{1, 2,2,2, 5, 6,7,8,9,10}

    resultado := MasDeLaMitad(arr)

    fmt.Println(resultado)
}
