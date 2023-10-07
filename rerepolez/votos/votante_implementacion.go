package votos

import (
	"rerepolez/errores"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni      int
	acciones TDAPila.Pila[accion]
}

type accion struct {
	lista     int
	voto      TipoVoto
	impugnado bool
}

func CrearVotante(dni int) Votante {
	pila := TDAPila.CrearPilaDinamica[accion]()
	return &votanteImplementacion{dni, pila}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int, votantes *TDALista.Lista[Votante]) error {

	dni_p := votante.LeerDNI()
	if ya_voto(dni_p, *votantes) {
		return errores.ErrorVotanteFraudulento{Dni: dni_p}
	}

	if votante.acciones.EstaVacia() || !votante.acciones.VerTope().impugnado {
		if alternativa == 0 {
			votante.acciones.Apilar(accion{impugnado: true})
		} else {
			votante.acciones.Apilar(accion{alternativa, tipo, false})
		}
	}

	return nil
}

func (votante *votanteImplementacion) Deshacer(votantes *TDALista.Lista[Votante]) error {
	dni_p := votante.LeerDNI()
	if ya_voto(dni_p, *votantes) {
		return errores.ErrorVotanteFraudulento{Dni: dni_p}
	}
	if votante.acciones.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votante.acciones.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto(votantes *TDALista.Lista[Votante]) (Voto, error) {
	voto := Voto{[CANT_VOTACION]int{0, 0, 0}, false}

	dni_p := votante.LeerDNI()
	if ya_voto(dni_p, *votantes) {
		return Voto{}, errores.ErrorVotanteFraudulento{Dni: dni_p}
	}

	if votante.acciones.EstaVacia() {
		return voto, nil
	}

	if votante.acciones.VerTope().impugnado {
		return Voto{[CANT_VOTACION]int{0, 0, 0}, true}, nil
	}

	for !votante.acciones.EstaVacia() {
		dato := votante.acciones.Desapilar()
		if voto.VotoPorTipo[dato.voto] == 0 {
			voto.VotoPorTipo[dato.voto] = dato.lista
		}
	}
	return voto, nil
}

func ya_voto(dni int, votantes TDALista.Lista[Votante]) bool {
	for iter := votantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().LeerDNI() == dni {
			return true
		}
	}
	return false
}
