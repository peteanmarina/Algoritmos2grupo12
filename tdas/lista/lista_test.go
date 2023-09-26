package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	mensaje_generico_lista_vacia  = "La lista esta vacia"
	mensaje_generico_no_elementos = "No hay mas elementos en la lista"
	mensaje_iterador              = "El iterador termino de iterar"
)

func TestRecienCreada(t *testing.T) {
	t.Log("Verificamos que lista recien creada se comporte como vacia")
	lista := TDALista.CrearListaEnlazada[string]()
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

	t.Log("Insertamos dos elementos y los borramos dejando lista vacia")
	lista.InsertarPrimero("uno")
	lista.InsertarUltimo("diez")
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

func TestComportamiento(t *testing.T) {
	t.Log("Verificamos que insertar pocos elementos y borrarlos se realiza correctamente")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	lista.InsertarUltimo(5)
	require.Equal(t, 5, lista.VerUltimo())

	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 5, lista.BorrarPrimero())

	require.Equal(t, true, lista.EstaVacia())
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())

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
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, mensaje_generico_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

/////////////////////////////ITERADOR INTERNO////////////////////////////////

func TestIteradorInterno(t *testing.T) {

	var (
		suma    int16  = 0
		puntero *int16 = &suma
		lista          = TDALista.CrearListaEnlazada[int16]()
	)

	lista.InsertarPrimero(int16(2))
	lista.InsertarPrimero(int16(5))
	lista.InsertarPrimero(int16(7))
	lista.InsertarPrimero(int16(11))
	lista.InsertarPrimero(int16(14))
	lista.InsertarPrimero(int16(21))

	t.Log("Test con funcion que suma todos los elementos de la lista")
	lista.Iterar(func(v int16) bool {
		*puntero += v
		return true
	})

	require.Equal(t, int16(60), *puntero)

	*puntero = 0

	t.Log("Test con funcion que suma todos los numeros de la lista impares exceptuando el 7")
	lista.Iterar(func(v int16) bool {
		if v == 7 {
			return false
		}
		if v%2 == 0 {
			*puntero += v
		}
		return true
	})
	require.Equal(t, int16(14), *puntero)
}

/////////////////////////////ITERADOR EXTERNO////////////////////////////////

func TestInsertarEnMedioIterador(t *testing.T) {
	t.Log("En lista vacia insertamos muchos elementos en distintos lugares. Iteramos con otro iterador")

	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(1)
	require.Equal(t, 1, iterador.VerActual())
	iterador.Insertar(2)
	require.Equal(t, 2, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, 1, iterador.VerActual())
	iterador.Insertar(3)
	require.Equal(t, 3, iterador.VerActual())
	iterador.Insertar(4)
	require.Equal(t, 4, iterador.VerActual())
	iterador.Insertar(5)
	require.Equal(t, 5, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, 4, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, 3, iterador.VerActual())
	iterador.Siguiente()
	require.Equal(t, 1, iterador.VerActual())
	iterador.Siguiente()
	iterador.Insertar(6)
	require.Equal(t, 6, iterador.VerActual())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

	iterador2 := lista.Iterador()

	require.Equal(t, 2, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, 5, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, 4, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, 3, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, 1, iterador2.VerActual())
	iterador2.Siguiente()
	require.Equal(t, 6, iterador2.VerActual())
	iterador2.Siguiente()

	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())

}

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

	require.PanicsWithValue(t, mensaje_iterador, func() { iter.Siguiente() })
	require.PanicsWithValue(t, mensaje_iterador, func() { iter.VerActual() })
}

func TestCasosLimitesIterador(t *testing.T) {
	t.Log("Borramos estando el primer y ultimo elemento")
	t.Log("verificamos que podamos insertar elementos cuando iteramos una lista vacia")
	lista1 := TDALista.CrearListaEnlazada[int]()

	iter1 := lista1.Iterador()
	require.PanicsWithValue(t, mensaje_iterador, func() { iter1.Borrar() })
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

func TestVerElementosIterador(t *testing.T) {
	t.Log("Verificamos que la lista se itere correctamente")

	elementos := [4]int{10, 17, 22, 26}
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(int(10))
	lista.InsertarUltimo(int(17))
	lista.InsertarUltimo(int(22))
	lista.InsertarUltimo(int(26))

	iter := lista.Iterador()

	for _, elemento := range elementos {
		require.True(t, iter.VerActual() == elemento)
		if iter.HaySiguiente() {
			iter.Siguiente()
		}
	}
}

func TestBuscarElementoIterador(t *testing.T) {
	t.Log("Recorremos la lista buscando un elemento")

	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarUltimo("hola")
	lista.InsertarUltimo("como")
	lista.InsertarUltimo("estas")
	lista.InsertarUltimo("?")

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual() == "estas" {
			require.True(t, iter.VerActual() == "estas")
		} else {
			require.False(t, iter.VerActual() == "estas")
		}
		iter.Siguiente()
	}
}

func TestExtremosListaAlIterar(t *testing.T) {
	t.Log("Verificamos que principio y fin se modifiquen al insertar y borrar elementos")

	lista := TDALista.CrearListaEnlazada[string]()

	iter := lista.Iterador()
	iter.Insertar("1")
	iter.Insertar("2")
	iter.Insertar("3")
	iter.Insertar("4")
	iter.Insertar("5")
	iter.Insertar("6")
	require.Equal(t, "6", lista.VerPrimero())
	require.Equal(t, "1", lista.VerUltimo())

	iter.Borrar()
	iter.Borrar()
	iter.Borrar()

	require.Equal(t, "3", lista.VerPrimero())
	require.Equal(t, "1", lista.VerUltimo())
}
