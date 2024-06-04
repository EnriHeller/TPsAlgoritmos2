package tp2

import (
	"fmt"
	"strconv"
	"strings"
	Dic "tdas/diccionario"
)

var arrInstrucciones = []string{"agregar_archivo", "ver_visitantes", "ver_mas_visitados"}

type Lector struct {
	instrucciones Dic.Diccionario[string, bool]
	ips           Dic.DiccionarioOrdenado[string, bool]
}

func CrearLector() Lector {
	instrucciones := Dic.CrearHash[string, bool]()

	for _, instruccion := range arrInstrucciones {
		instrucciones.Guardar(instruccion, true)
	}

	return Lector{
		instrucciones: instrucciones,
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
		desde, hasta := elementos[1],elementos[2]
		l.verVisitantes(desde, hasta)
	
	case "ver_mas_visitados":
		n, err := strconv.Atoi(elementos[1])

		if err != nil{
			return resultado, err
		}
		
		l.verMasVisitados(n)
	}

	return resultado, nil
}

func (l *Lector) agregarArchivo(archivo string) []string {
	var resultado []string
	return resultado
}

func (l *Lector) verVisitantes(desde string, hasta string) []string {
	var resultado []string
	return resultado
}

func (l *Lector) verMasVisitados(n int) []string {
	var resultado []string
	return resultado
}

