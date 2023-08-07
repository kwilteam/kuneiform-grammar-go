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
		"", "':'", "';'", "'('", "'{'", "')'", "'}'", "','", "'$'", "'#'", "'@'",
		"'.'", "'='", "'database'", "'use'", "'as'", "'table'", "'action'",
		"'init'", "'public'", "'private'", "'view'", "'mustsign'", "'int'",
		"'text'", "'min'", "'max'", "'minlen'", "'maxlen'", "'notnull'", "'primary'",
		"'default'", "'unique'", "'index'", "'foreign_key'", "'fk'", "'references'",
		"'ref'", "'on_update'", "'on_delete'", "'do'", "'no_action'", "'cascade'",
		"'set_null'", "'set_default'", "'restrict'",
	}
	staticData.symbolicNames = []string{
		"", "COL", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA",
		"DOLLAR", "HASH", "AT", "PERIOD", "EQ", "DATABASE_", "USE_", "AS_",
		"TABLE_", "ACTION_", "INIT_", "PUBLIC_", "PRIVATE_", "VIEW_", "MUSTSIGN_",
		"INT_", "TEXT_", "MIN_", "MAX_", "MIN_LEN_", "MAX_LEN_", "NOT_NULL_",
		"PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_", "FOREIGN_KEY_", "FOREIGN_KEY_ABBR_",
		"REFERENCES_", "REFERENCES_ABBR_", "ACTION_ON_UPDATE_", "ACTION_ON_DELETE_",
		"ACTION_DO_", "ACTION_DO_NO_ACTION_", "ACTION_DO_CASCADE_", "ACTION_DO_SET_NULL_",
		"ACTION_DO_SET_DEFAULT_", "ACTION_DO_RESTRICT_", "SELECT_", "INSERT_",
		"UPDATE_", "DELETE_", "WITH_", "IDENTIFIER", "INDEX_NAME", "PARAM_OR_VAR",
		"BLOCK_VAR", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL", "STRING_LITERAL",
		"SQL_KEYWORDS", "SQL_STMT", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"source_unit", "database_directive", "extension_directive", "ext_config_list",
		"ext_config", "table_decl", "column_def", "column_def_list", "column_type",
		"column_constraint", "literal_value", "number_value", "index_def", "foreign_key_action",
		"foreign_key_def", "action_visibility", "action_mutability", "action_auxiliary",
		"action_attr_list", "action_decl", "param_list", "parameter", "database_name",
		"extension_name", "ext_config_name", "table_name", "action_name", "column_name",
		"column_name_list", "index_name", "ext_config_value", "init_decl", "action_stmt_list",
		"action_stmt", "sql_stmt", "variable", "block_var", "ext_call_name",
		"callee_name", "call_receivers", "call_stmt", "call_body", "fn_arg_list",
		"fn_arg",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 362, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 5, 0, 92, 8, 0, 10, 0, 12, 0,
		95, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 100, 8, 0, 10, 0, 12, 0, 103, 9, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 116,
		8, 2, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3,
		127, 8, 3, 10, 3, 12, 3, 130, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 143, 8, 5, 5, 5, 145, 8, 5, 10, 5,
		12, 5, 148, 9, 5, 1, 5, 3, 5, 151, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		5, 6, 158, 8, 6, 10, 6, 12, 6, 161, 9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 166,
		8, 7, 10, 7, 12, 7, 169, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		201, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 3, 13, 215, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 229, 8, 14,
		10, 14, 12, 14, 232, 9, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 5, 18, 243, 8, 18, 10, 18, 12, 18, 246, 9, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 3,
		20, 259, 8, 20, 1, 20, 1, 20, 5, 20, 263, 8, 20, 10, 20, 12, 20, 266, 9,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28, 285, 8, 28, 10,
		28, 12, 28, 288, 9, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 4, 32, 302, 8, 32, 11, 32, 12, 32, 303,
		1, 33, 1, 33, 3, 33, 308, 8, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 3, 37, 320, 8, 37, 1, 38, 1, 38, 3, 38,
		324, 8, 38, 1, 39, 1, 39, 1, 39, 5, 39, 329, 8, 39, 10, 39, 12, 39, 332,
		9, 39, 1, 40, 1, 40, 1, 40, 3, 40, 337, 8, 40, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 3, 42, 348, 8, 42, 1, 42, 1, 42,
		5, 42, 352, 8, 42, 10, 42, 12, 42, 355, 9, 42, 1, 43, 1, 43, 1, 43, 3,
		43, 360, 8, 43, 1, 43, 0, 0, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 0, 8, 1, 0,
		23, 24, 2, 0, 55, 55, 57, 57, 2, 0, 30, 30, 32, 33, 1, 0, 38, 39, 1, 0,
		41, 45, 1, 0, 34, 35, 1, 0, 36, 37, 1, 0, 19, 20, 354, 0, 88, 1, 0, 0,
		0, 2, 106, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 123, 1, 0, 0, 0, 8, 131,
		1, 0, 0, 0, 10, 135, 1, 0, 0, 0, 12, 154, 1, 0, 0, 0, 14, 162, 1, 0, 0,
		0, 16, 170, 1, 0, 0, 0, 18, 200, 1, 0, 0, 0, 20, 202, 1, 0, 0, 0, 22, 204,
		1, 0, 0, 0, 24, 206, 1, 0, 0, 0, 26, 212, 1, 0, 0, 0, 28, 218, 1, 0, 0,
		0, 30, 233, 1, 0, 0, 0, 32, 235, 1, 0, 0, 0, 34, 237, 1, 0, 0, 0, 36, 244,
		1, 0, 0, 0, 38, 247, 1, 0, 0, 0, 40, 258, 1, 0, 0, 0, 42, 267, 1, 0, 0,
		0, 44, 269, 1, 0, 0, 0, 46, 271, 1, 0, 0, 0, 48, 273, 1, 0, 0, 0, 50, 275,
		1, 0, 0, 0, 52, 277, 1, 0, 0, 0, 54, 279, 1, 0, 0, 0, 56, 281, 1, 0, 0,
		0, 58, 289, 1, 0, 0, 0, 60, 291, 1, 0, 0, 0, 62, 293, 1, 0, 0, 0, 64, 301,
		1, 0, 0, 0, 66, 307, 1, 0, 0, 0, 68, 309, 1, 0, 0, 0, 70, 312, 1, 0, 0,
		0, 72, 314, 1, 0, 0, 0, 74, 316, 1, 0, 0, 0, 76, 323, 1, 0, 0, 0, 78, 325,
		1, 0, 0, 0, 80, 336, 1, 0, 0, 0, 82, 341, 1, 0, 0, 0, 84, 347, 1, 0, 0,
		0, 86, 359, 1, 0, 0, 0, 88, 89, 3, 2, 1, 0, 89, 93, 5, 2, 0, 0, 90, 92,
		3, 4, 2, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0,
		93, 94, 1, 0, 0, 0, 94, 101, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 100, 3,
		10, 5, 0, 97, 100, 3, 38, 19, 0, 98, 100, 3, 62, 31, 0, 99, 96, 1, 0, 0,
		0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0,
		0, 0, 104, 105, 5, 0, 0, 1, 105, 1, 1, 0, 0, 0, 106, 107, 5, 13, 0, 0,
		107, 108, 3, 44, 22, 0, 108, 3, 1, 0, 0, 0, 109, 110, 5, 14, 0, 0, 110,
		115, 3, 46, 23, 0, 111, 112, 5, 4, 0, 0, 112, 113, 3, 6, 3, 0, 113, 114,
		5, 6, 0, 0, 114, 116, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115, 116, 1, 0,
		0, 0, 116, 119, 1, 0, 0, 0, 117, 118, 5, 15, 0, 0, 118, 120, 3, 46, 23,
		0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		122, 5, 2, 0, 0, 122, 5, 1, 0, 0, 0, 123, 128, 3, 8, 4, 0, 124, 125, 5,
		7, 0, 0, 125, 127, 3, 8, 4, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1, 0, 0,
		0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 7, 1, 0, 0, 0, 130,
		128, 1, 0, 0, 0, 131, 132, 3, 48, 24, 0, 132, 133, 5, 1, 0, 0, 133, 134,
		3, 60, 30, 0, 134, 9, 1, 0, 0, 0, 135, 136, 5, 16, 0, 0, 136, 137, 3, 50,
		25, 0, 137, 138, 5, 4, 0, 0, 138, 146, 3, 14, 7, 0, 139, 142, 5, 7, 0,
		0, 140, 143, 3, 24, 12, 0, 141, 143, 3, 28, 14, 0, 142, 140, 1, 0, 0, 0,
		142, 141, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144, 139, 1, 0, 0, 0, 145,
		148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 150,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 151, 5, 7, 0, 0, 150, 149, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 5, 6, 0, 0,
		153, 11, 1, 0, 0, 0, 154, 155, 3, 54, 27, 0, 155, 159, 3, 16, 8, 0, 156,
		158, 3, 18, 9, 0, 157, 156, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 13, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 162, 167, 3, 12, 6, 0, 163, 164, 5, 7, 0, 0, 164, 166, 3, 12, 6,
		0, 165, 163, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 15, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 7,
		0, 0, 0, 171, 17, 1, 0, 0, 0, 172, 201, 5, 30, 0, 0, 173, 201, 5, 29, 0,
		0, 174, 201, 5, 32, 0, 0, 175, 176, 5, 31, 0, 0, 176, 177, 5, 3, 0, 0,
		177, 178, 3, 20, 10, 0, 178, 179, 5, 5, 0, 0, 179, 201, 1, 0, 0, 0, 180,
		181, 5, 25, 0, 0, 181, 182, 5, 3, 0, 0, 182, 183, 3, 22, 11, 0, 183, 184,
		5, 5, 0, 0, 184, 201, 1, 0, 0, 0, 185, 186, 5, 26, 0, 0, 186, 187, 5, 3,
		0, 0, 187, 188, 3, 22, 11, 0, 188, 189, 5, 5, 0, 0, 189, 201, 1, 0, 0,
		0, 190, 191, 5, 27, 0, 0, 191, 192, 5, 3, 0, 0, 192, 193, 3, 22, 11, 0,
		193, 194, 5, 5, 0, 0, 194, 201, 1, 0, 0, 0, 195, 196, 5, 28, 0, 0, 196,
		197, 5, 3, 0, 0, 197, 198, 3, 22, 11, 0, 198, 199, 5, 5, 0, 0, 199, 201,
		1, 0, 0, 0, 200, 172, 1, 0, 0, 0, 200, 173, 1, 0, 0, 0, 200, 174, 1, 0,
		0, 0, 200, 175, 1, 0, 0, 0, 200, 180, 1, 0, 0, 0, 200, 185, 1, 0, 0, 0,
		200, 190, 1, 0, 0, 0, 200, 195, 1, 0, 0, 0, 201, 19, 1, 0, 0, 0, 202, 203,
		7, 1, 0, 0, 203, 21, 1, 0, 0, 0, 204, 205, 5, 55, 0, 0, 205, 23, 1, 0,
		0, 0, 206, 207, 3, 58, 29, 0, 207, 208, 7, 2, 0, 0, 208, 209, 5, 3, 0,
		0, 209, 210, 3, 56, 28, 0, 210, 211, 5, 5, 0, 0, 211, 25, 1, 0, 0, 0, 212,
		214, 7, 3, 0, 0, 213, 215, 5, 40, 0, 0, 214, 213, 1, 0, 0, 0, 214, 215,
		1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 7, 4, 0, 0, 217, 27, 1, 0,
		0, 0, 218, 219, 7, 5, 0, 0, 219, 220, 5, 3, 0, 0, 220, 221, 3, 56, 28,
		0, 221, 222, 5, 5, 0, 0, 222, 223, 7, 6, 0, 0, 223, 224, 3, 50, 25, 0,
		224, 225, 5, 3, 0, 0, 225, 226, 3, 56, 28, 0, 226, 230, 5, 5, 0, 0, 227,
		229, 3, 26, 13, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228,
		1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 29, 1, 0, 0, 0, 232, 230, 1, 0,
		0, 0, 233, 234, 7, 7, 0, 0, 234, 31, 1, 0, 0, 0, 235, 236, 5, 21, 0, 0,
		236, 33, 1, 0, 0, 0, 237, 238, 5, 22, 0, 0, 238, 35, 1, 0, 0, 0, 239, 243,
		3, 30, 15, 0, 240, 243, 3, 32, 16, 0, 241, 243, 3, 34, 17, 0, 242, 239,
		1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 246, 1, 0,
		0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 37, 1, 0, 0, 0,
		246, 244, 1, 0, 0, 0, 247, 248, 5, 17, 0, 0, 248, 249, 3, 52, 26, 0, 249,
		250, 5, 3, 0, 0, 250, 251, 3, 40, 20, 0, 251, 252, 5, 5, 0, 0, 252, 253,
		3, 36, 18, 0, 253, 254, 5, 4, 0, 0, 254, 255, 3, 64, 32, 0, 255, 256, 5,
		6, 0, 0, 256, 39, 1, 0, 0, 0, 257, 259, 3, 42, 21, 0, 258, 257, 1, 0, 0,
		0, 258, 259, 1, 0, 0, 0, 259, 264, 1, 0, 0, 0, 260, 261, 5, 7, 0, 0, 261,
		263, 3, 42, 21, 0, 262, 260, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262,
		1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 41, 1, 0, 0, 0, 266, 264, 1, 0,
		0, 0, 267, 268, 5, 53, 0, 0, 268, 43, 1, 0, 0, 0, 269, 270, 5, 51, 0, 0,
		270, 45, 1, 0, 0, 0, 271, 272, 5, 51, 0, 0, 272, 47, 1, 0, 0, 0, 273, 274,
		5, 51, 0, 0, 274, 49, 1, 0, 0, 0, 275, 276, 5, 51, 0, 0, 276, 51, 1, 0,
		0, 0, 277, 278, 5, 51, 0, 0, 278, 53, 1, 0, 0, 0, 279, 280, 5, 51, 0, 0,
		280, 55, 1, 0, 0, 0, 281, 286, 3, 54, 27, 0, 282, 283, 5, 7, 0, 0, 283,
		285, 3, 54, 27, 0, 284, 282, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 57, 1, 0, 0, 0, 288, 286, 1, 0,
		0, 0, 289, 290, 5, 52, 0, 0, 290, 59, 1, 0, 0, 0, 291, 292, 3, 20, 10,
		0, 292, 61, 1, 0, 0, 0, 293, 294, 5, 18, 0, 0, 294, 295, 5, 3, 0, 0, 295,
		296, 5, 5, 0, 0, 296, 297, 5, 4, 0, 0, 297, 298, 3, 64, 32, 0, 298, 299,
		5, 6, 0, 0, 299, 63, 1, 0, 0, 0, 300, 302, 3, 66, 33, 0, 301, 300, 1, 0,
		0, 0, 302, 303, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0,
		304, 65, 1, 0, 0, 0, 305, 308, 3, 68, 34, 0, 306, 308, 3, 80, 40, 0, 307,
		305, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 67, 1, 0, 0, 0, 309, 310, 5,
		59, 0, 0, 310, 311, 5, 2, 0, 0, 311, 69, 1, 0, 0, 0, 312, 313, 5, 53, 0,
		0, 313, 71, 1, 0, 0, 0, 314, 315, 5, 54, 0, 0, 315, 73, 1, 0, 0, 0, 316,
		319, 5, 51, 0, 0, 317, 318, 5, 11, 0, 0, 318, 320, 5, 51, 0, 0, 319, 317,
		1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 75, 1, 0, 0, 0, 321, 324, 3, 74,
		37, 0, 322, 324, 3, 52, 26, 0, 323, 321, 1, 0, 0, 0, 323, 322, 1, 0, 0,
		0, 324, 77, 1, 0, 0, 0, 325, 330, 3, 70, 35, 0, 326, 327, 5, 7, 0, 0, 327,
		329, 3, 70, 35, 0, 328, 326, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328,
		1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 79, 1, 0, 0, 0, 332, 330, 1, 0,
		0, 0, 333, 334, 3, 78, 39, 0, 334, 335, 5, 12, 0, 0, 335, 337, 1, 0, 0,
		0, 336, 333, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338,
		339, 3, 82, 41, 0, 339, 340, 5, 2, 0, 0, 340, 81, 1, 0, 0, 0, 341, 342,
		3, 76, 38, 0, 342, 343, 5, 3, 0, 0, 343, 344, 3, 84, 42, 0, 344, 345, 5,
		5, 0, 0, 345, 83, 1, 0, 0, 0, 346, 348, 3, 86, 43, 0, 347, 346, 1, 0, 0,
		0, 347, 348, 1, 0, 0, 0, 348, 353, 1, 0, 0, 0, 349, 350, 5, 7, 0, 0, 350,
		352, 3, 86, 43, 0, 351, 349, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351,
		1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 85, 1, 0, 0, 0, 355, 353, 1, 0,
		0, 0, 356, 360, 3, 20, 10, 0, 357, 360, 3, 70, 35, 0, 358, 360, 3, 72,
		36, 0, 359, 356, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0,
		360, 87, 1, 0, 0, 0, 28, 93, 99, 101, 115, 119, 128, 142, 146, 150, 159,
		167, 200, 214, 230, 242, 244, 258, 264, 286, 303, 307, 319, 323, 330, 336,
		347, 353, 359,
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
	KuneiformParserEOF                     = antlr.TokenEOF
	KuneiformParserCOL                     = 1
	KuneiformParserSCOL                    = 2
	KuneiformParserL_PAREN                 = 3
	KuneiformParserL_BRACE                 = 4
	KuneiformParserR_PAREN                 = 5
	KuneiformParserR_BRACE                 = 6
	KuneiformParserCOMMA                   = 7
	KuneiformParserDOLLAR                  = 8
	KuneiformParserHASH                    = 9
	KuneiformParserAT                      = 10
	KuneiformParserPERIOD                  = 11
	KuneiformParserEQ                      = 12
	KuneiformParserDATABASE_               = 13
	KuneiformParserUSE_                    = 14
	KuneiformParserAS_                     = 15
	KuneiformParserTABLE_                  = 16
	KuneiformParserACTION_                 = 17
	KuneiformParserINIT_                   = 18
	KuneiformParserPUBLIC_                 = 19
	KuneiformParserPRIVATE_                = 20
	KuneiformParserVIEW_                   = 21
	KuneiformParserMUSTSIGN_               = 22
	KuneiformParserINT_                    = 23
	KuneiformParserTEXT_                   = 24
	KuneiformParserMIN_                    = 25
	KuneiformParserMAX_                    = 26
	KuneiformParserMIN_LEN_                = 27
	KuneiformParserMAX_LEN_                = 28
	KuneiformParserNOT_NULL_               = 29
	KuneiformParserPRIMARY_                = 30
	KuneiformParserDEFAULT_                = 31
	KuneiformParserUNIQUE_                 = 32
	KuneiformParserINDEX_                  = 33
	KuneiformParserFOREIGN_KEY_            = 34
	KuneiformParserFOREIGN_KEY_ABBR_       = 35
	KuneiformParserREFERENCES_             = 36
	KuneiformParserREFERENCES_ABBR_        = 37
	KuneiformParserACTION_ON_UPDATE_       = 38
	KuneiformParserACTION_ON_DELETE_       = 39
	KuneiformParserACTION_DO_              = 40
	KuneiformParserACTION_DO_NO_ACTION_    = 41
	KuneiformParserACTION_DO_CASCADE_      = 42
	KuneiformParserACTION_DO_SET_NULL_     = 43
	KuneiformParserACTION_DO_SET_DEFAULT_  = 44
	KuneiformParserACTION_DO_RESTRICT_     = 45
	KuneiformParserSELECT_                 = 46
	KuneiformParserINSERT_                 = 47
	KuneiformParserUPDATE_                 = 48
	KuneiformParserDELETE_                 = 49
	KuneiformParserWITH_                   = 50
	KuneiformParserIDENTIFIER              = 51
	KuneiformParserINDEX_NAME              = 52
	KuneiformParserPARAM_OR_VAR            = 53
	KuneiformParserBLOCK_VAR               = 54
	KuneiformParserUNSIGNED_NUMBER_LITERAL = 55
	KuneiformParserSIGNED_NUMBER_LITERAL   = 56
	KuneiformParserSTRING_LITERAL          = 57
	KuneiformParserSQL_KEYWORDS            = 58
	KuneiformParserSQL_STMT                = 59
	KuneiformParserWS                      = 60
	KuneiformParserTERMINATOR              = 61
	KuneiformParserBLOCK_COMMENT           = 62
	KuneiformParserLINE_COMMENT            = 63
)

