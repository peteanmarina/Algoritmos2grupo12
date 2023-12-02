package main

import (
	"algogram/errores"
	"algogram/utilidades"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
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
	dict_post := utilidades.Procesar_posts()
	dict_comandos := TDADiccionario.CrearHash[string, func(TDADiccionario.Diccionario[string, utilidades.Usuario], TDADiccionario.Diccionario[int, utilidades.Post], []string, utilidades.Usuario) (utilidades.Usuario, error)]()

	inicializarDiccionarioComandos(dict_comandos)

	input := bufio.NewReader(os.Stdin)
	var fin bool
	var conectado utilidades.Usuario
	var error error
	for !fin {
		str_comando, _ := input.ReadString('\n')
		partes_comando := strings.Fields(str_comando)
		var comando string
		var parametros []string

		if len(partes_comando) > 0 {
			comando = partes_comando[0]
			parametros = partes_comando[1:]
		}

		if !dict_comandos.Pertenece(comando) {
			fin = true
		} else {
			conectado, error = dict_comandos.Obtener(comando)(dict_usuarios, dict_post, parametros, conectado)
			if error != nil {
				fmt.Println(error.Error())
			}
		}
	}
}

func inicializarDiccionarioComandos(dict_comandos TDADiccionario.Diccionario[string, func(TDADiccionario.Diccionario[string, utilidades.Usuario], TDADiccionario.Diccionario[int, utilidades.Post], []string, utilidades.Usuario) (utilidades.Usuario, error)]) {
	dict_comandos.Guardar("login", login)
	dict_comandos.Guardar("logout", logout)
	dict_comandos.Guardar("publicar", publicar)
	dict_comandos.Guardar("ver_siguiente_feed", verSiguiente)
	dict_comandos.Guardar("likear_post", likearPost)
	dict_comandos.Guardar("mostrar_likes", mostrarLikes)
}

func login(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {
	var usuario utilidades.Usuario
	nombre := strings.Join(parametros, " ")
	if dict_usuarios.Pertenece(nombre) {
		usuario = dict_usuarios.Obtener(nombre)
		err := usuario.Loguear(conectado)
		if err != nil {
			return conectado, err
		}
	} else {
		return conectado, errores.ErrorUsuarioInexistente{}
	}
	conectado = usuario
	fmt.Println("Hola", usuario.VerNombre())
	return conectado, nil
}

func logout(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {
	if conectado == nil {
		e := errores.ErrorUsuarioNoLogeado{}
		return conectado, e
	}
	err := conectado.Desloguear()
	if err != nil {
		return conectado, err
	}
	conectado = nil
	return conectado, nil
}

func publicar(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {
	if conectado == nil {
		return conectado, errores.ErrorUsuarioNoLogeado{}
	}
	contenido := strings.Join(parametros, " ")
	conectado.Publicar(dict_post, dict_usuarios, contenido)
	fmt.Println("Post publicado")
	return conectado, nil
}

func verSiguiente(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {
	if conectado == nil {
		return conectado, errores.ErrorPostInexistente_UsuarioNoLogeado{}
	}
	contenido, err := conectado.VerPostFeed()
	if err != nil {
		return conectado, err
	}
	fmt.Println(contenido)
	return conectado, nil
}

func likearPost(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {
	if conectado == nil {
		return conectado, errores.ErrorNoPost_UsuarioNoLogeado{}
	}
	id, _ := strconv.Atoi(parametros[0])
	if !dict_post.Pertenece(id) {
		return conectado, errores.ErrorNoPost_UsuarioNoLogeado{}
	}
	post := dict_post.Obtener(id)
	post.Lickear(conectado)
	fmt.Println("Post likeado")
	return conectado, nil
}

func mostrarLikes(dict_usuarios TDADiccionario.Diccionario[string, utilidades.Usuario], dict_post TDADiccionario.Diccionario[int, utilidades.Post], parametros []string, conectado utilidades.Usuario) (utilidades.Usuario, error) {

	id, _ := strconv.Atoi(parametros[0])
	if !dict_post.Pertenece(id) {
		return conectado, errores.ErrorPostInexistente{}
	}
	post := dict_post.Obtener(id)
	if post.VerLikes() == 0 {
		return conectado, errores.ErrorPostInexistente{}
	}
	post.MostrarLikes()
	return conectado, nil
}
