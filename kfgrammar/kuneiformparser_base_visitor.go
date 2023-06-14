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

func (v *BaseKuneiformParserVisitor) VisitIndex_def_list(ctx *Index_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_decl(ctx *Action_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_param_list(ctx *Action_param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
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

func (v *BaseKuneiformParserVisitor) VisitSql_keywords(ctx *Sql_keywordsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_stmt(ctx *Action_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKuneiformParserVisitor) VisitAction_stmt_list(ctx *Action_stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}
