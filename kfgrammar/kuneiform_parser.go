// Code generated from KuneiformParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar // KuneiformParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type KuneiformParser struct {
	*antlr.BaseParser
}

var kuneiformparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kuneiformparserParserInit() {
	staticData := &kuneiformparserParserStaticData
	staticData.literalNames = []string{
		"", "':'", "';'", "'('", "'{'", "')'", "'}'", "','", "'$'", "'#'", "'database'",
		"'use'", "'as'", "'table'", "'action'", "'public'", "'private'", "'view'",
		"'init'", "'int'", "'text'", "'min'", "'max'", "'minlen'", "'maxlen'",
		"'notnull'", "'primary'", "'default'", "'unique'", "'index'", "'foreign_key'",
		"'fk'", "'references'", "'ref'", "'on_update'", "'on_delete'", "'do'",
		"'no_action'", "'cascade'", "'set_null'", "'set_default'", "'restrict'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'='", "'+'", "'.'", "", "", "'@'",
	}
	staticData.symbolicNames = []string{
		"", "COL", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA",
		"DOLLAR", "HASH", "DATABASE_", "USE_", "AS_", "TABLE_", "ACTION_", "PUBLIC_",
		"PRIVATE_", "VIEW_", "INIT_", "INT_", "TEXT_", "MIN_", "MAX_", "MIN_LEN_",
		"MAX_LEN_", "NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_",
		"FOREIGN_KEY_", "FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_",
		"ACTION_ON_UPDATE_", "ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_",
		"ACTION_DO_CASCADE_", "ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_",
		"ACTION_DO_RESTRICT_", "SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_",
		"ACTION_OPEN", "INIT_OPEN", "IDENTIFIER", "INDEX_NAME", "PARAMETER",
		"UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL", "STRING_LITERAL",
		"WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT", "ACTION_CLOSE",
		"EQ", "PLUS", "PERIOD", "A_COMMA", "A_DOLLAR", "A_AT", "A_L_PAREN",
		"A_R_PAREN", "A_STMT_END", "SQL_KEYWORDS", "A_IDENTIFIER", "A_VARIABLE",
		"A_REF", "A_UNSIGNED_NUMBER_LITERAL", "A_SIGNED_NUMBER_LITERAL", "A_STRING_LITERAL",
		"A_SQL_STMT", "A_WS", "A_TERMINATOR", "A_BLOCK_COMMENT", "A_LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"source_unit", "database_directive", "extension_directive", "ext_config_list",
		"ext_config", "table_decl", "column_def", "column_def_list", "column_type",
		"column_constraint", "literal_value", "number_value", "index_def", "foreign_key_action",
		"foreign_key_def", "action_decl", "param_list", "database_name", "extension_name",
		"ext_config_name", "table_name", "action_name", "column_name", "column_name_list",
		"index_name", "ext_config_value", "init_decl", "action_stmt_list", "action_stmt",
		"a_sql_stmt", "a_variable_name", "a_block_variable_name", "a_literal_value",
		"a_fn_name", "a_call_receivers", "a_call_stmt", "a_call_body", "a_fn_arg_list",
		"a_fn_arg",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 330, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 5, 0, 82, 8, 0, 10, 0, 12,
		0, 85, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		106, 8, 2, 1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5,
		3, 117, 8, 3, 10, 3, 12, 3, 120, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 133, 8, 5, 5, 5, 135, 8, 5, 10,
		5, 12, 5, 138, 9, 5, 1, 5, 3, 5, 141, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 5, 6, 148, 8, 6, 10, 6, 12, 6, 151, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 156,
		8, 7, 10, 7, 12, 7, 159, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		191, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 3, 13, 205, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 219, 8, 14,
		10, 14, 12, 14, 222, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 3, 16, 234, 8, 16, 1, 16, 1, 16, 5, 16, 238, 8,
		16, 10, 16, 12, 16, 241, 9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 258,
		8, 23, 10, 23, 12, 23, 261, 9, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 4, 27, 272, 8, 27, 11, 27, 12, 27, 273, 1, 28,
		1, 28, 3, 28, 278, 8, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 292, 8, 33, 1, 34, 1, 34,
		1, 34, 5, 34, 297, 8, 34, 10, 34, 12, 34, 300, 9, 34, 1, 35, 1, 35, 1,
		35, 3, 35, 305, 8, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 3, 37, 316, 8, 37, 1, 37, 1, 37, 5, 37, 320, 8, 37, 10, 37,
		12, 37, 323, 9, 37, 1, 38, 1, 38, 1, 38, 3, 38, 328, 8, 38, 1, 38, 0, 0,
		39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 0, 8, 1, 0, 19, 20, 2, 0, 52, 52, 54, 54, 2, 0, 26, 26, 28,
		29, 1, 0, 34, 35, 1, 0, 37, 41, 1, 0, 30, 31, 1, 0, 32, 33, 2, 0, 73, 73,
		75, 75, 323, 0, 78, 1, 0, 0, 0, 2, 96, 1, 0, 0, 0, 4, 99, 1, 0, 0, 0, 6,
		113, 1, 0, 0, 0, 8, 121, 1, 0, 0, 0, 10, 125, 1, 0, 0, 0, 12, 144, 1, 0,
		0, 0, 14, 152, 1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18, 190, 1, 0, 0, 0, 20,
		192, 1, 0, 0, 0, 22, 194, 1, 0, 0, 0, 24, 196, 1, 0, 0, 0, 26, 202, 1,
		0, 0, 0, 28, 208, 1, 0, 0, 0, 30, 223, 1, 0, 0, 0, 32, 233, 1, 0, 0, 0,
		34, 242, 1, 0, 0, 0, 36, 244, 1, 0, 0, 0, 38, 246, 1, 0, 0, 0, 40, 248,
		1, 0, 0, 0, 42, 250, 1, 0, 0, 0, 44, 252, 1, 0, 0, 0, 46, 254, 1, 0, 0,
		0, 48, 262, 1, 0, 0, 0, 50, 264, 1, 0, 0, 0, 52, 266, 1, 0, 0, 0, 54, 271,
		1, 0, 0, 0, 56, 277, 1, 0, 0, 0, 58, 279, 1, 0, 0, 0, 60, 282, 1, 0, 0,
		0, 62, 284, 1, 0, 0, 0, 64, 286, 1, 0, 0, 0, 66, 288, 1, 0, 0, 0, 68, 293,
		1, 0, 0, 0, 70, 304, 1, 0, 0, 0, 72, 309, 1, 0, 0, 0, 74, 315, 1, 0, 0,
		0, 76, 327, 1, 0, 0, 0, 78, 79, 3, 2, 1, 0, 79, 83, 5, 2, 0, 0, 80, 82,
		3, 4, 2, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		83, 84, 1, 0, 0, 0, 84, 91, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 90, 3,
		10, 5, 0, 87, 90, 3, 30, 15, 0, 88, 90, 3, 52, 26, 0, 89, 86, 1, 0, 0,
		0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0,
		94, 95, 5, 0, 0, 1, 95, 1, 1, 0, 0, 0, 96, 97, 5, 10, 0, 0, 97, 98, 3,
		34, 17, 0, 98, 3, 1, 0, 0, 0, 99, 100, 5, 11, 0, 0, 100, 105, 3, 36, 18,
		0, 101, 102, 5, 4, 0, 0, 102, 103, 3, 6, 3, 0, 103, 104, 5, 6, 0, 0, 104,
		106, 1, 0, 0, 0, 105, 101, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 109,
		1, 0, 0, 0, 107, 108, 5, 12, 0, 0, 108, 110, 3, 36, 18, 0, 109, 107, 1,
		0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 2, 0,
		0, 112, 5, 1, 0, 0, 0, 113, 118, 3, 8, 4, 0, 114, 115, 5, 7, 0, 0, 115,
		117, 3, 8, 4, 0, 116, 114, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 7, 1, 0, 0, 0, 120, 118, 1, 0, 0,
		0, 121, 122, 3, 38, 19, 0, 122, 123, 5, 1, 0, 0, 123, 124, 3, 50, 25, 0,
		124, 9, 1, 0, 0, 0, 125, 126, 5, 13, 0, 0, 126, 127, 3, 40, 20, 0, 127,
		128, 5, 4, 0, 0, 128, 136, 3, 14, 7, 0, 129, 132, 5, 7, 0, 0, 130, 133,
		3, 24, 12, 0, 131, 133, 3, 28, 14, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1,
		0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 129, 1, 0, 0, 0, 135, 138, 1, 0, 0,
		0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138,
		136, 1, 0, 0, 0, 139, 141, 5, 7, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 5, 6, 0, 0, 143, 11, 1, 0,
		0, 0, 144, 145, 3, 44, 22, 0, 145, 149, 3, 16, 8, 0, 146, 148, 3, 18, 9,
		0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 13, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 157, 3,
		12, 6, 0, 153, 154, 5, 7, 0, 0, 154, 156, 3, 12, 6, 0, 155, 153, 1, 0,
		0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 15, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 7, 0, 0, 0, 161, 17,
		1, 0, 0, 0, 162, 191, 5, 26, 0, 0, 163, 191, 5, 25, 0, 0, 164, 191, 5,
		28, 0, 0, 165, 166, 5, 27, 0, 0, 166, 167, 5, 3, 0, 0, 167, 168, 3, 20,
		10, 0, 168, 169, 5, 5, 0, 0, 169, 191, 1, 0, 0, 0, 170, 171, 5, 21, 0,
		0, 171, 172, 5, 3, 0, 0, 172, 173, 3, 22, 11, 0, 173, 174, 5, 5, 0, 0,
		174, 191, 1, 0, 0, 0, 175, 176, 5, 22, 0, 0, 176, 177, 5, 3, 0, 0, 177,
		178, 3, 22, 11, 0, 178, 179, 5, 5, 0, 0, 179, 191, 1, 0, 0, 0, 180, 181,
		5, 23, 0, 0, 181, 182, 5, 3, 0, 0, 182, 183, 3, 22, 11, 0, 183, 184, 5,
		5, 0, 0, 184, 191, 1, 0, 0, 0, 185, 186, 5, 24, 0, 0, 186, 187, 5, 3, 0,
		0, 187, 188, 3, 22, 11, 0, 188, 189, 5, 5, 0, 0, 189, 191, 1, 0, 0, 0,
		190, 162, 1, 0, 0, 0, 190, 163, 1, 0, 0, 0, 190, 164, 1, 0, 0, 0, 190,
		165, 1, 0, 0, 0, 190, 170, 1, 0, 0, 0, 190, 175, 1, 0, 0, 0, 190, 180,
		1, 0, 0, 0, 190, 185, 1, 0, 0, 0, 191, 19, 1, 0, 0, 0, 192, 193, 7, 1,
		0, 0, 193, 21, 1, 0, 0, 0, 194, 195, 5, 52, 0, 0, 195, 23, 1, 0, 0, 0,
		196, 197, 3, 48, 24, 0, 197, 198, 7, 2, 0, 0, 198, 199, 5, 3, 0, 0, 199,
		200, 3, 46, 23, 0, 200, 201, 5, 5, 0, 0, 201, 25, 1, 0, 0, 0, 202, 204,
		7, 3, 0, 0, 203, 205, 5, 36, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0,
		0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 7, 4, 0, 0, 207, 27, 1, 0, 0, 0,
		208, 209, 7, 5, 0, 0, 209, 210, 5, 3, 0, 0, 210, 211, 3, 46, 23, 0, 211,
		212, 5, 5, 0, 0, 212, 213, 7, 6, 0, 0, 213, 214, 3, 40, 20, 0, 214, 215,
		5, 3, 0, 0, 215, 216, 3, 46, 23, 0, 216, 220, 5, 5, 0, 0, 217, 219, 3,
		26, 13, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 221, 1, 0, 0, 0, 221, 29, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		223, 224, 5, 14, 0, 0, 224, 225, 3, 42, 21, 0, 225, 226, 5, 3, 0, 0, 226,
		227, 3, 32, 16, 0, 227, 228, 5, 5, 0, 0, 228, 229, 5, 47, 0, 0, 229, 230,
		3, 54, 27, 0, 230, 231, 5, 59, 0, 0, 231, 31, 1, 0, 0, 0, 232, 234, 5,
		51, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 239, 1, 0, 0,
		0, 235, 236, 5, 7, 0, 0, 236, 238, 5, 51, 0, 0, 237, 235, 1, 0, 0, 0, 238,
		241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 33, 1,
		0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 243, 5, 49, 0, 0, 243, 35, 1, 0, 0,
		0, 244, 245, 5, 49, 0, 0, 245, 37, 1, 0, 0, 0, 246, 247, 5, 49, 0, 0, 247,
		39, 1, 0, 0, 0, 248, 249, 5, 49, 0, 0, 249, 41, 1, 0, 0, 0, 250, 251, 5,
		49, 0, 0, 251, 43, 1, 0, 0, 0, 252, 253, 5, 49, 0, 0, 253, 45, 1, 0, 0,
		0, 254, 259, 3, 44, 22, 0, 255, 256, 5, 7, 0, 0, 256, 258, 3, 44, 22, 0,
		257, 255, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 47, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 5,
		50, 0, 0, 263, 49, 1, 0, 0, 0, 264, 265, 3, 20, 10, 0, 265, 51, 1, 0, 0,
		0, 266, 267, 5, 48, 0, 0, 267, 268, 3, 54, 27, 0, 268, 269, 5, 59, 0, 0,
		269, 53, 1, 0, 0, 0, 270, 272, 3, 56, 28, 0, 271, 270, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 55, 1,
		0, 0, 0, 275, 278, 3, 58, 29, 0, 276, 278, 3, 70, 35, 0, 277, 275, 1, 0,
		0, 0, 277, 276, 1, 0, 0, 0, 278, 57, 1, 0, 0, 0, 279, 280, 5, 76, 0, 0,
		280, 281, 5, 68, 0, 0, 281, 59, 1, 0, 0, 0, 282, 283, 5, 71, 0, 0, 283,
		61, 1, 0, 0, 0, 284, 285, 5, 72, 0, 0, 285, 63, 1, 0, 0, 0, 286, 287, 7,
		7, 0, 0, 287, 65, 1, 0, 0, 0, 288, 291, 5, 70, 0, 0, 289, 290, 5, 62, 0,
		0, 290, 292, 5, 70, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292,
		67, 1, 0, 0, 0, 293, 298, 3, 60, 30, 0, 294, 295, 5, 63, 0, 0, 295, 297,
		3, 60, 30, 0, 296, 294, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1,
		0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 69, 1, 0, 0, 0, 300, 298, 1, 0, 0,
		0, 301, 302, 3, 68, 34, 0, 302, 303, 5, 60, 0, 0, 303, 305, 1, 0, 0, 0,
		304, 301, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306,
		307, 3, 72, 36, 0, 307, 308, 5, 68, 0, 0, 308, 71, 1, 0, 0, 0, 309, 310,
		3, 66, 33, 0, 310, 311, 5, 66, 0, 0, 311, 312, 3, 74, 37, 0, 312, 313,
		5, 67, 0, 0, 313, 73, 1, 0, 0, 0, 314, 316, 3, 76, 38, 0, 315, 314, 1,
		0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 321, 1, 0, 0, 0, 317, 318, 5, 63, 0,
		0, 318, 320, 3, 76, 38, 0, 319, 317, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 75, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 324, 328, 3, 64, 32, 0, 325, 328, 3, 60, 30, 0, 326, 328, 3,
		62, 31, 0, 327, 324, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 326, 1, 0,
		0, 0, 328, 77, 1, 0, 0, 0, 25, 83, 89, 91, 105, 109, 118, 132, 136, 140,
		149, 157, 190, 204, 220, 233, 239, 259, 273, 277, 291, 298, 304, 315, 321,
		327,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// KuneiformParserInit initializes any static state used to implement KuneiformParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKuneiformParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KuneiformParserInit() {
	staticData := &kuneiformparserParserStaticData
	staticData.once.Do(kuneiformparserParserInit)
}

// NewKuneiformParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKuneiformParser(input antlr.TokenStream) *KuneiformParser {
	KuneiformParserInit()
	this := new(KuneiformParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &kuneiformparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "KuneiformParser.g4"

	return this
}

// KuneiformParser tokens.
const (
	KuneiformParserEOF                       = antlr.TokenEOF
	KuneiformParserCOL                       = 1
	KuneiformParserSCOL                      = 2
	KuneiformParserL_PAREN                   = 3
	KuneiformParserL_BRACE                   = 4
	KuneiformParserR_PAREN                   = 5
	KuneiformParserR_BRACE                   = 6
	KuneiformParserCOMMA                     = 7
	KuneiformParserDOLLAR                    = 8
	KuneiformParserHASH                      = 9
	KuneiformParserDATABASE_                 = 10
	KuneiformParserUSE_                      = 11
	KuneiformParserAS_                       = 12
	KuneiformParserTABLE_                    = 13
	KuneiformParserACTION_                   = 14
	KuneiformParserPUBLIC_                   = 15
	KuneiformParserPRIVATE_                  = 16
	KuneiformParserVIEW_                     = 17
	KuneiformParserINIT_                     = 18
	KuneiformParserINT_                      = 19
	KuneiformParserTEXT_                     = 20
	KuneiformParserMIN_                      = 21
	KuneiformParserMAX_                      = 22
	KuneiformParserMIN_LEN_                  = 23
	KuneiformParserMAX_LEN_                  = 24
	KuneiformParserNOT_NULL_                 = 25
	KuneiformParserPRIMARY_                  = 26
	KuneiformParserDEFAULT_                  = 27
	KuneiformParserUNIQUE_                   = 28
	KuneiformParserINDEX_                    = 29
	KuneiformParserFOREIGN_KEY_              = 30
	KuneiformParserFOREIGN_KEY_ABBR_         = 31
	KuneiformParserREFERENCES_               = 32
	KuneiformParserREFERENCES_ABBR_          = 33
	KuneiformParserACTION_ON_UPDATE_         = 34
	KuneiformParserACTION_ON_DELETE_         = 35
	KuneiformParserACTION_DO_                = 36
	KuneiformParserACTION_DO_NO_ACTION_      = 37
	KuneiformParserACTION_DO_CASCADE_        = 38
	KuneiformParserACTION_DO_SET_NULL_       = 39
	KuneiformParserACTION_DO_SET_DEFAULT_    = 40
	KuneiformParserACTION_DO_RESTRICT_       = 41
	KuneiformParserSELECT_                   = 42
	KuneiformParserINSERT_                   = 43
	KuneiformParserUPDATE_                   = 44
	KuneiformParserDELETE_                   = 45
	KuneiformParserWITH_                     = 46
	KuneiformParserACTION_OPEN               = 47
	KuneiformParserINIT_OPEN                 = 48
	KuneiformParserIDENTIFIER                = 49
	KuneiformParserINDEX_NAME                = 50
	KuneiformParserPARAMETER                 = 51
	KuneiformParserUNSIGNED_NUMBER_LITERAL   = 52
	KuneiformParserSIGNED_NUMBER_LITERAL     = 53
	KuneiformParserSTRING_LITERAL            = 54
	KuneiformParserWS                        = 55
	KuneiformParserTERMINATOR                = 56
	KuneiformParserBLOCK_COMMENT             = 57
	KuneiformParserLINE_COMMENT              = 58
	KuneiformParserACTION_CLOSE              = 59
	KuneiformParserEQ                        = 60
	KuneiformParserPLUS                      = 61
	KuneiformParserPERIOD                    = 62
	KuneiformParserA_COMMA                   = 63
	KuneiformParserA_DOLLAR                  = 64
	KuneiformParserA_AT                      = 65
	KuneiformParserA_L_PAREN                 = 66
	KuneiformParserA_R_PAREN                 = 67
	KuneiformParserA_STMT_END                = 68
	KuneiformParserSQL_KEYWORDS              = 69
	KuneiformParserA_IDENTIFIER              = 70
	KuneiformParserA_VARIABLE                = 71
	KuneiformParserA_REF                     = 72
	KuneiformParserA_UNSIGNED_NUMBER_LITERAL = 73
	KuneiformParserA_SIGNED_NUMBER_LITERAL   = 74
	KuneiformParserA_STRING_LITERAL          = 75
	KuneiformParserA_SQL_STMT                = 76
	KuneiformParserA_WS                      = 77
	KuneiformParserA_TERMINATOR              = 78
	KuneiformParserA_BLOCK_COMMENT           = 79
	KuneiformParserA_LINE_COMMENT            = 80
)

// KuneiformParser rules.
const (
	KuneiformParserRULE_source_unit           = 0
	KuneiformParserRULE_database_directive    = 1
	KuneiformParserRULE_extension_directive   = 2
	KuneiformParserRULE_ext_config_list       = 3
	KuneiformParserRULE_ext_config            = 4
	KuneiformParserRULE_table_decl            = 5
	KuneiformParserRULE_column_def            = 6
	KuneiformParserRULE_column_def_list       = 7
	KuneiformParserRULE_column_type           = 8
	KuneiformParserRULE_column_constraint     = 9
	KuneiformParserRULE_literal_value         = 10
	KuneiformParserRULE_number_value          = 11
	KuneiformParserRULE_index_def             = 12
	KuneiformParserRULE_foreign_key_action    = 13
	KuneiformParserRULE_foreign_key_def       = 14
	KuneiformParserRULE_action_decl           = 15
	KuneiformParserRULE_param_list            = 16
	KuneiformParserRULE_database_name         = 17
	KuneiformParserRULE_extension_name        = 18
	KuneiformParserRULE_ext_config_name       = 19
	KuneiformParserRULE_table_name            = 20
	KuneiformParserRULE_action_name           = 21
	KuneiformParserRULE_column_name           = 22
	KuneiformParserRULE_column_name_list      = 23
	KuneiformParserRULE_index_name            = 24
	KuneiformParserRULE_ext_config_value      = 25
	KuneiformParserRULE_init_decl             = 26
	KuneiformParserRULE_action_stmt_list      = 27
	KuneiformParserRULE_action_stmt           = 28
	KuneiformParserRULE_a_sql_stmt            = 29
	KuneiformParserRULE_a_variable_name       = 30
	KuneiformParserRULE_a_block_variable_name = 31
	KuneiformParserRULE_a_literal_value       = 32
	KuneiformParserRULE_a_fn_name             = 33
	KuneiformParserRULE_a_call_receivers      = 34
	KuneiformParserRULE_a_call_stmt           = 35
	KuneiformParserRULE_a_call_body           = 36
	KuneiformParserRULE_a_fn_arg_list         = 37
	KuneiformParserRULE_a_fn_arg              = 38
)

// ISource_unitContext is an interface to support dynamic dispatch.
type ISource_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Database_directive() IDatabase_directiveContext
	SCOL() antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllExtension_directive() []IExtension_directiveContext
	Extension_directive(i int) IExtension_directiveContext
	AllTable_decl() []ITable_declContext
	Table_decl(i int) ITable_declContext
	AllAction_decl() []IAction_declContext
	Action_decl(i int) IAction_declContext
	AllInit_decl() []IInit_declContext
	Init_decl(i int) IInit_declContext

	// IsSource_unitContext differentiates from other interfaces.
	IsSource_unitContext()
}

