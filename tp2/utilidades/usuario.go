package utilidades

// Post modela una publicacion generica con sus metodos para que los usuarios interactuen con el
type Usuario interface {

	//VerNombre devuelve el nombre de usuario asociado a donde se llama
	VerNombre() string

	//VerPostFeed te permite ver el siguiente post en tu feed teniendo en cuenta los usuarios con mayor vinculo
	VerPostFeed() string

	VerAfinidad() int
}
