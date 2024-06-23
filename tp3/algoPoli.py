from grafo import Grafo
import biblioteca as b

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

def main(archivo_datos):
    grafo = construir_grafo(archivo_datos)
    entrada = input("Ingrese un comando: ")
    comando, params_arr, k = parsearComando(entrada)
    formato = ""
    if comando == "min_seguimientos":
        origen, destino = params_arr
        formato = min_seguimientos(grafo, origen, destino)
    elif comando == "mas_imp":
        cantidad = int(params_arr[0])
        formato = mas_imp(grafo, cantidad)
    elif comando == "persecucion":
        formato = persecucion(grafo, params_arr, k)
        
    elif comando == "comunidades":
        comunidades = b.obtener_comunidades(grafo, params_arr[0])

        for i,c in enumerate(comunidades):
            respuesta = f"Comunidad {i+1}: "
            vertices = ", ".join(c)
            print(respuesta + vertices)

    elif comando == "divulgar":
        pass
    elif comando == "divulgar_ciclo":
        pass
    elif comando == "cfc":
        pass
    else:
        print("Comando no valido.")
    
    if formato:
        print(formato)

main(mensajes_minimos)
