package main

import (
	"algogram/errores"
	"algogram/utilidades"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PARAMETROS_INICIALES_ESPERADOS = 2
)

func main() {

	var args = os.Args
	if len(os.Args) != PARAMETROS_INICIALES_ESPERADOS {
		e_dni := errores.ErrorParametros{}
		fmt.Println(e_dni.Error())
		return
	}

	pwd_usuarios := args[1]
	dict_usuarios, err := utilidades.Procesar_usuarios(pwd_usuarios)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := bufio.NewReader(os.Stdin)
	var fin bool
	for !fin {
		str_comando, _ := input.ReadString('\n')
		partes_comando := strings.Fields(str_comando)
		var comando string
		var parametros []string

		if len(partes_comando) > 0 {
			comando = partes_comando[0]
			parametros = partes_comando[1:]
		}

		switch comando {
		// hacer la logica para que la variable EnLinea este fuera del main (porpuesta: en el archivo utilidades)
		case "login":
			if utilidades.EnLinea != nil {
				e := errores.ErrorUsuarioLogeado{}
				fmt.Println(e.Error())
				continue
			}

			if dict_usuarios.Pertenece(parametros[0]) {
				usuario := dict_usuarios.Obtener(parametros[0])
				utilidades.Logear(usuario)
			} else {
				e := errores.ErrorUsuarioInexistente{}
				fmt.Println(e.Error())
				continue
			}

			fmt.Println("Hola ", utilidades.EnLinea.VerNombre())
		case "logout":
			if utilidades.EnLinea == nil {
				e := errores.ErrorUsuarioNoLogeado{}
				fmt.Println(e.Error())
				continue
			}
			utilidades.EnLinea = nil
			fmt.Println("Adios")
		case "publicar":
			fmt.Println(comando)
			fmt.Println(parametros)
		case "ver_siguiente_feed":
			fmt.Println(comando)
			fmt.Println(parametros)
		case "likear_post":
			fmt.Println(comando)
			fmt.Println(parametros)
		case "mostrar_likes":
			fmt.Println(comando)
			fmt.Println(parametros)
		default:
			fin = true
		}
	}

}
