package main

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
)

type nodoArbol struct {
	clave          int
	hijo_izquierdo *nodoArbol
	hijo_derecho   *nodoArbol
}

func calcularAltura(raiz *nodoArbol) int {
	if raiz == nil {
		return 0
	}
	alturaIzquierda := calcularAltura(raiz.hijo_izquierdo)
	alturaDerecha := calcularAltura(raiz.hijo_derecho)
	if alturaIzquierda > alturaDerecha {
		return alturaIzquierda + 1
	} else {
		return alturaDerecha + 1
	}
}

func obtenerCantidadHojas(raiz *nodoArbol) int {
	if raiz == nil {
		return 0
	}
	if raiz.hijo_derecho == nil && raiz.hijo_izquierdo == nil {
		return 1
	}
	return obtenerCantidadHojas(raiz.hijo_derecho) + obtenerCantidadHojas(raiz.hijo_izquierdo)
}

func inOrder(node *nodoArbol) {
	if node == nil {
		return
	}
	inOrder(node.hijo_izquierdo) //izq
	fmt.Print(node.clave, " ")   //raiz
	inOrder(node.hijo_derecho)   //der
}

func preOrder(node *nodoArbol) {
	if node == nil {
		return
	}

	fmt.Print(node.clave, " ")    //raiz
	preOrder(node.hijo_izquierdo) //izq
	preOrder(node.hijo_derecho)   //der
}

func postOrder(node *nodoArbol) {
	if node == nil {
		return
	}
	postOrder(node.hijo_izquierdo) //izq
	postOrder(node.hijo_derecho)   //der
	fmt.Print(node.clave, " ")     //raiz
}

func levelOrder(raiz *nodoArbol) {
	if raiz == nil {
		return
	}
	lista := TDALista.CrearListaEnlazada[*nodoArbol]()
	lista.InsertarUltimo(raiz)

	for lista.Largo() > 0 {
		nodo := lista.VerPrimero()
		lista.BorrarPrimero()

		fmt.Print(nodo.clave, " ")

		if nodo.hijo_izquierdo != nil {
			lista.InsertarUltimo(nodo.hijo_izquierdo)
		}

		if nodo.hijo_derecho != nil {
			lista.InsertarUltimo(nodo.hijo_derecho)
		}
	}
}

func DiccionarioInvertirUnico(hash TDADiccionario.Diccionario[string, string]) TDADiccionario.Diccionario[string, string] {
	hash_nuevo := TDADiccionario.CrearHash[string, string]()
	i := hash.Iterador()
	for i.VerActual(); i.HaySiguiente(); i.Siguiente() {
		clave, dato := i.VerActual()
		if hash_nuevo.Pertenece(dato) {
			panic("Hay valores repetidos")
		}
		hash_nuevo.Guardar(dato, clave)
	}
	return hash_nuevo
}

func nivelesInverso(raiz *nodoArbol) {
	if raiz == nil {
		return
	}
	lista := TDALista.CrearListaEnlazada[*nodoArbol]()
	pila := TDAPila.CrearPilaDinamica[*nodoArbol]()
	lista.InsertarUltimo(raiz)

	for lista.Largo() > 0 {
		nodo := lista.VerPrimero()
		lista.BorrarPrimero()
		pila.Apilar(nodo)

		if nodo.hijo_izquierdo != nil {
			lista.InsertarUltimo(nodo.hijo_izquierdo)
		}

		if nodo.hijo_derecho != nil {
			lista.InsertarUltimo(nodo.hijo_derecho)
		}
	}
	for !pila.EstaVacia() {
		fmt.Print(pila.Desapilar().clave, " ")
	}
}

type Resultado struct {
	pais      string
	resultado string // "v" si ganamos, "d" perdimos y "e" empatamos
}
type Resumen struct {
	pais    string
	ventaja int // de partidos de ventaja sobre el contrincante
}

func paternidad(resultados []Resultado) []Resumen {
	hash := TDADiccionario.CrearHash[string, int]()
	for i := 0; i < len(resultados); i++ {
		pais := resultados[i].pais
		result := resultados[i].resultado
		var valor int
		if hash.Pertenece(pais) {
			valor = hash.Obtener(pais)
		}
		if result == "v" {
			hash.Guardar(pais, valor+1)
		} else if result == "d" {
			hash.Guardar(pais, valor-1)
		} else {
			hash.Guardar(pais, valor)
		}
	}
	i := hash.Iterador()
	resumen := make([]Resumen, hash.Cantidad())
	indice := 0
	for i.VerActual(); i.HaySiguiente(); i.Siguiente() {
		resumen[indice].pais, resumen[indice].ventaja = i.VerActual()
		indice++
	}
	return resumen
}

func esHeapDeMinimos(arr []int) bool {
	return esHeapRecursivo(arr, 0)
}

