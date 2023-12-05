package grafo

type Grafo[V comparable] interface {
	//Recibe una cadena y la agrega como vertice al grafo
	AgregarVertice(V)

	//Recibe un vertice y lo elimina del grafo
	SacarVertice(V)

	//Recibe dos vertices y crea una arista entre ellos
	AgregarArista(V, V)

	//Recibe un vertice y un adyacente al mismo y elimina la arista del grafo
	SacarArista(V, V)

	////Recibe una arista y devuelve True si existe en el grafo, False en caso contrario
	ExisteArista(V, V) bool

	//Recibe un vertice y devuelve True si existe en el grafo, False en caso contrario
	ExisteVertice(V) bool

	// Devuelve un slice con los vertices del grafo
	ObtenerVertices() []V

	// Recibe un vertice y devuelve un slice de sus adyacentes
	ObtenerAdyacentes(V) []V
}
