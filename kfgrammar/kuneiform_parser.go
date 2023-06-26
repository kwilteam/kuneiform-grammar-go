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
		"source_unit", "database_directive", "table_decl", "column_def", "column_def_list",
		"column_type", "column_constraint", "literal_value", "number_value",
		"index_def", "foreign_key_action", "foreign_key_def", "action_decl",
		"action_param_list", "database_name", "table_name", "action_name", "column_name",
		"column_name_list", "index_name", "sql_keywords", "sql_stmt", "action_stmt",
		"action_stmt_list",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 57, 214, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 53,
		8, 0, 10, 0, 12, 0, 56, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 5, 2, 72, 8, 2, 10, 2, 12,
		2, 75, 9, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3,
		85, 8, 3, 10, 3, 12, 3, 88, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 93, 8, 4, 10,
		4, 12, 4, 96, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 128, 8,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 3, 10, 142, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 156, 8, 11, 10, 11, 12, 11, 159,
		9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 3, 13, 172, 8, 13, 1, 13, 1, 13, 5, 13, 176, 8, 13, 10, 13,
		12, 13, 179, 9, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 5, 18, 192, 8, 18, 10, 18, 12, 18, 195, 9, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 5, 23, 209, 8, 23, 10, 23, 12, 23, 212, 9, 23, 1, 23, 0, 0,
		24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 0, 9, 1, 0, 16, 17, 2, 0, 47, 47, 49, 49, 2, 0,
		23, 23, 25, 26, 1, 0, 31, 32, 1, 0, 34, 38, 1, 0, 27, 28, 1, 0, 29, 30,
		1, 0, 14, 15, 1, 0, 39, 43, 209, 0, 48, 1, 0, 0, 0, 2, 59, 1, 0, 0, 0,
		4, 62, 1, 0, 0, 0, 6, 81, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 97, 1, 0,
		0, 0, 12, 127, 1, 0, 0, 0, 14, 129, 1, 0, 0, 0, 16, 131, 1, 0, 0, 0, 18,
		133, 1, 0, 0, 0, 20, 139, 1, 0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 160, 1,
		0, 0, 0, 26, 171, 1, 0, 0, 0, 28, 180, 1, 0, 0, 0, 30, 182, 1, 0, 0, 0,
		32, 184, 1, 0, 0, 0, 34, 186, 1, 0, 0, 0, 36, 188, 1, 0, 0, 0, 38, 196,
		1, 0, 0, 0, 40, 198, 1, 0, 0, 0, 42, 200, 1, 0, 0, 0, 44, 204, 1, 0, 0,
		0, 46, 206, 1, 0, 0, 0, 48, 49, 3, 2, 1, 0, 49, 54, 5, 1, 0, 0, 50, 53,
		3, 4, 2, 0, 51, 53, 3, 24, 12, 0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0,
		0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57,
		1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 58, 5, 0, 0, 1, 58, 1, 1, 0, 0, 0,
		59, 60, 5, 11, 0, 0, 60, 61, 3, 28, 14, 0, 61, 3, 1, 0, 0, 0, 62, 63, 5,
		12, 0, 0, 63, 64, 3, 30, 15, 0, 64, 65, 5, 3, 0, 0, 65, 73, 3, 8, 4, 0,
		66, 69, 5, 6, 0, 0, 67, 70, 3, 18, 9, 0, 68, 70, 3, 22, 11, 0, 69, 67,
		1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0,
		72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 77, 1,
		0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 78, 5, 6, 0, 0, 77, 76, 1, 0, 0, 0, 77,
		78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 5, 0, 0, 80, 5, 1, 0, 0,
		0, 81, 82, 3, 34, 17, 0, 82, 86, 3, 10, 5, 0, 83, 85, 3, 12, 6, 0, 84,
		83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0,
		0, 87, 7, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 94, 3, 6, 3, 0, 90, 91, 5,
		6, 0, 0, 91, 93, 3, 6, 3, 0, 92, 90, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 9, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 97, 98, 7, 0, 0, 0, 98, 11, 1, 0, 0, 0, 99, 128, 5, 23, 0, 0, 100, 128,
		5, 22, 0, 0, 101, 128, 5, 25, 0, 0, 102, 103, 5, 24, 0, 0, 103, 104, 5,
		2, 0, 0, 104, 105, 3, 14, 7, 0, 105, 106, 5, 4, 0, 0, 106, 128, 1, 0, 0,
		0, 107, 108, 5, 18, 0, 0, 108, 109, 5, 2, 0, 0, 109, 110, 3, 16, 8, 0,
		110, 111, 5, 4, 0, 0, 111, 128, 1, 0, 0, 0, 112, 113, 5, 19, 0, 0, 113,
		114, 5, 2, 0, 0, 114, 115, 3, 16, 8, 0, 115, 116, 5, 4, 0, 0, 116, 128,
		1, 0, 0, 0, 117, 118, 5, 20, 0, 0, 118, 119, 5, 2, 0, 0, 119, 120, 3, 16,
		8, 0, 120, 121, 5, 4, 0, 0, 121, 128, 1, 0, 0, 0, 122, 123, 5, 21, 0, 0,
		123, 124, 5, 2, 0, 0, 124, 125, 3, 16, 8, 0, 125, 126, 5, 4, 0, 0, 126,
		128, 1, 0, 0, 0, 127, 99, 1, 0, 0, 0, 127, 100, 1, 0, 0, 0, 127, 101, 1,
		0, 0, 0, 127, 102, 1, 0, 0, 0, 127, 107, 1, 0, 0, 0, 127, 112, 1, 0, 0,
		0, 127, 117, 1, 0, 0, 0, 127, 122, 1, 0, 0, 0, 128, 13, 1, 0, 0, 0, 129,
		130, 7, 1, 0, 0, 130, 15, 1, 0, 0, 0, 131, 132, 5, 47, 0, 0, 132, 17, 1,
		0, 0, 0, 133, 134, 3, 38, 19, 0, 134, 135, 7, 2, 0, 0, 135, 136, 5, 2,
		0, 0, 136, 137, 3, 36, 18, 0, 137, 138, 5, 4, 0, 0, 138, 19, 1, 0, 0, 0,
		139, 141, 7, 3, 0, 0, 140, 142, 5, 33, 0, 0, 141, 140, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 7, 4, 0, 0, 144, 21, 1,
		0, 0, 0, 145, 146, 7, 5, 0, 0, 146, 147, 5, 2, 0, 0, 147, 148, 3, 36, 18,
		0, 148, 149, 5, 4, 0, 0, 149, 150, 7, 6, 0, 0, 150, 151, 3, 30, 15, 0,
		151, 152, 5, 2, 0, 0, 152, 153, 3, 36, 18, 0, 153, 157, 5, 4, 0, 0, 154,
		156, 3, 20, 10, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 23, 1, 0, 0, 0, 159, 157, 1, 0,
		0, 0, 160, 161, 5, 13, 0, 0, 161, 162, 3, 32, 16, 0, 162, 163, 5, 2, 0,
		0, 163, 164, 3, 26, 13, 0, 164, 165, 5, 4, 0, 0, 165, 166, 7, 7, 0, 0,
		166, 167, 5, 3, 0, 0, 167, 168, 3, 46, 23, 0, 168, 169, 5, 5, 0, 0, 169,
		25, 1, 0, 0, 0, 170, 172, 5, 46, 0, 0, 171, 170, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 177, 1, 0, 0, 0, 173, 174, 5, 6, 0, 0, 174, 176, 5, 46,
		0, 0, 175, 173, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0,
		177, 178, 1, 0, 0, 0, 178, 27, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 181,
		5, 44, 0, 0, 181, 29, 1, 0, 0, 0, 182, 183, 5, 44, 0, 0, 183, 31, 1, 0,
		0, 0, 184, 185, 5, 44, 0, 0, 185, 33, 1, 0, 0, 0, 186, 187, 5, 44, 0, 0,
		187, 35, 1, 0, 0, 0, 188, 193, 3, 34, 17, 0, 189, 190, 5, 6, 0, 0, 190,
		192, 3, 34, 17, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 37, 1, 0, 0, 0, 195, 193, 1, 0,
		0, 0, 196, 197, 5, 45, 0, 0, 197, 39, 1, 0, 0, 0, 198, 199, 7, 8, 0, 0,
		199, 41, 1, 0, 0, 0, 200, 201, 3, 40, 20, 0, 201, 202, 5, 57, 0, 0, 202,
		203, 5, 55, 0, 0, 203, 43, 1, 0, 0, 0, 204, 205, 3, 42, 21, 0, 205, 45,
		1, 0, 0, 0, 206, 210, 3, 44, 22, 0, 207, 209, 3, 44, 22, 0, 208, 207, 1,
		0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0,
		0, 211, 47, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 14, 52, 54, 69, 73, 77, 86,
		94, 127, 141, 157, 171, 177, 193, 210,
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
	KuneiformParserSCOL                    = 1
	KuneiformParserL_PAREN                 = 2
	KuneiformParserL_BRACE                 = 3
	KuneiformParserR_PAREN                 = 4
	KuneiformParserR_BRACE                 = 5
	KuneiformParserCOMMA                   = 6
	KuneiformParserPERIOD                  = 7
	KuneiformParserDOLLAR                  = 8
	KuneiformParserAT                      = 9
	KuneiformParserHASH                    = 10
	KuneiformParserDATABASE_               = 11
	KuneiformParserTABLE_                  = 12
	KuneiformParserACTION_                 = 13
	KuneiformParserPUBLIC_                 = 14
	KuneiformParserPRIVATE_                = 15
	KuneiformParserINT_                    = 16
	KuneiformParserTEXT_                   = 17
	KuneiformParserMIN_                    = 18
	KuneiformParserMAX_                    = 19
	KuneiformParserMIN_LEN_                = 20
	KuneiformParserMAX_LEN_                = 21
	KuneiformParserNOT_NULL_               = 22
	KuneiformParserPRIMARY_                = 23
	KuneiformParserDEFAULT_                = 24
	KuneiformParserUNIQUE_                 = 25
	KuneiformParserINDEX_                  = 26
	KuneiformParserFOREIGN_KEY_            = 27
	KuneiformParserFOREIGN_KEY_ABBR_       = 28
	KuneiformParserREFERENCES_             = 29
	KuneiformParserREFERENCES_ABBR_        = 30
	KuneiformParserACTION_ON_UPDATE_       = 31
	KuneiformParserACTION_ON_DELETE_       = 32
	KuneiformParserACTION_DO_              = 33
	KuneiformParserACTION_DO_NO_ACTION_    = 34
	KuneiformParserACTION_DO_CASCADE_      = 35
	KuneiformParserACTION_DO_SET_NULL_     = 36
	KuneiformParserACTION_DO_SET_DEFAULT_  = 37
	KuneiformParserACTION_DO_RESTRICT_     = 38
	KuneiformParserSELECT_                 = 39
	KuneiformParserINSERT_                 = 40
	KuneiformParserUPDATE_                 = 41
	KuneiformParserDELETE_                 = 42
	KuneiformParserWITH_                   = 43
	KuneiformParserIDENTIFIER              = 44
	KuneiformParserINDEX_NAME              = 45
	KuneiformParserACTION_PARAMETER        = 46
	KuneiformParserUNSIGNED_NUMBER_LITERAL = 47
	KuneiformParserSIGNED_NUMBER_LITERAL   = 48
	KuneiformParserSTRING_LITERAL          = 49
	KuneiformParserWS                      = 50
	KuneiformParserTERMINATOR              = 51
	KuneiformParserBLOCK_COMMENT           = 52
	KuneiformParserLINE_COMMENT            = 53
	KuneiformParserUNEXPECTED_CHAR         = 54
	KuneiformParserSQL_END_SCOL            = 55
	KuneiformParserSQL_NL                  = 56
	KuneiformParserSQL_STMT                = 57
)

