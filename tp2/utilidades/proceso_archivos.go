package utilidades

import (
	"algogram/errores"
	"bufio"
	"os"
	TDADiccionario "tdas/diccionario"
)

func Procesar_usuarios(pwd string) (TDADiccionario.Diccionario[string, Usuario], error) {
	dict_usuarios := TDADiccionario.CrearHash[string, Usuario]()
	usuarios, err := os.Open(pwd)
	if err != nil {
		e := errores.ErrorLeerArchivo{}
		return nil, e
	}
	defer usuarios.Close()
	afinidad := 0
	s := bufio.NewScanner(usuarios)
	for s.Scan() {
		dict_usuarios.Guardar(s.Text(), CrearUsuario(s.Text(), afinidad))
		afinidad++
	}
	return dict_usuarios, nil
}

/* func VerLogeado() Usuario {
	return EnLinea
}
*/
