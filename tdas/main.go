package main

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

func main()  {
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearHash[string, string]()
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)

		dic.Guardar(claves[i], valores[i])
		fmt.Println(dic.Cantidad())
	}

	//fmt.Println(dic.Cantidad())
}