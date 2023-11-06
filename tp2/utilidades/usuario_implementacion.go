package utilidades

import (
	"algogram/errores"
	"fmt"
	"math"
	TDACola_Prioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

type usuario struct {
	nombre   string
	feed     TDACola_Prioridad.ColaPrioridad[Post]
	afinidad int
}

var EnLinea Usuario
var dueñoFeed Usuario

func CrearUsuario(nombre string, afinidad int) Usuario {
	cola_p := TDACola_Prioridad.CrearHeap[Post](cmp_afinidad)
	return &usuario{nombre, cola_p, afinidad}
}

func (u *usuario) VerNombre() string {
	return u.nombre
}

func (u *usuario) VerPostFeed() (string, error) {
	if u.feed.EstaVacia() {
		return "", errores.ErrorPostInexistente_UsuarioNoLogeado{}
	}

	if u.nombre == "barbara" {
		for !u.feed.EstaVacia() {
			fmt.Println(u.feed.Desencolar())
		}
	}
	post := u.feed.Desencolar()

	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", post.VerID(), post.VerAutor().VerNombre(), post.VerContenido(), post.VerLikes()), nil
}

func (u *usuario) VerAfinidad() int {
	return u.afinidad
}

func (u *usuario) ActualizarFeed(post Post) {
	u.feed.Encolar(post)
}

func (u *usuario) Publicar(posts TDADiccionario.Diccionario[int, Post], usuarios TDADiccionario.Diccionario[string, Usuario], contenido string) {
	nuevo_post := CrearPost(contenido, EnLinea, posts.Cantidad())
	posts.Guardar(posts.Cantidad(), nuevo_post)
	usuarios.Iterar(func(clave string, valor Usuario) bool {
		if clave != EnLinea.VerNombre() {
			dueñoFeed = valor
			valor.ActualizarFeed(nuevo_post)
		}

		return true
	})
	dueñoFeed = EnLinea
}

func (u *usuario) Loguear() error {
	if EnLinea != nil {
		e := errores.ErrorUsuarioLogeado{}
		return e
	}

	EnLinea = u
	return nil
}

func Desloguear() error {
	if EnLinea == nil {
		e := errores.ErrorUsuarioNoLogeado{}
		return e
	}
	EnLinea = nil
	fmt.Println("Adios")
	return nil
}

func cmp_afinidad(nuevo, actual Post) int {
	afinidadNueva := nuevo.VerAutor().VerAfinidad()
	afinidadActual := actual.VerAutor().VerAfinidad()
	distanciaNuevo := math.Abs(float64(afinidadNueva - dueñoFeed.VerAfinidad()))
	distanciaActual := math.Abs(float64(afinidadActual - dueñoFeed.VerAfinidad()))

	if distanciaNuevo < distanciaActual {
		return 1
	} else if distanciaNuevo > distanciaActual {
		return -1
	} else {
		if actual.VerID() > nuevo.VerID() {
			return 1
		}

		return -1
	}

}
