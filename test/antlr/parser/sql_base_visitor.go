// Code generated from Sql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSqlVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSingleExpression(ctx *SingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitStatementDefault(ctx *StatementDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitWith(ctx *WithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTableElement(ctx *TableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLikeClause(ctx *LikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQueryTermDefault(ctx *QueryTermDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitGroupBy(ctx *GroupByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRollup(ctx *RollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCube(ctx *CubeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitGroupingExpressions(ctx *GroupingExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNamedQuery(ctx *NamedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRelationDefault(ctx *RelationDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoinType(ctx *JoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSampledRelation(ctx *SampledRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSampleType(ctx *SampleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSubqueryRelation(ctx *SubqueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUnnest(ctx *UnnestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLateral(ctx *LateralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBooleanDefault(ctx *BooleanDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPredicated(ctx *PredicatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQuantifiedComparison(ctx *QuantifiedComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLike(ctx *LikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNullPredicate(ctx *NullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDistinctFrom(ctx *DistinctFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTypeConstructor(ctx *TypeConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSubstring(ctx *SubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPosition(ctx *PositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBasicStringLiteral(ctx *BasicStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitComparisonQuantifier(ctx *ComparisonQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTypeSql(ctx *TypeSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDoubleLiteral(ctx *DoubleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
