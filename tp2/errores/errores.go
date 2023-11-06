package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type ErrorUsuarioLogeado struct{}

func (e ErrorUsuarioLogeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLogeado struct{}

func (e ErrorUsuarioNoLogeado) Error() string {
	return "Error: no habia usuario loggeado"
}
