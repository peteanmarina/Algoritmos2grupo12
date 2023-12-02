package grafo

//
type Grafo interface {
	//
	AgregarVertice(string)
	//
	SacarVertice(string)
	//
	AgregarArista(string, string)
	//
	SacarArista(string, string)
	//
	ExisteArista(string, string) bool
	//
	ExisteVertice(string) bool
	//
	ObtenerVertices() []string
	//
	ObtenerAdyacentes(string) []string
}
