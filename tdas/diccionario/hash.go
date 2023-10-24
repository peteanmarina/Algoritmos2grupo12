package diccionario

import (
	"encoding/binary"
	"fmt"
)

const (
	TAMANIO_INICIAL      = 13
	ES_GUARDAR           = true
	NO_ES_GUARDAR        = false
	VALOR_AUMENTO        = 2
	VALOR_REDUCCION      = 2
	FACTOR_AUMENTO       = 0.8
	FACTOR_REDUCCION     = 0.3
	PANIC_NO_ENCONTRADO  = "La clave no pertenece al diccionario"
	PANIC_TERMINO_ITERAR = "El iterador termino de iterar"
)

type estado int

const (
	vacio estado = iota
	ocupado
	borrado
)

type elementoHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []elementoHash[K, V]
	cantidad int
}

type iteradorDiccionario[K comparable, V any] struct {
	indice_actual int
	hash          *hashCerrado[K, V]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{crearTabla[K, V](TAMANIO_INICIAL), 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func crearTabla[K comparable, V any](tamanio int) []elementoHash[K, V] {
	tabla := make([]elementoHash[K, V], tamanio)
	for i := range tabla {
		tabla[i].estado = vacio
	}
	return tabla
}

func hashear(data []byte, largo_tabla int) int { // MurmurHash 2

	const (
		seed uint32 = 0xc70f6907
		m    uint32 = 0x5bd1e995
		r    uint32 = 24
	)

	hash := seed ^ uint32(len(data))
	for i := 0; i < len(data); i += 4 {
		if i+3 >= len(data) {
			break
		}
		k := binary.LittleEndian.Uint32(data[i:])
		k *= m
		k ^= k >> r
		k *= m

		hash *= m
		hash ^= k
	}
	hash ^= hash >> 13
	hash *= m
	hash ^= hash >> 15
	resultado := int(hash) % largo_tabla
	return resultado
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice, encontrado := hash.obtenerIndiceBuscado(clave, ES_GUARDAR)
	hash.tabla[indice].dato = dato
	hash.tabla[indice].clave = clave
	hash.tabla[indice].estado = ocupado
	if !encontrado {
		hash.cantidad++
		hash.redimensionarDeSerNecesario()
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, pertenece := hash.obtenerIndiceBuscado(clave, ES_GUARDAR)
	return pertenece
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	indice, _ := hash.obtenerIndiceBuscado(clave, NO_ES_GUARDAR)
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	indice, _ := hash.obtenerIndiceBuscado(clave, NO_ES_GUARDAR)
	retornado := hash.tabla[indice].dato
	hash.tabla[indice].estado = borrado
	hash.cantidad--
	hash.redimensionarDeSerNecesario()
	return retornado
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) obtenerIndiceBuscado(clave K, esGuardar bool) (int, bool) {
	bytes := convertirABytes[K](clave)
	clave_hasheada := hashear(bytes, len(hash.tabla))
	indice, encontrado := recorrerTablaCeldas(hash.tabla, clave_hasheada, clave, esGuardar)
	if !encontrado && !esGuardar {
		panic(PANIC_NO_ENCONTRADO)
	}
	return indice, encontrado
}

func recorrerTablaCeldas[K comparable, V any](tabla []elementoHash[K, V], indice int, clave K, es_guardar bool) (int, bool) {

	if indice == len(tabla) {
		return recorrerTablaCeldas(tabla, 0, clave, es_guardar)
	}

	if tabla[indice].estado == vacio {
		return indice, false
	}

	if tabla[indice].estado == ocupado && tabla[indice].clave == clave {
		return indice, true
	}

	return recorrerTablaCeldas(tabla, indice+1, clave, es_guardar)
}

func (hash *hashCerrado[K, V]) redimensionarDeSerNecesario() {
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
	if capacidadNueva <= TAMANIO_INICIAL {
		capacidadNueva = TAMANIO_INICIAL
	}
	hash.reacomodarCeldas(crearTabla[K, V](capacidadNueva))
}

func (hash *hashCerrado[K, V]) reacomodarCeldas(tabla_nueva []elementoHash[K, V]) {
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

func (hash *hashCerrado[K, V]) buscarElementoOcupado(indice int) int {
	for i := indice; i < len(hash.tabla); i++ {
		if hash.tabla[i].estado == ocupado {
			return i
		}
	}
	return -1
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorDiccionario[K, V]{hash.buscarElementoOcupado(0), hash}
}

func (hash *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for i := 0; i < len(hash.tabla); i++ {
		if hash.tabla[i].estado != ocupado {
			continue
		}
		if !f(hash.tabla[i].clave, hash.tabla[i].dato) {
			return
		}
	}
}

func (i *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return i.indice_actual != -1 && i.indice_actual < len(i.hash.tabla)
}

func (i *iteradorDiccionario[K, V]) VerActual() (K, V) {
	i.lanzarPanicTerminoIterar()
	return i.hash.tabla[i.indice_actual].clave, i.hash.tabla[i.indice_actual].dato
}

func (i *iteradorDiccionario[K, V]) Siguiente() {
	i.lanzarPanicTerminoIterar()
	i.indice_actual++
	i.indice_actual = i.hash.buscarElementoOcupado(i.indice_actual)
}

func (i *iteradorDiccionario[K, V]) lanzarPanicTerminoIterar() {
	if !i.HaySiguiente() {
		panic(PANIC_TERMINO_ITERAR)
	}
}
