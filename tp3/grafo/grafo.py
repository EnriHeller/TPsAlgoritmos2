class Grafo:

    def __init__(self, dirigido = False, vertices=[]):

        self.dirigido= dirigido
        self.vertices = {v:{} for v in vertices}

    def obtener_vertices(self):
        return list(self.vertices.keys())
    
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v]= {}

    def adyacentes(self, v):
        return list(self.vertices[v].keys())
    
    def agregar_arista(self, v, w, peso):
        self.vertices[v][w] = peso
    
    def peso_arista(self, v, w):
        return self.vertices[v][w]
    
    def estan_unidos(self, v, w):
        if w in self.vertices[v]:
            return True
        return  False
    
    def borrar_vertice(self, v):
        self.vertices.pop(v)
        for vert in self.vertices:
            if v in vert:
                vert.pop(v)
    
    