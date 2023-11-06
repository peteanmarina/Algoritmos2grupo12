package utilidades

// Post modela una publicacion generica con sus metodos para que los usuarios interactuen con el
type Post interface {

	//VerContenido devuelve una cadena con lo que tenga escrito el post
	VerContenido() string

	//Lickear te permite indicar que te gusta un post
	Lickear(int)

	//VerLickes muestra cuantos y quienes le dieron me gusta al post
	VerLickes()

	VerAutor() Usuario
}
