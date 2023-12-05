package main

import (
	"bufio"
	"fmt"
	"netstats/errores"
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

type comando struct {
	Nombre  string
	Funcion func([]string, TDAGrafo.Grafo[string], []comando)
}

func main() {
	var args = os.Args

	if len(args) != PARAMETROS_INICIALES_ESPERADOS {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}

	archivo, _ := os.Open(args[1])
	defer archivo.Close()

	grafo := TDAGrafo.CrearGrafo[string]()
	s := bufio.NewScanner((archivo))
	for s.Scan() {
		resultado := strings.Split(s.Text(), "	")
		grafo.AgregarVertice(resultado[0])
		for i := 1; i < len(resultado); i++ {
			if !grafo.ExisteVertice(resultado[i]) {
				grafo.AgregarVertice(resultado[i])
			}
			grafo.AgregarArista(resultado[0], resultado[i])
		}
		fmt.Println(grafo.ObtenerAdyacentes(resultado[0]))
	}

	comandos := []comando{
		{"camino", camino},
		{"diametro", diametro},
		{"rango", rango},
	}

	dict_comandos := TDADiccionario.CrearHash[string, func([]string, TDAGrafo.Grafo[string], []comando)]()

	inicializarDiccionarioComandos(comandos, dict_comandos)

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
			dict_comandos.Obtener(comando)(parametros, grafo, comandos)
			if error != nil {
				fmt.Println(error.Error())
			}
		}
	}
}

func inicializarDiccionarioComandos(comandos []comando, dict_comandos TDADiccionario.Diccionario[string, func([]string, TDAGrafo.Grafo[string], []comando)]) {
	for _, c := range comandos {
		dict_comandos.Guardar(c.Nombre, c.Funcion)
	}
}

func camino(parametros []string, grafo TDAGrafo.Grafo[string], comandos []comando) {
	aux := strings.Join(parametros, " ")
	parametros = strings.Split(aux, ",")

	if len(parametros) != 2 {
		panic("FALTAN PARAMETROS")
	}

	inicio := parametros[0]
	destino := parametros[1]

	padres, _ := utilidades.BFS(grafo, inicio, destino)
	camino, costo := utilidades.ReconstruirCamino(padres, inicio, destino)
	println()
	for _, vertice := range camino {
		println(vertice)
	}
	println("Costo: ", costo)
}

func diametro(parametros []string, grafo TDAGrafo.Grafo[string], comandos []comando) {
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

func rango(parametros []string, grafo TDAGrafo.Grafo[string], comandos []comando) {
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
