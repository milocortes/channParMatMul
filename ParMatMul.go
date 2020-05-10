package main

import (
	"fmt"
	)

func main() {

	// Definimos nuestros canales
	norte:=make(chan int, 0)
	este:=make(chan int, 0)
	oeste:=make(chan int, 0)

	// Definimos un slice que agregará los resultados de las multiplicaciones
	var result []int
	// Definimos un slice con los valores de la última columna de la matriz que posmultiplica
	// Definimos un slice conmpuesto por tres ceros


	var matB=[][]int{{1,0,1},{0,1,0},{2,2,0}}

	for i := 0; i <=2; i++ {
		result=append(result,realizarMul(matB[:][i], este, norte, oeste))
	}

	fmt.Println(result)
}

func multiplier(entero int,este chan int, norte chan int, oeste chan int )  {
	go func() {
		oeste<-(entero*(<-norte))+(<-este)
	}()
}

func actulizarCanal(entero int, canal chan int)  {
	go func() {
		canal<-entero
	}()
}

func realizarMul(miSlice []int,este chan int, norte chan int, oeste chan int ) int {

	// Inicializamos los valores de los canales
	actulizarCanal(miSlice[2], norte)
	actulizarCanal(0, este)

	multiplier(9, este,norte,oeste )

	// Actualizamos los valores de los canales
	actulizarCanal(<-oeste, este)
	actulizarCanal(miSlice[1], norte)

	multiplier(8, este,norte,oeste )

	// Actualizamos los valores de los canales
	actulizarCanal(<-oeste, este)
	actulizarCanal(miSlice[0], norte)

	multiplier(7, este,norte,oeste )
	return(<-oeste)

}
