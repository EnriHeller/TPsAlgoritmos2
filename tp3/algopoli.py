#!/usr/bin/env python3

from grafo import Grafo
import biblioteca as b
import sys

mensajes = "./data/mensajes.tsv"
mensajes_minimos = "./data/minimo.tsv"

def construir_grafo(archivo_datos):
    grafo_mensajes = Grafo(True)
    with open(archivo_datos, "r") as archivo:
        for linea in archivo:
            v, w = linea.strip().split("\t")
            if v not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(v)
            if w not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(w)
            grafo_mensajes.agregar_arista(v, w, 1)
    return grafo_mensajes

def parsear_comando(entrada):
    elementos = entrada.split(" ")
    params_arr = []
    k = 0
    if len(elementos) > 1:
        params_arr = elementos[1].split(",")
    if len(elementos) > 2:
        k = int(elementos[2])
    return elementos[0], params_arr, k

def min_seguimientos(grafo, origen, destino):
    camino_arr = b.camino_minimo(grafo, origen, destino)
    return " -> ".join(camino_arr)

def divulgar(grafo, vertice, n):
    res = b.divulgar(grafo, vertice, n)
    return ", ".join(res)

def mas_imp(grafo, cantidad):
    res = b.obtenerNMasCentrales(grafo, cantidad)
    return ", ".join(res)

def persecucion(grafo, vertices, k):
    res = b.caminos_mas_rapidos(grafo, vertices, k)
    return " -> ".join(res)

def divulgar_ciclo(grafo, vertice):
    res = b.ciclo_mas_corto(grafo, vertice)
    return " -> ".join(res)

def main():
    if len(sys.argv) < 2:
        sys.exit("Error: No se pasan todos los parámetros")

    archivo_datos = sys.argv[1]
    archivo_comandos = sys.argv[2] if len(sys.argv) > 2 else None

    grafo = construir_grafo(archivo_datos)

    if archivo_comandos:
        with open(archivo_comandos, 'r') as archivo:
            for entrada in archivo:
                procesar_entrada(grafo, entrada.strip())
    else:
        for entrada in sys.stdin:
            procesar_entrada(grafo, entrada.strip())

def procesar_entrada(grafo, entrada):
    comando, params_arr, k = parsear_comando(entrada)
    resultado = ""
    if comando == "min_seguimientos":
        origen, destino = params_arr
        resultado = min_seguimientos(grafo, origen, destino)
    elif comando == "cfc":
        cfcs = b.cfcs_grafo(grafo)
        for i, cfc in enumerate(cfcs):
            respuesta = f"CFC {i+1}: "
            vertices = ", ".join(cfc)
            print(respuesta + vertices)
    elif comando == "divulgar":
        resultado = divulgar(grafo, params_arr[0], k)
    elif comando == "divulgar_ciclo":
        resultado = divulgar_ciclo(grafo, params_arr[0])
    elif comando == "comunidades":
        comunidades = b.obtener_comunidades(grafo, params_arr[0])
        for i, c in enumerate(comunidades):
            respuesta = f"Comunidad {i+1}: "
            vertices = ", ".join(c)
            print(respuesta + vertices)
    elif comando == "mas_imp":
        cantidad = int(params_arr[0])
        resultado = mas_imp(grafo, cantidad)
    elif comando == "persecucion":
        resultado = persecucion(grafo, params_arr, k)
    else:
        print("Comando no valido.")
    if resultado:
        print(resultado)

if __name__ == "__main__":
    main()
