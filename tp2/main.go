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

		instruccion, resultado, err := lectorDs.Procesar(comando)

		if err != nil {
			fmt.Println("ERROR")
		} else {
			//imprimir por salida estandar
			fmt.Println(resultado)
		}

		switch instruccion {
		case "agregar_archivo":
			for _, ip := range resultado {
				fmt.Println("Dos: " + ip)
			}
			fmt.Println("OK")

		case "ver_visitantes":
			fmt.Println("Visitantes: ")
			for _, ip := range resultado {
				fmt.Println("\t" + ip)
			}
			fmt.Println("OK")

		case "ver_mas_visitados":
			fmt.Println("Sitios mas visitados:")
			for _, sitio := range resultado {
				fmt.Println("\t" + sitio)
			}
			fmt.Println("OK")

		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}
