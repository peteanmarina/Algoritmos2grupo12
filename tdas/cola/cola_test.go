package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Desencolar y ver primero a una cola vacia")
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestAlgunosElementos(t *testing.T) {
	t.Log("Encolar y desencolar algunos elementos, quedando una cola vacia al final")
	cola := TDACola.CrearColaEnlazada[bool]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
	cola.Encolar(true)
	require.False(t, cola.EstaVacia())
	cola.Encolar(false)
	require.False(t, cola.EstaVacia())
	require.Equal(t, true, cola.VerPrimero())
	require.Equal(t, true, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, false, cola.VerPrimero())
	require.Equal(t, false, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestMuchosElementos(t *testing.T) {
	t.Log("Encolar y desencolar muchos elementos, quedando una cola vacia al final")
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
		require.True(t, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	}

}
