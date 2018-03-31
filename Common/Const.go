package Common

type Type int32

const (
	_ Type = iota
	INT
	DOUBLE
	STRING
)

type Operator int32

const (
	_ Operator = iota
	INTERSECT
	UNION
	EXCEPT
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