// KuneiformParser rules.
const (
	KuneiformParserRULE_source_unit         = 0
	KuneiformParserRULE_database_directive  = 1
	KuneiformParserRULE_extension_directive = 2
	KuneiformParserRULE_ext_config_list     = 3
	KuneiformParserRULE_ext_config          = 4
	KuneiformParserRULE_table_decl          = 5
	KuneiformParserRULE_column_def          = 6
	KuneiformParserRULE_column_def_list     = 7
	KuneiformParserRULE_column_type         = 8
	KuneiformParserRULE_column_constraint   = 9
	KuneiformParserRULE_literal_value       = 10
	KuneiformParserRULE_number_value        = 11
	KuneiformParserRULE_index_def           = 12
	KuneiformParserRULE_foreign_key_action  = 13
	KuneiformParserRULE_foreign_key_def     = 14
	KuneiformParserRULE_action_visibility   = 15
	KuneiformParserRULE_action_mutability   = 16
	KuneiformParserRULE_action_auxiliary    = 17
	KuneiformParserRULE_action_attr_list    = 18
	KuneiformParserRULE_action_decl         = 19
	KuneiformParserRULE_param_list          = 20
	KuneiformParserRULE_parameter           = 21
	KuneiformParserRULE_database_name       = 22
	KuneiformParserRULE_extension_name      = 23
	KuneiformParserRULE_ext_config_name     = 24
	KuneiformParserRULE_table_name          = 25
	KuneiformParserRULE_action_name         = 26
	KuneiformParserRULE_column_name         = 27
	KuneiformParserRULE_column_name_list    = 28
	KuneiformParserRULE_index_name          = 29
	KuneiformParserRULE_ext_config_value    = 30
	KuneiformParserRULE_init_decl           = 31
	KuneiformParserRULE_action_stmt_list    = 32
	KuneiformParserRULE_action_stmt         = 33
	KuneiformParserRULE_sql_stmt            = 34
	KuneiformParserRULE_variable            = 35
	KuneiformParserRULE_block_var           = 36
	KuneiformParserRULE_ext_call_name       = 37
	KuneiformParserRULE_callee_name         = 38
	KuneiformParserRULE_call_receivers      = 39
	KuneiformParserRULE_call_stmt           = 40
	KuneiformParserRULE_call_body           = 41
	KuneiformParserRULE_fn_arg_list         = 42
	KuneiformParserRULE_fn_arg              = 43
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
		p.SetState(88)
		p.Database_directive()
	}
	{
		p.SetState(89)
		p.Match(KuneiformParserSCOL)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserUSE_ {
		{
			p.SetState(90)
			p.Extension_directive()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&458752) != 0 {
		p.SetState(99)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserTABLE_:
			{
				p.SetState(96)
				p.Table_decl()
			}

		case KuneiformParserACTION_:
			{
				p.SetState(97)
				p.Action_decl()
			}

		case KuneiformParserINIT_:
			{
				p.SetState(98)
				p.Init_decl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
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
		p.SetState(106)
		p.Match(KuneiformParserDATABASE_)
	}
	{
		p.SetState(107)
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
		p.SetState(109)
		p.Match(KuneiformParserUSE_)
	}
	{
		p.SetState(110)
		p.Extension_name()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserL_BRACE {
		{
			p.SetState(111)
			p.Match(KuneiformParserL_BRACE)
		}
		{
			p.SetState(112)
			p.Ext_config_list()
		}
		{
			p.SetState(113)
			p.Match(KuneiformParserR_BRACE)
		}

	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserAS_ {
		{
			p.SetState(117)
			p.Match(KuneiformParserAS_)
		}
		{
			p.SetState(118)
			p.Extension_name()
		}

	}
	{
		p.SetState(121)
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
		p.SetState(123)
		p.Ext_config()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(124)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(125)
			p.Ext_config()
		}

		p.SetState(130)
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
		p.SetState(131)
		p.Ext_config_name()
	}
	{
		p.SetState(132)
		p.Match(KuneiformParserCOL)
	}
	{
		p.SetState(133)
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
		p.SetState(135)
		p.Match(KuneiformParserTABLE_)
	}
	{
		p.SetState(136)
		p.Table_name()
	}
	{
		p.SetState(137)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(138)
		p.Column_def_list()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				p.Match(KuneiformParserCOMMA)
			}
			p.SetState(142)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case KuneiformParserINDEX_NAME:
				{
					p.SetState(140)
					p.Index_def()
				}

			case KuneiformParserFOREIGN_KEY_, KuneiformParserFOREIGN_KEY_ABBR_:
				{
					p.SetState(141)
					p.Foreign_key_def()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserCOMMA {
		{
			p.SetState(149)
			p.Match(KuneiformParserCOMMA)
		}

	}
	{
		p.SetState(152)
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
		p.SetState(154)
		p.Column_name()
	}
	{
		p.SetState(155)
		p.Column_type()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8556380160) != 0 {
		{
			p.SetState(156)
			p.Column_constraint()
		}

		p.SetState(161)
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
		p.SetState(162)
		p.Column_def()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(163)
				p.Match(KuneiformParserCOMMA)
			}
			{
				p.SetState(164)
				p.Column_def()
			}

		}
		p.SetState(169)
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
		p.SetState(170)
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

	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserPRIMARY_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(KuneiformParserPRIMARY_)
		}

	case KuneiformParserNOT_NULL_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Match(KuneiformParserNOT_NULL_)
		}

	case KuneiformParserUNIQUE_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Match(KuneiformParserUNIQUE_)
		}

	case KuneiformParserDEFAULT_:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(175)
			p.Match(KuneiformParserDEFAULT_)
		}
		{
			p.SetState(176)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(177)
			p.Literal_value()
		}
		{
			p.SetState(178)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(180)
			p.Match(KuneiformParserMIN_)
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

	case KuneiformParserMAX_:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)
			p.Match(KuneiformParserMAX_)
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

	case KuneiformParserMIN_LEN_:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(190)
			p.Match(KuneiformParserMIN_LEN_)
		}
		{
			p.SetState(191)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(192)
			p.Number_value()
		}
		{
			p.SetState(193)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_LEN_:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(195)
			p.Match(KuneiformParserMAX_LEN_)
		}
		{
			p.SetState(196)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(197)
			p.Number_value()
		}
		{
			p.SetState(198)
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
		p.SetState(202)
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
		p.SetState(204)
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
		p.SetState(206)
		p.Index_name()
	}
	{
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13958643712) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(208)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(209)
		p.Column_name_list()
	}
	{
		p.SetState(210)
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
		p.SetState(212)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserACTION_DO_ {
		{
			p.SetState(213)
			p.Match(KuneiformParserACTION_DO_)
		}

	}
	{
		p.SetState(216)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68169720922112) != 0) {
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
		p.SetState(218)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserFOREIGN_KEY_ || _la == KuneiformParserFOREIGN_KEY_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(219)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(220)
		p.Column_name_list()
	}
	{
		p.SetState(221)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(222)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserREFERENCES_ || _la == KuneiformParserREFERENCES_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(223)
		p.Table_name()
	}
	{
		p.SetState(224)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(225)
		p.Column_name_list()
	}
	{
		p.SetState(226)
		p.Match(KuneiformParserR_PAREN)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_ {
		{
			p.SetState(227)
			p.Foreign_key_action()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_visibilityContext is an interface to support dynamic dispatch.
type IAction_visibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUBLIC_() antlr.TerminalNode
	PRIVATE_() antlr.TerminalNode

	// IsAction_visibilityContext differentiates from other interfaces.
	IsAction_visibilityContext()
}

type Action_visibilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_visibilityContext() *Action_visibilityContext {
	var p = new(Action_visibilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_visibility
	return p
}

func (*Action_visibilityContext) IsAction_visibilityContext() {}

func NewAction_visibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_visibilityContext {
	var p = new(Action_visibilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_visibility

	return p
}

func (s *Action_visibilityContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_visibilityContext) PUBLIC_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPUBLIC_, 0)
}

func (s *Action_visibilityContext) PRIVATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIVATE_, 0)
}

