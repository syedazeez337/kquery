// Code generated from grammar/KQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KQuery
import "github.com/antlr4-go/antlr/v4"

// BaseKQueryListener is a complete listener for a parse tree produced by KQueryParser.
type BaseKQueryListener struct{}

var _ KQueryListener = &BaseKQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseKQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseKQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseKQueryListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseKQueryListener) ExitResource(ctx *ResourceContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseKQueryListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseKQueryListener) ExitFieldList(ctx *FieldListContext) {}

// EnterField is called when production field is entered.
func (s *BaseKQueryListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseKQueryListener) ExitField(ctx *FieldContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseKQueryListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseKQueryListener) ExitExpr(ctx *ExprContext) {}