type Source_unitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySource_unitContext() *Source_unitContext {
	var p = new(Source_unitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_source_unit
	return p
}

func (*Source_unitContext) IsSource_unitContext() {}

func NewSource_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Source_unitContext {
	var p = new(Source_unitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_source_unit

	return p
}

func (s *Source_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Source_unitContext) Database_directive() IDatabase_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabase_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabase_directiveContext)
}

func (s *Source_unitContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Source_unitContext) EOF() antlr.TerminalNode {
	return s.GetToken(KuneiformParserEOF, 0)
}

func (s *Source_unitContext) AllExtension_directive() []IExtension_directiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtension_directiveContext); ok {
			len++
		}
	}

	tst := make([]IExtension_directiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtension_directiveContext); ok {
			tst[i] = t.(IExtension_directiveContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Extension_directive(i int) IExtension_directiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_directiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_directiveContext)
}

func (s *Source_unitContext) AllTable_decl() []ITable_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_declContext); ok {
			len++
		}
	}

	tst := make([]ITable_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_declContext); ok {
			tst[i] = t.(ITable_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Table_decl(i int) ITable_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_declContext)
}

func (s *Source_unitContext) AllAction_decl() []IAction_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_declContext); ok {
			len++
		}
	}

	tst := make([]IAction_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_declContext); ok {
			tst[i] = t.(IAction_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Action_decl(i int) IAction_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_declContext)
}

