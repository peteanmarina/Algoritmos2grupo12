package main

import (
	"algogram/errores"
	"algogram/utilidades"
	"bufio"
	"fmt"
	"os"
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
	dict_comandos := TDADiccionario.CrearHash[string, func(TDADiccionario.Diccionario[string, utilidades.Usuario], TDADiccionario.Diccionario[int, *utilidades.Post], []string, utilidades.Usuario) (utilidades.Usuario, error)]()

	utilidades.InicializarDiccionarioComandos(dict_comandos)

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
