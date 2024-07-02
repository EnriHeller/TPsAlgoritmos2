import random

class Grafo:

    def __init__(self, dirigido=False, vertices=[]):
        self.dirigido = dirigido
        self.vertices = {v: {} for v in vertices}
        self.entrantes = {v: set() for v in vertices}
        self.centrales = []

    def obtener_vertices(self):
        return self.vertices.keys()
    
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}
        
        if v not in self.entrantes:
            self.entrantes[v] = set()

    def adyacentes(self, v):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        return self.vertices[v].keys()
    
    def agregar_arista(self, v, w, peso=1):
        if v not in self.vertices or w not in self.vertices:
            raise ValueError("Una de las claves no pertenece al grafo")
        self.vertices[v][w] = peso
        self.entrantes[w].add(v)

        if not self.dirigido:
            self.vertices[w][v] = peso
            self.entrantes[v].add(w)

    
    def peso_arista(self, v, w):
        if self.estan_unidos(v, w):
            return self.vertices[v].get(w)
    
    def estan_unidos(self, v, w):
        if v not in self.vertices or w not in self.vertices:
            raise ValueError("Una de las claves no pertenece al grafo")
        return w in self.vertices[v]
    
    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        for w in self.vertices[v].keys():
            self.borrar_arista(v, w)
        del self.vertices[v]
        del self.entrantes[v]
    
    def borrar_arista(self, v, w):
        if self.estan_unidos(v, w):
            del self.vertices[v][w]
            del self.entrantes[w][v]
            if not self.dirigido and v in self.vertices[w]:
                del self.vertices[w][v]
                del self.entrantes[v][w]

    def obtener_entrantes(self,v):
        return self.entrantes[v]

    def vertice_aleatorio(self):
        vertices = list(self.vertices.keys())
        return random.choice(vertices)

    def __str__(self):
        resultado = ""
        for v in self.vertices:
            for w in self.vertices[v]:
                resultado += str(v) + " -> " + str(w) + "\n"
        return resultado
    
    def __iter__(self):
	    return iter(self.vertices.keys())