// KuneiformParser rules.
const (
	KuneiformParserRULE_source_unit        = 0
	KuneiformParserRULE_database_directive = 1
	KuneiformParserRULE_table_decl         = 2
	KuneiformParserRULE_column_def         = 3
	KuneiformParserRULE_column_def_list    = 4
	KuneiformParserRULE_column_type        = 5
	KuneiformParserRULE_column_constraint  = 6
	KuneiformParserRULE_literal_value      = 7
	KuneiformParserRULE_number_value       = 8
	KuneiformParserRULE_index_def          = 9
	KuneiformParserRULE_foreign_key_action = 10
	KuneiformParserRULE_foreign_key_def    = 11
	KuneiformParserRULE_action_decl        = 12
	KuneiformParserRULE_action_param_list  = 13
	KuneiformParserRULE_database_name      = 14
	KuneiformParserRULE_table_name         = 15
	KuneiformParserRULE_action_name        = 16
	KuneiformParserRULE_column_name        = 17
	KuneiformParserRULE_column_name_list   = 18
	KuneiformParserRULE_index_name         = 19
	KuneiformParserRULE_sql_keywords       = 20
	KuneiformParserRULE_sql_stmt           = 21
	KuneiformParserRULE_action_stmt        = 22
	KuneiformParserRULE_action_stmt_list   = 23
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
	AllTable_decl() []ITable_declContext
	Table_decl(i int) ITable_declContext
	AllAction_decl() []IAction_declContext
	Action_decl(i int) IAction_declContext

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
		p.SetState(48)
		p.Database_directive()
	}
	{
		p.SetState(49)
		p.Match(KuneiformParserSCOL)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserTABLE_ || _la == KuneiformParserACTION_ {
		p.SetState(52)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserTABLE_:
			{
				p.SetState(50)
				p.Table_decl()
			}

		case KuneiformParserACTION_:
			{
				p.SetState(51)
				p.Action_decl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
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
		p.SetState(59)
		p.Match(KuneiformParserDATABASE_)
	}
	{
		p.SetState(60)
		p.Database_name()
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
	p.EnterRule(localctx, 4, KuneiformParserRULE_table_decl)
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
		p.SetState(62)
		p.Match(KuneiformParserTABLE_)
	}
	{
		p.SetState(63)
		p.Table_name()
	}
	{
		p.SetState(64)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(65)
		p.Column_def_list()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(66)
				p.Match(KuneiformParserCOMMA)
			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case KuneiformParserINDEX_NAME:
				{
					p.SetState(67)
					p.Index_def()
				}

			case KuneiformParserFOREIGN_KEY_, KuneiformParserFOREIGN_KEY_ABBR_:
				{
					p.SetState(68)
					p.Foreign_key_def()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserCOMMA {
		{
			p.SetState(76)
			p.Match(KuneiformParserCOMMA)
		}

	}
	{
		p.SetState(79)
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
	p.EnterRule(localctx, 6, KuneiformParserRULE_column_def)
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
		p.SetState(81)
		p.Column_name()
	}
	{
		p.SetState(82)
		p.Column_type()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66846720) != 0 {
		{
			p.SetState(83)
			p.Column_constraint()
		}

		p.SetState(88)
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
	p.EnterRule(localctx, 8, KuneiformParserRULE_column_def_list)

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
		p.SetState(89)
		p.Column_def()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(90)
				p.Match(KuneiformParserCOMMA)
			}
			{
				p.SetState(91)
				p.Column_def()
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, KuneiformParserRULE_column_type)
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
		p.SetState(97)
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
	p.EnterRule(localctx, 12, KuneiformParserRULE_column_constraint)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserPRIMARY_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(KuneiformParserPRIMARY_)
		}

	case KuneiformParserNOT_NULL_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(KuneiformParserNOT_NULL_)
		}

	case KuneiformParserUNIQUE_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Match(KuneiformParserUNIQUE_)
		}

	case KuneiformParserDEFAULT_:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Match(KuneiformParserDEFAULT_)
		}
		{
			p.SetState(103)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(104)
			p.Literal_value()
		}
		{
			p.SetState(105)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(KuneiformParserMIN_)
		}
		{
			p.SetState(108)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(109)
			p.Number_value()
		}
		{
			p.SetState(110)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(112)
			p.Match(KuneiformParserMAX_)
		}
		{
			p.SetState(113)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(114)
			p.Number_value()
		}
		{
			p.SetState(115)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_LEN_:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.Match(KuneiformParserMIN_LEN_)
		}
		{
			p.SetState(118)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(119)
			p.Number_value()
		}
		{
			p.SetState(120)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_LEN_:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(122)
			p.Match(KuneiformParserMAX_LEN_)
		}
		{
			p.SetState(123)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(124)
			p.Number_value()
		}
		{
			p.SetState(125)
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
	p.EnterRule(localctx, 14, KuneiformParserRULE_literal_value)
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
		p.SetState(129)
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
	p.EnterRule(localctx, 16, KuneiformParserRULE_number_value)

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
	p.EnterRule(localctx, 18, KuneiformParserRULE_index_def)
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
		p.SetState(133)
		p.Index_name()
	}
	{
		p.SetState(134)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&109051904) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(135)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(136)
		p.Column_name_list()
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 20, KuneiformParserRULE_foreign_key_action)
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
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserACTION_DO_ {
		{
			p.SetState(140)
			p.Match(KuneiformParserACTION_DO_)
		}

	}
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&532575944704) != 0) {
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
	p.EnterRule(localctx, 22, KuneiformParserRULE_foreign_key_def)
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
		p.SetState(145)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserFOREIGN_KEY_ || _la == KuneiformParserFOREIGN_KEY_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(146)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(147)
		p.Column_name_list()
	}
	{
		p.SetState(148)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserREFERENCES_ || _la == KuneiformParserREFERENCES_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(150)
		p.Table_name()
	}
	{
		p.SetState(151)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(152)
		p.Column_name_list()
	}
	{
		p.SetState(153)
		p.Match(KuneiformParserR_PAREN)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_ {
		{
			p.SetState(154)
			p.Foreign_key_action()
		}

		p.SetState(159)
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
	Action_param_list() IAction_param_listContext
	R_PAREN() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	R_BRACE() antlr.TerminalNode
	PUBLIC_() antlr.TerminalNode
	PRIVATE_() antlr.TerminalNode

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

func (s *Action_declContext) Action_param_list() IAction_param_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_param_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_param_listContext)
}

