package algebra

import "math"

func add(a, b float64) float64 {
	return a + b
}

func Add(a, b iExpression) oExpression {
	return oExpression{left: a, right: b, opSymbol: "+", opMethod: add}
}

func Sine(a iExpression) fExpression {
	return fExpression{e: a, fName: "sin", fMethod: math.Sin}
}
