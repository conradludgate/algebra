package algebra

func Operation(a, b iExpression, operator string) oExpression {
	return oExpression{a, b, operator}
}

func Function(function string, expressions ...iExpression) fExpression {
	return fExpression{expressions, function}
}

// func Add(a, b iExpression) oExpression {
// 	return Operation(a, b, "+",
// 		func(a, b float64) float64 {
// 			return a + b
// 		})
// }

// func Subtract(a, b iExpression) oExpression {
// 	return Operation(a, b, "-",
// 		func(a, b float64) float64 {
// 			return a - b
// 		})
// }

// func Multiply(a, b iExpression) oExpression {
// 	return Operation(a, b, "*",
// 		func(a, b float64) float64 {
// 			return a * b
// 		})
// }

// func Divide(a, b iExpression) oExpression {
// 	return Operation(a, b, "/",
// 		func(a, b float64) float64 {
// 			return a / b
// 		})
// }

// func Power(a, b iExpression) oExpression {
// 	return Operation(a, b, "^", math.Pow)
// }

// func Log(a, b iExpression) oExpression {
// 	return Operation(a, b, "log",
// 		func(a, b float64) float64 {
// 			return math.Log(a) / math.Log(b)
// 		})
// }

// func Sine(a iExpression) fExpression {
// 	//return fExpression{e: a, fName: "sin", fMethod: math.Sin}
// 	return Function(a, "sin", math.Sin)
// }

var Operators map[string](func(a, b float64) float64) {
	"+": (func(a, b float64) float64 { return a + b }),
	"-": (func(a, b float64) float64 { return a - b })
}

var Functions map[string]func(values []float64) float64