func (s *Action_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
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

func (s *Action_declContext) PUBLIC_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPUBLIC_, 0)
}

func (s *Action_declContext) PRIVATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIVATE_, 0)
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
	p.EnterRule(localctx, 24, KuneiformParserRULE_action_decl)
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
		p.Match(KuneiformParserACTION_)
	}
	{
		p.SetState(161)
		p.Action_name()
	}
	{
		p.SetState(162)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(163)
		p.Action_param_list()
	}
	{
		p.SetState(164)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(165)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserPUBLIC_ || _la == KuneiformParserPRIVATE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(166)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(167)
		p.Action_stmt_list()
	}
	{
		p.SetState(168)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IAction_param_listContext is an interface to support dynamic dispatch.
type IAction_param_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllACTION_PARAMETER() []antlr.TerminalNode
	ACTION_PARAMETER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAction_param_listContext differentiates from other interfaces.
	IsAction_param_listContext()
}

type Action_param_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_param_listContext() *Action_param_listContext {
	var p = new(Action_param_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_param_list
	return p
}

func (*Action_param_listContext) IsAction_param_listContext() {}

func NewAction_param_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_param_listContext {
	var p = new(Action_param_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_param_list

	return p
}

func (s *Action_param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_param_listContext) AllACTION_PARAMETER() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserACTION_PARAMETER)
}

