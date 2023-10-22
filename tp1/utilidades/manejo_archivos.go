package utilidades

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
	TDALista "tdas/lista"
)

const (
	MAX_DNI                        = 99999999
	PARAMETROS_INICIALES_ESPERADOS = 3
	CANT_DATOS_ESPERADOS_PARTIDO   = 4
)

func AbrirArchivos(comandos []string) ([]votos.Votante, TDALista.Lista[votos.Partido], error) {
	partidos := TDALista.CrearListaEnlazada[votos.Partido]()
	dnis := make([]votos.Votante, 0)

	archivo_lista, err2 := os.Open(os.Args[1])
	archivo_dni, err1 := os.Open(os.Args[2])
	defer archivo_lista.Close()
	defer archivo_dni.Close()

	if err1 != nil || err2 != nil {
		e := errores.ErrorLeerArchivo{}
		return nil, nil, e
	}

	s1 := bufio.NewScanner(archivo_dni)
	for s1.Scan() {
		linea, err := strconv.Atoi(s1.Text())
		if err != nil || linea > MAX_DNI {
			e := errores.ErrorLeerArchivo{}
			return nil, nil, e
		}
		dnis = append(dnis, votos.CrearVotante(linea, false))
	}

	dnis = RadixSort(dnis, MAX_DNI)

	partidos.InsertarPrimero(votos.CrearPartidoEnBlanco())
	s2 := bufio.NewScanner(archivo_lista)

	for s2.Scan() {
		datos_partido := strings.Split(s2.Text(), ",")
		if len(datos_partido) != CANT_DATOS_ESPERADOS_PARTIDO {
			e := errores.ErrorLeerArchivo{}
			return nil, nil, e
		}
		nombre_lista := datos_partido[0]
		presidente := datos_partido[1]
		gobernador := datos_partido[2]
		intendente := datos_partido[3]
		candidatos := votos.CrearArregloCandidato([3]string{presidente, gobernador, intendente})
		partido := votos.CrearPartido(nombre_lista, candidatos)
		partidos.InsertarUltimo(partido)
	}

	return dnis, partidos, nil
}
