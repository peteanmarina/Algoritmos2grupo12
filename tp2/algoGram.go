package main

import (
	"algogram/errores"
	"algogram/utilidades"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PARAMETROS_INICIALES_ESPERADOS = 2
)

var EnLinea utilidades.Usuario

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
	dict_post := utilidades.Procesar_posts()

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
		case "login":
			var usuario utilidades.Usuario
			nombre := strings.Join(parametros, " ")
			if dict_usuarios.Pertenece(nombre) {
				usuario = dict_usuarios.Obtener(nombre)
				err := usuario.Loguear()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
			} else {
				e := errores.ErrorUsuarioInexistente{}
				fmt.Println(e.Error())
				continue
			}
			EnLinea = usuario
			fmt.Println("Hola", usuario.VerNombre())
		case "logout":
			err := utilidades.Desloguear()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			EnLinea = nil
		case "publicar":
			if EnLinea == nil {
				e := errores.ErrorUsuarioNoLogeado{}
				fmt.Println(e.Error())
				continue
			}
			contenido := strings.Join(parametros, " ")
			EnLinea.Publicar(dict_post, dict_usuarios, contenido)
			fmt.Println("Post publicado")
		case "ver_siguiente_feed":
			if EnLinea == nil {
				e := errores.ErrorPostInexistente_UsuarioNoLogeado{}
				fmt.Println(e.Error())
				continue
			}
			contenido, err := EnLinea.VerPostFeed()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(contenido)
		case "likear_post":
			if EnLinea == nil {
				e := errores.ErrorNoPost_UsuarioNoLogeado{}
				fmt.Println(e.Error())
				continue
			}
			id, _ := strconv.Atoi(parametros[0])
			if !dict_post.Pertenece(id) {
				e := errores.ErrorNoPost_UsuarioNoLogeado{}
				fmt.Println(e.Error())
				continue
			}
			post := dict_post.Obtener(id)
			post.Lickear()
			fmt.Println("Post likeado")
		case "mostrar_likes":
			id, _ := strconv.Atoi(parametros[0])
			if !dict_post.Pertenece(id) {
				e := errores.ErrorPostInexistente{}
				fmt.Println(e.Error())
				continue
			}
			post := dict_post.Obtener(id)
			if post.VerLikes() == 0 {
				e := errores.ErrorPostInexistente{}
				fmt.Println(e.Error())
				continue
			}
			post.MostrarLikes()
		default:
			fin = true
		}
	}

}
