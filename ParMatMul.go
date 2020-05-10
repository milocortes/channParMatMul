package main

import (
	"fmt"
	)

func main() {

	// Definimos nuestros canales
	norte:=make(chan int, 0)
	este:=make(chan int, 0)
	oeste:=make(chan int, 0)

	// Definimos la matriz que posmultiplica
	var matB=[][]int{{1,0,1},{0,1,0},{2,2,0}}
	// Definimos la matriz que posmultiplica
	var matA=[][]int{{1,2,3},{4,5,6},{7,8,9}}


	// Definimos un slice que agregará los resultados de las multiplicaciones
	var resultUno []int

	for i := 0; i <=2; i++ {
		resultUno=append(resultUno,realizarMul(matA[0][:],matB[:][i], este, norte, oeste))
	}

	fmt.Println(resultUno)

	// Definimos un slice que agregará los resultados de las multiplicaciones
	var resultDos []int

	for i := 0; i <=2; i++ {
		resultDos=append(resultDos,realizarMul(matA[1][:],matB[:][i], este, norte, oeste))
	}

	fmt.Println(resultDos)

	// Definimos un slice que agregará los resultados de las multiplicaciones
	var resultTres []int

	for i := 0; i <=2; i++ {
		resultTres=append(resultTres,realizarMul(matA[2][:],matB[:][i], este, norte, oeste))
	}

	fmt.Println(resultTres)

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

func realizarMul(vectorA []int,vectorB []int,este chan int, norte chan int, oeste chan int ) int {

	// Inicializamos los valores de los canales
	actulizarCanal(vectorB[2], norte)
	actulizarCanal(0, este)

	multiplier(vectorA[2], este,norte,oeste )

	// Actualizamos los valores de los canales
	actulizarCanal(<-oeste, este)
	actulizarCanal(vectorB[1], norte)

	multiplier(vectorA[1], este,norte,oeste )

	// Actualizamos los valores de los canales
	actulizarCanal(<-oeste, este)
	actulizarCanal(vectorB[0], norte)

	multiplier(vectorA[0], este,norte,oeste )
	return(<-oeste)

}
