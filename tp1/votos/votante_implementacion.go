package votos

import (
	"fmt"
	"rerepolez/errores"
	TDACola "tdas/cola"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni      int
	acciones TDAPila.Pila[accion]
	ya_voto  bool
}

type accion struct {
	lista     int
	voto      TipoVoto
	impugnado bool
}

func CrearVotante(dni int, ya_voto bool) Votante {
	pila := TDAPila.CrearPilaDinamica[accion]()
	return &votanteImplementacion{dni, pila, ya_voto}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	dni_p := votante.LeerDNI()
	if votante.ya_voto {
		return errores.ErrorVotanteFraudulento{Dni: dni_p}
	}

	if alternativa == 0 {
		votante.acciones.Apilar(accion{impugnado: true})
	} else {
		votante.acciones.Apilar(accion{alternativa, tipo, false})
	}

	return nil
}

func (votante *votanteImplementacion) Deshacer(fila TDACola.Cola[Votante]) error {
	dni_p := votante.LeerDNI()
	if votante.ya_voto {
		fila.Desencolar()
		return errores.ErrorVotanteFraudulento{Dni: dni_p}
	}

	if votante.acciones.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votante.acciones.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	voto := Voto{[CANT_VOTACION]int{0, 0, 0}, false}

	dni_p := votante.LeerDNI()
	if votante.ya_voto {
		return Voto{}, errores.ErrorVotanteFraudulento{Dni: dni_p}
	}

	votante.ya_voto = true

	if votante.acciones.EstaVacia() {
		return voto, nil
	}

	for !votante.acciones.EstaVacia() {
		if votante.acciones.VerTope().impugnado {
			return Voto{[CANT_VOTACION]int{0, 0, 0}, true}, nil
		}
		dato := votante.acciones.Desapilar()
		if voto.VotoPorTipo[dato.voto] == 0 {
			voto.VotoPorTipo[dato.voto] = dato.lista
		}
	}
	return voto, nil
}

func (votante *votanteImplementacion) Prueba() {
	fmt.Println(votante.ya_voto)
}