func (s *Action_visibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_visibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_visibilityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_visibility(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_visibility() (localctx IAction_visibilityContext) {
	this := p
	_ = this

	localctx = NewAction_visibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KuneiformParserRULE_action_visibility)
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
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserPUBLIC_ || _la == KuneiformParserPRIVATE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAction_mutabilityContext is an interface to support dynamic dispatch.
type IAction_mutabilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VIEW_() antlr.TerminalNode

	// IsAction_mutabilityContext differentiates from other interfaces.
	IsAction_mutabilityContext()
}

type Action_mutabilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_mutabilityContext() *Action_mutabilityContext {
	var p = new(Action_mutabilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_mutability
	return p
}

func (*Action_mutabilityContext) IsAction_mutabilityContext() {}

func NewAction_mutabilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_mutabilityContext {
	var p = new(Action_mutabilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_mutability

	return p
}

func (s *Action_mutabilityContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_mutabilityContext) VIEW_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserVIEW_, 0)
}

func (s *Action_mutabilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_mutabilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_mutabilityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_mutability(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_mutability() (localctx IAction_mutabilityContext) {
	this := p
	_ = this

	localctx = NewAction_mutabilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KuneiformParserRULE_action_mutability)

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
		p.SetState(235)
		p.Match(KuneiformParserVIEW_)
	}

	return localctx
}

// IAction_auxiliaryContext is an interface to support dynamic dispatch.
type IAction_auxiliaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUSTSIGN_() antlr.TerminalNode

	// IsAction_auxiliaryContext differentiates from other interfaces.
	IsAction_auxiliaryContext()
}