func (s *Source_unitContext) AllInit_decl() []IInit_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInit_declContext); ok {
			len++
		}
	}

	tst := make([]IInit_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInit_declContext); ok {
			tst[i] = t.(IInit_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Init_decl(i int) IInit_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInit_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInit_declContext)
}

func (s *Source_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Source_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Source_unitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSource_unit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Source_unit() (localctx ISource_unitContext) {
	this := p
	_ = this

	localctx = NewSource_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KuneiformParserRULE_source_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Database_directive()
	}
	{
		p.SetState(79)
		p.Match(KuneiformParserSCOL)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserUSE_ {
		{
			p.SetState(80)
			p.Extension_directive()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281474976735232) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserTABLE_:
			{
				p.SetState(86)
				p.Table_decl()
			}

		case KuneiformParserACTION_:
			{
				p.SetState(87)
				p.Action_decl()
			}

		case KuneiformParserINIT_OPEN:
			{
				p.SetState(88)
				p.Init_decl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(KuneiformParserEOF)
	}

	return localctx
}

// IDatabase_directiveContext is an interface to support dynamic dispatch.
type IDatabase_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATABASE_() antlr.TerminalNode
	Database_name() IDatabase_nameContext

	// IsDatabase_directiveContext differentiates from other interfaces.
	IsDatabase_directiveContext()
}

type Database_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabase_directiveContext() *Database_directiveContext {
	var p = new(Database_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_database_directive
	return p
}

func (*Database_directiveContext) IsDatabase_directiveContext() {}

func NewDatabase_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Database_directiveContext {
	var p = new(Database_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_database_directive

	return p
}

func (s *Database_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Database_directiveContext) DATABASE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDATABASE_, 0)
}

func (s *Database_directiveContext) Database_name() IDatabase_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabase_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabase_nameContext)
}

func (s *Database_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Database_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Database_directiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitDatabase_directive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Database_directive() (localctx IDatabase_directiveContext) {
	this := p
	_ = this

	localctx = NewDatabase_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KuneiformParserRULE_database_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(KuneiformParserDATABASE_)
	}
	{
		p.SetState(97)
		p.Database_name()
	}

	return localctx
}

// IExtension_directiveContext is an interface to support dynamic dispatch.
type IExtension_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE_() antlr.TerminalNode
	AllExtension_name() []IExtension_nameContext
	Extension_name(i int) IExtension_nameContext
	SCOL() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	Ext_config_list() IExt_config_listContext
	R_BRACE() antlr.TerminalNode
	AS_() antlr.TerminalNode

	// IsExtension_directiveContext differentiates from other interfaces.
	IsExtension_directiveContext()
}

