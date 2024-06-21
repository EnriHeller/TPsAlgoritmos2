import heapq
from collections import deque
from pila import Pila


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

        if v == destino and dist_v != 0:
            break

        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v, w)
            dist_w = dist_v + peso

            if dist_w < distancias[w]:
                distancias[w] = dist_w
                padres[w] = v
                heapq.heappush(heap, (dist_w, w))

    resultado = []
    actual = destino
    while actual is not None:
        resultado.append(actual)
        actual = padres[actual]

    return resultado[::-1] if distancias[destino] != float("inf") else []

#Divulgación de rumor: 
# Obtengo lista de todos los vertices que se pueden visitar a partir del vertice pasado por parámetro, a un radio de n.

def divulgar(grafo, v, n):
    cola = deque()
    orden = {v: 0}
    visitados = set(v)
    cola.append(v)

    while cola:
        actual = cola.popleft()

        if orden[actual] == n:
            break

        for w in grafo.adyacentes(actual):
            if w not in visitados:
                visitados.add(actual)
                orden[w] = orden[actual] + 1
                cola.append(w)

    return list(visitados)

#Delincuentes más importantes: 
# Obtener los n mas importantes (centralidad) usando pageRank
def pageRank(grafo, v):
    d = 0.6
    N = len(grafo.obtener_vertices())
    sumatoria = 0

    for w in grafo.adyacentes(v):
        sumatoria += pageRank(grafo, w)/len(grafo.adyacentes(w))

    res = (1- d)/N + d*sumatoria

    return res

def obtenerNMasCentrales(grafo,n):

    if len(grafo.centrales) == 0:
        resultado = []
        for v in grafo.obtener_vertices():
            pr = pageRank(grafo, v)
            resultado.append((v, pr))

        sorted(resultado)
        grafo.centrales = resultado

        largoLista = len(resultado)
        return resultado[largoLista-n:]
    else:
        resultado = grafo.centrales
        largoLista = len(resultado)
        return resultado[largoLista-n:]


#Persecución rápida:
# Dado un vertice en concreto, quiero el camino minimo entre los k vertices mas importantes. En caso de tener caminos de igual largo, priorizar los que vayan a un vertice más importante. Esto se aplica para una lista de vertices concretos
def caminos_mas_rapidos(grafo, vertices, k):
    kMasImportantes = obtenerNMasCentrales(grafo, k)
    resultado = []

    for v in vertices:
        caminos_minimos = []
        for w in kMasImportantes:
            cm = camino_minimo(grafo,v,w)
            caminos_minimos.append((len(cm), cm))
        sorted(caminos_minimos)
        resultado.append((v, caminos_minimos[0][1]))

    sorted(resultado)
    return resultado[0][1]

#Comunidades: 
# Obtener comunidades de al menos n integrantes usando labelPropagation
def labelPropagation(grafo):
    vertices = grafo.obtener_vertices()
    Label = {v: v for v in vertices}

    for v in Label:
        frecuencias = {w: 0 for w in grafo.adyacentes(v)}
        #Calculo maximas frecuencias
        visitados = set(v)
        max_freq(grafo, v, Label, frecuencias, visitados)
        freq_ordenadas = sorted(frecuencias, key= lambda clave: frecuencias[clave], reverse= True)
        Label[v] = freq_ordenadas[0]

def max_freq(grafo, v, Label, frecuencias, visitados):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            valor = Label[w]
            frecuencias[valor] += 1
            max_freq(grafo, w, Label,frecuencias, visitados)

#Ciclo más corto:
#  Se pasa un vertice por parámetro y se busca el camino más corto donde se empiece y termine por este vertice. Si no hay ciclo, se envía "No se encontro recorrido".


def cicloMasCorto(grafo, v):
    # Buscar el ciclo más corto que comienza y termina en 'v'
    res = camino_minimo(grafo, v, v)
    
    # Si el camino encontrado tiene menos de 2 vértices, no es un ciclo válido
    if len(res) < 2:
        return "No se encontró recorrido"
    
    return res


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




