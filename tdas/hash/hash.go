package diccionario

import (
	"fmt"
)

type estado int

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
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]celdaHash[K, V], 1) //tamaño inicial??
	for i := range tabla {
		tabla[i].estado = vacio
	}
	return &hashCerrado[K, V]{tabla, 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (*hashCerrado[K, V]) Borrar(clave K) V {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Cantidad() int {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Guardar(clave K, dato V) {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Obtener(clave K) V {
	panic("unimplemented")
}

func (*hashCerrado[K, V]) Pertenece(clave K) bool {
	panic("unimplemented")
}
