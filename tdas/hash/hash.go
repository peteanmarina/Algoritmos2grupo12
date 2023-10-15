package diccionario

import (
	"crypto/sha256"
	"fmt"
	"math"
)

//go test -bench=. -benchmem

type estado int

const (
	TAMANIO_INICIAL  = 10
	ES_GUARDAR       = true
	NO_ES_GUARDAR    = false
	VALOR_AUMENTO    = 2
	VALOR_REDUCCION  = 2
	FACTOR_AUMENTO   = 0.7
	FACTOR_REDUCCION = 0.2
)

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
	cantidad int
}

type iterador[K comparable, V any] struct {
	actual_indice int
	ultima_indice int
	hash          *hashCerrado[K, V]
}

func (i *iterador[K, V]) HaySiguiente() bool {
	if i.actual_indice == -1 {
		return false
	}
	return i.actual_indice <= i.ultima_indice
}

func (i *iterador[K, V]) VerActual() (K, V) {
	if i.actual_indice > i.ultima_indice || i.actual_indice == -1 {
		panic("El iterador termino de iterar")
	}
	return i.hash.tabla[i.actual_indice].clave, i.hash.tabla[i.actual_indice].dato
}

func (i *iterador[K, V]) Siguiente() {
	if i.actual_indice > i.ultima_indice || i.actual_indice == -1 {
		panic("El iterador termino de iterar")
	}
	i.actual_indice++
	for i.actual_indice <= i.ultima_indice && i.hash.tabla[i.actual_indice].estado != ocupado {
		i.actual_indice++
	}

}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]celdaHash[K, V], TAMANIO_INICIAL)
	for i := range tabla {
		tabla[i].estado = vacio
	}
	return &hashCerrado[K, V]{tabla, 0}
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
	indice, encontrado := recorrerTablaCeldas(hash.tabla, clave_hasheada, clave, NO_ES_GUARDAR)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	retornado := hash.tabla[indice].dato
	hash.tabla[indice].estado = borrado
	hash.cantidad--
	hash.redimensionar()
	return retornado
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	pertenece := hash.Pertenece(clave)
	if pertenece {
		hash.Borrar(clave)
	}
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	indice, _ := recorrerTablaCeldas(hash.tabla, clave_hasheada, clave, ES_GUARDAR)
	hash.tabla[indice].dato = dato
	hash.tabla[indice].clave = clave
	hash.tabla[indice].estado = ocupado
	hash.cantidad++
	hash.redimensionar()
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	indice, encontrado := recorrerTablaCeldas(hash.tabla, clave_hasheada, clave, NO_ES_GUARDAR)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	clave_hasheada := hashear(convertirABytes[K](clave), len(hash.tabla))
	_, encontrado := recorrerTablaCeldas(hash.tabla, clave_hasheada, clave, NO_ES_GUARDAR)
	return encontrado
}

func recorrerTablaCeldas[K comparable, V any](tabla []celdaHash[K, V], indice int, clave K, es_guardar bool) (int, bool) {

	if indice == len(tabla) {
		return recorrerTablaCeldas(tabla, 0, clave, es_guardar)
	}

	if tabla[indice].estado == vacio { //funciona
		return indice, false
	}

	if tabla[indice].estado == borrado {
		if es_guardar {
			return indice, false
		}
		return recorrerTablaCeldas(tabla, indice+1, clave, es_guardar)
	}

	if tabla[indice].estado == ocupado && tabla[indice].clave == clave {
		return indice, true
	}

	return recorrerTablaCeldas(tabla, indice+1, clave, es_guardar)
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iterador[K, V]{hash.buscar_elemento_ocupado(0), len(hash.tabla) - 1, hash}
}

func (hash *hashCerrado[K, V]) buscar_elemento_ocupado(indice int) int {
	for i := indice; i < len(hash.tabla); i++ {
		if hash.tabla[i].estado == ocupado {
			return i
		}
	}
	return -1
}

func (hash *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for i := 0; i < len(hash.tabla); i++ {
		if hash.tabla[i].estado == ocupado {
			if !f(hash.tabla[i].clave, hash.tabla[i].dato) {
				break
			}
		}
	}
}

func (hash *hashCerrado[K, V]) redimensionar() {
	largo := len(hash.tabla)
	var capacidadNueva int
	factor_carga := float64(hash.cantidad) / float64(largo)
	if factor_carga >= FACTOR_AUMENTO {
		capacidadNueva = largo * VALOR_AUMENTO
	} else if factor_carga <= FACTOR_REDUCCION {
		capacidadNueva = largo / VALOR_REDUCCION
	} else {
		return
	}
	tablaNueva := make([]celdaHash[K, V], capacidadNueva)
	hash.reacomodarCeldas(tablaNueva)
}

func (hash *hashCerrado[K, V]) reacomodarCeldas(tabla_nueva []celdaHash[K, V]) {

	for _, celda := range hash.tabla {
		if celda.estado == ocupado {
			clave_hasheada := hashear(convertirABytes[K](celda.clave), len(tabla_nueva))
			indice, _ := recorrerTablaCeldas(tabla_nueva, clave_hasheada, celda.clave, ES_GUARDAR)
			tabla_nueva[indice].dato = celda.dato
			tabla_nueva[indice].clave = celda.clave
			tabla_nueva[indice].estado = ocupado
		}
	}
	hash.tabla = tabla_nueva
}
