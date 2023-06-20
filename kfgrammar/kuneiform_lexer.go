// Code generated from KuneiformLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type KuneiformLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var kuneiformlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kuneiformlexerLexerInit() {
	staticData := &kuneiformlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE", "SQL_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'('", "'{'", "')'", "'}'", "','", "'.'", "'$'", "'@'", "'#'",
		"'database'", "'table'", "'action'", "'public'", "'private'", "'int'",
		"'text'", "'min'", "'max'", "'minlen'", "'maxlen'", "'notnull'", "'primary'",
		"'default'", "'unique'", "'index'", "'foreign_key'", "'fk'", "'references'",
		"'ref'", "'on_update'", "'on_delete'", "'do'", "'no_action'", "'cascade'",
		"'set_null'", "'set_default'", "'restrict'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA", "PERIOD",
		"DOLLAR", "AT", "HASH", "DATABASE_", "TABLE_", "ACTION_", "PUBLIC_",
		"PRIVATE_", "INT_", "TEXT_", "MIN_", "MAX_", "MIN_LEN_", "MAX_LEN_",
		"NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_", "FOREIGN_KEY_",
		"FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_", "ACTION_ON_UPDATE_",
		"ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_", "ACTION_DO_CASCADE_",
		"ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_", "ACTION_DO_RESTRICT_",
		"SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "IDENTIFIER", "INDEX_NAME",
		"ACTION_PARAMETER", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
		"UNEXPECTED_CHAR", "SQL_END_SCOL", "SQL_NL", "SQL_STMT",
	}
	staticData.ruleNames = []string{
		"SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA", "PERIOD",
		"DOLLAR", "AT", "HASH", "DATABASE_", "TABLE_", "ACTION_", "PUBLIC_",
		"PRIVATE_", "INT_", "TEXT_", "MIN_", "MAX_", "MIN_LEN_", "MAX_LEN_",
		"NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_", "FOREIGN_KEY_",
		"FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_", "ACTION_ON_UPDATE_",
		"ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_", "ACTION_DO_CASCADE_",
		"ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_", "ACTION_DO_RESTRICT_",
		"SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "IDENTIFIER", "INDEX_NAME",
		"ACTION_PARAMETER", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
		"UNEXPECTED_CHAR", "DIGIT", "DOUBLE_QUOTE_STRING_CHAR", "SINGLE_QUOTE_STRING_CHAR",
		"DOUBLE_QUOTE_STRING", "SINGLE_QUOTE_STRING", "SQL_END_SCOL", "SQL_NL",
		"SQL_STMT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 57, 512, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2,
		25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30,
		7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7,
		35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40,
		2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2,
		46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51,
		7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7,
		56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 5, 43, 398, 8, 43, 10, 43, 12,
		43, 401, 9, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 4, 46,
		410, 8, 46, 11, 46, 12, 46, 411, 1, 47, 3, 47, 415, 8, 47, 1, 47, 4, 47,
		418, 8, 47, 11, 47, 12, 47, 419, 1, 48, 1, 48, 3, 48, 424, 8, 48, 1, 49,
		4, 49, 427, 8, 49, 11, 49, 12, 49, 428, 1, 49, 1, 49, 1, 50, 4, 50, 434,
		8, 50, 11, 50, 12, 50, 435, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 5,
		51, 444, 8, 51, 10, 51, 12, 51, 447, 9, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 458, 8, 52, 10, 52, 12, 52, 461,
		9, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 3,
		55, 472, 8, 55, 1, 56, 1, 56, 1, 56, 3, 56, 477, 8, 56, 1, 57, 1, 57, 5,
		57, 481, 8, 57, 10, 57, 12, 57, 484, 9, 57, 1, 57, 1, 57, 1, 58, 1, 58,
		5, 58, 490, 8, 58, 10, 58, 12, 58, 493, 9, 58, 1, 58, 1, 58, 1, 59, 1,
		59, 1, 59, 1, 59, 1, 60, 4, 60, 502, 8, 60, 11, 60, 12, 60, 503, 1, 60,
		1, 60, 1, 61, 4, 61, 509, 8, 61, 11, 61, 12, 61, 510, 1, 445, 0, 62, 2,
		1, 4, 2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11,
		24, 12, 26, 13, 28, 14, 30, 15, 32, 16, 34, 17, 36, 18, 38, 19, 40, 20,
		42, 21, 44, 22, 46, 23, 48, 24, 50, 25, 52, 26, 54, 27, 56, 28, 58, 29,
		60, 30, 62, 31, 64, 32, 66, 33, 68, 34, 70, 35, 72, 36, 74, 37, 76, 38,
		78, 39, 80, 40, 82, 41, 84, 42, 86, 43, 88, 44, 90, 45, 92, 46, 94, 47,
		96, 48, 98, 49, 100, 50, 102, 51, 104, 52, 106, 53, 108, 54, 110, 0, 112,
		0, 114, 0, 116, 0, 118, 0, 120, 55, 122, 56, 124, 57, 2, 0, 1, 24, 2, 0,
		83, 83, 115, 115, 2, 0, 69, 69, 101, 101, 2, 0, 76, 76, 108, 108, 2, 0,
		67, 67, 99, 99, 2, 0, 84, 84, 116, 116, 2, 0, 73, 73, 105, 105, 2, 0, 78,
		78, 110, 110, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 80,
		80, 112, 112, 2, 0, 68, 68, 100, 100, 2, 0, 65, 65, 97, 97, 2, 0, 87, 87,
		119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 43, 43, 45, 45, 2, 0, 9, 9, 32, 32, 2, 0,
		10, 10, 13, 13, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4,
		0, 10, 10, 13, 13, 39, 39, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 59,
		59, 520, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1,
		0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16,
		1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0,
		24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0,
		0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0,
		0, 0, 40, 1, 0, 0, 0, 0, 42, 1, 0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 0,
		0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 1,
		0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0, 62,
		1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0, 68, 1, 0, 0, 0, 0,
		70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0, 0,
		0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0, 0, 0, 84, 1, 0, 0,
		0, 0, 86, 1, 0, 0, 0, 0, 88, 1, 0, 0, 0, 0, 90, 1, 0, 0, 0, 0, 92, 1, 0,
		0, 0, 0, 94, 1, 0, 0, 0, 0, 96, 1, 0, 0, 0, 0, 98, 1, 0, 0, 0, 0, 100,
		1, 0, 0, 0, 0, 102, 1, 0, 0, 0, 0, 104, 1, 0, 0, 0, 0, 106, 1, 0, 0, 0,
		0, 108, 1, 0, 0, 0, 1, 120, 1, 0, 0, 0, 1, 122, 1, 0, 0, 0, 1, 124, 1,
		0, 0, 0, 2, 126, 1, 0, 0, 0, 4, 128, 1, 0, 0, 0, 6, 130, 1, 0, 0, 0, 8,
		132, 1, 0, 0, 0, 10, 134, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 138, 1,
		0, 0, 0, 16, 140, 1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 144, 1, 0, 0, 0,
		22, 146, 1, 0, 0, 0, 24, 155, 1, 0, 0, 0, 26, 161, 1, 0, 0, 0, 28, 168,
		1, 0, 0, 0, 30, 175, 1, 0, 0, 0, 32, 183, 1, 0, 0, 0, 34, 187, 1, 0, 0,
		0, 36, 192, 1, 0, 0, 0, 38, 196, 1, 0, 0, 0, 40, 200, 1, 0, 0, 0, 42, 207,
		1, 0, 0, 0, 44, 214, 1, 0, 0, 0, 46, 222, 1, 0, 0, 0, 48, 230, 1, 0, 0,
		0, 50, 238, 1, 0, 0, 0, 52, 245, 1, 0, 0, 0, 54, 251, 1, 0, 0, 0, 56, 263,
		1, 0, 0, 0, 58, 266, 1, 0, 0, 0, 60, 277, 1, 0, 0, 0, 62, 281, 1, 0, 0,
		0, 64, 291, 1, 0, 0, 0, 66, 301, 1, 0, 0, 0, 68, 304, 1, 0, 0, 0, 70, 314,
		1, 0, 0, 0, 72, 322, 1, 0, 0, 0, 74, 331, 1, 0, 0, 0, 76, 343, 1, 0, 0,
		0, 78, 352, 1, 0, 0, 0, 80, 361, 1, 0, 0, 0, 82, 370, 1, 0, 0, 0, 84, 379,
		1, 0, 0, 0, 86, 388, 1, 0, 0, 0, 88, 395, 1, 0, 0, 0, 90, 402, 1, 0, 0,
		0, 92, 405, 1, 0, 0, 0, 94, 409, 1, 0, 0, 0, 96, 414, 1, 0, 0, 0, 98, 423,
		1, 0, 0, 0, 100, 426, 1, 0, 0, 0, 102, 433, 1, 0, 0, 0, 104, 439, 1, 0,
		0, 0, 106, 453, 1, 0, 0, 0, 108, 464, 1, 0, 0, 0, 110, 466, 1, 0, 0, 0,
		112, 471, 1, 0, 0, 0, 114, 476, 1, 0, 0, 0, 116, 478, 1, 0, 0, 0, 118,
		487, 1, 0, 0, 0, 120, 496, 1, 0, 0, 0, 122, 501, 1, 0, 0, 0, 124, 508,
		1, 0, 0, 0, 126, 127, 5, 59, 0, 0, 127, 3, 1, 0, 0, 0, 128, 129, 5, 40,
		0, 0, 129, 5, 1, 0, 0, 0, 130, 131, 5, 123, 0, 0, 131, 7, 1, 0, 0, 0, 132,
		133, 5, 41, 0, 0, 133, 9, 1, 0, 0, 0, 134, 135, 5, 125, 0, 0, 135, 11,
		1, 0, 0, 0, 136, 137, 5, 44, 0, 0, 137, 13, 1, 0, 0, 0, 138, 139, 5, 46,
		0, 0, 139, 15, 1, 0, 0, 0, 140, 141, 5, 36, 0, 0, 141, 17, 1, 0, 0, 0,
		142, 143, 5, 64, 0, 0, 143, 19, 1, 0, 0, 0, 144, 145, 5, 35, 0, 0, 145,
		21, 1, 0, 0, 0, 146, 147, 5, 100, 0, 0, 147, 148, 5, 97, 0, 0, 148, 149,
		5, 116, 0, 0, 149, 150, 5, 97, 0, 0, 150, 151, 5, 98, 0, 0, 151, 152, 5,
		97, 0, 0, 152, 153, 5, 115, 0, 0, 153, 154, 5, 101, 0, 0, 154, 23, 1, 0,
		0, 0, 155, 156, 5, 116, 0, 0, 156, 157, 5, 97, 0, 0, 157, 158, 5, 98, 0,
		0, 158, 159, 5, 108, 0, 0, 159, 160, 5, 101, 0, 0, 160, 25, 1, 0, 0, 0,
		161, 162, 5, 97, 0, 0, 162, 163, 5, 99, 0, 0, 163, 164, 5, 116, 0, 0, 164,
		165, 5, 105, 0, 0, 165, 166, 5, 111, 0, 0, 166, 167, 5, 110, 0, 0, 167,
		27, 1, 0, 0, 0, 168, 169, 5, 112, 0, 0, 169, 170, 5, 117, 0, 0, 170, 171,
		5, 98, 0, 0, 171, 172, 5, 108, 0, 0, 172, 173, 5, 105, 0, 0, 173, 174,
		5, 99, 0, 0, 174, 29, 1, 0, 0, 0, 175, 176, 5, 112, 0, 0, 176, 177, 5,
		114, 0, 0, 177, 178, 5, 105, 0, 0, 178, 179, 5, 118, 0, 0, 179, 180, 5,
		97, 0, 0, 180, 181, 5, 116, 0, 0, 181, 182, 5, 101, 0, 0, 182, 31, 1, 0,
		0, 0, 183, 184, 5, 105, 0, 0, 184, 185, 5, 110, 0, 0, 185, 186, 5, 116,
		0, 0, 186, 33, 1, 0, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 101, 0,
		0, 189, 190, 5, 120, 0, 0, 190, 191, 5, 116, 0, 0, 191, 35, 1, 0, 0, 0,
		192, 193, 5, 109, 0, 0, 193, 194, 5, 105, 0, 0, 194, 195, 5, 110, 0, 0,
		195, 37, 1, 0, 0, 0, 196, 197, 5, 109, 0, 0, 197, 198, 5, 97, 0, 0, 198,
		199, 5, 120, 0, 0, 199, 39, 1, 0, 0, 0, 200, 201, 5, 109, 0, 0, 201, 202,
		5, 105, 0, 0, 202, 203, 5, 110, 0, 0, 203, 204, 5, 108, 0, 0, 204, 205,
		5, 101, 0, 0, 205, 206, 5, 110, 0, 0, 206, 41, 1, 0, 0, 0, 207, 208, 5,
		109, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 120, 0, 0, 210, 211, 5,
		108, 0, 0, 211, 212, 5, 101, 0, 0, 212, 213, 5, 110, 0, 0, 213, 43, 1,
		0, 0, 0, 214, 215, 5, 110, 0, 0, 215, 216, 5, 111, 0, 0, 216, 217, 5, 116,
		0, 0, 217, 218, 5, 110, 0, 0, 218, 219, 5, 117, 0, 0, 219, 220, 5, 108,
		0, 0, 220, 221, 5, 108, 0, 0, 221, 45, 1, 0, 0, 0, 222, 223, 5, 112, 0,
		0, 223, 224, 5, 114, 0, 0, 224, 225, 5, 105, 0, 0, 225, 226, 5, 109, 0,
		0, 226, 227, 5, 97, 0, 0, 227, 228, 5, 114, 0, 0, 228, 229, 5, 121, 0,
		0, 229, 47, 1, 0, 0, 0, 230, 231, 5, 100, 0, 0, 231, 232, 5, 101, 0, 0,
		232, 233, 5, 102, 0, 0, 233, 234, 5, 97, 0, 0, 234, 235, 5, 117, 0, 0,
		235, 236, 5, 108, 0, 0, 236, 237, 5, 116, 0, 0, 237, 49, 1, 0, 0, 0, 238,
		239, 5, 117, 0, 0, 239, 240, 5, 110, 0, 0, 240, 241, 5, 105, 0, 0, 241,
		242, 5, 113, 0, 0, 242, 243, 5, 117, 0, 0, 243, 244, 5, 101, 0, 0, 244,
		51, 1, 0, 0, 0, 245, 246, 5, 105, 0, 0, 246, 247, 5, 110, 0, 0, 247, 248,
		5, 100, 0, 0, 248, 249, 5, 101, 0, 0, 249, 250, 5, 120, 0, 0, 250, 53,
		1, 0, 0, 0, 251, 252, 5, 102, 0, 0, 252, 253, 5, 111, 0, 0, 253, 254, 5,
		114, 0, 0, 254, 255, 5, 101, 0, 0, 255, 256, 5, 105, 0, 0, 256, 257, 5,
		103, 0, 0, 257, 258, 5, 110, 0, 0, 258, 259, 5, 95, 0, 0, 259, 260, 5,
		107, 0, 0, 260, 261, 5, 101, 0, 0, 261, 262, 5, 121, 0, 0, 262, 55, 1,
		0, 0, 0, 263, 264, 5, 102, 0, 0, 264, 265, 5, 107, 0, 0, 265, 57, 1, 0,
		0, 0, 266, 267, 5, 114, 0, 0, 267, 268, 5, 101, 0, 0, 268, 269, 5, 102,
		0, 0, 269, 270, 5, 101, 0, 0, 270, 271, 5, 114, 0, 0, 271, 272, 5, 101,
		0, 0, 272, 273, 5, 110, 0, 0, 273, 274, 5, 99, 0, 0, 274, 275, 5, 101,
		0, 0, 275, 276, 5, 115, 0, 0, 276, 59, 1, 0, 0, 0, 277, 278, 5, 114, 0,
		0, 278, 279, 5, 101, 0, 0, 279, 280, 5, 102, 0, 0, 280, 61, 1, 0, 0, 0,
		281, 282, 5, 111, 0, 0, 282, 283, 5, 110, 0, 0, 283, 284, 5, 95, 0, 0,
		284, 285, 5, 117, 0, 0, 285, 286, 5, 112, 0, 0, 286, 287, 5, 100, 0, 0,
		287, 288, 5, 97, 0, 0, 288, 289, 5, 116, 0, 0, 289, 290, 5, 101, 0, 0,
		290, 63, 1, 0, 0, 0, 291, 292, 5, 111, 0, 0, 292, 293, 5, 110, 0, 0, 293,
		294, 5, 95, 0, 0, 294, 295, 5, 100, 0, 0, 295, 296, 5, 101, 0, 0, 296,
		297, 5, 108, 0, 0, 297, 298, 5, 101, 0, 0, 298, 299, 5, 116, 0, 0, 299,
		300, 5, 101, 0, 0, 300, 65, 1, 0, 0, 0, 301, 302, 5, 100, 0, 0, 302, 303,
		5, 111, 0, 0, 303, 67, 1, 0, 0, 0, 304, 305, 5, 110, 0, 0, 305, 306, 5,
		111, 0, 0, 306, 307, 5, 95, 0, 0, 307, 308, 5, 97, 0, 0, 308, 309, 5, 99,
		0, 0, 309, 310, 5, 116, 0, 0, 310, 311, 5, 105, 0, 0, 311, 312, 5, 111,
		0, 0, 312, 313, 5, 110, 0, 0, 313, 69, 1, 0, 0, 0, 314, 315, 5, 99, 0,
		0, 315, 316, 5, 97, 0, 0, 316, 317, 5, 115, 0, 0, 317, 318, 5, 99, 0, 0,
		318, 319, 5, 97, 0, 0, 319, 320, 5, 100, 0, 0, 320, 321, 5, 101, 0, 0,
		321, 71, 1, 0, 0, 0, 322, 323, 5, 115, 0, 0, 323, 324, 5, 101, 0, 0, 324,
		325, 5, 116, 0, 0, 325, 326, 5, 95, 0, 0, 326, 327, 5, 110, 0, 0, 327,
		328, 5, 117, 0, 0, 328, 329, 5, 108, 0, 0, 329, 330, 5, 108, 0, 0, 330,
		73, 1, 0, 0, 0, 331, 332, 5, 115, 0, 0, 332, 333, 5, 101, 0, 0, 333, 334,
		5, 116, 0, 0, 334, 335, 5, 95, 0, 0, 335, 336, 5, 100, 0, 0, 336, 337,
		5, 101, 0, 0, 337, 338, 5, 102, 0, 0, 338, 339, 5, 97, 0, 0, 339, 340,
		5, 117, 0, 0, 340, 341, 5, 108, 0, 0, 341, 342, 5, 116, 0, 0, 342, 75,
		1, 0, 0, 0, 343, 344, 5, 114, 0, 0, 344, 345, 5, 101, 0, 0, 345, 346, 5,
		115, 0, 0, 346, 347, 5, 116, 0, 0, 347, 348, 5, 114, 0, 0, 348, 349, 5,
		105, 0, 0, 349, 350, 5, 99, 0, 0, 350, 351, 5, 116, 0, 0, 351, 77, 1, 0,
		0, 0, 352, 353, 7, 0, 0, 0, 353, 354, 7, 1, 0, 0, 354, 355, 7, 2, 0, 0,
		355, 356, 7, 1, 0, 0, 356, 357, 7, 3, 0, 0, 357, 358, 7, 4, 0, 0, 358,
		359, 1, 0, 0, 0, 359, 360, 6, 38, 0, 0, 360, 79, 1, 0, 0, 0, 361, 362,
		7, 5, 0, 0, 362, 363, 7, 6, 0, 0, 363, 364, 7, 0, 0, 0, 364, 365, 7, 1,
		0, 0, 365, 366, 7, 7, 0, 0, 366, 367, 7, 4, 0, 0, 367, 368, 1, 0, 0, 0,
		368, 369, 6, 39, 0, 0, 369, 81, 1, 0, 0, 0, 370, 371, 7, 8, 0, 0, 371,
		372, 7, 9, 0, 0, 372, 373, 7, 10, 0, 0, 373, 374, 7, 11, 0, 0, 374, 375,
		7, 4, 0, 0, 375, 376, 7, 1, 0, 0, 376, 377, 1, 0, 0, 0, 377, 378, 6, 40,
		0, 0, 378, 83, 1, 0, 0, 0, 379, 380, 7, 10, 0, 0, 380, 381, 7, 1, 0, 0,
		381, 382, 7, 2, 0, 0, 382, 383, 7, 1, 0, 0, 383, 384, 7, 4, 0, 0, 384,
		385, 7, 1, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 6, 41, 0, 0, 387, 85,
		1, 0, 0, 0, 388, 389, 7, 12, 0, 0, 389, 390, 7, 5, 0, 0, 390, 391, 7, 4,
		0, 0, 391, 392, 7, 13, 0, 0, 392, 393, 1, 0, 0, 0, 393, 394, 6, 42, 0,
		0, 394, 87, 1, 0, 0, 0, 395, 399, 7, 14, 0, 0, 396, 398, 7, 15, 0, 0, 397,
		396, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400,
		1, 0, 0, 0, 400, 89, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402, 403, 5, 35,
		0, 0, 403, 404, 3, 88, 43, 0, 404, 91, 1, 0, 0, 0, 405, 406, 5, 36, 0,
		0, 406, 407, 3, 88, 43, 0, 407, 93, 1, 0, 0, 0, 408, 410, 3, 110, 54, 0,
		409, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411,
		412, 1, 0, 0, 0, 412, 95, 1, 0, 0, 0, 413, 415, 7, 16, 0, 0, 414, 413,
		1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 417, 1, 0, 0, 0, 416, 418, 3, 110,
		54, 0, 417, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0,
		419, 420, 1, 0, 0, 0, 420, 97, 1, 0, 0, 0, 421, 424, 3, 116, 57, 0, 422,
		424, 3, 118, 58, 0, 423, 421, 1, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424, 99,
		1, 0, 0, 0, 425, 427, 7, 17, 0, 0, 426, 425, 1, 0, 0, 0, 427, 428, 1, 0,
		0, 0, 428, 426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0,
		430, 431, 6, 49, 1, 0, 431, 101, 1, 0, 0, 0, 432, 434, 7, 18, 0, 0, 433,
		432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 435, 436,
		1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 438, 6, 50, 1, 0, 438, 103, 1, 0,
		0, 0, 439, 440, 5, 47, 0, 0, 440, 441, 5, 42, 0, 0, 441, 445, 1, 0, 0,
		0, 442, 444, 9, 0, 0, 0, 443, 442, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0, 445,
		446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 448, 1, 0, 0, 0, 447, 445,
		1, 0, 0, 0, 448, 449, 5, 42, 0, 0, 449, 450, 5, 47, 0, 0, 450, 451, 1,
		0, 0, 0, 451, 452, 6, 51, 1, 0, 452, 105, 1, 0, 0, 0, 453, 454, 5, 47,
		0, 0, 454, 455, 5, 47, 0, 0, 455, 459, 1, 0, 0, 0, 456, 458, 8, 18, 0,
		0, 457, 456, 1, 0, 0, 0, 458, 461, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459,
		460, 1, 0, 0, 0, 460, 462, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 462, 463,
		6, 52, 1, 0, 463, 107, 1, 0, 0, 0, 464, 465, 9, 0, 0, 0, 465, 109, 1, 0,
		0, 0, 466, 467, 7, 19, 0, 0, 467, 111, 1, 0, 0, 0, 468, 472, 8, 20, 0,
		0, 469, 470, 5, 92, 0, 0, 470, 472, 9, 0, 0, 0, 471, 468, 1, 0, 0, 0, 471,
		469, 1, 0, 0, 0, 472, 113, 1, 0, 0, 0, 473, 477, 8, 21, 0, 0, 474, 475,
		5, 92, 0, 0, 475, 477, 9, 0, 0, 0, 476, 473, 1, 0, 0, 0, 476, 474, 1, 0,
		0, 0, 477, 115, 1, 0, 0, 0, 478, 482, 5, 34, 0, 0, 479, 481, 3, 112, 55,
		0, 480, 479, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 482,
		483, 1, 0, 0, 0, 483, 485, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 485, 486,
		5, 34, 0, 0, 486, 117, 1, 0, 0, 0, 487, 491, 5, 39, 0, 0, 488, 490, 3,
		114, 56, 0, 489, 488, 1, 0, 0, 0, 490, 493, 1, 0, 0, 0, 491, 489, 1, 0,
		0, 0, 491, 492, 1, 0, 0, 0, 492, 494, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0,
		494, 495, 5, 39, 0, 0, 495, 119, 1, 0, 0, 0, 496, 497, 3, 2, 0, 0, 497,
		498, 1, 0, 0, 0, 498, 499, 6, 59, 2, 0, 499, 121, 1, 0, 0, 0, 500, 502,
		7, 22, 0, 0, 501, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 501, 1, 0,
		0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 506, 6, 60, 3, 0,
		506, 123, 1, 0, 0, 0, 507, 509, 8, 23, 0, 0, 508, 507, 1, 0, 0, 0, 509,
		510, 1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511, 125,
		1, 0, 0, 0, 17, 0, 1, 399, 411, 414, 419, 423, 428, 435, 445, 459, 471,
		476, 482, 491, 503, 510, 4, 2, 1, 0, 0, 1, 0, 2, 0, 0, 6, 0, 0,
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

// KuneiformLexerInit initializes any static state used to implement KuneiformLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewKuneiformLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KuneiformLexerInit() {
	staticData := &kuneiformlexerLexerStaticData
	staticData.once.Do(kuneiformlexerLexerInit)
}

// NewKuneiformLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewKuneiformLexer(input antlr.CharStream) *KuneiformLexer {
	KuneiformLexerInit()
	l := new(KuneiformLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &kuneiformlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "KuneiformLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KuneiformLexer tokens.
const (
	KuneiformLexerSCOL                    = 1
	KuneiformLexerL_PAREN                 = 2
	KuneiformLexerL_BRACE                 = 3
	KuneiformLexerR_PAREN                 = 4
	KuneiformLexerR_BRACE                 = 5
	KuneiformLexerCOMMA                   = 6
	KuneiformLexerPERIOD                  = 7
	KuneiformLexerDOLLAR                  = 8
	KuneiformLexerAT                      = 9
	KuneiformLexerHASH                    = 10
	KuneiformLexerDATABASE_               = 11
	KuneiformLexerTABLE_                  = 12
	KuneiformLexerACTION_                 = 13
	KuneiformLexerPUBLIC_                 = 14
	KuneiformLexerPRIVATE_                = 15
	KuneiformLexerINT_                    = 16
	KuneiformLexerTEXT_                   = 17
	KuneiformLexerMIN_                    = 18
	KuneiformLexerMAX_                    = 19
	KuneiformLexerMIN_LEN_                = 20
	KuneiformLexerMAX_LEN_                = 21
	KuneiformLexerNOT_NULL_               = 22
	KuneiformLexerPRIMARY_                = 23
	KuneiformLexerDEFAULT_                = 24
	KuneiformLexerUNIQUE_                 = 25
	KuneiformLexerINDEX_                  = 26
	KuneiformLexerFOREIGN_KEY_            = 27
	KuneiformLexerFOREIGN_KEY_ABBR_       = 28
	KuneiformLexerREFERENCES_             = 29
	KuneiformLexerREFERENCES_ABBR_        = 30
	KuneiformLexerACTION_ON_UPDATE_       = 31
	KuneiformLexerACTION_ON_DELETE_       = 32
	KuneiformLexerACTION_DO_              = 33
	KuneiformLexerACTION_DO_NO_ACTION_    = 34
	KuneiformLexerACTION_DO_CASCADE_      = 35
	KuneiformLexerACTION_DO_SET_NULL_     = 36
	KuneiformLexerACTION_DO_SET_DEFAULT_  = 37
	KuneiformLexerACTION_DO_RESTRICT_     = 38
	KuneiformLexerSELECT_                 = 39
	KuneiformLexerINSERT_                 = 40
	KuneiformLexerUPDATE_                 = 41
	KuneiformLexerDELETE_                 = 42
	KuneiformLexerWITH_                   = 43
	KuneiformLexerIDENTIFIER              = 44
	KuneiformLexerINDEX_NAME              = 45
	KuneiformLexerACTION_PARAMETER        = 46
	KuneiformLexerUNSIGNED_NUMBER_LITERAL = 47
	KuneiformLexerSIGNED_NUMBER_LITERAL   = 48
	KuneiformLexerSTRING_LITERAL          = 49
	KuneiformLexerWS                      = 50
	KuneiformLexerTERMINATOR              = 51
	KuneiformLexerBLOCK_COMMENT           = 52
	KuneiformLexerLINE_COMMENT            = 53
	KuneiformLexerUNEXPECTED_CHAR         = 54
	KuneiformLexerSQL_END_SCOL            = 55
	KuneiformLexerSQL_NL                  = 56
	KuneiformLexerSQL_STMT                = 57
)

// KuneiformLexerSQL_MODE is the KuneiformLexer mode.
const KuneiformLexerSQL_MODE = 1