type Action_auxiliaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_auxiliaryContext() *Action_auxiliaryContext {
	var p = new(Action_auxiliaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_auxiliary
	return p
}

func (*Action_auxiliaryContext) IsAction_auxiliaryContext() {}

func NewAction_auxiliaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_auxiliaryContext {
	var p = new(Action_auxiliaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_auxiliary

	return p
}

func (s *Action_auxiliaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_auxiliaryContext) MUSTSIGN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMUSTSIGN_, 0)
}

func (s *Action_auxiliaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_auxiliaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_auxiliaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_auxiliary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_auxiliary() (localctx IAction_auxiliaryContext) {
	this := p
	_ = this

	localctx = NewAction_auxiliaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KuneiformParserRULE_action_auxiliary)

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
		p.SetState(237)
		p.Match(KuneiformParserMUSTSIGN_)
	}

	return localctx
}

// IAction_attr_listContext is an interface to support dynamic dispatch.
type IAction_attr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAction_visibility() []IAction_visibilityContext
	Action_visibility(i int) IAction_visibilityContext
	AllAction_mutability() []IAction_mutabilityContext
	Action_mutability(i int) IAction_mutabilityContext
	AllAction_auxiliary() []IAction_auxiliaryContext
	Action_auxiliary(i int) IAction_auxiliaryContext

	// IsAction_attr_listContext differentiates from other interfaces.
	IsAction_attr_listContext()
}

