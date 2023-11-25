package main

import (
	"fmt"
	"strconv"
)

func main() {

	var tema1 string = "tema1 variables, funciones y paquetes"
	fmt.Println(tema1)

	// podemos declarar mas variables de una sola sentada

	var nombre, apellido string = "santiago", "guillen"
	var anios int = 33
	fmt.Println("Mi nombre es", nombre, ", mi apellido es", apellido, "y tengo ", anios)

	// esta es la manera mas normal de hacer en go a la hora de declarar

	fistname, lastname := "Santiagos", "Guillens"
	age := 33
	fmt.Println(fistname, lastname, age)

	// Constantes

	const HTTPSstatusOK = 200

	// podemos declarar conmstantes que no vayamos a usar.
	const (
		STATTUSOK             = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	// Tipos de DATOS: Numero enteros

	// Existen variuos tipos de int pero no seran muy relevantes hasta el momento

	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807

	fmt.Println(integer8, integer16, integer32, integer64)

	// Como es evidente, no se pueden sumar diferentes datos de diferentes ints

	// Seran imporantes las runes, que es un alias para el tipo de datos int32 --> usado para representar caracter Unicode.

	rune := 'G'
	fmt.Println(rune)

	// Numeros de punto flotante

	// var float32 float32 = 2147483647
	// var float64 float64 = 9223372036854775807
	// fmt.Println(float32, float64)

	// Comando para ver el maximo de valores que adminte float32 y float64
	// fmt.Println(math.MaxFloat32, math.MaxFloat64)

	// esto para mates esta ptm
	const e = 2.71828

	// Valores booleanos

	// Los valores booleanos son o true o false

	var FeatureFlag bool = true

	fmt.Println(FeatureFlag)

	// Cadenas

	// es lo que llevamos usando hace un rato, pero hay que tener en cuenta los caracteres de escape.

	// A veces necesitará usar caracteres de escape. Para incluirlos en Go, use una barra diagonal inversa (\) antes del carácter.
	//  A continuación se muestran los ejemplos más comunes del uso de caracteres de escape:

	/*
		\n para líneas nuevas
		\r para retornos de carro
		\t para pestañas (nos da un poco mas de espacio)
		\' para comillas simples
		\" para comillas dobles
		\\ para barras diagonales inversas
	*/

	fullName := "Santiagos Guillen (alias \"flagpe\")\n"

	fmt.Println(fullName)

	// VALORES PREDETERMIDANDOS

	/*
		A continuación se presenta una lista con algunos valores predeterminados para los tipos que hemos explorado hasta ahora:

			0 para los tipos int (y todos sus subtipos, como int64)
			+0.000000e+000 para los tipos float32 y float64
			false para los tipos bool
			Un valor vacío para los tipos string

	*/

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// Conversion de Tipos.

	var interger16 int16 = 127
	var interger32 int32 = 32767

	fmt.Println(interger16 + int16(interger32))

	// Otra manera de conversion es usar el paquete strconv

	i, _ := strconv.Atoi("-42") /* requiere 2 valores*/
	s := strconv.Itoa(-42)

	fmt.Println(i, s)
	/*
		Observe que hay un carácter de subrayado (_) que se usa como el nombre de una variable en el código anterior.
		En Go, el objeto _ significa que no vamos a usar el valor de esa variable y que queremos omitirlo.
		De lo contrario, el programa no se compilará porque necesitamos usar todas las variables declaradas.
	*/

}
