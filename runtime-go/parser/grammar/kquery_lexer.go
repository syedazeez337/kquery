// Code generated from grammar/KQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type KQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var KQueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kquerylexerLexerInit() {
	staticData := &KQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'.'", "'from'", "'where'", "'select'", "'limit'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "FROM", "WHERE", "SELECT", "LIMIT", "IDENT", "INT", "WS",
		"COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "FROM", "WHERE", "SELECT", "LIMIT", "IDENT", "INT",
		"WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 79, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 52, 8, 6, 10, 6, 12, 6, 55, 9, 6,
		1, 7, 4, 7, 58, 8, 7, 11, 7, 12, 7, 59, 1, 8, 4, 8, 63, 8, 8, 11, 8, 12,
		8, 64, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 73, 8, 9, 10, 9, 12, 9,
		76, 9, 9, 1, 9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 82, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1,
		21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 30, 1, 0, 0, 0,
		9, 36, 1, 0, 0, 0, 11, 43, 1, 0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 57, 1, 0,
		0, 0, 17, 62, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 22, 5, 44, 0, 0, 22,
		2, 1, 0, 0, 0, 23, 24, 5, 46, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 102,
		0, 0, 26, 27, 5, 114, 0, 0, 27, 28, 5, 111, 0, 0, 28, 29, 5, 109, 0, 0,
		29, 6, 1, 0, 0, 0, 30, 31, 5, 119, 0, 0, 31, 32, 5, 104, 0, 0, 32, 33,
		5, 101, 0, 0, 33, 34, 5, 114, 0, 0, 34, 35, 5, 101, 0, 0, 35, 8, 1, 0,
		0, 0, 36, 37, 5, 115, 0, 0, 37, 38, 5, 101, 0, 0, 38, 39, 5, 108, 0, 0,
		39, 40, 5, 101, 0, 0, 40, 41, 5, 99, 0, 0, 41, 42, 5, 116, 0, 0, 42, 10,
		1, 0, 0, 0, 43, 44, 5, 108, 0, 0, 44, 45, 5, 105, 0, 0, 45, 46, 5, 109,
		0, 0, 46, 47, 5, 105, 0, 0, 47, 48, 5, 116, 0, 0, 48, 12, 1, 0, 0, 0, 49,
		53, 7, 0, 0, 0, 50, 52, 7, 1, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0,
		0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 14, 1, 0, 0, 0, 55, 53,
		1, 0, 0, 0, 56, 58, 7, 2, 0, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0,
		59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 16, 1, 0, 0, 0, 61, 63, 7,
		3, 0, 0, 62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 6, 8, 0, 0, 67, 18, 1, 0, 0,
		0, 68, 69, 5, 47, 0, 0, 69, 70, 5, 47, 0, 0, 70, 74, 1, 0, 0, 0, 71, 73,
		8, 4, 0, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 6,
		9, 0, 0, 78, 20, 1, 0, 0, 0, 5, 0, 53, 59, 64, 74, 1, 6, 0, 0,
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

// KQueryLexerInit initializes any static state used to implement KQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewKQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KQueryLexerInit() {
	staticData := &KQueryLexerLexerStaticData
	staticData.once.Do(kquerylexerLexerInit)
}

// NewKQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewKQueryLexer(input antlr.CharStream) *KQueryLexer {
	KQueryLexerInit()
	l := new(KQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &KQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "KQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KQueryLexer tokens.
const (
	KQueryLexerT__0    = 1
	KQueryLexerT__1    = 2
	KQueryLexerFROM    = 3
	KQueryLexerWHERE   = 4
	KQueryLexerSELECT  = 5
	KQueryLexerLIMIT   = 6
	KQueryLexerIDENT   = 7
	KQueryLexerINT     = 8
	KQueryLexerWS      = 9
	KQueryLexerCOMMENT = 10
)
