// Code generated from ./Sql.g4 by ANTLR 4.7.1. DO NOT EDIT.

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

// EnterStatement is called when production statement is entered.
func (s *BaseSqlListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSqlListener) ExitStatement(ctx *StatementContext) {}

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

// EnterQueryTerm is called when production queryTerm is entered.
func (s *BaseSqlListener) EnterQueryTerm(ctx *QueryTermContext) {}

// ExitQueryTerm is called when production queryTerm is exited.
func (s *BaseSqlListener) ExitQueryTerm(ctx *QueryTermContext) {}

// EnterQueryPrimary is called when production queryPrimary is entered.
func (s *BaseSqlListener) EnterQueryPrimary(ctx *QueryPrimaryContext) {}

// ExitQueryPrimary is called when production queryPrimary is exited.
func (s *BaseSqlListener) ExitQueryPrimary(ctx *QueryPrimaryContext) {}

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

// EnterGroupingElement is called when production groupingElement is entered.
func (s *BaseSqlListener) EnterGroupingElement(ctx *GroupingElementContext) {}

// ExitGroupingElement is called when production groupingElement is exited.
func (s *BaseSqlListener) ExitGroupingElement(ctx *GroupingElementContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseSqlListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseSqlListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *BaseSqlListener) EnterSelectItem(ctx *SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *BaseSqlListener) ExitSelectItem(ctx *SelectItemContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseSqlListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseSqlListener) ExitRelation(ctx *RelationContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseSqlListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseSqlListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseSqlListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseSqlListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterSampleType is called when production sampleType is entered.
func (s *BaseSqlListener) EnterSampleType(ctx *SampleTypeContext) {}

// ExitSampleType is called when production sampleType is exited.
func (s *BaseSqlListener) ExitSampleType(ctx *SampleTypeContext) {}

// EnterSampledRelation is called when production sampledRelation is entered.
func (s *BaseSqlListener) EnterSampledRelation(ctx *SampledRelationContext) {}

// ExitSampledRelation is called when production sampledRelation is exited.
func (s *BaseSqlListener) ExitSampledRelation(ctx *SampledRelationContext) {}

// EnterRelationPrimary is called when production relationPrimary is entered.
func (s *BaseSqlListener) EnterRelationPrimary(ctx *RelationPrimaryContext) {}

// ExitRelationPrimary is called when production relationPrimary is exited.
func (s *BaseSqlListener) ExitRelationPrimary(ctx *RelationPrimaryContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSqlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSqlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBooleanExpression is called when production booleanExpression is entered.
func (s *BaseSqlListener) EnterBooleanExpression(ctx *BooleanExpressionContext) {}

// ExitBooleanExpression is called when production booleanExpression is exited.
func (s *BaseSqlListener) ExitBooleanExpression(ctx *BooleanExpressionContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseSqlListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseSqlListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSqlListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSqlListener) ExitPredicate(ctx *PredicateContext) {}

// EnterValueExpression is called when production valueExpression is entered.
func (s *BaseSqlListener) EnterValueExpression(ctx *ValueExpressionContext) {}

// ExitValueExpression is called when production valueExpression is exited.
func (s *BaseSqlListener) ExitValueExpression(ctx *ValueExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSqlListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSqlListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseSqlListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseSqlListener) ExitStringValue(ctx *StringValueContext) {}

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

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSqlListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSqlListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSqlListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSqlListener) ExitNumber(ctx *NumberContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseSqlListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseSqlListener) ExitNonReserved(ctx *NonReservedContext) {}