type Action_attr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_attr_listContext() *Action_attr_listContext {
	var p = new(Action_attr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_attr_list
	return p
}

func (*Action_attr_listContext) IsAction_attr_listContext() {}

func NewAction_attr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_attr_listContext {
	var p = new(Action_attr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_attr_list

	return p
}

func (s *Action_attr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_attr_listContext) AllAction_visibility() []IAction_visibilityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_visibilityContext); ok {
			len++
		}
	}

	tst := make([]IAction_visibilityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_visibilityContext); ok {
			tst[i] = t.(IAction_visibilityContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_visibility(i int) IAction_visibilityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_visibilityContext); ok {
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

	return t.(IAction_visibilityContext)
}

func (s *Action_attr_listContext) AllAction_mutability() []IAction_mutabilityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_mutabilityContext); ok {
			len++
		}
	}

	tst := make([]IAction_mutabilityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_mutabilityContext); ok {
			tst[i] = t.(IAction_mutabilityContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_mutability(i int) IAction_mutabilityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_mutabilityContext); ok {
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

	return t.(IAction_mutabilityContext)
}

func (s *Action_attr_listContext) AllAction_auxiliary() []IAction_auxiliaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_auxiliaryContext); ok {
			len++
		}
	}

	tst := make([]IAction_auxiliaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_auxiliaryContext); ok {
			tst[i] = t.(IAction_auxiliaryContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_auxiliary(i int) IAction_auxiliaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_auxiliaryContext); ok {
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

	return t.(IAction_auxiliaryContext)
}

func (s *Action_attr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_attr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_attr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_attr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_attr_list() (localctx IAction_attr_listContext) {
	this := p
	_ = this

	localctx = NewAction_attr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KuneiformParserRULE_action_attr_list)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0 {
		p.SetState(242)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserPUBLIC_, KuneiformParserPRIVATE_:
			{
				p.SetState(239)
				p.Action_visibility()
			}

		case KuneiformParserVIEW_:
			{
				p.SetState(240)
				p.Action_mutability()
			}

		case KuneiformParserMUSTSIGN_:
			{
				p.SetState(241)
				p.Action_auxiliary()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(246)
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
	Action_attr_list() IAction_attr_listContext
	L_BRACE() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	R_BRACE() antlr.TerminalNode

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

func (s *Action_declContext) Action_attr_list() IAction_attr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_attr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_attr_listContext)
}

func (s *Action_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
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

func (s *Action_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
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
	p.EnterRule(localctx, 38, KuneiformParserRULE_action_decl)

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
		p.SetState(247)
		p.Match(KuneiformParserACTION_)
	}
	{
		p.SetState(248)
		p.Action_name()
	}
	{
		p.SetState(249)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(250)
		p.Param_list()
	}
	{
		p.SetState(251)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(252)
		p.Action_attr_list()
	}
	{
		p.SetState(253)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(254)
		p.Action_stmt_list()
	}
	{
		p.SetState(255)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
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

func (s *Param_listContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *Param_listContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
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
	p.EnterRule(localctx, 40, KuneiformParserRULE_param_list)
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
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPARAM_OR_VAR {
		{
			p.SetState(257)
			p.Parameter()
		}

	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(260)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(261)
			p.Parameter()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_OR_VAR() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) PARAM_OR_VAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPARAM_OR_VAR, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KuneiformParserRULE_parameter)

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
		p.SetState(267)
		p.Match(KuneiformParserPARAM_OR_VAR)
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
	p.EnterRule(localctx, 44, KuneiformParserRULE_database_name)

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
		p.SetState(269)
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
	p.EnterRule(localctx, 46, KuneiformParserRULE_extension_name)

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
		p.SetState(271)
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
	p.EnterRule(localctx, 48, KuneiformParserRULE_ext_config_name)

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
		p.SetState(273)
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
	p.EnterRule(localctx, 50, KuneiformParserRULE_table_name)

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
		p.SetState(275)
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
	p.EnterRule(localctx, 52, KuneiformParserRULE_action_name)

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
		p.SetState(277)
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
	p.EnterRule(localctx, 54, KuneiformParserRULE_column_name)

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
	p.EnterRule(localctx, 56, KuneiformParserRULE_column_name_list)
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
		p.SetState(281)
		p.Column_name()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(282)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(283)
			p.Column_name()
		}

		p.SetState(288)
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
	p.EnterRule(localctx, 58, KuneiformParserRULE_index_name)

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
		p.SetState(289)
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
	p.EnterRule(localctx, 60, KuneiformParserRULE_ext_config_value)

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
		p.SetState(291)
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
	INIT_() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	R_BRACE() antlr.TerminalNode

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

func (s *Init_declContext) INIT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINIT_, 0)
}

func (s *Init_declContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Init_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Init_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
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

func (s *Init_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
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
	p.EnterRule(localctx, 62, KuneiformParserRULE_init_decl)

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
		p.Match(KuneiformParserINIT_)
	}
	{
		p.SetState(294)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(295)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(296)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(297)
		p.Action_stmt_list()
	}
	{
		p.SetState(298)
		p.Match(KuneiformParserR_BRACE)
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
	p.EnterRule(localctx, 64, KuneiformParserRULE_action_stmt_list)
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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&587719751371849728) != 0) {
		{
			p.SetState(300)
			p.Action_stmt()
		}

		p.SetState(303)
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
	Sql_stmt() ISql_stmtContext
	Call_stmt() ICall_stmtContext

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

func (s *Action_stmtContext) Sql_stmt() ISql_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_stmtContext)
}

