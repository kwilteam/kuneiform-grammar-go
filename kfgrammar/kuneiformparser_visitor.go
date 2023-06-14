// Code generated from KuneiformParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar // KuneiformParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by KuneiformParser.
type KuneiformParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KuneiformParser#source_unit.
	VisitSource_unit(ctx *Source_unitContext) interface{}

	// Visit a parse tree produced by KuneiformParser#database_directive.
	VisitDatabase_directive(ctx *Database_directiveContext) interface{}

	// Visit a parse tree produced by KuneiformParser#table_decl.
	VisitTable_decl(ctx *Table_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_def.
	VisitColumn_def(ctx *Column_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_def_list.
	VisitColumn_def_list(ctx *Column_def_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_type.
	VisitColumn_type(ctx *Column_typeContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_constraint.
	VisitColumn_constraint(ctx *Column_constraintContext) interface{}

	// Visit a parse tree produced by KuneiformParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#number_value.
	VisitNumber_value(ctx *Number_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#index_def.
	VisitIndex_def(ctx *Index_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#index_def_list.
	VisitIndex_def_list(ctx *Index_def_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_decl.
	VisitAction_decl(ctx *Action_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_param_list.
	VisitAction_param_list(ctx *Action_param_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_name.
	VisitAction_name(ctx *Action_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_name_list.
	VisitColumn_name_list(ctx *Column_name_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#sql_keywords.
	VisitSql_keywords(ctx *Sql_keywordsContext) interface{}

	// Visit a parse tree produced by KuneiformParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt.
	VisitAction_stmt(ctx *Action_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt_list.
	VisitAction_stmt_list(ctx *Action_stmt_listContext) interface{}
}
