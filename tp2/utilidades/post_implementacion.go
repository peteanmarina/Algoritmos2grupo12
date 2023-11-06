package utilidades

import (
	TDADiccionario "tdas/diccionario"
)

type post struct {
	contenido string
	likes     TDADiccionario.DiccionarioOrdenado[string, Usuario]
	autor     Usuario
}

func (p *post) VerContenido() string {
	return p.contenido
}

//Lickear te permite indicar que te gusta un post
func (p *post) Lickear(id int) {

}

//VerLickes muestra cuantos y quienes le dieron me gusta al post
func (p *post) VerLickes() {

}

func (p *post) VerAutor() Usuario {
	return p.autor
}
