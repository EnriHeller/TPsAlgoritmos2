package main

import (
	"bufio"
	"fmt"
	"os"
	lector "tp2/lector"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		comando := entrada.Text()
		lectorDs := lector.CrearLector()

		resultado, err := lectorDs.Procesar(comando)

		if err != nil {
			fmt.Println("ERROR")
		} else {
			//imprimir por salida estandar
			fmt.Println(resultado)
		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}