func (s *Action_param_listContext) ACTION_PARAMETER(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_PARAMETER, i)
}

func (s *Action_param_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Action_param_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Action_param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_param_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_param_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_param_list() (localctx IAction_param_listContext) {
	this := p
	_ = this

	localctx = NewAction_param_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KuneiformParserRULE_action_param_list)
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
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserACTION_PARAMETER {
		{
			p.SetState(170)
			p.Match(KuneiformParserACTION_PARAMETER)
		}

	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(173)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(174)
			p.Match(KuneiformParserACTION_PARAMETER)
		}

		p.SetState(179)
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
	p.EnterRule(localctx, 28, KuneiformParserRULE_database_name)

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
		p.SetState(180)
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
	p.EnterRule(localctx, 30, KuneiformParserRULE_table_name)

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
		p.SetState(182)
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
	p.EnterRule(localctx, 32, KuneiformParserRULE_action_name)

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
		p.SetState(184)
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
	p.EnterRule(localctx, 34, KuneiformParserRULE_column_name)

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
		p.SetState(186)
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
	p.EnterRule(localctx, 36, KuneiformParserRULE_column_name_list)
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
		p.SetState(188)
		p.Column_name()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(189)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(190)
			p.Column_name()
		}

		p.SetState(195)
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
	p.EnterRule(localctx, 38, KuneiformParserRULE_index_name)

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
		p.Match(KuneiformParserINDEX_NAME)
	}

	return localctx
}

