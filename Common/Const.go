package Common

type Type int32

const (
	_ Type = iota
	INT
	DOUBLE
	STRING
)

type JoinType int32

const (
	_ JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	INNERJOIN
)

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
)

type ComparisonOperator int32

const (
	_ ComparisonOperator = iota
	EQ
	NEQ
	LT
	LTE
	GT
	GTE
)

type Quantifier int32

const (
	_ Quantifier = iota
	ALL
	DISTINCT
	SOME
	ANY
)

type Order int32

const (
	_ Order = iota
	ASC
	DESC
	FIRST
	LAST
)

type Not int32
type As int32

type FuncType int32

const (
	_ FuncType = iota
	AGGREGATE
	NORMAL
)
