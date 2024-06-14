package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Heap "tdas/cola_prioridad"
	Dic "tdas/diccionario"
	"time"
)

var arrInstrucciones = []string{"agregar_archivo", "ver_visitantes", "ver_mas_visitados"}

type sitiosVisitados struct {
	nombre   string
	cantidad int
}

type lector struct {
	instrucciones Dic.Diccionario[string, bool]
	ips           Dic.DiccionarioOrdenado[string, bool]
	sitios        Dic.Diccionario[string, int]
}

type solicitud struct {
	ultimoSitio string
	ultimaFecha string
	contador    int
}

func CrearLector() lector {
	instrucciones := Dic.CrearHash[string, bool]()
	ips := Dic.CrearABB[string, bool](compararIps)
	sitios := Dic.CrearHash[string, int]()

	for _, instruccion := range arrInstrucciones {
		instrucciones.Guardar(instruccion, true)
	}

	return lector{
		instrucciones: instrucciones,
		ips:           ips,
		sitios:        sitios,
	}
}

func (l *lector) Procesar(comando string) (string, []string, error) {

	var resultado []string
	elementos := strings.Fields(comando)
	instruccion := elementos[0]

	if !l.instrucciones.Pertenece(instruccion) {
		return instruccion, resultado, fmt.Errorf("comando no valido")
	}

	switch instruccion {

	case "agregar_archivo":
		nombreArchivo := elementos[1]
		resultado = l.agregarArchivo(nombreArchivo)

	case "ver_visitantes":
		desde, hasta := elementos[1], elementos[2]
		resultado = l.verVisitantes(desde, hasta)

	case "ver_mas_visitados":
		n, err := strconv.Atoi(elementos[1])

		if err != nil {
			return instruccion, resultado, err
		}

		resultado = l.verMasVisitados(n)

	}

	return instruccion, resultado, nil
}

func (l *lector) agregarArchivo(ruta string) []string {

	var res []string
	entradas := Dic.CrearHash[string, solicitud]()
	archivo, err := os.Open(ruta)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
		return res
	}

	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	for s.Scan() {
		linea := strings.Split(s.Text(), "\t")
		ip, fecha, visitado := linea[0], linea[1], linea[3]
		if !entradas.Pertenece(ip) {
			entradas.Guardar(ip, solicitud{ultimoSitio: visitado, ultimaFecha: fecha, contador: 1})

		}
		//Guardo cantidad de veces que se visitó un sitio
		l.guardarSitios(visitado)

		if !l.ips.Pertenece(ip) {
			l.ips.Guardar(ip, true)
		}

		datos := entradas.Obtener(ip)
		sitioAnterior, fechaAnterior, contador := datos.ultimoSitio, datos.ultimaFecha, datos.contador

		diferencia := obtenerDiferencia(fechaAnterior, fecha)

		if diferencia <= 2 && visitado == sitioAnterior {
			entradas.Guardar(ip, solicitud{ultimoSitio: visitado, ultimaFecha: fecha, contador: contador + 1})
		} else {
			entradas.Guardar(ip, solicitud{ultimoSitio: visitado, ultimaFecha: fecha, contador: 1})
		}

		if contador+1 == 5 {
			res = append(res, ip)
			entradas.Guardar(ip, solicitud{ultimoSitio: visitado, ultimaFecha: fecha, contador: 1})
		}

	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}

	return res
}

func (l *lector) guardarSitios(visitado string) {
	if !l.sitios.Pertenece(visitado) {
		l.sitios.Guardar(visitado, 1)
	} else {
		valorActual := l.sitios.Obtener(visitado)
		l.sitios.Guardar(visitado, valorActual+1)
	}
}

func (l *lector) verVisitantes(desde string, hasta string) []string {

	var resultado []string
	for iter := l.ips.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
	}
	return resultado
}

func (l *lector) verMasVisitados(n int) []string {
	return l.TopKStream(n)
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

	for i := range arr1 {
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

func (l *lector) TopKStream(k int) []string {

	sitiosArr := make([]sitiosVisitados, k)

	for iter := l.sitios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		sitiosArr = append(sitiosArr, sitiosVisitados{nombre: clave, cantidad: valor})
	}

	cp := Heap.CrearHeapArr(sitiosArr[:k], func(s1, s2 sitiosVisitados) int {
		c1, c2 := s1.cantidad, s2.cantidad
		return c2 - c1
	})

	for _, elem := range sitiosArr[k:] {
		if elem.cantidad > cp.VerMax().cantidad {
			cp.Desencolar()
			cp.Encolar(elem)
		}
	}

	top := make([]string, k)

	for i := 0; !cp.EstaVacia(); i++ {
		tope := cp.Desencolar()
		sitio := tope.nombre
		cant := tope.cantidad

		if cant != 0 {
			top[k-i-1] = sitio + " - " + strconv.Itoa(cant)
		}

	}
	return top
}
