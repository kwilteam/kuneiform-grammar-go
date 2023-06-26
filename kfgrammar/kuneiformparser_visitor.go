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

	// Visit a parse tree produced by KuneiformParser#extension_directive.
	VisitExtension_directive(ctx *Extension_directiveContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config_list.
	VisitExt_config_list(ctx *Ext_config_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config.
	VisitExt_config(ctx *Ext_configContext) interface{}

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

	// Visit a parse tree produced by KuneiformParser#foreign_key_action.
	VisitForeign_key_action(ctx *Foreign_key_actionContext) interface{}

	// Visit a parse tree produced by KuneiformParser#foreign_key_def.
	VisitForeign_key_def(ctx *Foreign_key_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_decl.
	VisitAction_decl(ctx *Action_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#param_list.
	VisitParam_list(ctx *Param_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#extension_name.
	VisitExtension_name(ctx *Extension_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config_name.
	VisitExt_config_name(ctx *Ext_config_nameContext) interface{}

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

	// Visit a parse tree produced by KuneiformParser#ext_config_value.
	VisitExt_config_value(ctx *Ext_config_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt_list.
	VisitAction_stmt_list(ctx *Action_stmt_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt.
	VisitAction_stmt(ctx *Action_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_sql_stmt.
	VisitA_sql_stmt(ctx *A_sql_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_variable_name.
	VisitA_variable_name(ctx *A_variable_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_block_variable_name.
	VisitA_block_variable_name(ctx *A_block_variable_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_literal_value.
	VisitA_literal_value(ctx *A_literal_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_fn_name.
	VisitA_fn_name(ctx *A_fn_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_call_receivers.
	VisitA_call_receivers(ctx *A_call_receiversContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_call_stmt.
	VisitA_call_stmt(ctx *A_call_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_call_body.
	VisitA_call_body(ctx *A_call_bodyContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_fn_arg_list.
	VisitA_fn_arg_list(ctx *A_fn_arg_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#a_fn_arg.
	VisitA_fn_arg(ctx *A_fn_argContext) interface{}
}
