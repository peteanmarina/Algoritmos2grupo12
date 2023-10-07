package main

import (
	"bufio"
	"fmt"
	"os"
	errores "rerepolez/errores"
	votos "rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

const MAX_DNI = 99999999

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
		e_dni := errores.ErrorParametros{}
		fmt.Println(e_dni.Error())
		return
	}

	archivo_lista, err2 := os.Open(os.Args[1])
	archivo_dni, err1 := os.Open(os.Args[2])

	if err1 != nil || err2 != nil {
		e := errores.ErrorLeerArchivo{}
		fmt.Println(e.Error())
		return
	}

	dnis := make([]int, 0)

	s1 := bufio.NewScanner(archivo_dni)
	//se podria plantear un error si no se puede escanear el archivo
	/* if s1.Err() != nil {
		return
	} */
	for s1.Scan() {
		linea, err := strconv.Atoi(s1.Text())
		if err != nil || linea > MAX_DNI {
			e := errores.ErrorLeerArchivo{}
			fmt.Println(e.Error())
			return
		}
		dnis = append(dnis, linea)
	}

	//todos los partidos con sus datos. Lista enlazada con partido blanco y los demas
	partidos := TDALista.CrearListaEnlazada[votos.Partido]()
	partidos.InsertarPrimero(votos.CrearPartidoEnBlanco())

	s2 := bufio.NewScanner(archivo_lista)
	//se podria plantear un error si no se puede escanear el archivo
	/* if s2.Err() != nil {
		return
	} */
	for s2.Scan() {
		nombres := strings.Split(s2.Text(), ",")
		if len(nombres) != 4 {
			e := errores.ErrorLeerArchivo{}
			fmt.Println(e.Error())
			return
		}
		nombre_lista := nombres[0]
		presidente := nombres[1]
		gobernador := nombres[2]
		intendente := nombres[3]
		candidatos := votos.CrearArregloCandidato([3]string{presidente, gobernador, intendente})
		partido := votos.CrearPartido(nombre_lista, candidatos)
		partidos.InsertarUltimo(partido)
	}

	archivo_dni.Close()
	archivo_lista.Close()

	// proceso de pre-votacion
	dnis = radixSort(dnis)
	enfilados := TDACola.CrearColaEnlazada[votos.Votante]()
	votantes := TDALista.CrearListaEnlazada[votos.Votante]()
	votos_realizados := TDALista.CrearListaEnlazada[votos.Voto]()

	// proceso de votacion
	input := bufio.NewReader(os.Stdin)
	var fin bool
	for !fin {

		str_comando, err := input.ReadString('\n')

		//poco probable que este error surga pero cree este nuevo tipo de error
		//por hacer algo
		if err != nil || len(str_comando) == 1 {
			e := errores.ErrorDesconocido{}
			fmt.Println(e.Error())
			continue
		}

		str_comando = strings.TrimSpace(str_comando)
		partes := strings.Fields(str_comando)

		switch partes[0] {
		case "ingresar":
			//fijarse que el usuario ponga cosas coherentes
			if len(partes) < 2 {
				e := errores.ErrorParametros{}
				fmt.Println(e.Error())
				continue
			}

			dni_p, err := strconv.Atoi(partes[1])
			if err != nil {
				e := errores.DNIError{}
				fmt.Println(e.Error())
				continue
			}
			if busquedaBinaria(dnis, 0, len(dnis)-1, dni_p) == -1 {
				e := errores.DNIFueraPadron{}
				fmt.Println(e.Error())
				continue
			}
			votante := votos.CrearVotante(dni_p)
			enfilados.Encolar(votante)

		case "votar":

			if enfilados.EstaVacia() {
				e := errores.FilaVacia{}
				fmt.Println(e.Error())
				continue
			}

			if len(partes) < 3 {
				e := errores.ErrorParametros{}
				fmt.Println(e.Error())
				continue
			}

			t_voto := partes[1] //no hace falta pasarlo a lower, que escriban bien.
			var alternativa votos.TipoVoto

			switch t_voto {
			case "Presidente":
				alternativa = votos.PRESIDENTE
			case "Gobernador":
				alternativa = votos.GOBERNADOR
			case "Intendente":
				alternativa = votos.INTENDENTE
			default:
				e := errores.ErrorTipoVoto{}
				fmt.Println(e.Error())
				continue
			}

			nro_lista, err := strconv.Atoi(partes[2])
			if err != nil || (nro_lista > partidos.Largo() && nro_lista < 0) {
				e := errores.ErrorAlternativaInvalida{}
				fmt.Println(e.Error())
				continue
			}

			votante := enfilados.VerPrimero()
			err = votante.Votar(alternativa, nro_lista, &votantes)
			if err != nil {
				fmt.Println(err.Error())
			}

		case "deshacer":
			if len(partes) != 1 {
				e := errores.ErrorDesconocido{}
				fmt.Println(e.Error())
			}

			err := enfilados.VerPrimero().Deshacer(&votantes)
			if err != nil {
				fmt.Println(err.Error())
			}

			//( armar una idea de como es votante.votar() ) idea -> desapilar
		case "fin-votar":
			//si ya voto, nada, da error si vota 2 veces nomas (y se muestra el error de fraude)

			if enfilados.EstaVacia() {
				fin = true
				continue
			}

			votante := enfilados.VerPrimero()
			voto, err := votante.FinVoto(&votantes)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			votos_realizados.InsertarUltimo(voto)
			votantes.InsertarPrimero(enfilados.Desencolar())

		default:
			fmt.Println("COMANDO INVALIDO")
			if enfilados.EstaVacia() {
				fin = true
			} else {
				e := errores.ErrorCiudadanosSinVotar{}
				fmt.Println(e.Error())
			}
		}

	}

	//procedo de post-votacion
	var impugnados int
	for iter_votos := votos_realizados.Iterador(); iter_votos.HaySiguiente(); iter_votos.Siguiente() {
		voto := iter_votos.VerActual()
		if voto.Impugnado {
			impugnados++
			continue
		}
		k := 0
		for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
			partido := iter_par.VerActual()
			if k == voto.VotoPorTipo[votos.PRESIDENTE] {
				partido.VotadoPara(votos.PRESIDENTE)
			}
			if k == voto.VotoPorTipo[votos.GOBERNADOR] {
				partido.VotadoPara(votos.GOBERNADOR)
			}
			if k == voto.VotoPorTipo[votos.INTENDENTE] {
				partido.VotadoPara(votos.INTENDENTE)
			}
			k++
		}
	}

	//mostrarResultados(partidos)
	fmt.Println("Presidente:")
	for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
		partido := iter_par.VerActual()
		fmt.Println(partido.ObtenerResultado(votos.PRESIDENTE))
	}
	fmt.Println("\nGobernador:")
	for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
		partido := iter_par.VerActual()
		fmt.Println(partido.ObtenerResultado(votos.GOBERNADOR))
	}
	fmt.Println("\nIntendente:")
	for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
		partido := iter_par.VerActual()
		fmt.Println(partido.ObtenerResultado(votos.INTENDENTE))
	}
	if impugnados == 1 {
		fmt.Printf("\nVotos impugnados: %d votos\n", impugnados)
	} else {
		fmt.Printf("\nVotos impugnados: %d votos\n", impugnados)
	}
}
