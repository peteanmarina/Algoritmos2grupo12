package main

import (
	"bufio"
	"os"
	"rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
	//"rerepolez/errores" descomentar cuando descubramos como funca
)

const MAX_DNI = 99999999

// EJEMPLO DE USO
// go build rerepolez.go
// ./rerepolez 04_partidos 04_padron
// ingresar 12345678
// ingresar 1234567
// fin-votar
// fin-votar
// fin-votar

func busquedaBinaria(arreglo []int, inicio, fin, elemento int) int {
	//condicion inicial: elementos ordenados de menor a mayor
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if arreglo[medio] == elemento {
		return medio
	} else if arreglo[medio] < elemento {
		return busquedaBinaria(arreglo, medio+1, fin, elemento)
	} else {
		return busquedaBinaria(arreglo, inicio, medio-1, elemento)
	}
}
func countingSort(arreglo []int, cifra int) []int {
	n := len(arreglo)
	nuevo_arreglo := make([]int, n)
	contadores := make([]int, 10)
	for i := 0; i < n; i++ {
		indice := (arreglo[i] / cifra) % 10
		contadores[indice]++
	}
	acumuladores := make([]int, 10)
	for i := 1; i < 10; i++ {
		acumuladores[i] = acumuladores[i-1] + contadores[i-1]
	}
	for i := 0; i < n; i++ {
		indice := (arreglo[i] / cifra) % 10
		nuevo_arreglo[acumuladores[indice]] = arreglo[i]
		acumuladores[indice]++
	}
	for i := 0; i < n; i++ {
		arreglo[i] = nuevo_arreglo[i]
	}
	return arreglo
}
func radixSort(arr []int) []int {
	for cifra := 1; MAX_DNI/cifra > 0; cifra *= 10 {
		arr = countingSort(arr, cifra)
	}
	return arr
}

func main() {
	if len(os.Args) != 3 { //no se reciben los archivos necesarios
		return
	}
	archivo_lista, err2 := os.Open("tests/" + os.Args[1])
	archivo_dni, err1 := os.Open("tests/" + os.Args[2])

	if err1 != nil {
		//errores.ErrorLeerArchivo()
	}
	if err2 != nil {
		//errores.ErrorLeerArchivo()
	}

	defer archivo_dni.Close()
	defer archivo_lista.Close()

	dnis := make([]int, 0)
	s1 := bufio.NewScanner(archivo_dni)

	for s1.Scan() {
		linea, _ := strconv.Atoi(s1.Text()) //manejar error si hay problema con archivo
		dnis = append(dnis, linea)
	}

	dnis = radixSort(dnis)
	votantes := TDACola.CrearColaEnlazada[votos.Votante]()

	//todos los partidos con sus datos. Lista enlazada con partido blanco y los demas
	partidos := TDALista.CrearListaEnlazada[votos.Partido]()
	partidos.InsertarPrimero(votos.CrearVotosEnBlanco())

	s2 := bufio.NewScanner(archivo_lista)
	for s2.Scan() {
		nombres := strings.Split(s2.Text(), ",")
		candidatos := votos.CrearArregloCandidato([3]string{nombres[1], nombres[2], nombres[3]})
		partido := votos.CrearPartido(nombres[0], candidatos)
		votos.CrearVotosEnBlanco()
		partidos.InsertarUltimo(partido)
	}

	// err = s2.Err()

	// if err != nil {go
	// 	fmt.Println(err)
	// }

	fin := false

	reader := bufio.NewReader(os.Stdin)

	for !fin {

		comando, err := reader.ReadString('\n')

		if err != nil {
			//dios sabe
			continue
		}

		comando = strings.TrimSpace(comando)
		partes := strings.Fields(comando)

		switch partes[0] {
		case "ingresar":
			//fijarse q esta en el padron
			//creas votante
			//encolar

			// ingresar dni
			buscado, _ /*err*/ := strconv.Atoi(partes[1]) //error formato invalido para un dni
			//errores.DNIError() no se como funca
			dni := busquedaBinaria(dnis, 0, len(dnis)-1, buscado)

			if dni != -1 {
				votantes.Encolar(votos.CrearVotante(dni))
			}
		case "votar":
			//(VER QUE MIERDA HACER) apilar
			if !votantes.EstaVacia() { //VER QUE ONDA LAS EXCEPCIONES
				//tipo, _:=strconv.Atoi(partes[1])
				//lista, _:=strconv.Atoi(partes[2])
				//votantes.VerPrimero().Votar(tipo, lista)

				//aca no tengo TipoVoto, solucionar

			} else {
				//errores.FilaVacia()
			}

		case "deshacer":
			//(VER QUE MIERDA HACER)desapilar
		case "fin-votar":
			//si ya voto, nada, da error si vota 2 veces nomas (y se muestra el error de fraude)

			if votantes.EstaVacia() {
				fin = true
			} else {
				votantes.Desencolar()
			}
		}

	}

	//mostrarResultados(partidos)

	//presidente
	//de la lista, de cada partido, cada presidente ObtenerResultado( presidente  (0) )

	//gober

	//inten

	// lista(parti parti parti parti)
	// partido -> presi gober intendente

	//recorrer toda la lista, y guardar los mensajes para caaaaada candidato de caada partido
	//mostrarlo en el orden q pides
}