type Extension_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_directiveContext() *Extension_directiveContext {
	var p = new(Extension_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_extension_directive
	return p
}

func (*Extension_directiveContext) IsExtension_directiveContext() {}

func NewExtension_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_directiveContext {
	var p = new(Extension_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_extension_directive

	return p
}

func (s *Extension_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_directiveContext) USE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUSE_, 0)
}

func (s *Extension_directiveContext) AllExtension_name() []IExtension_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtension_nameContext); ok {
			len++
		}
	}

	tst := make([]IExtension_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtension_nameContext); ok {
			tst[i] = t.(IExtension_nameContext)
			i++
		}
	}

	return tst
}

func (s *Extension_directiveContext) Extension_name(i int) IExtension_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_nameContext)
}

func (s *Extension_directiveContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Extension_directiveContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Extension_directiveContext) Ext_config_list() IExt_config_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_listContext)
}

func (s *Extension_directiveContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Extension_directiveContext) AS_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserAS_, 0)
}

func (s *Extension_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_directiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExtension_directive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Extension_directive() (localctx IExtension_directiveContext) {
	this := p
	_ = this

	localctx = NewExtension_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KuneiformParserRULE_extension_directive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(KuneiformParserUSE_)
	}
	{
		p.SetState(100)
		p.Extension_name()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserL_BRACE {
		{
			p.SetState(101)
			p.Match(KuneiformParserL_BRACE)
		}
		{
			p.SetState(102)
			p.Ext_config_list()
		}
		{
			p.SetState(103)
			p.Match(KuneiformParserR_BRACE)
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserAS_ {
		{
			p.SetState(107)
			p.Match(KuneiformParserAS_)
		}
		{
			p.SetState(108)
			p.Extension_name()
		}

	}
	{
		p.SetState(111)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// IExt_config_listContext is an interface to support dynamic dispatch.
type IExt_config_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExt_config() []IExt_configContext
	Ext_config(i int) IExt_configContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExt_config_listContext differentiates from other interfaces.
	IsExt_config_listContext()
}

type Ext_config_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_listContext() *Ext_config_listContext {
	var p = new(Ext_config_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_list
	return p
}

func (*Ext_config_listContext) IsExt_config_listContext() {}

func NewExt_config_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_listContext {
	var p = new(Ext_config_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_list

	return p
}

func (s *Ext_config_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_listContext) AllExt_config() []IExt_configContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExt_configContext); ok {
			len++
		}
	}

	tst := make([]IExt_configContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExt_configContext); ok {
			tst[i] = t.(IExt_configContext)
			i++
		}
	}

	return tst
}

func (s *Ext_config_listContext) Ext_config(i int) IExt_configContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_configContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_configContext)
}

func (s *Ext_config_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Ext_config_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Ext_config_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_list() (localctx IExt_config_listContext) {
	this := p
	_ = this

	localctx = NewExt_config_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KuneiformParserRULE_ext_config_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Ext_config()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(114)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(115)
			p.Ext_config()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExt_configContext is an interface to support dynamic dispatch.
type IExt_configContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ext_config_name() IExt_config_nameContext
	COL() antlr.TerminalNode
	Ext_config_value() IExt_config_valueContext

	// IsExt_configContext differentiates from other interfaces.
	IsExt_configContext()
}

type Ext_configContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_configContext() *Ext_configContext {
	var p = new(Ext_configContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config
	return p
}

func (*Ext_configContext) IsExt_configContext() {}

func NewExt_configContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_configContext {
	var p = new(Ext_configContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config

	return p
}

func (s *Ext_configContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_configContext) Ext_config_name() IExt_config_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_nameContext)
}

func (s *Ext_configContext) COL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOL, 0)
}

func (s *Ext_configContext) Ext_config_value() IExt_config_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_valueContext)
}

func (s *Ext_configContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_configContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_configContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config() (localctx IExt_configContext) {
	this := p
	_ = this

	localctx = NewExt_configContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KuneiformParserRULE_ext_config)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Ext_config_name()
	}
	{
		p.SetState(122)
		p.Match(KuneiformParserCOL)
	}
	{
		p.SetState(123)
		p.Ext_config_value()
	}

	return localctx
}

// ITable_declContext is an interface to support dynamic dispatch.
type ITable_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TABLE_() antlr.TerminalNode
	Table_name() ITable_nameContext
	L_BRACE() antlr.TerminalNode
	Column_def_list() IColumn_def_listContext
	R_BRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllIndex_def() []IIndex_defContext
	Index_def(i int) IIndex_defContext
	AllForeign_key_def() []IForeign_key_defContext
	Foreign_key_def(i int) IForeign_key_defContext

	// IsTable_declContext differentiates from other interfaces.
	IsTable_declContext()
}

type Table_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_declContext() *Table_declContext {
	var p = new(Table_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_table_decl
	return p
}

func (*Table_declContext) IsTable_declContext() {}

func NewTable_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_declContext {
	var p = new(Table_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_table_decl

	return p
}

func (s *Table_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_declContext) TABLE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserTABLE_, 0)
}

func (s *Table_declContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Table_declContext) Column_def_list() IColumn_def_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_def_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_def_listContext)
}

func (s *Table_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Table_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Table_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Table_declContext) AllIndex_def() []IIndex_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_defContext); ok {
			len++
		}
	}

	tst := make([]IIndex_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_defContext); ok {
			tst[i] = t.(IIndex_defContext)
			i++
		}
	}

	return tst
}

func (s *Table_declContext) Index_def(i int) IIndex_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_defContext)
}

func (s *Table_declContext) AllForeign_key_def() []IForeign_key_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForeign_key_defContext); ok {
			len++
		}
	}

	tst := make([]IForeign_key_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForeign_key_defContext); ok {
			tst[i] = t.(IForeign_key_defContext)
			i++
		}
	}

	return tst
}

func (s *Table_declContext) Foreign_key_def(i int) IForeign_key_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeign_key_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeign_key_defContext)
}

func (s *Table_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitTable_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Table_decl() (localctx ITable_declContext) {
	this := p
	_ = this

	localctx = NewTable_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KuneiformParserRULE_table_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(KuneiformParserTABLE_)
	}
	{
		p.SetState(126)
		p.Table_name()
	}
	{
		p.SetState(127)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(128)
		p.Column_def_list()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(129)
				p.Match(KuneiformParserCOMMA)
			}
			p.SetState(132)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case KuneiformParserINDEX_NAME:
				{
					p.SetState(130)
					p.Index_def()
				}

			case KuneiformParserFOREIGN_KEY_, KuneiformParserFOREIGN_KEY_ABBR_:
				{
					p.SetState(131)
					p.Foreign_key_def()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserCOMMA {
		{
			p.SetState(139)
			p.Match(KuneiformParserCOMMA)
		}

	}
	{
		p.SetState(142)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IColumn_defContext is an interface to support dynamic dispatch.
type IColumn_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Column_name() IColumn_nameContext
	Column_type() IColumn_typeContext
	AllColumn_constraint() []IColumn_constraintContext
	Column_constraint(i int) IColumn_constraintContext

	// IsColumn_defContext differentiates from other interfaces.
	IsColumn_defContext()
}

type Column_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_defContext() *Column_defContext {
	var p = new(Column_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_def
	return p
}

func (*Column_defContext) IsColumn_defContext() {}

func NewColumn_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_defContext {
	var p = new(Column_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_def

	return p
}

func (s *Column_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_defContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_defContext) Column_type() IColumn_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_typeContext)
}

func (s *Column_defContext) AllColumn_constraint() []IColumn_constraintContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_constraintContext); ok {
			len++
		}
	}

	tst := make([]IColumn_constraintContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_constraintContext); ok {
			tst[i] = t.(IColumn_constraintContext)
			i++
		}
	}

	return tst
}

func (s *Column_defContext) Column_constraint(i int) IColumn_constraintContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_constraintContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_constraintContext)
}

func (s *Column_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_def() (localctx IColumn_defContext) {
	this := p
	_ = this

	localctx = NewColumn_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KuneiformParserRULE_column_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Column_name()
	}
	{
		p.SetState(145)
		p.Column_type()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&534773760) != 0 {
		{
			p.SetState(146)
			p.Column_constraint()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumn_def_listContext is an interface to support dynamic dispatch.
type IColumn_def_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_def() []IColumn_defContext
	Column_def(i int) IColumn_defContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_def_listContext differentiates from other interfaces.
	IsColumn_def_listContext()
}

type Column_def_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_def_listContext() *Column_def_listContext {
	var p = new(Column_def_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_def_list
	return p
}

func (*Column_def_listContext) IsColumn_def_listContext() {}

func NewColumn_def_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_def_listContext {
	var p = new(Column_def_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_def_list

	return p
}

func (s *Column_def_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_def_listContext) AllColumn_def() []IColumn_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_defContext); ok {
			len++
		}
	}

	tst := make([]IColumn_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_defContext); ok {
			tst[i] = t.(IColumn_defContext)
			i++
		}
	}

	return tst
}

func (s *Column_def_listContext) Column_def(i int) IColumn_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_defContext)
}

