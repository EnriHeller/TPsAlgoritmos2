package tp2

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		solicitud := entrada.Text()
		fmt.Println(solicitud)
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}
