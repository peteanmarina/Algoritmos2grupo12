package main

import (
	"bufio"
	"fmt"
	"os"
	errores "rerepolez/errores"
	utilidades "rerepolez/utilidades"
	votos "rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
)

const (
	MAX_DNI                           = 99999999
	PARAMETROS_INICIALES_ESPERADOS    = 3
	CANT_DATOS_ESPERADOS_PARTIDO      = 4
	CANT_ELEMENTOS_ESPERADOS_INGRESAR = 2
	CANT_ELEMENTOS_ESPERADOS_VOTAR    = 3
)

func main() {

	if len(os.Args) != PARAMETROS_INICIALES_ESPERADOS {
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
	for s1.Scan() {
		linea, err := strconv.Atoi(s1.Text())
		if err != nil || linea > MAX_DNI {
			e := errores.ErrorLeerArchivo{}
			fmt.Println(e.Error())
			return
		}
		dnis = append(dnis, linea)
	}

	partidos := TDALista.CrearListaEnlazada[votos.Partido]()
	partidos.InsertarPrimero(votos.CrearPartidoEnBlanco())

	s2 := bufio.NewScanner(archivo_lista)

	for s2.Scan() {
		datos_partido := strings.Split(s2.Text(), ",")
		if len(datos_partido) != CANT_DATOS_ESPERADOS_PARTIDO {
			e := errores.ErrorLeerArchivo{}
			fmt.Println(e.Error())
			return
		}
		nombre_lista := datos_partido[0]
		presidente := datos_partido[1]
		gobernador := datos_partido[2]
		intendente := datos_partido[3]
		candidatos := votos.CrearArregloCandidato([3]string{presidente, gobernador, intendente})
		partido := votos.CrearPartido(nombre_lista, candidatos)
		partidos.InsertarUltimo(partido)
	}

	archivo_dni.Close()
	archivo_lista.Close()
	dnis = utilidades.RadixSort(dnis, MAX_DNI)
	enfilados := TDACola.CrearColaEnlazada[votos.Votante]()
	votantes := TDALista.CrearListaEnlazada[votos.Votante]()
	votos_realizados := TDALista.CrearListaEnlazada[votos.Voto]()

	input := bufio.NewReader(os.Stdin)
	var fin bool
	for !fin {

		str_comando, _ := input.ReadString('\n')
		str_comando = strings.TrimSpace(str_comando)
		comando_principal := ""
		if len(str_comando) > 0 {
			comando_principal = strings.Fields(str_comando)[0]
		}

		partes := strings.Fields(str_comando)

		switch comando_principal {
		case "ingresar":
			if len(partes) != CANT_ELEMENTOS_ESPERADOS_INGRESAR {
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
			if utilidades.BusquedaBinaria(dnis, 0, len(dnis)-1, dni_p) == utilidades.NO_ENCONTRADO {
				e := errores.DNIFueraPadron{}
				fmt.Println(e.Error())
				continue
			}
			votante := votos.CrearVotante(dni_p)
			enfilados.Encolar(votante)
			imprimirOK()

		case "votar":

			if enfilados.EstaVacia() {
				e := errores.FilaVacia{}
				fmt.Println(e.Error())
				continue
			}

			if len(partes) != CANT_ELEMENTOS_ESPERADOS_VOTAR {
				e := errores.ErrorParametros{}
				fmt.Println(e.Error())
				continue
			}

			t_voto := partes[1]
			var alternativa votos.TipoVoto

			switch t_voto {
			case "Presidente":
				alternativa = votos.PRESIDENTE
			case "Gobernador":
				alternativa = votos.GOBERNADOR
			case "Intendente":
				alternativa = votos.INTENDENTE
			default:
				fmt.Println(errores.ErrorTipoVoto{}.Error())
				continue
			}

			nro_lista, err := strconv.Atoi(partes[2])

			if err != nil || nro_lista+1 > partidos.Largo() || nro_lista < 0 {
				fmt.Println(errores.ErrorAlternativaInvalida{}.Error())
				continue
			}

			votante := enfilados.VerPrimero()
			err = votante.Votar(alternativa, nro_lista, &votantes)
			if err != nil {
				enfilados.Desencolar()
				fmt.Println(err.Error())
			} else {
				imprimirOK()
			}

		case "deshacer":

			if enfilados.EstaVacia() {
				fmt.Println(errores.FilaVacia{}.Error())
				continue
			}

			dni_p := enfilados.VerPrimero().LeerDNI()
			if votos.Ya_voto(dni_p, votantes) {
				enfilados.Desencolar()
				fmt.Println(errores.ErrorVotanteFraudulento{Dni: dni_p}.Error())
				continue
			}

			err := enfilados.VerPrimero().Deshacer(&votantes)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				imprimirOK()
			}

		case "fin-votar":
			if enfilados.EstaVacia() {
				fmt.Println(errores.FilaVacia{}.Error())
				continue
			}

			votante := enfilados.VerPrimero()
			voto, _ := votante.FinVoto(&votantes)
			imprimirOK()

			votos_realizados.InsertarUltimo(voto)
			votantes.InsertarPrimero(enfilados.Desencolar())

		default:
			fin = true
			if !enfilados.EstaVacia() {
				fmt.Println(errores.ErrorCiudadanosSinVotar{}.Error())
			}
		}

	}

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

	fmt.Println("Presidente:")
	imprimirResultado(partidos, votos.PRESIDENTE)
	fmt.Println("\nGobernador:")
	imprimirResultado(partidos, votos.GOBERNADOR)
	fmt.Println("\nIntendente:")
	imprimirResultado(partidos, votos.INTENDENTE)

	if impugnados == 1 {
		fmt.Printf("\nVotos Impugnados: %d voto\n", impugnados)
	} else {
		fmt.Printf("\nVotos Impugnados: %d votos\n", impugnados)
	}
}
func imprimirOK() {
	fmt.Println("OK")
}

func imprimirResultado(partidos TDALista.Lista[votos.Partido], tipo votos.TipoVoto) {
	for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
		partido := iter_par.VerActual()
		fmt.Println(partido.ObtenerResultado(tipo))
	}
}
