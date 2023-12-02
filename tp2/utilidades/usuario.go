package utilidades

import (
	TDADiccionario "tdas/diccionario"
)

// Post modela una publicacion generica con sus metodos para que los usuarios interactuen con el
type Usuario interface {

	//VerNombre devuelve el nombre de usuario asociado a donde se llama
	VerNombre() string

	En_linea() bool

	//VerPostFeed te permite ver el siguiente post en tu feed teniendo en cuenta los usuarios con mayor vinculo
	VerPostFeed() (string, error)

	VerAfinidad() int

	Loguear(Usuario) error

	Desloguear() error

	Publicar(TDADiccionario.Diccionario[int, *Post], TDADiccionario.Diccionario[string, Usuario], string)

	ActualizarFeed(*Post)
}