// ISql_keywordsContext is an interface to support dynamic dispatch.
type ISql_keywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT_() antlr.TerminalNode
	INSERT_() antlr.TerminalNode
	UPDATE_() antlr.TerminalNode
	DELETE_() antlr.TerminalNode
	WITH_() antlr.TerminalNode

	// IsSql_keywordsContext differentiates from other interfaces.
	IsSql_keywordsContext()
}

type Sql_keywordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_keywordsContext() *Sql_keywordsContext {
	var p = new(Sql_keywordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_sql_keywords
	return p
}

func (*Sql_keywordsContext) IsSql_keywordsContext() {}

func NewSql_keywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_keywordsContext {
	var p = new(Sql_keywordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_sql_keywords

	return p
}

func (s *Sql_keywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_keywordsContext) SELECT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSELECT_, 0)
}

func (s *Sql_keywordsContext) INSERT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINSERT_, 0)
}

func (s *Sql_keywordsContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUPDATE_, 0)
}

func (s *Sql_keywordsContext) DELETE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDELETE_, 0)
}

func (s *Sql_keywordsContext) WITH_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserWITH_, 0)
}

func (s *Sql_keywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_keywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_keywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSql_keywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Sql_keywords() (localctx ISql_keywordsContext) {
	this := p
	_ = this

	localctx = NewSql_keywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KuneiformParserRULE_sql_keywords)
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
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17042430230528) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sql_keywords() ISql_keywordsContext
	SQL_STMT() antlr.TerminalNode
	SQL_END_SCOL() antlr.TerminalNode

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

func (s *Sql_stmtContext) Sql_keywords() ISql_keywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_keywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_keywordsContext)
}

func (s *Sql_stmtContext) SQL_STMT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_STMT, 0)
}

func (s *Sql_stmtContext) SQL_END_SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_END_SCOL, 0)
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
	p.EnterRule(localctx, 42, KuneiformParserRULE_sql_stmt)

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
		p.SetState(200)
		p.Sql_keywords()
	}
	{
		p.SetState(201)
		p.Match(KuneiformParserSQL_STMT)
	}
	{
		p.SetState(202)
		p.Match(KuneiformParserSQL_END_SCOL)
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
	p.EnterRule(localctx, 44, KuneiformParserRULE_action_stmt)

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
		p.Sql_stmt()
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
	p.EnterRule(localctx, 46, KuneiformParserRULE_action_stmt_list)
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
		p.Action_stmt()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17042430230528) != 0 {
		{
			p.SetState(207)
			p.Action_stmt()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
