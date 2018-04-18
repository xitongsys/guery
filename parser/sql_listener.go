// Code generated from ./Sql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SqlListener is a complete listener for a parse tree produced by SqlParser.
type SqlListener interface {
	antlr.ParseTreeListener

	// EnterSingleStatement is called when entering the singleStatement production.
	EnterSingleStatement(c *SingleStatementContext)

	// EnterSingleExpression is called when entering the singleExpression production.
	EnterSingleExpression(c *SingleExpressionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterTableElement is called when entering the tableElement production.
	EnterTableElement(c *TableElementContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterLikeClause is called when entering the likeClause production.
	EnterLikeClause(c *LikeClauseContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterQueryTerm is called when entering the queryTerm production.
	EnterQueryTerm(c *QueryTermContext)

	// EnterQueryPrimary is called when entering the queryPrimary production.
	EnterQueryPrimary(c *QueryPrimaryContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterGroupBy is called when entering the groupBy production.
	EnterGroupBy(c *GroupByContext)

	// EnterGroupingElement is called when entering the groupingElement production.
	EnterGroupingElement(c *GroupingElementContext)

	// EnterSetQuantifier is called when entering the setQuantifier production.
	EnterSetQuantifier(c *SetQuantifierContext)

	// EnterSelectItem is called when entering the selectItem production.
	EnterSelectItem(c *SelectItemContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterJoinType is called when entering the joinType production.
	EnterJoinType(c *JoinTypeContext)

	// EnterJoinCriteria is called when entering the joinCriteria production.
	EnterJoinCriteria(c *JoinCriteriaContext)

	// EnterSampleType is called when entering the sampleType production.
	EnterSampleType(c *SampleTypeContext)

	// EnterSampledRelation is called when entering the sampledRelation production.
	EnterSampledRelation(c *SampledRelationContext)

	// EnterRelationPrimary is called when entering the relationPrimary production.
	EnterRelationPrimary(c *RelationPrimaryContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBooleanExpression is called when entering the booleanExpression production.
	EnterBooleanExpression(c *BooleanExpressionContext)

	// EnterPredicated is called when entering the predicated production.
	EnterPredicated(c *PredicatedContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterComparisonQuantifier is called when entering the comparisonQuantifier production.
	EnterComparisonQuantifier(c *ComparisonQuantifierContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterTypeSql is called when entering the typeSql production.
	EnterTypeSql(c *TypeSqlContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterWhenClause is called when entering the whenClause production.
	EnterWhenClause(c *WhenClauseContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterOver is called when entering the over production.
	EnterOver(c *OverContext)

	// EnterPrivilege is called when entering the privilege production.
	EnterPrivilege(c *PrivilegeContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterNonReserved is called when entering the nonReserved production.
	EnterNonReserved(c *NonReservedContext)

	// ExitSingleStatement is called when exiting the singleStatement production.
	ExitSingleStatement(c *SingleStatementContext)

	// ExitSingleExpression is called when exiting the singleExpression production.
	ExitSingleExpression(c *SingleExpressionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitTableElement is called when exiting the tableElement production.
	ExitTableElement(c *TableElementContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitLikeClause is called when exiting the likeClause production.
	ExitLikeClause(c *LikeClauseContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitQueryTerm is called when exiting the queryTerm production.
	ExitQueryTerm(c *QueryTermContext)

	// ExitQueryPrimary is called when exiting the queryPrimary production.
	ExitQueryPrimary(c *QueryPrimaryContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitGroupBy is called when exiting the groupBy production.
	ExitGroupBy(c *GroupByContext)

	// ExitGroupingElement is called when exiting the groupingElement production.
	ExitGroupingElement(c *GroupingElementContext)

	// ExitSetQuantifier is called when exiting the setQuantifier production.
	ExitSetQuantifier(c *SetQuantifierContext)

	// ExitSelectItem is called when exiting the selectItem production.
	ExitSelectItem(c *SelectItemContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitJoinType is called when exiting the joinType production.
	ExitJoinType(c *JoinTypeContext)

	// ExitJoinCriteria is called when exiting the joinCriteria production.
	ExitJoinCriteria(c *JoinCriteriaContext)

	// ExitSampleType is called when exiting the sampleType production.
	ExitSampleType(c *SampleTypeContext)

	// ExitSampledRelation is called when exiting the sampledRelation production.
	ExitSampledRelation(c *SampledRelationContext)

	// ExitRelationPrimary is called when exiting the relationPrimary production.
	ExitRelationPrimary(c *RelationPrimaryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBooleanExpression is called when exiting the booleanExpression production.
	ExitBooleanExpression(c *BooleanExpressionContext)

	// ExitPredicated is called when exiting the predicated production.
	ExitPredicated(c *PredicatedContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitComparisonQuantifier is called when exiting the comparisonQuantifier production.
	ExitComparisonQuantifier(c *ComparisonQuantifierContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitTypeSql is called when exiting the typeSql production.
	ExitTypeSql(c *TypeSqlContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitWhenClause is called when exiting the whenClause production.
	ExitWhenClause(c *WhenClauseContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitOver is called when exiting the over production.
	ExitOver(c *OverContext)

	// ExitPrivilege is called when exiting the privilege production.
	ExitPrivilege(c *PrivilegeContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitNonReserved is called when exiting the nonReserved production.
	ExitNonReserved(c *NonReservedContext)
}
