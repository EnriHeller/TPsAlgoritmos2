package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Dic "tdas/diccionario"
	"time"
)

var arrInstrucciones = []string{"agregar_archivo", "ver_visitantes", "ver_mas_visitados"}

type Lector struct {
	instrucciones Dic.Diccionario[string, bool]
	ips           Dic.DiccionarioOrdenado[string, bool]
	sitios        Dic.DiccionarioOrdenado[string, int]
}

func CrearLector() Lector {
	instrucciones := Dic.CrearHash[string, bool]()
	ips := Dic.CrearABB[string, bool](compararIps)

	for _, instruccion := range arrInstrucciones {
		instrucciones.Guardar(instruccion, true)
	}

	return Lector{
		instrucciones: instrucciones,
		ips:           ips,
	}
}

func (l *Lector) Procesar(comando string) ([]string, error) {

	var resultado []string
	elementos := strings.Fields(comando)
	instruccion := elementos[0]

	if !l.instrucciones.Pertenece(instruccion) {
		return resultado, fmt.Errorf("comando no valido")
	}

	switch instruccion {

	case "agregar_archivo":
		nombreArchivo := elementos[1]
		l.agregarArchivo(nombreArchivo)

	case "ver_visitantes":
		desde, hasta := elementos[1], elementos[2]
		l.verVisitantes(desde, hasta)

	case "ver_mas_visitados":
		n, err := strconv.Atoi(elementos[1])

		if err != nil {
			return resultado, err
		}

		l.verMasVisitados(n)
	}

	return resultado, nil
}

func (l *Lector) agregarArchivo(ruta string) Dic.Diccionario[string, bool] {

	res := Dic.CrearHash[string, bool]()
	sitios := Dic.CrearHash[string, int]()
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
		return res
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	var fechaAnterior string
	var ipAnterior string
	var visitadoAnterior string
	contador := 0
	for s.Scan() {
		linea := strings.Fields(s.Text())
		ip, fecha, visitado := linea[0], linea[1], linea[2]

		if !l.ips.Pertenece(ip) {
			l.ips.Guardar(ip, true)
		}
		if fechaAnterior == "" && ipAnterior == "" && visitadoAnterior == "" {
			fechaAnterior = fecha
			ipAnterior = ip
			visitadoAnterior = visitado
			continue
		}
		diferencia := obtenerDiferencia(fechaAnterior, fecha)
		if diferencia <= 2 && visitado == visitadoAnterior && ip == ipAnterior {
			contador++
		} else {
			contador = 0
		}

		if contador == 5 {
			res.Guardar(ip, true)
			contador = 0
		}

		fechaAnterior = fecha

	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}

	return res
}

func (l *Lector) verVisitantes(desde string, hasta string) []string {

	var resultado []string
	for iter := l.ips.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
	}
	return resultado
}

func (l *Lector) verMasVisitados(n int) []string {

	return resultado
}

func obtenerDiferencia(anterior, actual string) int {

	var layout = time.RFC3339

	ant, _ := time.Parse(layout, anterior)
	act, _ := time.Parse(layout, actual)
	dif := act.Sub(ant)
	return int(dif.Seconds())
}

func compararIps(a, b string) int {

	arr1 := strings.Split(a, ".")
	arr2 := strings.Split(b, ".")
	nuevo1 := make([]int, len(arr1))
	nuevo2 := make([]int, len(arr2))

	for i, _ := range arr1 {
		num, _ := strconv.Atoi(arr1[i])
		nuevo1[i] = num
		num, _ = strconv.Atoi(arr2[i])
		nuevo2[i] = num
	}
	for i := 0; i < 4; i++ {
		if nuevo1[i] < nuevo2[i] {
			return -1
		} else if nuevo1[i] > nuevo2[i] {
			return 1
		}
	}
	return 0
}