func esHeapRecursivo(arr []int, indice int) bool {
	if indice >= len(arr) {
		return true
	}

	hijoIzquierdo := 2*indice + 1
	hijoDerecho := 2*indice + 2

	// Comprobar si el nodo actual es menor que sus hijos (si existen)
	if (hijoIzquierdo < len(arr) && arr[indice] > arr[hijoIzquierdo]) ||
		(hijoDerecho < len(arr) && arr[indice] > arr[hijoDerecho]) {
		return false
	}

	// Verificar recursivamente los subárboles izquierdo y derecho
	return esHeapRecursivo(arr, hijoIzquierdo) && esHeapRecursivo(arr, hijoDerecho)
}

func (abb *nodoArbol) CantidadConNdescendientes(N int) int {
	cant, _ := abb.ContarNodosConNHijos(N)
	return cant
}

func (abb *nodoArbol) ContarNodosConNHijos(N int) (int, int) {
	if abb == nil {
		return 0, 0
	}

	cantidadIzq, hijosIzq := abb.hijo_izquierdo.ContarNodosConNHijos(N)
	cantidadDer, hijosDer := abb.hijo_derecho.ContarNodosConNHijos(N)

	cantidadTotal := cantidadIzq + cantidadDer
	hijosActual := hijosIzq + hijosDer

	if abb.hijo_izquierdo != nil {
		hijosActual++
	}
	if abb.hijo_derecho != nil {
		hijosActual++
	}

	if hijosActual == N {
		return 1 + cantidadTotal, hijosActual
	}

	return cantidadTotal, hijosActual
}

func mostrarHeap() {
	// var arr []int = []int{7, 5, 6, 4, 0, -3, 4, 2}
	// fmt.Println(esHeapDeMinimos(arr))
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println(esHeapDeMinimos(arr))
}

func mostrarHash() {
	hash := TDADiccionario.CrearHash[string, string]()
	hash.Guardar("clave1", "valor1")
	hash.Guardar("clave3", "valor3")
	hash.Guardar("clave2", "valor2")
	hash.Guardar("clave4", "valor4")
	hash_nuevo := DiccionarioInvertirUnico(hash)
	i := hash.Iterador()
	for i.VerActual(); i.HaySiguiente(); i.Siguiente() {
		clave, dato := i.VerActual()
		println(clave, dato)
		println()
		hash.Guardar(dato, clave)
	}
	i = hash_nuevo.Iterador()
	for i.VerActual(); i.HaySiguiente(); i.Siguiente() {
		clave, dato := i.VerActual()
		println(clave, dato)
		println()

		hash_nuevo.Guardar(dato, clave)
	}
}

func main() {
	//mostrarHeap()
	mostrarArbol()
	//mostrarHash()

}

func (ab *nodoArbol) EsAbb() bool {
	if ab == nil {
		return true
	}
	if ab.hijo_izquierdo == nil || ab.clave > ab.hijo_izquierdo.clave && (ab.hijo_derecho == nil || ab.clave < ab.hijo_derecho.clave) { //ab mayor a izquierdo y menor a derecho
		izquierdoAbb := ab.hijo_izquierdo.EsAbb() //sub arbol izquierdo es abb?
		derechoAbb := ab.hijo_derecho.EsAbb()     //sub arbol derecho es abb?
		if izquierdoAbb && derechoAbb {           //ambos son abb?
			return true
		}
		return false
	}
	return false
}

func mostrarArbol() {

	raiz := &nodoArbol{clave: 10}
	raiz.hijo_izquierdo = &nodoArbol{clave: 7}
	raiz.hijo_derecho = &nodoArbol{clave: 20}
	raiz.hijo_izquierdo.hijo_izquierdo = &nodoArbol{clave: 5}
	raiz.hijo_izquierdo.hijo_derecho = &nodoArbol{clave: 9}
	raiz.hijo_derecho.hijo_izquierdo = &nodoArbol{clave: 16}
	raiz.hijo_derecho.hijo_derecho = &nodoArbol{clave: 22}

	fmt.Println("Recorrido en Orden:")
	inOrder(raiz)
	fmt.Println()

	fmt.Println("Recorrido en Preorden:")
	preOrder(raiz)
	fmt.Println()

	fmt.Println("Recorrido en Postorden:")
	postOrder(raiz)
	fmt.Println()

	fmt.Println("Recorrido por Niveles:")
	levelOrder(raiz)
	fmt.Println()

	fmt.Println("Recorrido por Niveles Inverso:")
	nivelesInverso(raiz)
	fmt.Println()

	fmt.Println("Cantidad de hojas:")
	fmt.Println(obtenerCantidadHojas(raiz))
	fmt.Println()

	fmt.Println("Cantidad con 6 descendientes:")
	fmt.Println(raiz.CantidadConNdescendientes(6))
	fmt.Println()

	print(raiz.EsAbb())
}