func (s *Action_stmtContext) Call_stmt() ICall_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_stmtContext)
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
	p.EnterRule(localctx, 66, KuneiformParserRULE_action_stmt)

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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserSQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Sql_stmt()
		}

	case KuneiformParserIDENTIFIER, KuneiformParserPARAM_OR_VAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.Call_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SQL_STMT() antlr.TerminalNode
	SCOL() antlr.TerminalNode

	// IsSql_stmtContext differentiates from other interfaces.
	IsSql_stmtContext()
}

type Sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_sql_stmt
	return p
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_sql_stmt

	return p
}

func (s *Sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmtContext) SQL_STMT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_STMT, 0)
}

func (s *Sql_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Sql_stmt() (localctx ISql_stmtContext) {
	this := p
	_ = this

	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KuneiformParserRULE_sql_stmt)

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
		p.Match(KuneiformParserSQL_STMT)
	}
	{
		p.SetState(310)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_OR_VAR() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) PARAM_OR_VAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPARAM_OR_VAR, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KuneiformParserRULE_variable)

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
		p.SetState(312)
		p.Match(KuneiformParserPARAM_OR_VAR)
	}

	return localctx
}

// IBlock_varContext is an interface to support dynamic dispatch.
type IBlock_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_VAR() antlr.TerminalNode

	// IsBlock_varContext differentiates from other interfaces.
	IsBlock_varContext()
}

