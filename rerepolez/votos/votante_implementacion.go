package votos

import (
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni      int
	acciones *TDAPila.Pila[accion]
}

type accion struct {
	lista int
	voto  TipoVoto
}

func CrearVotante(dni int) Votante {
	pila := TDAPila.CrearPilaDinamica[accion]()
	return &votanteImplementacion{dni, &pila}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
