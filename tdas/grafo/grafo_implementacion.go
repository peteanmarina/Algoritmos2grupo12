package grafo

type grafo struct {
	dic      map[string][]string
	cantidad int
}

func CrearGrafo() Grafo {
	return &grafo{make(map[string][]string), 0}
}

func (graf *grafo) AgregarVertice(v string) {
	if graf.ExisteVertice(v) {
		panic("YA EXISTIA ESTE VERTICE")
	}
	graf.dic[v] = make([]string, 0)
	graf.cantidad++
}

func (graf *grafo) SacarVertice(v string) {
	if !graf.ExisteVertice(v) {
		panic("NO HABIA TAL VERTICE")
	}
	delete(graf.dic, v)
	graf.cantidad--
	for _, value := range graf.dic {
		for i := 0; i < len(value); i++ {
			if value[i] == v {
				value[i], value[0] = value[0], value[i]
				value = value[1:]
			}
		}
	}
}

func (graf *grafo) AgregarArista(v, w string) {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	graf.dic[v] = append(graf.dic[v], w)
}

func (graf *grafo) SacarArista(v, w string) {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	var habia bool
	for i := 0; i < len(graf.dic[v]); i++ {
		if graf.dic[v][i] == w {
			graf.dic[v][i], graf.dic[v][0] = graf.dic[v][0], graf.dic[v][i]
			graf.dic[v] = graf.dic[v][1:]
			habia = true
		}
	}

	if !habia {
		panic("NO HABIA TAL ARISTA")
	}

}

func (graf *grafo) ExisteArista(v, w string) bool {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}

	for i := 0; i < len(graf.dic[v]); i++ {
		if graf.dic[v][i] == w {
			return true
		}
	}
	return false
}

func (graf *grafo) ExisteVertice(v string) bool {
	_, esta_v := graf.dic[v]
	return esta_v
}

func (graf *grafo) ObtenerVertices() []string {
	vertices := make([]string, graf.cantidad)
	i := 0
	for vertice := range graf.dic {
		vertices[i] = vertice
		i++
	}
	return vertices
}

func (graf *grafo) ObtenerAdyacentes(v string) []string {
	if !graf.ExisteVertice(v) {
		panic("NO HABIAN TAL VERTICES")
	}
	adyacentes := make([]string, 0)
	adyacentes = append(adyacentes, graf.dic[v]...)
	return adyacentes
}
