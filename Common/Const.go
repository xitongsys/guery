package Common

type Type int32

const (
	_ Type = iota
	INT
	DOUBLE
	STRING
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
	case "ASTERISK":
		res = ASTERISK
	case "SLASH":
		res = SLASH
	case "PERCENT":
		res = PERCENT
	case "PLUS":
		res = PLUS
	case "MINUS":
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

type JoinType int32

const (
	_ JoinType = iota
	INNER
	LEFT
	RIGHT
	FULL
)

type Not int32
type As int32
