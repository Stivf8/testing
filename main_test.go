//creamos nuestro test de funcion Suma

package main

import "testing"

/*//recibimos tipo testing.T, go verifica que el test este marchando bien.
//test unitario
func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	//validamos resultado, enviando los resultados y el esperado
	if total != 10 {
		t.Errorf("Error, valor %d esperado %d", total, 10)
	}
}*/
//test de varios casos
func TestSuma(t *testing.T) {
	//creamos slices de structs tablaCases con 3 valores,
	tableCases := []struct {
		//espera dos valores enteros la funcion Suma
		a int
		b int
		//devuelve un valor entero
		n int
	}{
		//definimos los atributos
		//ejemplo, 1 es a dos es b y n es el esperado
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tableCases {
		total := Suma(item.a, item.b)
		if total != item.n {
			t.Errorf("Failed: Incorrect sum, got %d expected %d", total, item.n)
		}
	}
}
