package algebra

import "strconv"

type iExpression interface {
	Parse(values map[string]float64) float64
	String() string
}

type (
	Variable string
	Number   float64
)

// Operator Expression
type oExpression struct {
	left, right iExpression
	opSymbol    string
	opMethod    func(a, b float64) float64
}

// Function Expression
type fExpression struct {
	e       iExpression
	fName   string
	fMethod func(values float64) float64
}

// Parse
func (V Variable) Parse(values map[string]float64) float64 {
	return values[string(V)]
}

func (N Number) Parse(values map[string]float64) float64 {
	return float64(N)
}

func (O oExpression) Parse(values map[string]float64) float64 {
	return O.opMethod(O.left.Parse(values), O.right.Parse(values))
}

func (F fExpression) Parse(values map[string]float64) float64 {
	return F.fMethod(F.Parse(values))
}

// String
func (V Variable) String() string {
	return string(V)
}

func (N Number) String() string {
	return strconv.FormatFloat(float64(N), 'f', -1, 64)
}

func (o oExpression) String() string {
	return "(" + o.left.String() + " " + o.opSymbol + " " + o.right.String() + ")"
}

func (F fExpression) String() string {
	return F.fName + "(" + F.e.String() + ")"
}
