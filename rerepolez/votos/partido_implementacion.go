package votos

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

type partidoImplementacion struct { //en principio la lista es por orden en el archivo, 0 reservado para voto blanco
	nombre_partido string
	candidatos     [CANT_VOTACION]candidato //PRESIDENTE, GOBERNADOR, INTENDENTE
}

type partidoEnBlanco struct {
	candidatos [CANT_VOTACION]int //votos x tipo de voto
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]candidato) Partido {
	return &partidoImplementacion{nombre, candidatos}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) { //VER
	partido.candidatos[tipo].votos++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	// MOSTRAR CANTIDAD DE VOTOS POR TIPO
	return ""
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) { //VER
	blanco.candidatos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
