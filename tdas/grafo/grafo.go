package grafo

//
type Grafo interface {
	//Recibe una cadena y la agrega como vertice al grafo
	AgregarVertice(string)
	
	//Recibe un vertice y lo elimina del grafo
	SacarVertice(string)
	
	//Recibe dos vertices y crea una arista entre ellos
	AgregarArista(string, string)
	
	//Recibe un vertice y un adyacente al mismo y elimina la arista del grafo
	SacarArista(string, string)
	
	////Recibe una arista y devuelve True si existe en el grafo, False en caso contrario
	ExisteArista(string, string) bool
	
	//Recibe un vertice y devuelve True si existe en el grafo, False en caso contrario
	ExisteVertice(string) bool
	
	// Devuelve un slice con los vertices del grafo
	ObtenerVertices() []string

	// Recibe un vertice y devuelve un slice de sus adyacentes
	ObtenerAdyacentes(string) []string
}
