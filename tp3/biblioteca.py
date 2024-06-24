import heapq
from collections import deque
from pila import Pila
from grafo import Grafo

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
def pageRank(grafo, v, pr = {}, iteracion = 0,max_iter = 100, d=0.85):
    if v in pr:
        return pr[v]
    
    if iteracion >= max_iter:
        return pr.get(v, (1 - d) / len(grafo.obtener_vertices()))

    N = len(grafo.obtener_vertices())
    sumatoria = 0

    for w in grafo.adyacentes(v):
        sumatoria += pageRank(grafo, w, pr, iteracion + 1)/len(grafo.adyacentes(w))

    res = (1- d)/N + d*sumatoria
    pr[v] = res

    return res

def obtenerNMasCentrales(grafo,n):
    if len(grafo.centrales) == 0:
        resultado = []
        for v in grafo.obtener_vertices():
            pr = pageRank(grafo, v)
            resultado.append((v, pr))

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