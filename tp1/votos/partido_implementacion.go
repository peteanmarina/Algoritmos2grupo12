package votos

import "fmt"

type candidato struct {
	nombre_completo string
	votos           int
}

func CrearArregloCandidato(nombres [CANT_VOTACION]string) [CANT_VOTACION]candidato {
	candidatos := [CANT_VOTACION]candidato{
		{nombres[0], 0},
		{nombres[1], 0},
		{nombres[2], 0},
	}
	return candidatos
}

type partidoImplementacion struct {
	nombre_partido string
	candidatos     [CANT_VOTACION]candidato
}

type partidoEnBlanco struct {
	candidatos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]candidato) Partido {
	return &partidoImplementacion{nombre, candidatos}
}

func CrearPartidoEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.candidatos[tipo].votos++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.candidatos[tipo].votos == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido.nombre_partido, partido.candidatos[tipo].nombre_completo, partido.candidatos[tipo].votos)
	}
	return fmt.Sprintf("%s - %s: %d votos", partido.nombre_partido, partido.candidatos[tipo].nombre_completo, partido.candidatos[tipo].votos)
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.candidatos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	retornado := fmt.Sprintf("Votos en Blanco: %d votos", blanco.candidatos[tipo])
	if blanco.candidatos[tipo] == 1 {
		retornado = retornado[:len(retornado)-1]
	}
	return retornado
}
