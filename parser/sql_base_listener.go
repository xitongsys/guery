// Code generated from Sql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSqlListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlListener struct{}

var _ SqlListener = &BaseSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseSqlListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseSqlListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseSqlListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseSqlListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterStatementDefault is called when production statementDefault is entered.
func (s *BaseSqlListener) EnterStatementDefault(ctx *StatementDefaultContext) {}

// ExitStatementDefault is called when production statementDefault is exited.
func (s *BaseSqlListener) ExitStatementDefault(ctx *StatementDefaultContext) {}

// EnterWith is called when production with is entered.
func (s *BaseSqlListener) EnterWith(ctx *WithContext) {}

// ExitWith is called when production with is exited.
func (s *BaseSqlListener) ExitWith(ctx *WithContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseSqlListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseSqlListener) ExitTableElement(ctx *TableElementContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseSqlListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseSqlListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterLikeClause is called when production likeClause is entered.
func (s *BaseSqlListener) EnterLikeClause(ctx *LikeClauseContext) {}

// ExitLikeClause is called when production likeClause is exited.
func (s *BaseSqlListener) ExitLikeClause(ctx *LikeClauseContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseSqlListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseSqlListener) ExitProperties(ctx *PropertiesContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseSqlListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseSqlListener) ExitProperty(ctx *PropertyContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSqlListener) ExitQuery(ctx *QueryContext) {}

// EnterQueryTermDefault is called when production queryTermDefault is entered.
func (s *BaseSqlListener) EnterQueryTermDefault(ctx *QueryTermDefaultContext) {}

// ExitQueryTermDefault is called when production queryTermDefault is exited.
func (s *BaseSqlListener) ExitQueryTermDefault(ctx *QueryTermDefaultContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseSqlListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseSqlListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseSqlListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseSqlListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterTable is called when production table is entered.
func (s *BaseSqlListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSqlListener) ExitTable(ctx *TableContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseSqlListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseSqlListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseSqlListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseSqlListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseSqlListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseSqlListener) ExitSortItem(ctx *SortItemContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseSqlListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseSqlListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterGroupBy is called when production groupBy is entered.
func (s *BaseSqlListener) EnterGroupBy(ctx *GroupByContext) {}

// ExitGroupBy is called when production groupBy is exited.
func (s *BaseSqlListener) ExitGroupBy(ctx *GroupByContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseSqlListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseSqlListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseSqlListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseSqlListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseSqlListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseSqlListener) ExitCube(ctx *CubeContext) {}

// EnterGroupingExpressions is called when production groupingExpressions is entered.
func (s *BaseSqlListener) EnterGroupingExpressions(ctx *GroupingExpressionsContext) {}

// ExitGroupingExpressions is called when production groupingExpressions is exited.
func (s *BaseSqlListener) ExitGroupingExpressions(ctx *GroupingExpressionsContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseSqlListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseSqlListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseSqlListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseSqlListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseSqlListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseSqlListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseSqlListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseSqlListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterRelationDefault is called when production relationDefault is entered.
func (s *BaseSqlListener) EnterRelationDefault(ctx *RelationDefaultContext) {}

// ExitRelationDefault is called when production relationDefault is exited.
func (s *BaseSqlListener) ExitRelationDefault(ctx *RelationDefaultContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseSqlListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseSqlListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseSqlListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseSqlListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseSqlListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseSqlListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterSampledRelation is called when production sampledRelation is entered.
func (s *BaseSqlListener) EnterSampledRelation(ctx *SampledRelationContext) {}

// ExitSampledRelation is called when production sampledRelation is exited.
func (s *BaseSqlListener) ExitSampledRelation(ctx *SampledRelationContext) {}

// EnterSampleType is called when production sampleType is entered.
func (s *BaseSqlListener) EnterSampleType(ctx *SampleTypeContext) {}

// ExitSampleType is called when production sampleType is exited.
func (s *BaseSqlListener) ExitSampleType(ctx *SampleTypeContext) {}

// EnterAliasedRelation is called when production aliasedRelation is entered.
func (s *BaseSqlListener) EnterAliasedRelation(ctx *AliasedRelationContext) {}

// ExitAliasedRelation is called when production aliasedRelation is exited.
func (s *BaseSqlListener) ExitAliasedRelation(ctx *AliasedRelationContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseSqlListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseSqlListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlListener) ExitTableName(ctx *TableNameContext) {}

// EnterSubqueryRelation is called when production subqueryRelation is entered.
func (s *BaseSqlListener) EnterSubqueryRelation(ctx *SubqueryRelationContext) {}

// ExitSubqueryRelation is called when production subqueryRelation is exited.
func (s *BaseSqlListener) ExitSubqueryRelation(ctx *SubqueryRelationContext) {}

// EnterUnnest is called when production unnest is entered.
func (s *BaseSqlListener) EnterUnnest(ctx *UnnestContext) {}

// ExitUnnest is called when production unnest is exited.
func (s *BaseSqlListener) ExitUnnest(ctx *UnnestContext) {}

// EnterLateral is called when production lateral is entered.
func (s *BaseSqlListener) EnterLateral(ctx *LateralContext) {}

// ExitLateral is called when production lateral is exited.
func (s *BaseSqlListener) ExitLateral(ctx *LateralContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseSqlListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseSqlListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSqlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSqlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseSqlListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseSqlListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterBooleanDefault is called when production booleanDefault is entered.
func (s *BaseSqlListener) EnterBooleanDefault(ctx *BooleanDefaultContext) {}

// ExitBooleanDefault is called when production booleanDefault is exited.
func (s *BaseSqlListener) ExitBooleanDefault(ctx *BooleanDefaultContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseSqlListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseSqlListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseSqlListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseSqlListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSqlListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSqlListener) ExitComparison(ctx *ComparisonContext) {}

// EnterQuantifiedComparison is called when production quantifiedComparison is entered.
func (s *BaseSqlListener) EnterQuantifiedComparison(ctx *QuantifiedComparisonContext) {}

// ExitQuantifiedComparison is called when production quantifiedComparison is exited.
func (s *BaseSqlListener) ExitQuantifiedComparison(ctx *QuantifiedComparisonContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseSqlListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseSqlListener) ExitBetween(ctx *BetweenContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseSqlListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseSqlListener) ExitInList(ctx *InListContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseSqlListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseSqlListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterLike is called when production like is entered.
func (s *BaseSqlListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseSqlListener) ExitLike(ctx *LikeContext) {}

// EnterNullPredicate is called when production nullPredicate is entered.
func (s *BaseSqlListener) EnterNullPredicate(ctx *NullPredicateContext) {}

// ExitNullPredicate is called when production nullPredicate is exited.
func (s *BaseSqlListener) ExitNullPredicate(ctx *NullPredicateContext) {}

// EnterDistinctFrom is called when production distinctFrom is entered.
func (s *BaseSqlListener) EnterDistinctFrom(ctx *DistinctFromContext) {}

// ExitDistinctFrom is called when production distinctFrom is exited.
func (s *BaseSqlListener) ExitDistinctFrom(ctx *DistinctFromContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseSqlListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseSqlListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseSqlListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseSqlListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseSqlListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseSqlListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseSqlListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseSqlListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseSqlListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseSqlListener) ExitDereference(ctx *DereferenceContext) {}

// EnterTypeConstructor is called when production typeConstructor is entered.
func (s *BaseSqlListener) EnterTypeConstructor(ctx *TypeConstructorContext) {}

// ExitTypeConstructor is called when production typeConstructor is exited.
func (s *BaseSqlListener) ExitTypeConstructor(ctx *TypeConstructorContext) {}

// EnterSpecialDateTimeFunction is called when production specialDateTimeFunction is entered.
func (s *BaseSqlListener) EnterSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) {}

// ExitSpecialDateTimeFunction is called when production specialDateTimeFunction is exited.
func (s *BaseSqlListener) ExitSpecialDateTimeFunction(ctx *SpecialDateTimeFunctionContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BaseSqlListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BaseSqlListener) ExitSubstring(ctx *SubstringContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseSqlListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseSqlListener) ExitCast(ctx *CastContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseSqlListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseSqlListener) ExitLambda(ctx *LambdaContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseSqlListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseSqlListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseSqlListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseSqlListener) ExitParameter(ctx *ParameterContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSqlListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSqlListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSqlListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSqlListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseSqlListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseSqlListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseSqlListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseSqlListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseSqlListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseSqlListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseSqlListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseSqlListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseSqlListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseSqlListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseSqlListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseSqlListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseSqlListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseSqlListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseSqlListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseSqlListener) ExitExtract(ctx *ExtractContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSqlListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSqlListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseSqlListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseSqlListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSqlListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSqlListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseSqlListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseSqlListener) ExitExists(ctx *ExistsContext) {}

// EnterPosition is called when production position is entered.
func (s *BaseSqlListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BaseSqlListener) ExitPosition(ctx *PositionContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseSqlListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseSqlListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseSqlListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseSqlListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterBasicStringLiteral is called when production basicStringLiteral is entered.
func (s *BaseSqlListener) EnterBasicStringLiteral(ctx *BasicStringLiteralContext) {}

// ExitBasicStringLiteral is called when production basicStringLiteral is exited.
func (s *BaseSqlListener) ExitBasicStringLiteral(ctx *BasicStringLiteralContext) {}

// EnterUnicodeStringLiteral is called when production unicodeStringLiteral is entered.
func (s *BaseSqlListener) EnterUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) {}

// ExitUnicodeStringLiteral is called when production unicodeStringLiteral is exited.
func (s *BaseSqlListener) ExitUnicodeStringLiteral(ctx *UnicodeStringLiteralContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterComparisonQuantifier is called when production comparisonQuantifier is entered.
func (s *BaseSqlListener) EnterComparisonQuantifier(ctx *ComparisonQuantifierContext) {}

// ExitComparisonQuantifier is called when production comparisonQuantifier is exited.
func (s *BaseSqlListener) ExitComparisonQuantifier(ctx *ComparisonQuantifierContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseSqlListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseSqlListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterTypeSql is called when production typeSql is entered.
func (s *BaseSqlListener) EnterTypeSql(ctx *TypeSqlContext) {}

// ExitTypeSql is called when production typeSql is exited.
func (s *BaseSqlListener) ExitTypeSql(ctx *TypeSqlContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseSqlListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseSqlListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseSqlListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseSqlListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseSqlListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseSqlListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseSqlListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseSqlListener) ExitFilter(ctx *FilterContext) {}

// EnterOver is called when production over is entered.
func (s *BaseSqlListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseSqlListener) ExitOver(ctx *OverContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseSqlListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseSqlListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseSqlListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseSqlListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseSqlListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseSqlListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseSqlListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseSqlListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseSqlListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseSqlListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseSqlListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseSqlListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseSqlListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseSqlListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSqlListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSqlListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseSqlListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseSqlListener) ExitNonReserved(ctx *NonReservedContext) {}