type Block_varContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_varContext() *Block_varContext {
	var p = new(Block_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_block_var
	return p
}

func (*Block_varContext) IsBlock_varContext() {}

func NewBlock_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_varContext {
	var p = new(Block_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_block_var

	return p
}

func (s *Block_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_varContext) BLOCK_VAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserBLOCK_VAR, 0)
}

func (s *Block_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitBlock_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Block_var() (localctx IBlock_varContext) {
	this := p
	_ = this

	localctx = NewBlock_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KuneiformParserRULE_block_var)

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
		p.SetState(314)
		p.Match(KuneiformParserBLOCK_VAR)
	}

	return localctx
}

// IExt_call_nameContext is an interface to support dynamic dispatch.
type IExt_call_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsExt_call_nameContext differentiates from other interfaces.
	IsExt_call_nameContext()
}

type Ext_call_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_call_nameContext() *Ext_call_nameContext {
	var p = new(Ext_call_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_call_name
	return p
}

func (*Ext_call_nameContext) IsExt_call_nameContext() {}

func NewExt_call_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_call_nameContext {
	var p = new(Ext_call_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_call_name

	return p
}

func (s *Ext_call_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_call_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserIDENTIFIER)
}

func (s *Ext_call_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, i)
}

func (s *Ext_call_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPERIOD, 0)
}

func (s *Ext_call_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_call_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_call_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_call_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_call_name() (localctx IExt_call_nameContext) {
	this := p
	_ = this

	localctx = NewExt_call_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KuneiformParserRULE_ext_call_name)
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
		p.SetState(316)
		p.Match(KuneiformParserIDENTIFIER)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPERIOD {
		{
			p.SetState(317)
			p.Match(KuneiformParserPERIOD)
		}
		{
			p.SetState(318)
			p.Match(KuneiformParserIDENTIFIER)
		}

	}

	return localctx
}

// ICallee_nameContext is an interface to support dynamic dispatch.
type ICallee_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ext_call_name() IExt_call_nameContext
	Action_name() IAction_nameContext

	// IsCallee_nameContext differentiates from other interfaces.
	IsCallee_nameContext()
}

type Callee_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallee_nameContext() *Callee_nameContext {
	var p = new(Callee_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_callee_name
	return p
}

func (*Callee_nameContext) IsCallee_nameContext() {}

func NewCallee_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Callee_nameContext {
	var p = new(Callee_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_callee_name

	return p
}

func (s *Callee_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Callee_nameContext) Ext_call_name() IExt_call_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_call_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_call_nameContext)
}

func (s *Callee_nameContext) Action_name() IAction_nameContext {
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

func (s *Callee_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Callee_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Callee_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCallee_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Callee_name() (localctx ICallee_nameContext) {
	this := p
	_ = this

	localctx = NewCallee_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KuneiformParserRULE_callee_name)

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

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Ext_call_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Action_name()
		}

	}

	return localctx
}

