package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)


const (
	entero1 = 10
	entero2 = 17
	entero3 = 22
	entero4 = 26
	entero5 = 7
	entero6 = 10
	cadena1 = "hola!"
	cadena2 = "como"
	cadena3 = "estas"
	cadena4 = "??"
)

func TestRecienCreada(t *testing.T) {
	t.Log("Verificamos que lista recien creada se comporte como vacia")
	lista := TDALista.CrearListaEnlazada[string]()
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	t.Log("Insertamos dos elementos y los borramos dejando lista vacia")
	lista.InsertarPrimero(cadena1)
	lista.InsertarUltimo(cadena2)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

func TestComportamiento(t *testing.T) {
	t.Log("Verificamos que insertar pocos elementos y borrarlos se realiza correctamente")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(entero1)
	require.Equal(t, entero1, lista.VerPrimero())
	lista.InsertarUltimo(entero2)
	require.Equal(t, entero2, lista.VerUltimo())
	lista.InsertarPrimero(entero3)
	require.Equal(t, entero3, lista.VerPrimero())
	lista.InsertarUltimo(entero4)
	require.Equal(t, entero4, lista.VerUltimo())

	require.Equal(t, 4, lista.Largo())

	require.Equal(t, entero3, lista.BorrarPrimero())
	require.Equal(t, entero1, lista.BorrarPrimero())
	require.Equal(t, entero2, lista.BorrarPrimero())
	require.Equal(t, entero4, lista.BorrarPrimero())

	require.Equal(t, true, lista.EstaVacia())
	lista.InsertarUltimo(entero5)
	require.Equal(t, entero5, lista.VerUltimo())

}

func TestVolumen(t *testing.T) {
	t.Log("Insertamos gran cantidad de elementos al principio y final")
	lista := TDALista.CrearListaEnlazada[int64]()
	var i int64
	for i = 5000; i >= 0; i-- {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
	}
	for i = 5001; i < 10000; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
	}

	t.Log("Borramos el primer elemento hasta dejar lista vacia")
	for i = 0; i < 10000; i++ {
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.Equal(t, 0, lista.Largo())
}

/////////////////////////////ITERADOR INTERNO////////////////////////////////

func TestIteradorInterno(t *testing.T) {

	var (
		suma    int16  = 0
		puntero *int16 = &suma
		lista          = TDALista.CrearListaEnlazada[int16]()
	)

	lista.InsertarPrimero(int16(entero1))
	lista.InsertarPrimero(int16(entero2))
	lista.InsertarPrimero(int16(entero3))
	lista.InsertarPrimero(int16(entero4))
	lista.InsertarPrimero(int16(entero5))
	lista.InsertarPrimero(int16(entero6))

	t.Log("Test con funcion que suma todos los elementos de la lista")
	lista.Iterar(func(v int16) bool {
		*puntero += v
		return true
	})

	require.Equal(t, int16(entero1+entero2+entero3+entero4+entero5+entero6), *puntero)

	*puntero = 0

	t.Log("Test con funcion que itera la lista, sumando todos los numeros pares y cortando la iteración al llegar al 7")

	lista.Iterar(func(v int16) bool {
		if v == 7 {
			return false
		}
		if v%2 == 0 {
			*puntero += v
		}
		return true
	})
	require.Equal(t, int16(entero6), *puntero)
}

/////////////////////////////ITERADOR EXTERNO////////////////////////////////

func TestComportamientoIterador(t *testing.T) {

	t.Log("Verificamos el comportamiento general del iterador")
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	iter.Insertar(1.0)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1.0, lista.VerPrimero())
	require.Equal(t, 1.0, lista.VerUltimo())
	require.Equal(t, 1.0, iter.VerActual())

	iter.Insertar(2.0)
	require.Equal(t, true, iter.HaySiguiente())
	iter.Siguiente()
	require.Equal(t, 1.0, iter.VerActual())
	require.Equal(t, 1.0, lista.VerUltimo())
	require.Equal(t, 2.0, lista.VerPrimero())
	require.Equal(t, true, iter.HaySiguiente())
	require.Equal(t, 2, lista.Largo())
	iter.Siguiente()

	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter.VerActual() })
}

func TestInsertarElementosAlIterar(t *testing.T) {
	t.Log("En lista vacia insertamos muchos elementos en distintos lugares. Iteramos con otro iterador")
	t.Log("Se verifica que el principio y fin de la lista se actualicen correctamente")
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(entero1)
	require.Equal(t, entero1, iterador.VerActual())
	require.Equal(t, entero1, lista.VerPrimero())
	require.Equal(t, entero1, lista.VerUltimo())
	iterador.Insertar(entero2)
	require.Equal(t, entero2, lista.VerPrimero())
	require.Equal(t, entero2, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, entero1, iterador.VerActual())
	iterador.Insertar(entero3)
	require.Equal(t, entero3, iterador.VerActual())
	iterador.Insertar(entero4)
	require.Equal(t, entero4, iterador.VerActual())
	iterador.Insertar(entero5)
	require.Equal(t, entero5, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, entero4, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, entero3, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, entero1, iterador.VerActual())
	iterador.Siguiente()
	iterador.Insertar(entero6)
	require.Equal(t, entero6, iterador.VerActual())
	require.Equal(t, entero6, lista.VerUltimo())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

	iterador2 := lista.Iterador()

	require.Equal(t, entero2, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, entero5, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, entero4, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, entero3, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, entero1, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, entero6, iterador2.VerActual())
	iterador2.Siguiente()

	require.Equal(t, entero2, lista.VerPrimero())
	require.Equal(t, entero6, lista.VerUltimo())

}

func TestCasosLimitesIterador(t *testing.T) {
	t.Log("Borramos estando el primer y ultimo elemento")
	t.Log("Verificamos que podamos insertar elementos cuando iteramos una lista vacia")
	lista1 := TDALista.CrearListaEnlazada[int]()

	iter1 := lista1.Iterador()
	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter1.Borrar() })
	iter1.Insertar(1)
	require.Equal(t, 1, lista1.Largo())
	iter1.Borrar()
	require.Equal(t, 0, lista1.Largo())

	lista2 := TDALista.CrearListaEnlazada[int]()
	iter2 := lista2.Iterador()

	iter2.Insertar(1)
	require.Equal(t, 1, iter2.VerActual())
	iter2.Insertar(2)
	require.Equal(t, 2, iter2.VerActual())
	iter2.Insertar(3)
	require.Equal(t, 3, iter2.VerActual())

	require.Equal(t, true, iter2.HaySiguiente())
	iter2.Siguiente()
	require.Equal(t, 2, iter2.VerActual())

	require.Equal(t, true, iter2.HaySiguiente())
	iter2.Siguiente()
	require.Equal(t, 1, iter2.VerActual())

	require.Equal(t, 1, lista2.VerUltimo())
	iter2.Borrar()
	require.Equal(t, 2, lista2.VerUltimo())
}
