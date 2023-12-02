package utilidades

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
)

type post struct {
	contenido string
	likes     TDADiccionario.DiccionarioOrdenado[string, Usuario]
	autor     Usuario
	id        int
}

func cmp_nombre(usu1, usu2 string) int {
	if usu1 > usu2 {
		return 1
	} else if usu1 < usu2 {
		return -1
	}
	return 0
}

func CrearPost(contenido string, autor Usuario, id int) Post {
	likes := TDADiccionario.CrearABB[string, Usuario](cmp_nombre)
	return &post{contenido, likes, autor, id}
}

func (p *post) VerContenido() string {
	return p.contenido
}

func (p *post) Lickear(u Usuario) {
	p.likes.Guardar(u.VerNombre(), u)
}

func (p *post) MostrarLikes() {
	fmt.Printf("El post tiene %d likes:\n", p.VerLikes())
	p.likes.Iterar(func(clave string, dato Usuario) bool {
		fmt.Printf("	%s\n", clave)
		return true
	})
}

func (p *post) VerLikes() int {
	return p.likes.Cantidad()
}

func (p *post) VerAutor() Usuario {
	return p.autor
}

func (p *post) VerID() int {
	return p.id
}
