package algebra

import "strconv"

// Structure
type iExpression interface {
	Parse(values map[string]float64) (float64, error)
	String() string // For fmt/log/other print functions
}

type (
	Variable string
	Number   float64
)

type oExpression struct { // Operator Expression
	left, right iExpression
	opSymbol    string
	opMethod    func(a, b float64) float64
}

type fExpression struct { // Function Expression
	e       iExpression
	fName   string
	fMethod func(values float64) float64
}

// Parse
func (V Variable) Parse(values map[string]float64) (float64, error) {
	return 0, nil
}

func (N Number) Parse(values map[string]float64) (float64, error) {
	return float64(N), nil
}

func (o oExpression) Parse(values map[string]float64) (float64, error) {
	return 0, nil
}

func (F fExpression) Parse(values map[string]float64) (float64, error) {
	return 0, nil
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
