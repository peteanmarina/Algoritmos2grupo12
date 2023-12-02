package main

import (
	"bufio"
	"fmt"
	"netstats/utilidades"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

const (
	PARAMETROS_INICIALES_ESPERADOS = 2
)

func main() {
	grafo := TDAGrafo.CrearGrafo()
	dict_comandos := TDADiccionario.CrearHash[string, func([]string, TDAGrafo.Grafo)]()

	inicializarDiccionarioComandos(dict_comandos)
	input := bufio.NewReader(os.Stdin)
	var fin bool
	var error error
	for !fin {
		str_comando, _ := input.ReadString('\n')
		partes_comando := strings.Fields(str_comando)
		var comando string
		var parametros []string

		if len(partes_comando) > 0 {
			comando = partes_comando[0]
			parametros = partes_comando[1:]
		}

		if !dict_comandos.Pertenece(comando) {
			fin = true
		} else {
			dict_comandos.Obtener(comando)(parametros, grafo)
			if error != nil {
				fmt.Println(error.Error())
			}
		}
	}
}

func inicializarDiccionarioComandos(dict_comandos TDADiccionario.Diccionario[string, func([]string, TDAGrafo.Grafo)]) {
	dict_comandos.Guardar("camino", camino)
	dict_comandos.Guardar("lectura", lectura)
	dict_comandos.Guardar("diametro", diametro)
}

func camino(parametros []string, grafo TDAGrafo.Grafo) {
	inicio := parametros[0]
	destino := parametros[1]
	padres, _ := utilidades.BFS(grafo, inicio, destino)
	camino, costo := utilidades.ReconstruirCamino(padres, inicio, destino)
	println()
	for vertice := range camino {
		println(vertice)
	}
	println("Costo: ", costo)
}

func lectura(parametros []string, grafo TDAGrafo.Grafo) {

}

func diametro(parametros []string, grafo TDAGrafo.Grafo) {
	diametro := 0
	for _, vertice := range grafo.ObtenerVertices() {
		_, orden := utilidades.BFS(grafo, vertice, "")
		aux := utilidades.EncontrarMaximo(orden)
		if diametro < aux {
			diametro = aux
		}
	}
	println("El diametro es ", diametro)

}

func rango(parametros []string, grafo TDAGrafo.Grafo) {
	inicio := parametros[0]
	n, err := strconv.Atoi(parametros[1])

	if err != nil {
		return
	}

	padres, orden := utilidades.BFS(grafo, inicio, "")
	cont := 0
	for _, v := range padres {
		if orden[v] == n {
			cont++
		}
	}
	println(cont)
}

func ciclo(parametros []string, grafo TDAGrafo.Grafo) {
	v := parametros[0]
	n, err := strconv.Atoi(parametros[1])

	if err != nil {
		return
	}

	utilidades.ObtenerCiclo(grafo, v, n)
}
