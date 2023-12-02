package utilidades

import (
	"algogram/errores"
	"fmt"
	TDACola_Prioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
)

type usuario struct {
	nombre    string
<<<<<<< HEAD
	feed      TDACola_Prioridad.ColaPrioridad[Post]
=======
	feed      TDACola_Prioridad.ColaPrioridad[*Post]
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60
	afinidad  int
	conectado bool
}

func CrearUsuario(nombre string, afinidad int) Usuario {
<<<<<<< HEAD
	cola_p := TDACola_Prioridad.CrearHeap[Post](func(nuevo, actual Post) int {
		afinidadNueva := nuevo.VerAutor().VerAfinidad()
		afinidadActual := actual.VerAutor().VerAfinidad()
		distanciaNuevo := math.Abs(float64(afinidadNueva - afinidad))
		distanciaActual := math.Abs(float64(afinidadActual - afinidad))
=======
	cola_p := TDACola_Prioridad.CrearHeap[*Post](func(nuevo, actual *Post) int {
		afinidadNueva := (*nuevo).VerAutor().VerAfinidad()
		afinidadActual := (*actual).VerAutor().VerAfinidad()
		distanciaNuevo := valor_absoluto(afinidadNueva - afinidad)
		distanciaActual := valor_absoluto(afinidadActual - afinidad)
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60

		if distanciaNuevo < distanciaActual {
			return 1
		}
		if distanciaNuevo > distanciaActual {
			return -1
		}
<<<<<<< HEAD
		if actual.VerID() > nuevo.VerID() {
=======
		if (*actual).VerID() > (*nuevo).VerID() {
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60
			return 1
		}
		return -1

	})
	return &usuario{nombre, cola_p, afinidad, false}
}

func (u *usuario) VerNombre() string {
	return u.nombre
}

func (u *usuario) VerPostFeed() (string, error) {
	if u.feed.EstaVacia() {
		return "", errores.ErrorPostInexistente_UsuarioNoLogeado{}
	}
<<<<<<< HEAD
	var post Post
	post = u.feed.Desencolar()
=======

	post := *u.feed.Desencolar()
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60

	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", post.VerID(), post.VerAutor().VerNombre(), post.VerContenido(), post.VerLikes()), nil
}

func (u *usuario) VerAfinidad() int {
	return u.afinidad
}

func (u *usuario) ActualizarFeed(post *Post) {
	u.feed.Encolar(post)
}

<<<<<<< HEAD
func (u *usuario) Publicar(posts TDADiccionario.Diccionario[int, Post], usuarios TDADiccionario.Diccionario[string, Usuario], contenido string) {
	nuevo_post := CrearPost(contenido, u, posts.Cantidad())
	posts.Guardar(posts.Cantidad(), nuevo_post)
	usuarios.Iterar(func(clave string, usuario Usuario) bool {
		if clave != u.VerNombre() {
			usuario.ActualizarFeed(nuevo_post)
=======
func (u *usuario) Publicar(posts TDADiccionario.Diccionario[int, *Post], usuarios TDADiccionario.Diccionario[string, Usuario], contenido string) {
	nuevo_post := CrearPost(contenido, u, posts.Cantidad())
	posts.Guardar(posts.Cantidad(), &nuevo_post)
	usuarios.Iterar(func(clave string, usuario Usuario) bool {
		if clave != u.VerNombre() {
			usuario.ActualizarFeed(&nuevo_post)
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60
		}
		return true
	})
}

func (u *usuario) Loguear(conectado Usuario) error {
	if conectado != nil {
		e := errores.ErrorUsuarioLogeado{}
		return e
	}
	u.conectado = true
	return nil
}

func (u *usuario) En_linea() bool {
	return u.conectado
}

func (u *usuario) Desloguear() error {
	if !u.En_linea() {
		e := errores.ErrorUsuarioNoLogeado{}
		return e
	}
	u.conectado = false
	fmt.Println("Adios")
	return nil
}
<<<<<<< HEAD
=======

func valor_absoluto(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
>>>>>>> 16021adb3269bd4fe76cece0b48ca00e7dfbaf60
