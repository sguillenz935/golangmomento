package calculator

// 2 reglas importantes:
// Si quiere que algo sea privado, escriba su nombre con la primera letra en minúscula.
// Si quiere que algo sea público, escriba su nombre con la primera letra en mayúscula.

var logMessage = "[LOG]"

// Version of the calculaotr
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integers numbers

func Sum(number1, number2 int) int {
	return number1 + number2
}
