package utilidades

import (
	"algogram/errores"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func Procesar_posts() TDADiccionario.Diccionario[int, *Post] {
	return TDADiccionario.CrearHash[int, *Post]()
}

func InicializarDiccionarioComandos(dict_comandos TDADiccionario.Diccionario[string, func(TDADiccionario.Diccionario[string, Usuario], TDADiccionario.Diccionario[int, *Post], []string, Usuario) (Usuario, error)]) {
	dict_comandos.Guardar("login", login)
	dict_comandos.Guardar("logout", logout)
	dict_comandos.Guardar("publicar", publicar)
	dict_comandos.Guardar("ver_siguiente_feed", verSiguiente)
	dict_comandos.Guardar("likear_post", likearPost)
	dict_comandos.Guardar("mostrar_likes", mostrarLikes)
}

func login(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {
	var usuario Usuario
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

func logout(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {
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

func publicar(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {
	if conectado == nil {
		return conectado, errores.ErrorUsuarioNoLogeado{}
	}
	contenido := strings.Join(parametros, " ")
	conectado.Publicar(dict_post, dict_usuarios, contenido)
	fmt.Println("Post publicado")
	return conectado, nil
}

func verSiguiente(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {
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

func likearPost(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {
	if conectado == nil {
		return conectado, errores.ErrorNoPost_UsuarioNoLogeado{}
	}
	id, _ := strconv.Atoi(parametros[0])
	if !dict_post.Pertenece(id) {
		return conectado, errores.ErrorNoPost_UsuarioNoLogeado{}
	}
	post := *dict_post.Obtener(id)
	post.Lickear(conectado)
	fmt.Println("Post likeado")
	return conectado, nil
}

func mostrarLikes(dict_usuarios TDADiccionario.Diccionario[string, Usuario], dict_post TDADiccionario.Diccionario[int, *Post], parametros []string, conectado Usuario) (Usuario, error) {

	id, _ := strconv.Atoi(parametros[0])
	if !dict_post.Pertenece(id) {
		return conectado, errores.ErrorPostInexistente{}
	}
	post := *dict_post.Obtener(id)
	if post.VerLikes() == 0 {
		return conectado, errores.ErrorPostInexistente{}
	}
	post.MostrarLikes()
	return conectado, nil
}
