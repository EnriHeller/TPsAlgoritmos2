import heapq
from collections import deque

#NOTA: Puede haber vertices unidos consigo mismos.

#Mínimo seguimientos:
# Busco camino minimo entre vertice origen y destino. Imprimo lista de vertices que se recorren.

def camino_minimo(grafo, origen, destino):
    distancias = {v: float("inf") for v in grafo.obtener_vertices()}
    distancias[origen] = 0
    padres = {origen: None}

    heap = [(origen, 0)]
    heapq.heapify(heap)

    while heap:
        v, dist_v = heapq.heappop(heap)

        if v == destino:
            break
    
        distancias[v] = min(dist_v, distancias[v])
        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v,w)
            dist_w = dist_v + peso
            distancias[w] = dist_w
            padres[w] = v
            heapq.heappush(heap, (w, dist_w))
    
    resultado = []
    actual = destino

    while actual != None:
        resultado.append(actual)
        actual = padres[actual]

    return resultado[::-1]

#Divulgación de rumor: 
# Obtengo lista de todos los vertices que se pueden visitar a partir del vertice pasado por parámetro, a un radio de n.

def vertices_n_dist(grafo, v, n):
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


#Persecución rápida:
# Dado un vertice en concreto, quiero el camino minimo entre los k vertices mas importantes. En caso de tener caminos de igual largo, priorizar los que vayan a un vertice más importante. Esto se aplica para una lista de vertices concretos


#Comunidades: 
# Obtener comunidades de al menos n integrantes usando labelPropagation


#Ciclo más corto:
#  Se pasa un vertice por parámetro y se busca el camino más corto donde se empiece y termine por este vertice. Si no hay ciclo, se envía "No se encontro recorrido".

#CFC:
# Obtengo cada Componente fuertemente conexa del grafo