func (s *Column_def_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Column_def_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Column_def_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_def_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_def_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_def_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_def_list() (localctx IColumn_def_listContext) {
	this := p
	_ = this

	localctx = NewColumn_def_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KuneiformParserRULE_column_def_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Column_def()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(153)
				p.Match(KuneiformParserCOMMA)
			}
			{
				p.SetState(154)
				p.Column_def()
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IColumn_typeContext is an interface to support dynamic dispatch.
type IColumn_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_() antlr.TerminalNode
	TEXT_() antlr.TerminalNode

	// IsColumn_typeContext differentiates from other interfaces.
	IsColumn_typeContext()
}

type Column_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_typeContext() *Column_typeContext {
	var p = new(Column_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_type
	return p
}

func (*Column_typeContext) IsColumn_typeContext() {}

func NewColumn_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_typeContext {
	var p = new(Column_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_type

	return p
}

func (s *Column_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_typeContext) INT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINT_, 0)
}

func (s *Column_typeContext) TEXT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserTEXT_, 0)
}

func (s *Column_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_type() (localctx IColumn_typeContext) {
	this := p
	_ = this

	localctx = NewColumn_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KuneiformParserRULE_column_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserINT_ || _la == KuneiformParserTEXT_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IColumn_constraintContext is an interface to support dynamic dispatch.
type IColumn_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY_() antlr.TerminalNode
	NOT_NULL_() antlr.TerminalNode
	UNIQUE_() antlr.TerminalNode
	DEFAULT_() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Literal_value() ILiteral_valueContext
	R_PAREN() antlr.TerminalNode
	MIN_() antlr.TerminalNode
	Number_value() INumber_valueContext
	MAX_() antlr.TerminalNode
	MIN_LEN_() antlr.TerminalNode
	MAX_LEN_() antlr.TerminalNode

	// IsColumn_constraintContext differentiates from other interfaces.
	IsColumn_constraintContext()
}

type Column_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_constraintContext() *Column_constraintContext {
	var p = new(Column_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_constraint
	return p
}

func (*Column_constraintContext) IsColumn_constraintContext() {}

func NewColumn_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_constraintContext {
	var p = new(Column_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_constraint

	return p
}

func (s *Column_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_constraintContext) PRIMARY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIMARY_, 0)
}

func (s *Column_constraintContext) NOT_NULL_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserNOT_NULL_, 0)
}

func (s *Column_constraintContext) UNIQUE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNIQUE_, 0)
}

func (s *Column_constraintContext) DEFAULT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDEFAULT_, 0)
}

func (s *Column_constraintContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Column_constraintContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Column_constraintContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Column_constraintContext) MIN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMIN_, 0)
}

func (s *Column_constraintContext) Number_value() INumber_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_valueContext)
}

func (s *Column_constraintContext) MAX_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMAX_, 0)
}

func (s *Column_constraintContext) MIN_LEN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMIN_LEN_, 0)
}

func (s *Column_constraintContext) MAX_LEN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMAX_LEN_, 0)
}

func (s *Column_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_constraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_constraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_constraint() (localctx IColumn_constraintContext) {
	this := p
	_ = this

	localctx = NewColumn_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KuneiformParserRULE_column_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserPRIMARY_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(KuneiformParserPRIMARY_)
		}

	case KuneiformParserNOT_NULL_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(KuneiformParserNOT_NULL_)
		}

	case KuneiformParserUNIQUE_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.Match(KuneiformParserUNIQUE_)
		}

	case KuneiformParserDEFAULT_:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(165)
			p.Match(KuneiformParserDEFAULT_)
		}
		{
			p.SetState(166)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(167)
			p.Literal_value()
		}
		{
			p.SetState(168)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(170)
			p.Match(KuneiformParserMIN_)
		}
		{
			p.SetState(171)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(172)
			p.Number_value()
		}
		{
			p.SetState(173)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(175)
			p.Match(KuneiformParserMAX_)
		}
		{
			p.SetState(176)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(177)
			p.Number_value()
		}
		{
			p.SetState(178)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_LEN_:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(180)
			p.Match(KuneiformParserMIN_LEN_)
		}
		{
			p.SetState(181)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(182)
			p.Number_value()
		}
		{
			p.SetState(183)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_LEN_:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(185)
			p.Match(KuneiformParserMAX_LEN_)
		}
		{
			p.SetState(186)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(187)
			p.Number_value()
		}
		{
			p.SetState(188)
			p.Match(KuneiformParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_literal_value
	return p
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSTRING_LITERAL, 0)
}

func (s *Literal_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitLiteral_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Literal_value() (localctx ILiteral_valueContext) {
	this := p
	_ = this

	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KuneiformParserRULE_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserUNSIGNED_NUMBER_LITERAL || _la == KuneiformParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumber_valueContext is an interface to support dynamic dispatch.
type INumber_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsNumber_valueContext differentiates from other interfaces.
	IsNumber_valueContext()
}

type Number_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_valueContext() *Number_valueContext {
	var p = new(Number_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_number_value
	return p
}

func (*Number_valueContext) IsNumber_valueContext() {}

func NewNumber_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_valueContext {
	var p = new(Number_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_number_value

	return p
}

func (s *Number_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Number_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitNumber_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Number_value() (localctx INumber_valueContext) {
	this := p
	_ = this

	localctx = NewNumber_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KuneiformParserRULE_number_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(KuneiformParserUNSIGNED_NUMBER_LITERAL)
	}

	return localctx
}

// IIndex_defContext is an interface to support dynamic dispatch.
type IIndex_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Index_name() IIndex_nameContext
	L_PAREN() antlr.TerminalNode
	Column_name_list() IColumn_name_listContext
	R_PAREN() antlr.TerminalNode
	UNIQUE_() antlr.TerminalNode
	INDEX_() antlr.TerminalNode
	PRIMARY_() antlr.TerminalNode

	// IsIndex_defContext differentiates from other interfaces.
	IsIndex_defContext()
}

type Index_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_defContext() *Index_defContext {
	var p = new(Index_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_index_def
	return p
}

func (*Index_defContext) IsIndex_defContext() {}

func NewIndex_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_defContext {
	var p = new(Index_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_index_def

	return p
}

func (s *Index_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_defContext) Index_name() IIndex_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Index_defContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Index_defContext) Column_name_list() IColumn_name_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Index_defContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Index_defContext) UNIQUE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNIQUE_, 0)
}

func (s *Index_defContext) INDEX_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINDEX_, 0)
}

func (s *Index_defContext) PRIMARY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIMARY_, 0)
}

func (s *Index_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitIndex_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Index_def() (localctx IIndex_defContext) {
	this := p
	_ = this

	localctx = NewIndex_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KuneiformParserRULE_index_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Index_name()
	}
	{
		p.SetState(197)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&872415232) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(198)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(199)
		p.Column_name_list()
	}
	{
		p.SetState(200)
		p.Match(KuneiformParserR_PAREN)
	}

	return localctx
}

// IForeign_key_actionContext is an interface to support dynamic dispatch.
type IForeign_key_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_ON_UPDATE_() antlr.TerminalNode
	ACTION_ON_DELETE_() antlr.TerminalNode
	ACTION_DO_NO_ACTION_() antlr.TerminalNode
	ACTION_DO_RESTRICT_() antlr.TerminalNode
	ACTION_DO_SET_NULL_() antlr.TerminalNode
	ACTION_DO_SET_DEFAULT_() antlr.TerminalNode
	ACTION_DO_CASCADE_() antlr.TerminalNode
	ACTION_DO_() antlr.TerminalNode

	// IsForeign_key_actionContext differentiates from other interfaces.
	IsForeign_key_actionContext()
}

type Foreign_key_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeign_key_actionContext() *Foreign_key_actionContext {
	var p = new(Foreign_key_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_foreign_key_action
	return p
}

func (*Foreign_key_actionContext) IsForeign_key_actionContext() {}

func NewForeign_key_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Foreign_key_actionContext {
	var p = new(Foreign_key_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_foreign_key_action

	return p
}

func (s *Foreign_key_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Foreign_key_actionContext) ACTION_ON_UPDATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_ON_UPDATE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_ON_DELETE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_ON_DELETE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_NO_ACTION_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_NO_ACTION_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_RESTRICT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_RESTRICT_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_SET_NULL_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_SET_NULL_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_SET_DEFAULT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_SET_DEFAULT_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_CASCADE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_CASCADE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_, 0)
}

func (s *Foreign_key_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Foreign_key_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Foreign_key_actionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitForeign_key_action(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Foreign_key_action() (localctx IForeign_key_actionContext) {
	this := p
	_ = this

	localctx = NewForeign_key_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KuneiformParserRULE_foreign_key_action)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserACTION_DO_ {
		{
			p.SetState(203)
			p.Match(KuneiformParserACTION_DO_)
		}

	}
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4260607557632) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IForeign_key_defContext is an interface to support dynamic dispatch.
type IForeign_key_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllL_PAREN() []antlr.TerminalNode
	L_PAREN(i int) antlr.TerminalNode
	AllColumn_name_list() []IColumn_name_listContext
	Column_name_list(i int) IColumn_name_listContext
	AllR_PAREN() []antlr.TerminalNode
	R_PAREN(i int) antlr.TerminalNode
	Table_name() ITable_nameContext
	FOREIGN_KEY_() antlr.TerminalNode
	FOREIGN_KEY_ABBR_() antlr.TerminalNode
	REFERENCES_() antlr.TerminalNode
	REFERENCES_ABBR_() antlr.TerminalNode
	AllForeign_key_action() []IForeign_key_actionContext
	Foreign_key_action(i int) IForeign_key_actionContext

	// IsForeign_key_defContext differentiates from other interfaces.
	IsForeign_key_defContext()
}

