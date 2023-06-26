// Code generated from KuneiformParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar // KuneiformParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseKuneiformParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKuneiformParserVisitor) VisitSource_unit(ctx *Source_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitDatabase_directive(ctx *Database_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExtension_directive(ctx *Extension_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExt_config_list(ctx *Ext_config_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExt_config(ctx *Ext_configContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitTable_decl(ctx *Table_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_def(ctx *Column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_def_list(ctx *Column_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_type(ctx *Column_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_constraint(ctx *Column_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitNumber_value(ctx *Number_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitIndex_def(ctx *Index_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitForeign_key_action(ctx *Foreign_key_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitForeign_key_def(ctx *Foreign_key_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_decl(ctx *Action_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitParam_list(ctx *Param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExtension_name(ctx *Extension_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExt_config_name(ctx *Ext_config_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_name(ctx *Action_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitColumn_name_list(ctx *Column_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitExt_config_value(ctx *Ext_config_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_stmt_list(ctx *Action_stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_stmt(ctx *Action_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_sql_stmt(ctx *A_sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_variable_name(ctx *A_variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_block_variable_name(ctx *A_block_variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_literal_value(ctx *A_literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_fn_name(ctx *A_fn_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_call_receivers(ctx *A_call_receiversContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_call_stmt(ctx *A_call_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_call_body(ctx *A_call_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_fn_arg_list(ctx *A_fn_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitA_fn_arg(ctx *A_fn_argContext) interface{} {
	return v.VisitChildren(ctx)
}
