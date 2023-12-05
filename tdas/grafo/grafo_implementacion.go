package grafo

type grafo[V comparable] struct {
	dic      map[V][]V
	cantidad int
}

func CrearGrafo[V comparable]() Grafo[V] {
	return &grafo[V]{make(map[V][]V), 0}
}

func (graf *grafo[V]) AgregarVertice(v V) {
	/* if graf.ExisteVertice(v) {
		fmt.Println(v)
		panic("YA EXISTIA ESTE VERTICE")
	} */
	graf.dic[v] = make([]V, 0)
	graf.cantidad++
}

func (graf *grafo[V]) SacarVertice(v V) {
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

func (graf *grafo[V]) AgregarArista(v, w V) {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	graf.dic[v] = append(graf.dic[v], w)
}

func (graf *grafo[V]) SacarArista(v, w V) {
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

func (graf *grafo[V]) ExisteArista(v, w V) bool {
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

func (graf *grafo[V]) ExisteVertice(v V) bool {
	_, esta_v := graf.dic[v]
	return esta_v
}

func (graf *grafo[V]) ObtenerVertices() []V {
	vertices := make([]V, graf.cantidad)
	i := 0
	for vertice := range graf.dic {
		vertices[i] = vertice
		i++
	}
	return vertices
}

func (graf *grafo[V]) ObtenerAdyacentes(v V) []V {
	if !graf.ExisteVertice(v) {
		panic("NO HABIAN TAL VERTICES")
	}
	adyacentes := make([]V, 0)
	adyacentes = append(adyacentes, graf.dic[v]...)
	return adyacentes
}
