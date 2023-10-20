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

func proceso_voto(t_voto string, nro_lista_s string, largo int) (votos.TipoVoto, int, error) {
	var alternativa votos.TipoVoto

	switch t_voto {
	case "Presidente":
		alternativa = votos.PRESIDENTE
	case "Gobernador":
		alternativa = votos.GOBERNADOR
	case "Intendente":
		alternativa = votos.INTENDENTE
	default:
		return 0, 0, errores.ErrorTipoVoto{}
	}

	nro_lista, err := strconv.Atoi(nro_lista_s)

	if err != nil || nro_lista+1 > largo || nro_lista < 0 {
		return 0, 0, errores.ErrorAlternativaInvalida{}
	}
	return alternativa, nro_lista, nil
}

func procesoDni(dni_s string, dnis []votos.Votante) (int, error) {
	dni_p, err := strconv.Atoi(dni_s)
	if err != nil {
		return 0, errores.DNIError{}
	}

	indice := utilidades.BusquedaBinaria(dnis, 0, len(dnis)-1, dni_p)

	if indice == utilidades.NO_ENCONTRADO {
		return 0, errores.DNIFueraPadron{}
	}
	return indice, nil
}

func procesarResultados(votos_realizados TDALista.Lista[votos.Voto], partidos TDALista.Lista[votos.Partido]) int {
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
	return impugnados
}

func imprimirOK() {
	fmt.Println("OK")
}

func imprimir_todos_resultados(partidos TDALista.Lista[votos.Partido], impugnados int) {
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

func imprimirResultado(partidos TDALista.Lista[votos.Partido], tipo votos.TipoVoto) {
	for iter_par := partidos.Iterador(); iter_par.HaySiguiente(); iter_par.Siguiente() {
		partido := iter_par.VerActual()
		fmt.Println(partido.ObtenerResultado(tipo))
	}
}

func main() {
	if len(os.Args) != PARAMETROS_INICIALES_ESPERADOS {
		e_dni := errores.ErrorParametros{}
		fmt.Println(e_dni.Error())
		return
	}

	dnis, partidos, err := utilidades.AbrirArchivos(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	enfilados := TDACola.CrearColaEnlazada[votos.Votante]()
	votos_realizados := TDALista.CrearListaEnlazada[votos.Voto]()

	input := bufio.NewReader(os.Stdin)
	var fin bool
	for !fin {
		str_comando, _ := input.ReadString('\n')
		str_comando = strings.TrimSpace(str_comando)
		comando := ""
		if len(str_comando) > 0 {
			comando = strings.Fields(str_comando)[0]
		}
		partes_comando := strings.Fields(str_comando)
		switch comando {
		case "ingresar":
			if len(partes_comando) != CANT_ELEMENTOS_ESPERADOS_INGRESAR {
				e := errores.ErrorParametros{}
				fmt.Println(e.Error())
				continue
			}

			dni_i, err := procesoDni(partes_comando[1], dnis)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			enfilados.Encolar(dnis[dni_i])

			imprimirOK()

		case "votar":

			if enfilados.EstaVacia() {
				e := errores.FilaVacia{}
				fmt.Println(e.Error())
				continue
			}

			if len(partes_comando) != CANT_ELEMENTOS_ESPERADOS_VOTAR {
				e := errores.ErrorParametros{}
				fmt.Println(e.Error())
				continue
			}
			alternativa, nro_lista, err := proceso_voto(partes_comando[1], partes_comando[2], partidos.Largo())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			votante := enfilados.VerPrimero()
			err = votante.Votar(alternativa, nro_lista)
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

			err := enfilados.VerPrimero().Deshacer(enfilados)
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
			voto, _ := votante.FinVoto()
			imprimirOK()
			votos_realizados.InsertarUltimo(voto)
			enfilados.Desencolar()

		default:
			fin = true
			if !enfilados.EstaVacia() {
				fmt.Println(errores.ErrorCiudadanosSinVotar{}.Error())
			}
		}

	}

	impugnados := procesarResultados(votos_realizados, partidos)

	imprimir_todos_resultados(partidos, impugnados)
}
