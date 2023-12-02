package utilidades

import (
	TDACola "tdas/cola"
	TDAGrafo "tdas/grafo"
)

func BFS(grafo TDAGrafo.Grafo, inicio string, destino string) (map[string]string, map[string]int) {

	visitado := make(map[string]bool)
	cola := TDACola.CrearColaEnlazada[string]()
	padres := make(map[string]string)
	orden := make(map[string]int)

	cola.Encolar(inicio)
	visitado[inicio] = true
	orden[inicio] = 0

	for !cola.EstaVacia() {

		v := cola.Desencolar()

		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitado[w] {
				cola.Encolar(w)
				visitado[w] = true
				padres[w] = v
				orden[w] = orden[v] + 1

				if w == destino { // si no quieren usar destino se pasa "" por parametro
					return padres, orden
				}
			}
		}
	}
	return padres, orden
}

func ReconstruirCamino(padres map[string]string, inicio string, destino string) ([]string, int) {

	camino := []string{destino}

	actual := destino
	costo := 0

	for padres[actual] != "" {
		costo++
		camino = append(camino, actual)
		actual = padres[actual]
	}

	return camino, costo
}

func _DFS(grafo TDAGrafo.Grafo, v string, n int, visitados map[string]bool, ciclo []string) []string { 
	
	visitados[v] = true
	ciclo = append(ciclo, v)

	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados[w] {
			resultado := _DFS(grafo, w, n, visitados, ciclo)

			if resultado != nil {
				return resultado
			}
		} else { 
			if len(ciclo) == n {
				return ciclo
			}
		}
	}
	return nil
}

func ObtenerCiclo(grafo TDAGrafo.Grafo, v string, n int) []string {
	visitados := make(map[string]bool)
	var camino []string

	camino = make([]string, 0, n)
	resultado := _DFS(grafo, v, n, visitados, camino)

	if resultado != nil {
		return resultado
	}

	visitados = make(map[string]bool)

	return nil
}
