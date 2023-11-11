// Los paquetes son un grupo de funciones
package main

// FMT es una de las principales librerias de GOlang https://pkg.go.dev/fmt //
import (
	"fmt"
	"reflect"

	"rsc.io/quote"
)

// Si se nos complica descargar un poquete automaticamente, ejecutamos eset comando en la cmd en direccion a nuestra carpeta (Con vscode)

// LLamamos a la funcion para usar el print que almacena el modulo quote, el orden seria func>fmt>quote
func main()  {
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
	myint= myint + 4
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

	

}