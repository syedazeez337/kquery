// Code generated from grammar/KQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KQuery
import "github.com/antlr4-go/antlr/v4"

// KQueryListener is a complete listener for a parse tree produced by KQueryParser.
type KQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
