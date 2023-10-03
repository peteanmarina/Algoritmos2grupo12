package votos

type candidato struct {
	nombre_completo string
	votos           int
}

type partidoImplementacion struct {
	presidente     *candidato
	gobernador     *candidato
	intendente     *candidato
	lista          int
	nombre_partido string
}

type partidoEnBlanco struct {
	presidente int
	gobernador int
	intendente int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]candidato) Partido {
	return &partidoImplementacion{}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return ""
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {

}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
