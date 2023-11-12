// Los paquetes son un grupo de funciones
package main

// FMT es una de las principales librerias de GOlang https://pkg.go.dev/fmt //
import (
	"container/list"
	"fmt"
	"reflect"

	"rsc.io/quote"
)

// Si se nos complica descargar un poquete automaticamente, ejecutamos eset comando en la cmd en direccion a nuestra carpeta (Con vscode)

// LLamamos a la funcion para usar el print que almacena el modulo quote, el orden seria func>fmt>quote
func main() {
	fmt.Println(quote.Go())

	// Cuidado con las comillas simples, no vale para este caso y cumple otras funciones
	var myfiststring string = "adena de texto si soy"
	fmt.Println(myfiststring)

	// Se puede escribir o no el tipo de texto, pero si lo definimos no podremos cambiarlo
	/* var myfiststring string = "adena de texto si soy"
	fmt.Println(myfiststring)*/

	// Cambiamos la cadena, cuidado con ello
	myfiststring = "Che pibe, cambie de cadena"
	fmt.Println(myfiststring)

	var myint int = 55
	myint = myint + 4
	fmt.Println(myint)

	// Juntar 2 valores, no es muy comun pero a tope
	fmt.Println(myfiststring, myint)

	fmt.Println(reflect.TypeOf(myint))

	var myfloat float64 = 6.55

	fmt.Println(myfloat)
	fmt.Println(reflect.TypeOf(myfloat))

	// Cuidado con el orden
	fmt.Println(myfloat + float64(myint))

	var mybool bool = false
	mybool = true
	fmt.Println(mybool)

	// Varaible declara e inicializada de manera abreviada
	mystring3 := "esto es una cadena"
	fmt.Println(mystring3)

	// Constantess
	const myconst = "Esto es una constante"

	// Control de flujo

	myint = 10
	mystring3 = "Hola"

	if myint == 10 && mystring3 == "Hola" {
		fmt.Println("el valor es 10")
	} else if myint == 11 || mystring3 == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

	// estructuras de datos

	// arrays

	var myArray [3]int 
	myArray[0] = 1
	myArray[1] = 22
	myArray[2] = 33
	fmt.Println(myArray)

	// var myArray [3]float64 otro tipo

	// map

	myMap := make(map[string]int)
	myMap["santiagos"] = 33
	myMap["teton"] = 99
	myMap["flagpe"] = 03
	fmt.Println(myMap["teton"])
	
	// otra manera 
	
	myMap2 := map[string]int{"santiagos": 33, "teton": 99, "flagpe": 03 }
	fmt.Println(myMap2)

	//listas

	mylist := list.New()

	mylist.PushBack(1)
	mylist.PushBack(2)
	mylist.PushBack(3)
	fmt.Println(mylist.Back().Value)

	// bucles
	 
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}


	// funciones

	// myfunction()

	fmt.Println(myfunction())


	// estructura 

	type MyStruct struct {
		name string
		age int 
	}

	mystruct := MyStruct{"santiagos", 33}
	fmt.Println(mystruct)

}

func myfunction() string {
	return "Mi funcion"
}
