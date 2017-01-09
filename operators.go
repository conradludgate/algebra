package algebra

import "math"

func Operation(a, b iExpression, symbol string, method func(a, b float64) float64) oExpression {
	return oExpression{a, b, symbol, method}
}

func Function(a iExpression, name string, method func(a float64) float64) fExpression {
	return fExpression{a, name, method}
}

func Add(a, b iExpression) oExpression {
	return Operation(a, b, "+",
		func(a, b float64) float64 {
			return a + b
		})
}

func Subtract(a, b iExpression) oExpression {
	return Operation(a, b, "-",
		func(a, b float64) float64 {
			return a - b
		})
}

func Multiply(a, b iExpression) oExpression {
	return Operation(a, b, "*",
		func(a, b float64) float64 {
			return a * b
		})
}

func Divide(a, b iExpression) oExpression {
	return Operation(a, b, "/",
		func(a, b float64) float64 {
			return a / b
		})
}

func Power(a, b iExpression) oExpression {
	return Operation(a, b, "^", math.Pow)
}

func Log(a, b iExpression) oExpression {
	return Operation(a, b, "log",
		func(a, b float64) float64 {
			return math.Log(a) / math.Log(b)
		})
}

func Sine(a iExpression) fExpression {
	//return fExpression{e: a, fName: "sin", fMethod: math.Sin}
	return Function(a, "sin", math.Sin)
}
