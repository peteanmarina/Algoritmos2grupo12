package diccionario

import (
	"crypto/sha256"
	"fmt"
	"math"
)

//go test -bench=. -benchmem

type estado int

const TAMANIO_INICIAL = 10
const ES_GUARDAR = true
const NO_ES_GUARDAR = false

const (
	vacio estado = iota
	ocupado
	borrado
)

type celdaHash[K comparable, V any] struct {
	clave K
	dato  V
	estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	borrados int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]celdaHash[K, V], TAMANIO_INICIAL)
	for i := range tabla {
		tabla[i].estado = vacio
	}
	return &hashCerrado[K, V]{tabla, 0, 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashear(data []byte, largo_tabla int) int {
	hash := sha256.Sum256(data)
	hashInt := 0
	for i := 0; i < len(hash); i++ {
		hashInt = hashInt*256 + int(hash[i])
	}
	resultado := hashInt % largo_tabla
	return int(math.Abs(float64(resultado)))
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	indice, encontrado := hash.recorrerHash(clave_hasheada, clave, NO_ES_GUARDAR)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	hash.tabla[indice].estado = borrado
	hash.cantidad--
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	//si la clave ya existe, hay riesgo de que se guarde 2 veces... y pasa
	pertenece := hash.Pertenece(clave)
	if pertenece {
		hash.Borrar(clave)
	}
	//nosotros elegimos q si la clave ya esta, se sobreescribe
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	indice, _ := hash.recorrerHash(clave_hasheada, clave, ES_GUARDAR)
	hash.tabla[indice].dato = dato
	hash.tabla[indice].clave = clave
	hash.tabla[indice].estado = ocupado
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	panic("unimplemented")
}

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {
	panic("unimplemented")
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	indice, encontrado := hash.recorrerHash(clave_hasheada, clave, NO_ES_GUARDAR)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	_, encontrado := hash.recorrerHash(clave_hasheada, clave, NO_ES_GUARDAR)
	return encontrado
}

func (hash *hashCerrado[K, V]) recorrerHash(indice int, clave K, es_guardar bool) (int, bool) {

	if indice == len(hash.tabla) {
		return hash.recorrerHash(0, clave, es_guardar)
	}

	if hash.tabla[indice].estado == vacio { //funciona
		return indice, false
	}

	if hash.tabla[indice].estado == borrado {
		if es_guardar {
			return indice, false
		}
		return hash.recorrerHash(indice+1, clave, es_guardar)
	}

	if hash.tabla[indice].estado == ocupado && hash.tabla[indice].clave == clave {
		return indice, true
	}

	return hash.recorrerHash(indice+1, clave, es_guardar)
}
