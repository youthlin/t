package plurals

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/cockroachdb/errors"
)

var ErrAntlr = fmt.Errorf("err antlr")

// errorListener handler lexer/parser errors
type errorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.addError(errors.Errorf("SyntaxError(line %v:%v): %v", line, column, msg)) // SyntaxError 语法错误
}

func (d *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.addError(errors.Errorf("ReportAmbiguity")) // Ambiguity 歧义
}

func (d *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.addError(errors.Errorf("ReportAttemptingFullContext")) // SLL 冲突
}

func (d *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.addError(errors.Errorf("ReportContextSensitivity")) // 上下文相关
}

func (d *errorListener) addError(err error) {
	if d.err == nil {
		d.err = ErrAntlr
	}
	d.err = errors.WithSecondaryError(d.err, err)
}
