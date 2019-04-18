// Code generated from JsonQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JsonQueryParser.
type JsonQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonQueryParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#parenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#presentExp.
	VisitPresentExp(ctx *PresentExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#logicalExp.
	VisitLogicalExp(ctx *LogicalExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#long.
	VisitLong(ctx *LongContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfInts.
	VisitListOfInts(ctx *ListOfIntsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listInts.
	VisitListInts(ctx *ListIntsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfInts.
	VisitSubListOfInts(ctx *SubListOfIntsContext) interface{}
}