package algebra

import (
	"strconv"
)

type iExpression interface {
	Parse(values map[string]float64) float64
	String() string
	Simplify() iExpression
}

type (
	Variable string
	Number   float64
)

// Operator Expression
type oExpression struct {
	left, right iExpression
	operation   string
}

// Function Expression
type fExpression struct {
	expressions []iExpression
	function    string
}

// Parse
func (V Variable) Parse(values map[string]float64) float64 {
	return values[string(V)]
}

func (N Number) Parse(values map[string]float64) float64 {
	return float64(N)
}

func (O oExpression) Parse(values map[string]float64) float64 {
	return Operators[O.operation](O.left.Parse(values), O.right.Parse(values))
}

func (F fExpression) Parse(values map[string]float64) float64 {
	n := len(F.expressions)
	if n == 0 {
		return Functions[F.function]([]float64{})
	}
	if n == 1 {
		return Functions[F.function]([]float64{F.expressions[0].Parse(values)})
	}
	if n == 2 {
		return Functions[F.function]([]float64{F.expressions[0].Parse(values),
			F.expressions[1].Parse(values)})
	}

	var v []float64
	for i, e := range F.expressions {
		v[i] = e.Parse(values)
	}
	return Functions[F.function](v)
}

// String
func (V Variable) String() string {
	return string(V)
}

func (N Number) String() string {
	return strconv.FormatFloat(float64(N), 'f', -1, 64)
}

func (O oExpression) String() string {
	return "(" + O.left.String() + " " + O.operation + " " + O.right.String() + ")"
}

func (F fExpression) String() string {
	if len(F.expressions) == 0 {
		return F.function + "()"
	}
	s := F.function + "(" + F.expressions[0].String()

	for _, e := range F.expressions[1:] {
		s += ", " + e.String()
	}

	return s + ")"
}

// Simplify
func (V Variable) Simplify() iExpression {
	return V
}

func (N Number) Simplify() iExpression {
	return N
}

func (O oExpression) Simplify() iExpression {
	if O.left == O.right {

	}
}
