import heapq
from collections import deque
from pila import Pila
import random

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
        return None

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

    if v in visitados:
        visitados.pop(v)

    return list(visitados.keys())

#Delincuentes más importantes: 
# Obtener los n mas importantes (centralidad) usando pageRank

def pageRank(grafo, max_iter=20, d=0.85):
    vertices = list(grafo.obtener_vertices())
    N = len(vertices)
    if N == 0:
        return {}
    pr = {v: 1 / N for v in vertices}
    for _ in range(max_iter):
        nuevo_pr = {v: (1 - d) / N for v in vertices}
        for v in vertices:
            for x in obtener_entrantes(grafo, v):
                len_adyacentes_x = len(grafo.adyacentes(x))
                if len_adyacentes_x > 0:
                    nuevo_pr[v] += d * pr[x] / len_adyacentes_x
        pr = nuevo_pr

    return pr

def obtenerNMasCentrales(grafo, n):
    if not grafo.centrales:
        pr = pageRank(grafo)
        # Guardar todos los vértices con sus valores de PageRank ordenados
        grafo.centrales = sorted(pr.items(), key=lambda v: v[1], reverse=True)
    
    return [v for v, _ in grafo.centrales[:n]]

#Persecución rápida:
# Dado un vertice en concreto, quiero el camino minimo entre los k vertices mas importantes. En caso de tener caminos de igual largo, priorizar los que vayan a un vertice más importante. Esto se aplica para una lista de vertices concretos

def caminos_mas_rapidos(grafo, vertices, k):
    kMasImportantes = obtenerNMasCentrales(grafo, k)

    for v in vertices:
        camino_minimo_actual = None
        min_distancia_actual = float('inf')

        for w in kMasImportantes:
            cm = camino_minimo(grafo, v, w)
            if cm:
                distancia = len(cm)
                if (distancia < min_distancia_actual) or (
                    distancia == min_distancia_actual and 
                    kMasImportantes.index(w) < kMasImportantes.index(cm[-1])
                ):
                    camino_minimo_actual = cm
                    min_distancia_actual = distancia

        if camino_minimo_actual:
            return camino_minimo_actual
        else:
            return []

def obtener_entrantes(grafo, v):
    res = []
    for w in grafo.obtener_vertices():
        if v in grafo.adyacentes(w):
            res.append(w)
    return res

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
    label = {v: v for v in grafo.obtener_vertices()}
    max_iter = 10
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

def max_freq(label, LabelWs):
    frecuencias = {}
    for v in LabelWs:
        frecuencias[label[v]] = frecuencias.get(label[v], 0) + 1

    return max(frecuencias, key=frecuencias.get)

#Ciclo más corto:

#  Se pasa un vertice por parámetro y se busca el camino más corto donde se empiece y termine por este vertice.
#  Si no hay ciclo, se envía "No se encontro recorrido".

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

    return None

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
    
    for v in grafo.obtener_vertices():
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