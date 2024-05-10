package diccionario_test

import (
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararCadenas(cad1, cad2 string) int {
	if cad1 < cad2 {
		return -1
	} else if cad1 > cad2 {
		return 1
	}
	return 0
}

func comparaEnteros(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0

}

func TestDiccAbbVacio(t *testing.T) {

	t.Log("Prueba que el ABB esta vacio")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })

}

func TestUnElementoABB(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	dic.Guardar("A", "10")
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, "10", dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}
