package utilidades

import (
	"math"
	TDACola_Prioridad "tdas/cola_prioridad"
)

type usuario struct {
	nombre   string
	feed     TDACola_Prioridad.ColaPrioridad[Post]
	afinidad int
}

var EnLinea Usuario

func CrearUsuario(nombre string, afinidad int) Usuario {
	cola_p := TDACola_Prioridad.CrearHeap[Post](cmp_afinidad)
	return &usuario{nombre, cola_p, afinidad}
}

func (u *usuario) VerNombre() string {
	return u.nombre
}

func (u *usuario) VerPostFeed() string {
	if u.feed.EstaVacia() {
		return "No hay publicaciones en el feed"
	}
	return u.feed.Desencolar().VerContenido()
}

func (u *usuario) VerAfinidad() int {
	return u.afinidad
}

func Logear(u Usuario) {
	EnLinea = u
}

func cmp_afinidad(nuevo, actual Post) int {
	if math.Abs(float64(nuevo.VerAutor().VerAfinidad()-EnLinea.VerAfinidad())) > math.Abs(float64(actual.VerAutor().VerAfinidad()-EnLinea.VerAfinidad())) {
		return -1
	} else if math.Abs(float64(nuevo.VerAutor().VerAfinidad()-EnLinea.VerAfinidad())) < math.Abs(float64(actual.VerAutor().VerAfinidad()-EnLinea.VerAfinidad())) {
		return 1
	}
	return 0
}
