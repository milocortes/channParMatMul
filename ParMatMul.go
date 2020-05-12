package main

import "fmt"


/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Definimos el proceso Zero que inicializa la suma parcial como una función que recibe un tipo de canal
que es send-only de int (dado que chan<-int).
*/

func zero(entero int,out chan<-int)  {
	out<-entero
	//close(out)
}

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Definimos el proceso multiplier que recibe una suma parcial del este, y agrega a esta el resultado de
multiplicar su elemento por el valor recibido del norte. El resultado es enviado por un canal al oestr.
Además, el proceso envía al sur sobre un canal el valor recibido del norte.

Este proceso lo definimos como una función que recibe dos tipos de canales, "oeste" y "sur",
que son send-only de int (dado que chan<-int) y dos tipos de canales, "este" y "norte" , que son
read-only de int (dado que <-chan int). Además, recibe un entero el cual es su elemento.
*/

func multiplier(firstElement int,este <-chan int, norte <-chan int, oeste chan<- int , sur chan<- int  )  {

		secondElement:=(<-norte)
		sum:=(<-este)

		sum=(firstElement*(secondElement))+(sum)

		oeste<-sum
		sur<-secondElement
		//close(oeste)
}

/*
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Definimos el proceso Source el cual es una función que recibe un entero y un tipo de canal que
es send-only de int (dado que chan<-int). Envía el valor recibido al canal recibido.
*/

func source(entero int,out chan <- int)  {
		out<-entero
		//close(out)
}

func main()  {

	cz13:=make(chan int, 0)
	cz26:=make(chan int, 0)
	cz39:=make(chan int, 0)

	c23:=make(chan int, 0)
	c12:=make(chan int, 0)
	c56:=make(chan int, 0)
	c45:=make(chan int, 0)
	c89:=make(chan int, 0)
	c78:=make(chan int, 0)

	cr11:=make(chan int, 0)
	cr24:=make(chan int, 0)
	cr37:=make(chan int, 0)

	c14:=make(chan int, 0)
	c25:=make(chan int, 0)
	c36:=make(chan int, 0)
	c47:=make(chan int, 0)
	c58:=make(chan int, 0)
	c69:=make(chan int, 0)
	
	sink1:=make(chan int, 0)
	sink2:=make(chan int, 0)
	sink3:=make(chan int, 0)

	s1:=make(chan int, 0)
	s2:=make(chan int, 0)
	s3:=make(chan int, 0)

	var vecSoureUno=[]int{1,0,2}
	var vecSoureDos=[]int{0,1,2}
	var vecSoureTres=[]int{1,0,0}

	var resultadosUno=[]int{}
	var resultadosDos=[]int{}
	var resultadosTres=[]int{}

	for i := 0; i <=2; i++ {

		go zero(0,cz13)
		go zero(0,cz26)
		go zero(0,cz39)

		go source(vecSoureUno[i],s1)
		go source(vecSoureDos[i],s2)
		go source(vecSoureTres[i],s3)

		go multiplier(3, cz13, s3, c23,c36)
		go multiplier(2, c23, s2, c12,c25)
		go multiplier(1, c12, s1, cr11,c14)

		go multiplier(6, cz26, c36, c56,c69)
		go multiplier(5, c56, c25, c45,c58)
		go multiplier(4, c45, c14, cr24,c47)

		go multiplier(9, cz39, c69, c89,sink3)
		go multiplier(8, c89, c58, c78,sink2)
		go multiplier(7, c78, c47, cr37,sink1)

		resultadosUno=append(resultadosUno,<-cr11)
		resultadosDos=append(resultadosDos,<-cr24)
		resultadosTres=append(resultadosTres,<-cr37)
	}

	fmt.Println(resultadosUno)
	fmt.Println(resultadosDos)
	fmt.Println(resultadosTres)

}
