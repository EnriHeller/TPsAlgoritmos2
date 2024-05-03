package main

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

func main()  {
	dic := TDADiccionario.CrearHash[int,int]()
	for i := 0; i < 10; i++ {
		dic.Guardar(i, i)
		fmt.Println("primer for", dic.Obtener(i))
	}
	for i := 0; i < 10; i++ {
		dic.Guardar(i, 2*i)
		fmt.Println("segundo for", dic.Obtener(i))
	}
	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
		if !ok{
			fmt.Println("fallo en esta clave", i)
		}
	}

}