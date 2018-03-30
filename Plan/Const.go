package Plan

type Operator int32

const (
	INTERSECT Operator = 1
	UNION     Operator = 2
	EXCEPT    Operator = 3
)

type Quantifier int32

const (
	ALL      Quantifier = 1
	DISTINCT Quantifier = 2
)

type Order int32

const (
	ASC   Order = 1
	DESC  Order = 2
	FIRST Order = 3
	LAST  Order = 4
)

type JoinType int32

const (
	INNER JoinType = 1
	LEFT  JoinType = 2
	RIGHT JoinType = 3
	FULL  JoinType = 4
)
