package Util

///////////////////////
type Operator int32

func NewOperator(name string) *Operator {
	var res Operator
	switch name {
	case "INTERSECT":
		res = INTERSECT
	case "UNION":
		res = UNION
	case "EXCEPT":
		res = EXCEPT
	case "AND":
		res = AND
	case "OR":
		res = OR
	case "*":
		res = ASTERISK
	case "/":
		res = SLASH
	case "%":
		res = PERCENT
	case "+":
		res = PLUS
	case "-":
		res = MINUS
	case ">":
		res = GT
	case "<":
		res = LT
	case ">=":
		res = GTE
	case "<=":
		res = LTE
	case "=":
		res = EQ
	case "<>":
		res = NEQ
	}
	return &res
}

const (
	_ Operator = iota
	INTERSECT
	UNION
	EXCEPT
	AND
	OR
	ASTERISK
	SLASH
	PERCENT
	PLUS
	MINUS
	EQ
	NEQ
	LT
	LTE
	GT
	GTE
)
