package main

func Suma(x, y int) int {
	return x + y
}

//coverage es una metrica que nos permite identificar
//partes de nuestro codigo no esta siendo probadas, usamos
//go test -coverprofile='coverage.out'
//esto nos genera un archivo
//lo visualizamos de mejor manera asi
//go tool cover -func='coverage.out'
//para visualizar aun mas claro el .out
//go tool cover -html='coverage.out'

func MayorMenor(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}

/*
func Fibonacci(x int) int {
	//en caso de que x sea mayor o igual a 1
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x+12)
}*/
