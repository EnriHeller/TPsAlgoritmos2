from grafo import Grafo
import biblioteca as b
import sys

mensajes = "./data/mensajes.tsv"
mensajes_minimos = "./data/minimo.tsv"

def construir_grafo(archivo_datos):
    grafo_mensajes = Grafo(True)
    
    with open(archivo_datos, "r") as archivo:
        for linea in archivo:
            v,w = linea.strip().split("\t")

            if v not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(v)

            if w not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(w)
            
            grafo_mensajes.agregar_arista(v,w,1)
    
    return grafo_mensajes

def parsearComando(entrada):
    elementos = entrada.split(" ")

    params_arr = []
    k = 0

    if len(elementos) > 1:
        params_arr = elementos[1].split(",")
    if len(elementos) > 2:
        k = int(elementos[2])

    return elementos[0], params_arr, k

def min_seguimientos(grafo, origen, destino):
    caminoArr = b.camino_minimo(grafo, origen, destino)
    return " -> ".join(caminoArr)

def mas_imp(grafo, cantidad):
    res = b.obtenerNMasCentrales(grafo,cantidad)
    return ", ".join(res)

def persecucion(grafo, vertices, k):
    res = b.caminos_mas_rapidos(grafo,vertices, k)
    return " -> ".join(res)

def divulgar(grafo, vertice, n):
    res = b.divulgar(grafo, vertice, n)
    return ", ".join(res)

def divulgar_ciclo(grafo, vertice):
    res = b.ciclo_mas_corto(grafo, vertice)
    return " -> ".join(res)

def main():
    #archivo que pasa en el ejecutable
    #if len(sys.argv) < 3:
    #    sys.exit("Error: No se pasan todos los parámetros")

    #archivo_datos = sys.argv[1]
    archivo_datos = mensajes
    grafo = construir_grafo(archivo_datos)
    entrada = input("Ingrese un comando: ")
    comando, params_arr, k = parsearComando(entrada)
    resultado = ""

    if comando == "min_seguimientos":
        origen, destino = params_arr
        resultado = min_seguimientos(grafo, origen, destino)
    
    elif comando == "divulgar_ciclo":
        resultado = divulgar_ciclo(grafo,params_arr[0])
    
    elif comando == "cfc":
        cfcs = b.cfcs_grafo(grafo)
        for i,cfc in enumerate(cfcs):
            respuesta = f"CFC {i+1}: "
            vertices = ", ".join(cfc)
            print(respuesta + vertices)
    
    elif comando == "comunidades":
        comunidades = b.obtener_comunidades(grafo, params_arr[0])
        for i,c in enumerate(comunidades):
            respuesta = f"Comunidad {i+1}: "
            vertices = ", ".join(c)
            print(respuesta + vertices)
            
    elif comando == "divulgar":
        resultado = divulgar(grafo,params_arr[0], k)

    #A Reparar:
    elif comando == "mas_imp":
        cantidad = int(params_arr[0])
        resultado = mas_imp(grafo, cantidad)
    
    elif comando == "persecucion":
        resultado = persecucion(grafo, params_arr, k)

    else:
        print("Comando no valido.")
    
    if resultado:
        print(resultado)



main()