type Foreign_key_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeign_key_defContext() *Foreign_key_defContext {
	var p = new(Foreign_key_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_foreign_key_def
	return p
}

func (*Foreign_key_defContext) IsForeign_key_defContext() {}

func NewForeign_key_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Foreign_key_defContext {
	var p = new(Foreign_key_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_foreign_key_def

	return p
}

func (s *Foreign_key_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Foreign_key_defContext) AllL_PAREN() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserL_PAREN)
}

func (s *Foreign_key_defContext) L_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, i)
}

func (s *Foreign_key_defContext) AllColumn_name_list() []IColumn_name_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			len++
		}
	}

	tst := make([]IColumn_name_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_name_listContext); ok {
			tst[i] = t.(IColumn_name_listContext)
			i++
		}
	}

	return tst
}

func (s *Foreign_key_defContext) Column_name_list(i int) IColumn_name_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Foreign_key_defContext) AllR_PAREN() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserR_PAREN)
}

func (s *Foreign_key_defContext) R_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, i)
}

func (s *Foreign_key_defContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Foreign_key_defContext) FOREIGN_KEY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserFOREIGN_KEY_, 0)
}

func (s *Foreign_key_defContext) FOREIGN_KEY_ABBR_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserFOREIGN_KEY_ABBR_, 0)
}

func (s *Foreign_key_defContext) REFERENCES_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserREFERENCES_, 0)
}

func (s *Foreign_key_defContext) REFERENCES_ABBR_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserREFERENCES_ABBR_, 0)
}

func (s *Foreign_key_defContext) AllForeign_key_action() []IForeign_key_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForeign_key_actionContext); ok {
			len++
		}
	}

	tst := make([]IForeign_key_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForeign_key_actionContext); ok {
			tst[i] = t.(IForeign_key_actionContext)
			i++
		}
	}

	return tst
}

func (s *Foreign_key_defContext) Foreign_key_action(i int) IForeign_key_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeign_key_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeign_key_actionContext)
}

func (s *Foreign_key_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Foreign_key_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Foreign_key_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitForeign_key_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Foreign_key_def() (localctx IForeign_key_defContext) {
	this := p
	_ = this

	localctx = NewForeign_key_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KuneiformParserRULE_foreign_key_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserFOREIGN_KEY_ || _la == KuneiformParserFOREIGN_KEY_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(209)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(210)
		p.Column_name_list()
	}
	{
		p.SetState(211)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(212)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserREFERENCES_ || _la == KuneiformParserREFERENCES_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(213)
		p.Table_name()
	}
	{
		p.SetState(214)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(215)
		p.Column_name_list()
	}
	{
		p.SetState(216)
		p.Match(KuneiformParserR_PAREN)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_ {
		{
			p.SetState(217)
			p.Foreign_key_action()
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_declContext is an interface to support dynamic dispatch.
type IAction_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_() antlr.TerminalNode
	Action_name() IAction_nameContext
	L_PAREN() antlr.TerminalNode
	Param_list() IParam_listContext
	R_PAREN() antlr.TerminalNode
	ACTION_OPEN() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	ACTION_CLOSE() antlr.TerminalNode

	// IsAction_declContext differentiates from other interfaces.
	IsAction_declContext()
}

type Action_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_declContext() *Action_declContext {
	var p = new(Action_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_decl
	return p
}

func (*Action_declContext) IsAction_declContext() {}

func NewAction_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_declContext {
	var p = new(Action_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_decl

	return p
}

func (s *Action_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_declContext) ACTION_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_, 0)
}

func (s *Action_declContext) Action_name() IAction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_nameContext)
}

func (s *Action_declContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Action_declContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Action_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Action_declContext) ACTION_OPEN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_OPEN, 0)
}

func (s *Action_declContext) Action_stmt_list() IAction_stmt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmt_listContext)
}

func (s *Action_declContext) ACTION_CLOSE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_CLOSE, 0)
}

func (s *Action_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_decl() (localctx IAction_declContext) {
	this := p
	_ = this

	localctx = NewAction_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KuneiformParserRULE_action_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(KuneiformParserACTION_)
	}
	{
		p.SetState(224)
		p.Action_name()
	}
	{
		p.SetState(225)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(226)
		p.Param_list()
	}
	{
		p.SetState(227)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(228)
		p.Match(KuneiformParserACTION_OPEN)
	}
	{
		p.SetState(229)
		p.Action_stmt_list()
	}
	{
		p.SetState(230)
		p.Match(KuneiformParserACTION_CLOSE)
	}

	return localctx
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPARAMETER() []antlr.TerminalNode
	PARAMETER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_param_list
	return p
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllPARAMETER() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserPARAMETER)
}

func (s *Param_listContext) PARAMETER(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserPARAMETER, i)
}

