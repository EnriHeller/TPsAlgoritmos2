import heapq
from collections import deque
from pila import Pila
from grafo import Grafo
import random

#NOTA: Puede haber vertices unidos consigo mismos.

#Mínimos seguimientos:

# Busco camino minimo entre vertice origen y destino. Imprimo lista de vertices que se recorren.
def camino_minimo(grafo, origen, destino):
    distancias = {v: float("inf") for v in grafo.obtener_vertices()}
    distancias[origen] = 0
    padres = {origen: None}
    visitados = set()

    heap = [(0, origen)]

    while heap:
        dist_v, v = heapq.heappop(heap)

        if v in visitados:
            continue

        visitados.add(v)

        if v == destino:
            break

        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v, w)
            dist_w = dist_v + peso

            if dist_w < distancias[w]:
                distancias[w] = dist_w
                padres[w] = v
                heapq.heappush(heap, (dist_w, w))

    if destino not in padres:
        return []

    resultado = []
    actual = destino
    while actual is not None:
        resultado.append(actual)
        actual = padres[actual]

    return resultado[::-1]

#Divulgación de rumor: 
# Obtengo lista de todos los vertices que se pueden visitar a partir del vertice pasado por parámetro, a un radio de n.

def divulgar(grafo, v, n):
    visitados = {}
    cola = deque()
    cola.append((v, 0))

    while cola:
        vert, dist = cola.popleft()
        if dist == n:
            continue
        elif dist < n:
            for w in grafo.adyacentes(vert):
                if w not in visitados or visitados[w] > dist + 1:
                    visitados[w] = dist + 1
                    cola.append((w, dist+1))
    visitados.pop(v)
    return list(visitados.keys())

#Delincuentes más importantes: 

# Obtener los n mas importantes (centralidad) usando pageRank
def pageRank(grafo, v, pr = {}, iteracion = 0,max_iter = 100, d=0.85):
    if v in pr:
        return pr[v]
    
    if iteracion >= max_iter:
        return pr.get(v, (1 - d) / len(grafo.obtener_vertices()))

    N = len(grafo.obtener_vertices())
    sumatoria = 0
    entrantes = obtener_entrantes(grafo, v)
    for w in entrantes:
        sumatoria += pageRank(grafo, w, pr, iteracion + 1)/len(grafo.adyacentes(w))

    res = (1- d)/N + d*sumatoria
    pr[v] = res

    return res

def obtenerNMasCentrales(grafo,n):
    if len(grafo.centrales) == 0:
        resultado = []
        pr = {}
        for v in grafo.obtener_vertices():
            pr_v = pageRank(grafo, v, pr)
            pr[v] = pr_v
            resultado.append((v, pr_v))
        resultado.sort(key=lambda v:v[1])
        grafo.centrales = [tupla[0] for tupla in resultado]
        resultado = grafo.centrales
        largoLista = len(resultado)
        return resultado[largoLista-n:]
    else:
        resultado = grafo.centrales
        largoLista = len(resultado)
        return resultado[largoLista-n:]


#Persecución rápida:
# Dado un vertice en concreto, quiero el camino minimo entre los k vertices mas importantes. En caso de tener caminos de igual largo, priorizar los que vayan a un vertice más importante. Esto se aplica para una lista de vertices concretos
def caminos_mas_rapidos(grafo, vertices, k):
    kMasImportantes = [v for v, _ in obtenerNMasCentrales(grafo, k)]
    resultado = []

    for v in vertices:
        caminos_minimos = []
        for w in kMasImportantes:
            cm = camino_minimo(grafo, v, w)
            if cm: 
                caminos_minimos.append((len(cm), cm))
        
        if caminos_minimos:
            caminos_minimos.sort(key=lambda cm:(cm[0], cm[1]))
            resultado.append(caminos_minimos[0][1])
        else:
            resultado.append([])

    return resultado[0]

#Comunidades: 
# Obtener comunidades de al menos n integrantes usando labelPropagation

def obtener_comunidades(grafo, n):

    Label = label_propagation(grafo)
    comunidades = {}
    for vertice, numero in Label.items():
        if numero not in comunidades:
            comunidades[numero] = []
        comunidades[numero].append(vertice)
    
    # Filtrar comunidades con al menos n integrantes
    filtro_comunidades = {numero: arreglo for numero, arreglo in comunidades.items() if len(arreglo) >= int(n)}

    return list(filtro_comunidades.values())

def label_propagation(grafo):
    # Inicializar etiquetas
    label = {v: v for v in grafo.obtener_vertices()}
    max_iter = 100
    for _ in range(max_iter):
        vertices = list(label.keys())
        random.shuffle(vertices)  # Orden aleatorio
        cambios = False
        for v in vertices:
            entrantes = obtener_entrantes(grafo, v)
            nueva_etiqueta = max_freq(label, entrantes)
            if label[v] != nueva_etiqueta:
                label[v] = nueva_etiqueta
                cambios = True
        if not cambios:
            break  # Terminar si no hay cambios en las etiquetas

    return label

def obtener_entrantes(grafo, v):
    res = []
    for w in grafo.obtener_vertices():
        if v in grafo.adyacentes(w):
            res.append(w)
    return res

def max_freq(label, LabelWs):
    frecuencias = {}
    for v in LabelWs:
        frecuencias[label[v]] = frecuencias.get(label[v], 0) + 1

    return max(frecuencias, key=frecuencias.get)

#Ciclo más corto:
#  Se pasa un vertice por parámetro y se busca el camino más corto donde se empiece y termine por este vertice. Si no hay ciclo, se envía "No se encontro recorrido".

def ciclo_mas_corto(grafo, origen):
    cola = deque()
    cola.append(origen)
    visitados = set()
    padres = {origen: None}

    while cola:
        v = cola.popleft()
        visitados.add(v)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                cola.append(w)
            elif w in visitados and w == origen:
                return reconstruir_camino(padres, v, origen)

    return "No se encontro recorrido"

def reconstruir_camino(padres, final, origen):
    camino = [origen]
    actual = final
    while actual != None:
        camino.append(actual)
        actual = padres[actual]

    return camino[::-1]

#CFC:

def cfcs_grafo(grafo):
    resultados = []
    visitados = set()
    pila = Pila()
    
    for v in grafo:
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, {}, {}, pila, set(), resultados, [0])
    return resultados


def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True: 
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)


'''def main():
    grafo = Grafo(True, vertices=["a","b","c","d","e", "f"])

    grafo.agregar_arista("a" ,"b" ,1)
    grafo.agregar_arista("b" ,"c" ,1)
    grafo.agregar_arista("c" ,"d" ,1)
    grafo.agregar_arista( "d","a" ,1)
    grafo.agregar_arista( "d","e" ,1)
    grafo.agregar_arista( "e","a" ,1)
    grafo.agregar_arista( "c","a" ,1)
    grafo.agregar_arista( "d","f" ,1)
    grafo.agregar_arista( "f","e" ,1)
    grafo.agregar_arista( "f","a" ,1)
    #grafo.agregar_arista( "d","f" ,1)

    print(ciclo_mas_corto(grafo, "a"))

main()'''