// ICall_receiversContext is an interface to support dynamic dispatch.
type ICall_receiversContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCall_receiversContext differentiates from other interfaces.
	IsCall_receiversContext()
}

type Call_receiversContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_receiversContext() *Call_receiversContext {
	var p = new(Call_receiversContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_receivers
	return p
}

func (*Call_receiversContext) IsCall_receiversContext() {}

func NewCall_receiversContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_receiversContext {
	var p = new(Call_receiversContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_receivers

	return p
}

func (s *Call_receiversContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_receiversContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *Call_receiversContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *Call_receiversContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Call_receiversContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Call_receiversContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_receiversContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_receiversContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_receivers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_receivers() (localctx ICall_receiversContext) {
	this := p
	_ = this

	localctx = NewCall_receiversContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, KuneiformParserRULE_call_receivers)
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
		p.SetState(325)
		p.Variable()
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(326)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(327)
			p.Variable()
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICall_stmtContext is an interface to support dynamic dispatch.
type ICall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call_body() ICall_bodyContext
	SCOL() antlr.TerminalNode
	Call_receivers() ICall_receiversContext
	EQ() antlr.TerminalNode

	// IsCall_stmtContext differentiates from other interfaces.
	IsCall_stmtContext()
}

type Call_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_stmtContext() *Call_stmtContext {
	var p = new(Call_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_stmt
	return p
}

func (*Call_stmtContext) IsCall_stmtContext() {}

func NewCall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_stmtContext {
	var p = new(Call_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_stmt

	return p
}

func (s *Call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_stmtContext) Call_body() ICall_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_bodyContext)
}

func (s *Call_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Call_stmtContext) Call_receivers() ICall_receiversContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_receiversContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_receiversContext)
}

func (s *Call_stmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(KuneiformParserEQ, 0)
}

func (s *Call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_stmt() (localctx ICall_stmtContext) {
	this := p
	_ = this

	localctx = NewCall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, KuneiformParserRULE_call_stmt)
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
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPARAM_OR_VAR {
		{
			p.SetState(333)
			p.Call_receivers()
		}
		{
			p.SetState(334)
			p.Match(KuneiformParserEQ)
		}

	}
	{
		p.SetState(338)
		p.Call_body()
	}
	{
		p.SetState(339)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// ICall_bodyContext is an interface to support dynamic dispatch.
type ICall_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Callee_name() ICallee_nameContext
	L_PAREN() antlr.TerminalNode
	Fn_arg_list() IFn_arg_listContext
	R_PAREN() antlr.TerminalNode

	// IsCall_bodyContext differentiates from other interfaces.
	IsCall_bodyContext()
}

type Call_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_bodyContext() *Call_bodyContext {
	var p = new(Call_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_body
	return p
}

func (*Call_bodyContext) IsCall_bodyContext() {}

func NewCall_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_bodyContext {
	var p = new(Call_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_body

	return p
}

func (s *Call_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_bodyContext) Callee_name() ICallee_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallee_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallee_nameContext)
}

func (s *Call_bodyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Call_bodyContext) Fn_arg_list() IFn_arg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_arg_listContext)
}

func (s *Call_bodyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Call_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_body() (localctx ICall_bodyContext) {
	this := p
	_ = this

	localctx = NewCall_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, KuneiformParserRULE_call_body)

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
		p.SetState(341)
		p.Callee_name()
	}
	{
		p.SetState(342)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(343)
		p.Fn_arg_list()
	}
	{
		p.SetState(344)
		p.Match(KuneiformParserR_PAREN)
	}

	return localctx
}

// IFn_arg_listContext is an interface to support dynamic dispatch.
type IFn_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFn_arg() []IFn_argContext
	Fn_arg(i int) IFn_argContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFn_arg_listContext differentiates from other interfaces.
	IsFn_arg_listContext()
}

type Fn_arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_arg_listContext() *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_fn_arg_list
	return p
}

func (*Fn_arg_listContext) IsFn_arg_listContext() {}

func NewFn_arg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_fn_arg_list

	return p
}

func (s *Fn_arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_arg_listContext) AllFn_arg() []IFn_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_argContext); ok {
			len++
		}
	}

	tst := make([]IFn_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_argContext); ok {
			tst[i] = t.(IFn_argContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_listContext) Fn_arg(i int) IFn_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_argContext); ok {
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

	return t.(IFn_argContext)
}

func (s *Fn_arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Fn_arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Fn_arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitFn_arg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Fn_arg_list() (localctx IFn_arg_listContext) {
	this := p
	_ = this

	localctx = NewFn_arg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, KuneiformParserRULE_fn_arg_list)
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
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&207165582859042816) != 0 {
		{
			p.SetState(346)
			p.Fn_arg()
		}

	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(349)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(350)
			p.Fn_arg()
		}

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFn_argContext is an interface to support dynamic dispatch.
type IFn_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Variable() IVariableContext
	Block_var() IBlock_varContext

	// IsFn_argContext differentiates from other interfaces.
	IsFn_argContext()
}

type Fn_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_argContext() *Fn_argContext {
	var p = new(Fn_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_fn_arg
	return p
}

func (*Fn_argContext) IsFn_argContext() {}

func NewFn_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_argContext {
	var p = new(Fn_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_fn_arg

	return p
}

func (s *Fn_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_argContext) Literal_value() ILiteral_valueContext {
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

func (s *Fn_argContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Fn_argContext) Block_var() IBlock_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_varContext)
}

func (s *Fn_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_argContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitFn_arg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Fn_arg() (localctx IFn_argContext) {
	this := p
	_ = this

	localctx = NewFn_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, KuneiformParserRULE_fn_arg)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserUNSIGNED_NUMBER_LITERAL, KuneiformParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Literal_value()
		}

	case KuneiformParserPARAM_OR_VAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Variable()
		}

	case KuneiformParserBLOCK_VAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.Block_var()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
