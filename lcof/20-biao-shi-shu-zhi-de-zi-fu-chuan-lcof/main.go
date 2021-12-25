package main

import (
	"leetcode/helper"
)

// 表示数值的字符串
type State int
type CharType int

const (
	StateInitial State = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: {
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: {
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: {
			CharNumber: StateFraction,
		},
		StateFraction: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: {
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: {
			CharNumber: StateExpNumber,
		},
		StateExpNumber: {
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: {
			CharSpace: StateEnd,
		},
	}
	var ok bool
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if state, ok = transfer[state][typ]; !ok {
			return false
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction ||
		state == StateExpNumber || state == StateEnd
}

func main() {
	helper.Assert(isNumber("+100"))
	helper.Assert(isNumber("5e2"))
	helper.Assert(isNumber("-123"))
	helper.Assert(isNumber("3.1416"))
	helper.Assert(isNumber("-1E-16"))
	helper.Assert(isNumber("0123"))
	helper.Assert(!isNumber("12e"))
	helper.Assert(!isNumber("1a3.14"))
	helper.Assert(!isNumber("1.2.3"))
	helper.Assert(!isNumber("+-5"))
	helper.Assert(!isNumber("12e+5.4"))
}
