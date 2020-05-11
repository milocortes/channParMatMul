# Uso de canales para multiplicación paralela de matrices

Un ***canal*** conecta un proceso remitente con un proceso receptor. Los canales son ***tipados***, lo que significa que 
debemos declara el tipo de mensajes que pueden ser enviados en el canal. En lo que sigue, suponemos que los canales son síncronos.
En sistemas operativos, los canales son denominados ***pipes***; los cuales permiten construir programas conectando un conjunto ya existente de programas.


Con canales podemos realizar ***pipelined computation***. En este caso, usaremos canales para 
realizar una multiplicación paralela de matrices.

## Brevísimas notas de concurrencia en Go

#### Goroutines

En Go, cada actividad actividad que se ejecuta de manera conrrente es llamada una ***goroutine***. Cuando un programa comienza, hay una sola goroutine que llama a la función ```main``` , de manera que podemos llamarla la   ***main goroutine***. Podemos crear nuevas goroutines con la declaración  ```go```. De forma sintática, una declaración ```go``` es un método o función ordinaria que es precedido por la palabra reservada ```go```.  

#### Canales

Si las goroutines son las actividades de un programa concurrente en Go, los ***channels*** son las conexiones entre estas.  Un canal es un mecanismo de comunicación que permite a una goroutine enviar valores a otra goroutine. Cada canal es un conducto de valores de un tipo particular, denominado el ***elemento tipo*** del canal. El tipo de un canal cuyos elementos tienen tipo ```int``` se escribe ```chan int```.

Para crear un canal, usamos la función ```make```:
```go
ch:=make(chan int) // ch tiene tipo 'chan int'

```
Un canal es una **referencia** a la estructura de datos creada por ```make```. Cuando copiamos un canal o lo pasamos como un argumento a una función, estamos copiando una referencia, de manera que el **caller** y el **calle** refieren a la misma estructura de datos. 

Un canal tiene dos operaciones principales, ***send*** y ***receive***, conocidas colectivamente como ***comunicaciones***. Una declaración send transmite un valor a través del canal de una goroutine a otra goroutine que esté ejecutando una expresión receive. Ambas operaciones son escritas usando el operador  ```<-```. En una decaración send, el operador ```<-``` del canal y del valor operado. En una expresión receive, ```<-``` precede al canal.
Una expresión receive cuyo resultado no es usado es una declaración válida:

```go
ch<- x // Una declaración send
x=<- ch // Una expresión receive en una declaración de asignación
<-ch 	// Una declaración receive, el resultado es descartado
```

Los canales presentan una tercera operación, ***close***, que define una bandera que indica que no se enviarán más valores por el canal. Para cerrar un canal, usamos la función ```close```:
```go
close(ch)
```

Un canal creado con una llamada simple a ```make``` es denomidado como un canal ***unbuffered***, pero ```make``` acepta un segundo argumento, un entero denominado como la ***capacidad*** del canal. Si la capacidad es distinta de cero, ```make``` crea un canan ***buffered**:

```go
ch = make(chan int)// unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

#### Pipelines

Los canales pueden ser usados para conectar goroutines de manera que el output de una es el intput de la otra. Esto es llamado un ***pipeline***. El programa que se presenta abajo consiste de tres goroutines conectadas por dos canales, como muestra la siguiente imagen:

![](images/pipeline.png) 


La primera goroutine, ***counter***, genera los enteros 0,1,2,..., y los envía sobre el canal a la segunda goroutine, ***square***, que recibe cada valor, lo eleva al cuadrado, y envía el resultado sobre otro canal a la tercera goroutine, ***printer***, que recibe los valores y los imprime. 


```go
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()
	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

```
Si queremos enviar una cantidad finita de valores  a través del pipeline, es preciso indicarle a la goroutine remitente que se comunique con la goroutine receptora para notificarle que más valores no serán enviados por el canal, para así evitar que permanezca esperando un valor. Lo anterior es realizado ***cerrando*** el canal con la función ```close```.

En el siguiente pipeline, cuando la goroutine counter termina el loop, cierra el canal ```naturals```, provocando que la goroutine squarer finalice su loop y cierre el canal ```squares```. Por último, la goroutine main finaliza su loop y termina el programa.

```go
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}


```

 
 