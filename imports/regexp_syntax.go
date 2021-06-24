// this file was generated by gomacro command: import _b "regexp/syntax"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	syntax "regexp/syntax"
)

// reflection: allow interpreted code to import "regexp/syntax"
func init() {
	Packages["regexp/syntax"] = Package{
	Name: "syntax",
	Binds: map[string]Value{
		"ClassNL":	ValueOf(syntax.ClassNL),
		"Compile":	ValueOf(syntax.Compile),
		"DotNL":	ValueOf(syntax.DotNL),
		"EmptyBeginLine":	ValueOf(syntax.EmptyBeginLine),
		"EmptyBeginText":	ValueOf(syntax.EmptyBeginText),
		"EmptyEndLine":	ValueOf(syntax.EmptyEndLine),
		"EmptyEndText":	ValueOf(syntax.EmptyEndText),
		"EmptyNoWordBoundary":	ValueOf(syntax.EmptyNoWordBoundary),
		"EmptyOpContext":	ValueOf(syntax.EmptyOpContext),
		"EmptyWordBoundary":	ValueOf(syntax.EmptyWordBoundary),
		"ErrInternalError":	ValueOf(syntax.ErrInternalError),
		"ErrInvalidCharClass":	ValueOf(syntax.ErrInvalidCharClass),
		"ErrInvalidCharRange":	ValueOf(syntax.ErrInvalidCharRange),
		"ErrInvalidEscape":	ValueOf(syntax.ErrInvalidEscape),
		"ErrInvalidNamedCapture":	ValueOf(syntax.ErrInvalidNamedCapture),
		"ErrInvalidPerlOp":	ValueOf(syntax.ErrInvalidPerlOp),
		"ErrInvalidRepeatOp":	ValueOf(syntax.ErrInvalidRepeatOp),
		"ErrInvalidRepeatSize":	ValueOf(syntax.ErrInvalidRepeatSize),
		"ErrInvalidUTF8":	ValueOf(syntax.ErrInvalidUTF8),
		"ErrMissingBracket":	ValueOf(syntax.ErrMissingBracket),
		"ErrMissingParen":	ValueOf(syntax.ErrMissingParen),
		"ErrMissingRepeatArgument":	ValueOf(syntax.ErrMissingRepeatArgument),
		"ErrTrailingBackslash":	ValueOf(syntax.ErrTrailingBackslash),
		"ErrUnexpectedParen":	ValueOf(syntax.ErrUnexpectedParen),
		"FoldCase":	ValueOf(syntax.FoldCase),
		"InstAlt":	ValueOf(syntax.InstAlt),
		"InstAltMatch":	ValueOf(syntax.InstAltMatch),
		"InstCapture":	ValueOf(syntax.InstCapture),
		"InstEmptyWidth":	ValueOf(syntax.InstEmptyWidth),
		"InstFail":	ValueOf(syntax.InstFail),
		"InstMatch":	ValueOf(syntax.InstMatch),
		"InstNop":	ValueOf(syntax.InstNop),
		"InstRune":	ValueOf(syntax.InstRune),
		"InstRune1":	ValueOf(syntax.InstRune1),
		"InstRuneAny":	ValueOf(syntax.InstRuneAny),
		"InstRuneAnyNotNL":	ValueOf(syntax.InstRuneAnyNotNL),
		"IsWordChar":	ValueOf(syntax.IsWordChar),
		"Literal":	ValueOf(syntax.Literal),
		"MatchNL":	ValueOf(syntax.MatchNL),
		"NonGreedy":	ValueOf(syntax.NonGreedy),
		"OneLine":	ValueOf(syntax.OneLine),
		"OpAlternate":	ValueOf(syntax.OpAlternate),
		"OpAnyChar":	ValueOf(syntax.OpAnyChar),
		"OpAnyCharNotNL":	ValueOf(syntax.OpAnyCharNotNL),
		"OpBeginLine":	ValueOf(syntax.OpBeginLine),
		"OpBeginText":	ValueOf(syntax.OpBeginText),
		"OpCapture":	ValueOf(syntax.OpCapture),
		"OpCharClass":	ValueOf(syntax.OpCharClass),
		"OpConcat":	ValueOf(syntax.OpConcat),
		"OpEmptyMatch":	ValueOf(syntax.OpEmptyMatch),
		"OpEndLine":	ValueOf(syntax.OpEndLine),
		"OpEndText":	ValueOf(syntax.OpEndText),
		"OpLiteral":	ValueOf(syntax.OpLiteral),
		"OpNoMatch":	ValueOf(syntax.OpNoMatch),
		"OpNoWordBoundary":	ValueOf(syntax.OpNoWordBoundary),
		"OpPlus":	ValueOf(syntax.OpPlus),
		"OpQuest":	ValueOf(syntax.OpQuest),
		"OpRepeat":	ValueOf(syntax.OpRepeat),
		"OpStar":	ValueOf(syntax.OpStar),
		"OpWordBoundary":	ValueOf(syntax.OpWordBoundary),
		"POSIX":	ValueOf(syntax.POSIX),
		"Parse":	ValueOf(syntax.Parse),
		"Perl":	ValueOf(syntax.Perl),
		"PerlX":	ValueOf(syntax.PerlX),
		"Simple":	ValueOf(syntax.Simple),
		"UnicodeGroups":	ValueOf(syntax.UnicodeGroups),
		"WasDollar":	ValueOf(syntax.WasDollar),
	}, Types: map[string]Type{
		"EmptyOp":	TypeOf((*syntax.EmptyOp)(nil)).Elem(),
		"Error":	TypeOf((*syntax.Error)(nil)).Elem(),
		"ErrorCode":	TypeOf((*syntax.ErrorCode)(nil)).Elem(),
		"Flags":	TypeOf((*syntax.Flags)(nil)).Elem(),
		"Inst":	TypeOf((*syntax.Inst)(nil)).Elem(),
		"InstOp":	TypeOf((*syntax.InstOp)(nil)).Elem(),
		"Op":	TypeOf((*syntax.Op)(nil)).Elem(),
		"Prog":	TypeOf((*syntax.Prog)(nil)).Elem(),
		"Regexp":	TypeOf((*syntax.Regexp)(nil)).Elem(),
	}, 
	}
}