func (s *Param_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Param_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitParam_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Param_list() (localctx IParam_listContext) {
	this := p
	_ = this

	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KuneiformParserRULE_param_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPARAMETER {
		{
			p.SetState(232)
			p.Match(KuneiformParserPARAMETER)
		}

	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(235)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(236)
			p.Match(KuneiformParserPARAMETER)
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDatabase_nameContext is an interface to support dynamic dispatch.
type IDatabase_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsDatabase_nameContext differentiates from other interfaces.
	IsDatabase_nameContext()
}

type Database_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabase_nameContext() *Database_nameContext {
	var p = new(Database_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_database_name
	return p
}

func (*Database_nameContext) IsDatabase_nameContext() {}

func NewDatabase_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Database_nameContext {
	var p = new(Database_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_database_name

	return p
}

func (s *Database_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Database_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Database_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Database_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Database_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitDatabase_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Database_name() (localctx IDatabase_nameContext) {
	this := p
	_ = this

	localctx = NewDatabase_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KuneiformParserRULE_database_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IExtension_nameContext is an interface to support dynamic dispatch.
type IExtension_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsExtension_nameContext differentiates from other interfaces.
	IsExtension_nameContext()
}

type Extension_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_nameContext() *Extension_nameContext {
	var p = new(Extension_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_extension_name
	return p
}

func (*Extension_nameContext) IsExtension_nameContext() {}

func NewExtension_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_nameContext {
	var p = new(Extension_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_extension_name

	return p
}

func (s *Extension_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Extension_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExtension_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Extension_name() (localctx IExtension_nameContext) {
	this := p
	_ = this

	localctx = NewExtension_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KuneiformParserRULE_extension_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IExt_config_nameContext is an interface to support dynamic dispatch.
type IExt_config_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsExt_config_nameContext differentiates from other interfaces.
	IsExt_config_nameContext()
}

type Ext_config_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_nameContext() *Ext_config_nameContext {
	var p = new(Ext_config_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_name
	return p
}

func (*Ext_config_nameContext) IsExt_config_nameContext() {}

func NewExt_config_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_nameContext {
	var p = new(Ext_config_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_name

	return p
}

func (s *Ext_config_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Ext_config_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_name() (localctx IExt_config_nameContext) {
	this := p
	_ = this

	localctx = NewExt_config_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KuneiformParserRULE_ext_config_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitTable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Table_name() (localctx ITable_nameContext) {
	this := p
	_ = this

	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KuneiformParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IAction_nameContext is an interface to support dynamic dispatch.
type IAction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAction_nameContext differentiates from other interfaces.
	IsAction_nameContext()
}

type Action_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_nameContext() *Action_nameContext {
	var p = new(Action_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_name
	return p
}

func (*Action_nameContext) IsAction_nameContext() {}

func NewAction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_nameContext {
	var p = new(Action_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_name

	return p
}

func (s *Action_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Action_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_name() (localctx IAction_nameContext) {
	this := p
	_ = this

	localctx = NewAction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KuneiformParserRULE_action_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_name() (localctx IColumn_nameContext) {
	this := p
	_ = this

	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, KuneiformParserRULE_column_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IColumn_name_listContext is an interface to support dynamic dispatch.
type IColumn_name_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_name_listContext differentiates from other interfaces.
	IsColumn_name_listContext()
}

type Column_name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_listContext() *Column_name_listContext {
	var p = new(Column_name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_name_list
	return p
}

func (*Column_name_listContext) IsColumn_name_listContext() {}

func NewColumn_name_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_listContext {
	var p = new(Column_name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_name_list

	return p
}

func (s *Column_name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_listContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_name_listContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_name_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Column_name_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Column_name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_name_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_name_list() (localctx IColumn_name_listContext) {
	this := p
	_ = this

	localctx = NewColumn_name_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, KuneiformParserRULE_column_name_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Column_name()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(255)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(256)
			p.Column_name()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX_NAME() antlr.TerminalNode

	// IsIndex_nameContext differentiates from other interfaces.
	IsIndex_nameContext()
}

type Index_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_index_name
	return p
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) INDEX_NAME() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINDEX_NAME, 0)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitIndex_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Index_name() (localctx IIndex_nameContext) {
	this := p
	_ = this

	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, KuneiformParserRULE_index_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(KuneiformParserINDEX_NAME)
	}

	return localctx
}

// IExt_config_valueContext is an interface to support dynamic dispatch.
type IExt_config_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext

	// IsExt_config_valueContext differentiates from other interfaces.
	IsExt_config_valueContext()
}

type Ext_config_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_valueContext() *Ext_config_valueContext {
	var p = new(Ext_config_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_value
	return p
}

func (*Ext_config_valueContext) IsExt_config_valueContext() {}

func NewExt_config_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_valueContext {
	var p = new(Ext_config_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_value

	return p
}

func (s *Ext_config_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_valueContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Ext_config_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_value() (localctx IExt_config_valueContext) {
	this := p
	_ = this

	localctx = NewExt_config_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, KuneiformParserRULE_ext_config_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Literal_value()
	}

	return localctx
}

// IInit_declContext is an interface to support dynamic dispatch.
type IInit_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INIT_OPEN() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	ACTION_CLOSE() antlr.TerminalNode

	// IsInit_declContext differentiates from other interfaces.
	IsInit_declContext()
}

type Init_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInit_declContext() *Init_declContext {
	var p = new(Init_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_init_decl
	return p
}

func (*Init_declContext) IsInit_declContext() {}

func NewInit_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Init_declContext {
	var p = new(Init_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_init_decl

	return p
}

func (s *Init_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Init_declContext) INIT_OPEN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINIT_OPEN, 0)
}

func (s *Init_declContext) Action_stmt_list() IAction_stmt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmt_listContext)
}

func (s *Init_declContext) ACTION_CLOSE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_CLOSE, 0)
}

func (s *Init_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Init_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Init_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitInit_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Init_decl() (localctx IInit_declContext) {
	this := p
	_ = this

	localctx = NewInit_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, KuneiformParserRULE_init_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(KuneiformParserINIT_OPEN)
	}
	{
		p.SetState(267)
		p.Action_stmt_list()
	}
	{
		p.SetState(268)
		p.Match(KuneiformParserACTION_CLOSE)
	}

	return localctx
}

// IAction_stmt_listContext is an interface to support dynamic dispatch.
type IAction_stmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAction_stmt() []IAction_stmtContext
	Action_stmt(i int) IAction_stmtContext

	// IsAction_stmt_listContext differentiates from other interfaces.
	IsAction_stmt_listContext()
}

type Action_stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_stmt_listContext() *Action_stmt_listContext {
	var p = new(Action_stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_stmt_list
	return p
}

func (*Action_stmt_listContext) IsAction_stmt_listContext() {}

func NewAction_stmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_stmt_listContext {
	var p = new(Action_stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_stmt_list

	return p
}

func (s *Action_stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_stmt_listContext) AllAction_stmt() []IAction_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_stmtContext); ok {
			len++
		}
	}

	tst := make([]IAction_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_stmtContext); ok {
			tst[i] = t.(IAction_stmtContext)
			i++
		}
	}

	return tst
}

func (s *Action_stmt_listContext) Action_stmt(i int) IAction_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmtContext)
}

func (s *Action_stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_stmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_stmt_list() (localctx IAction_stmt_listContext) {
	this := p
	_ = this

	localctx = NewAction_stmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, KuneiformParserRULE_action_stmt_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-70)) & ^0x3f) == 0 && ((int64(1)<<(_la-70))&67) != 0) {
		{
			p.SetState(270)
			p.Action_stmt()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_stmtContext is an interface to support dynamic dispatch.
type IAction_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_sql_stmt() IA_sql_stmtContext
	A_call_stmt() IA_call_stmtContext

	// IsAction_stmtContext differentiates from other interfaces.
	IsAction_stmtContext()
}

type Action_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_stmtContext() *Action_stmtContext {
	var p = new(Action_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_stmt
	return p
}

func (*Action_stmtContext) IsAction_stmtContext() {}

func NewAction_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_stmtContext {
	var p = new(Action_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_stmt

	return p
}

func (s *Action_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_stmtContext) A_sql_stmt() IA_sql_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_sql_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_sql_stmtContext)
}

func (s *Action_stmtContext) A_call_stmt() IA_call_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_call_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_call_stmtContext)
}

func (s *Action_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_stmt() (localctx IAction_stmtContext) {
	this := p
	_ = this

	localctx = NewAction_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, KuneiformParserRULE_action_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserA_SQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.A_sql_stmt()
		}

	case KuneiformParserA_IDENTIFIER, KuneiformParserA_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.A_call_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IA_sql_stmtContext is an interface to support dynamic dispatch.
type IA_sql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_SQL_STMT() antlr.TerminalNode
	A_STMT_END() antlr.TerminalNode

	// IsA_sql_stmtContext differentiates from other interfaces.
	IsA_sql_stmtContext()
}

type A_sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_sql_stmtContext() *A_sql_stmtContext {
	var p = new(A_sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_sql_stmt
	return p
}

func (*A_sql_stmtContext) IsA_sql_stmtContext() {}

func NewA_sql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_sql_stmtContext {
	var p = new(A_sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_sql_stmt

	return p
}

func (s *A_sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *A_sql_stmtContext) A_SQL_STMT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_SQL_STMT, 0)
}

func (s *A_sql_stmtContext) A_STMT_END() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_STMT_END, 0)
}

func (s *A_sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_sql_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_sql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_sql_stmt() (localctx IA_sql_stmtContext) {
	this := p
	_ = this

	localctx = NewA_sql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, KuneiformParserRULE_a_sql_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(KuneiformParserA_SQL_STMT)
	}
	{
		p.SetState(280)
		p.Match(KuneiformParserA_STMT_END)
	}

	return localctx
}

// IA_variable_nameContext is an interface to support dynamic dispatch.
type IA_variable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_VARIABLE() antlr.TerminalNode

	// IsA_variable_nameContext differentiates from other interfaces.
	IsA_variable_nameContext()
}

type A_variable_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_variable_nameContext() *A_variable_nameContext {
	var p = new(A_variable_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_variable_name
	return p
}

func (*A_variable_nameContext) IsA_variable_nameContext() {}

func NewA_variable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_variable_nameContext {
	var p = new(A_variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_variable_name

	return p
}

func (s *A_variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *A_variable_nameContext) A_VARIABLE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_VARIABLE, 0)
}

func (s *A_variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_variable_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_variable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_variable_name() (localctx IA_variable_nameContext) {
	this := p
	_ = this

	localctx = NewA_variable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, KuneiformParserRULE_a_variable_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(KuneiformParserA_VARIABLE)
	}

	return localctx
}

// IA_block_variable_nameContext is an interface to support dynamic dispatch.
type IA_block_variable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_REF() antlr.TerminalNode

	// IsA_block_variable_nameContext differentiates from other interfaces.
	IsA_block_variable_nameContext()
}

type A_block_variable_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_block_variable_nameContext() *A_block_variable_nameContext {
	var p = new(A_block_variable_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_block_variable_name
	return p
}

func (*A_block_variable_nameContext) IsA_block_variable_nameContext() {}

func NewA_block_variable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_block_variable_nameContext {
	var p = new(A_block_variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_block_variable_name

	return p
}

func (s *A_block_variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *A_block_variable_nameContext) A_REF() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_REF, 0)
}

func (s *A_block_variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_block_variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_block_variable_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_block_variable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_block_variable_name() (localctx IA_block_variable_nameContext) {
	this := p
	_ = this

	localctx = NewA_block_variable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, KuneiformParserRULE_a_block_variable_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(KuneiformParserA_REF)
	}

	return localctx
}

// IA_literal_valueContext is an interface to support dynamic dispatch.
type IA_literal_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_STRING_LITERAL() antlr.TerminalNode
	A_UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsA_literal_valueContext differentiates from other interfaces.
	IsA_literal_valueContext()
}

type A_literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_literal_valueContext() *A_literal_valueContext {
	var p = new(A_literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_literal_value
	return p
}

func (*A_literal_valueContext) IsA_literal_valueContext() {}

func NewA_literal_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_literal_valueContext {
	var p = new(A_literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_literal_value

	return p
}

func (s *A_literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *A_literal_valueContext) A_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_STRING_LITERAL, 0)
}

func (s *A_literal_valueContext) A_UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_UNSIGNED_NUMBER_LITERAL, 0)
}

func (s *A_literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_literal_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_literal_value() (localctx IA_literal_valueContext) {
	this := p
	_ = this

	localctx = NewA_literal_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, KuneiformParserRULE_a_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserA_UNSIGNED_NUMBER_LITERAL || _la == KuneiformParserA_STRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IA_fn_nameContext is an interface to support dynamic dispatch.
type IA_fn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllA_IDENTIFIER() []antlr.TerminalNode
	A_IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsA_fn_nameContext differentiates from other interfaces.
	IsA_fn_nameContext()
}

type A_fn_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_fn_nameContext() *A_fn_nameContext {
	var p = new(A_fn_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_fn_name
	return p
}

func (*A_fn_nameContext) IsA_fn_nameContext() {}

func NewA_fn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_fn_nameContext {
	var p = new(A_fn_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_fn_name

	return p
}

func (s *A_fn_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *A_fn_nameContext) AllA_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserA_IDENTIFIER)
}

func (s *A_fn_nameContext) A_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_IDENTIFIER, i)
}

