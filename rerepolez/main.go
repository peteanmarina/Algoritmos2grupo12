package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

func main() {

	// ingresar dni
	votantes := TDACola.CrearColaEnlazada[votos.Votante]()
	votantes.Encolar(votos.CrearVotante(12345678))

	votantes.Desencolar() //cuando ya voto (fin-votar)

	params := os.Args[1:]
	cmd := params[0]
	data := params[1:]
	archivo_dni, err := os.Open(data[0])
	archivo_listas, err := os.Open(data[1])

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", data[0], err)
		return
	}

	defer archivo_dni.Close()
	dnis := make([]int, 0)
	s1 := bufio.NewScanner(archivo_dni)

	for s1.Scan() {
		linea, _ := strconv.Atoi(s1.Text()) //manejar error si hay problema con archivo
		dnis = append(dnis, linea)
	}

	//todos los dnis. slice de int
	//ordenar para poder hacer busqueda binaria, y si se encuentra el q ingresa se encola a la cola
	//si no esta error "no recuerdo cual"
	//si ya voto, nada, da error si vota 2 veces nomas (y se muestra el error de fraude)

	//todos los partidos con sus datos. Lista enlazada con partido blanco y los demas

	partidos := TDALista.CrearListaEnlazada[votos.Partido]()
	partidos.InsertarPrimero(votos.CrearVotosEnBlanco())
	s2 := bufio.NewScanner(archivo_listas)
	for s2.Scan() {
		nombres := strings.Split(s2.Text(), ",")
		candidatos := votos.CrearArregloCandidato([3]string{nombres[1], nombres[2], nombres[3]})
		partido := votos.CrearPartido(nombres[0], candidatos)
		votos.CrearVotosEnBlanco()
		partidos.InsertarUltimo(partido)
	}

	err = s2.Err()

	if err != nil {
		fmt.Println(err)
	}

	switch cmd {
	case "ingresar":
		//fijarse q esta en el padron
		//creas votante
		//encolar
	case "votar":
		//(VER QUE MIERDA HACER) apilar
	case "deshacer":
		//(VER QUE MIERDA HACER)desapilar
	case "fin":
		//(VER QUE MIERDA HACER)desapilamos y hacemos en ese orden las cosas

	}

	//presidente
	//de la lista, de cada partido, cada presidente ObtenerResultado( presidente  (0) )

	//gober

	//inten

	// lista(parti parti parti parti)
	// partido -> presi gober intendente

	//recorrer toda la lista, y guardar los mensajes para caaaaada candidato de caada partido
	//mostrarlo en el orden q pides
}