func (s *A_fn_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPERIOD, 0)
}

func (s *A_fn_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_fn_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_fn_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_fn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_fn_name() (localctx IA_fn_nameContext) {
	this := p
	_ = this

	localctx = NewA_fn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, KuneiformParserRULE_a_fn_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(KuneiformParserA_IDENTIFIER)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPERIOD {
		{
			p.SetState(289)
			p.Match(KuneiformParserPERIOD)
		}
		{
			p.SetState(290)
			p.Match(KuneiformParserA_IDENTIFIER)
		}

	}

	return localctx
}

// IA_call_receiversContext is an interface to support dynamic dispatch.
type IA_call_receiversContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllA_variable_name() []IA_variable_nameContext
	A_variable_name(i int) IA_variable_nameContext
	AllA_COMMA() []antlr.TerminalNode
	A_COMMA(i int) antlr.TerminalNode

	// IsA_call_receiversContext differentiates from other interfaces.
	IsA_call_receiversContext()
}

type A_call_receiversContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_call_receiversContext() *A_call_receiversContext {
	var p = new(A_call_receiversContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_call_receivers
	return p
}

func (*A_call_receiversContext) IsA_call_receiversContext() {}

func NewA_call_receiversContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_call_receiversContext {
	var p = new(A_call_receiversContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_call_receivers

	return p
}

func (s *A_call_receiversContext) GetParser() antlr.Parser { return s.parser }

func (s *A_call_receiversContext) AllA_variable_name() []IA_variable_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IA_variable_nameContext); ok {
			len++
		}
	}

	tst := make([]IA_variable_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IA_variable_nameContext); ok {
			tst[i] = t.(IA_variable_nameContext)
			i++
		}
	}

	return tst
}

func (s *A_call_receiversContext) A_variable_name(i int) IA_variable_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_variable_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_variable_nameContext)
}

func (s *A_call_receiversContext) AllA_COMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserA_COMMA)
}

func (s *A_call_receiversContext) A_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_COMMA, i)
}

func (s *A_call_receiversContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_call_receiversContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_call_receiversContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_call_receivers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_call_receivers() (localctx IA_call_receiversContext) {
	this := p
	_ = this

	localctx = NewA_call_receiversContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KuneiformParserRULE_a_call_receivers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.A_variable_name()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserA_COMMA {
		{
			p.SetState(294)
			p.Match(KuneiformParserA_COMMA)
		}
		{
			p.SetState(295)
			p.A_variable_name()
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IA_call_stmtContext is an interface to support dynamic dispatch.
type IA_call_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_call_body() IA_call_bodyContext
	A_STMT_END() antlr.TerminalNode
	A_call_receivers() IA_call_receiversContext
	EQ() antlr.TerminalNode

	// IsA_call_stmtContext differentiates from other interfaces.
	IsA_call_stmtContext()
}

type A_call_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_call_stmtContext() *A_call_stmtContext {
	var p = new(A_call_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_call_stmt
	return p
}

func (*A_call_stmtContext) IsA_call_stmtContext() {}

func NewA_call_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_call_stmtContext {
	var p = new(A_call_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_call_stmt

	return p
}

func (s *A_call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *A_call_stmtContext) A_call_body() IA_call_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_call_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_call_bodyContext)
}

func (s *A_call_stmtContext) A_STMT_END() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_STMT_END, 0)
}

func (s *A_call_stmtContext) A_call_receivers() IA_call_receiversContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_call_receiversContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_call_receiversContext)
}

func (s *A_call_stmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(KuneiformParserEQ, 0)
}

func (s *A_call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_call_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_call_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_call_stmt() (localctx IA_call_stmtContext) {
	this := p
	_ = this

	localctx = NewA_call_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KuneiformParserRULE_a_call_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserA_VARIABLE {
		{
			p.SetState(301)
			p.A_call_receivers()
		}
		{
			p.SetState(302)
			p.Match(KuneiformParserEQ)
		}

	}
	{
		p.SetState(306)
		p.A_call_body()
	}
	{
		p.SetState(307)
		p.Match(KuneiformParserA_STMT_END)
	}

	return localctx
}

// IA_call_bodyContext is an interface to support dynamic dispatch.
type IA_call_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_fn_name() IA_fn_nameContext
	A_L_PAREN() antlr.TerminalNode
	A_fn_arg_list() IA_fn_arg_listContext
	A_R_PAREN() antlr.TerminalNode

	// IsA_call_bodyContext differentiates from other interfaces.
	IsA_call_bodyContext()
}

type A_call_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_call_bodyContext() *A_call_bodyContext {
	var p = new(A_call_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_call_body
	return p
}

func (*A_call_bodyContext) IsA_call_bodyContext() {}

func NewA_call_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_call_bodyContext {
	var p = new(A_call_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_call_body

	return p
}

func (s *A_call_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *A_call_bodyContext) A_fn_name() IA_fn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_fn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_fn_nameContext)
}

func (s *A_call_bodyContext) A_L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_L_PAREN, 0)
}

func (s *A_call_bodyContext) A_fn_arg_list() IA_fn_arg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_fn_arg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_fn_arg_listContext)
}

func (s *A_call_bodyContext) A_R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_R_PAREN, 0)
}

func (s *A_call_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_call_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_call_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_call_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_call_body() (localctx IA_call_bodyContext) {
	this := p
	_ = this

	localctx = NewA_call_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KuneiformParserRULE_a_call_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.A_fn_name()
	}
	{
		p.SetState(310)
		p.Match(KuneiformParserA_L_PAREN)
	}
	{
		p.SetState(311)
		p.A_fn_arg_list()
	}
	{
		p.SetState(312)
		p.Match(KuneiformParserA_R_PAREN)
	}

	return localctx
}

// IA_fn_arg_listContext is an interface to support dynamic dispatch.
type IA_fn_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllA_fn_arg() []IA_fn_argContext
	A_fn_arg(i int) IA_fn_argContext
	AllA_COMMA() []antlr.TerminalNode
	A_COMMA(i int) antlr.TerminalNode

	// IsA_fn_arg_listContext differentiates from other interfaces.
	IsA_fn_arg_listContext()
}

type A_fn_arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_fn_arg_listContext() *A_fn_arg_listContext {
	var p = new(A_fn_arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_fn_arg_list
	return p
}

func (*A_fn_arg_listContext) IsA_fn_arg_listContext() {}

func NewA_fn_arg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_fn_arg_listContext {
	var p = new(A_fn_arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_fn_arg_list

	return p
}

func (s *A_fn_arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *A_fn_arg_listContext) AllA_fn_arg() []IA_fn_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IA_fn_argContext); ok {
			len++
		}
	}

	tst := make([]IA_fn_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IA_fn_argContext); ok {
			tst[i] = t.(IA_fn_argContext)
			i++
		}
	}

	return tst
}

func (s *A_fn_arg_listContext) A_fn_arg(i int) IA_fn_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_fn_argContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_fn_argContext)
}

func (s *A_fn_arg_listContext) AllA_COMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserA_COMMA)
}

func (s *A_fn_arg_listContext) A_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserA_COMMA, i)
}

func (s *A_fn_arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_fn_arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_fn_arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_fn_arg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_fn_arg_list() (localctx IA_fn_arg_listContext) {
	this := p
	_ = this

	localctx = NewA_fn_arg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KuneiformParserRULE_a_fn_arg_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&23) != 0 {
		{
			p.SetState(314)
			p.A_fn_arg()
		}

	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserA_COMMA {
		{
			p.SetState(317)
			p.Match(KuneiformParserA_COMMA)
		}
		{
			p.SetState(318)
			p.A_fn_arg()
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IA_fn_argContext is an interface to support dynamic dispatch.
type IA_fn_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A_literal_value() IA_literal_valueContext
	A_variable_name() IA_variable_nameContext
	A_block_variable_name() IA_block_variable_nameContext

	// IsA_fn_argContext differentiates from other interfaces.
	IsA_fn_argContext()
}

type A_fn_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_fn_argContext() *A_fn_argContext {
	var p = new(A_fn_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_a_fn_arg
	return p
}

func (*A_fn_argContext) IsA_fn_argContext() {}

func NewA_fn_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_fn_argContext {
	var p = new(A_fn_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_a_fn_arg

	return p
}

func (s *A_fn_argContext) GetParser() antlr.Parser { return s.parser }

func (s *A_fn_argContext) A_literal_value() IA_literal_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_literal_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_literal_valueContext)
}

func (s *A_fn_argContext) A_variable_name() IA_variable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_variable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_variable_nameContext)
}

func (s *A_fn_argContext) A_block_variable_name() IA_block_variable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IA_block_variable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IA_block_variable_nameContext)
}

func (s *A_fn_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_fn_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_fn_argContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitA_fn_arg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) A_fn_arg() (localctx IA_fn_argContext) {
	this := p
	_ = this

	localctx = NewA_fn_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KuneiformParserRULE_a_fn_arg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserA_UNSIGNED_NUMBER_LITERAL, KuneiformParserA_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.A_literal_value()
		}

	case KuneiformParserA_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.A_variable_name()
		}

	case KuneiformParserA_REF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.A_block_variable_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
