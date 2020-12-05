package lex

import (
	"calculater_interpreter/dfa_state"
	"calculater_interpreter/token"
	"calculater_interpreter/token_type"
	"calculater_interpreter/utils"
	"fmt"
)

type SimpleLexi struct {
	token  token.Token
	tokens []*token.Token // 最终的 token 结果
}

func NewSimpleLexi() SimpleLexi {
	return SimpleLexi{
		token:  token.Token{},
		tokens: make([]*token.Token, 0),
	}
}

func (sl *SimpleLexi) Tokenize(script string) token.TokenReader {
	var state int = dfa_state.Inital

	for k, v := range script {
		fmt.Printf("loop start k: %d, v: %c, state: %d ||| token: %+v\n",
			k, v, state, sl.token)
		state = sl.doAutomation(v, state)
		fmt.Printf("loop end k: %d, v: %c, state: %d ||| token: %+v\n\n",
			k, v, state, sl.token)
	}

	return token.NewSimpleTokenReader(sl.tokens)
}

func (sl *SimpleLexi) resetToken() {
	sl.token.Type = ""
	sl.token.Text = ""
}

func (sl *SimpleLexi) initToken(c rune) int {
	if sl.token.Text != "" {
		fmt.Printf("init append: token[ %+v ]\n", sl.token)
		var aToken *token.Token
		aToken = new(token.Token)
		aToken.Type = sl.token.Type
		aToken.Text = sl.token.Text

		sl.tokens = append(sl.tokens, aToken)
		// token = &Token{"", ""}
		sl.resetToken()
	}
	s := string(c)
	newState := dfa_state.Inital

	if utils.IsAlpha(c) {
		if c == 'i' {
			newState = dfa_state.Int1
			sl.token.Type = token_type.Int
		} else {
			newState = dfa_state.Id
			sl.token.Type = token_type.Identifier
		}
		sl.token.Text += s
	} else if utils.IsDigit(c) {
		newState = dfa_state.IntLiteral
		sl.token.Type = token_type.IntLiteral
		sl.token.Text += s
	} else if c == '>' {
		newState = dfa_state.GT
		sl.token.Type = token_type.GT
		sl.token.Text += s
	} else if c == '=' {
		newState = dfa_state.Equal
		sl.token.Type = token_type.Assignment
		sl.token.Text += s
	} else if c == '+' {
		newState = dfa_state.Plus
		sl.token.Type = token_type.Plus
		sl.token.Text += s
	} else if c == '-' {
		newState = dfa_state.Minus
		sl.token.Type = token_type.Minus
		sl.token.Text += s
	} else if c == '*' {
		newState = dfa_state.Star
		sl.token.Type = token_type.Star
		sl.token.Text += s
	} else if c == '/' {
		newState = dfa_state.Slash
		sl.token.Type = token_type.Slash
		sl.token.Text += s
	}

	fmt.Printf("after init: tokens[ %+v ]\n", sl.tokens)
	return newState
}

func (sl *SimpleLexi) doAutomation(c rune, state int) int {
	fmt.Printf("doAutomation start: token[ %+v ]\n", sl.token)
	s := string(c)
	switch state {
	case dfa_state.Inital:
		state = sl.initToken(c)
		break
	case dfa_state.Id:
		if utils.IsAlpha(c) || utils.IsDigit(c) {
			sl.token.Text += s
		} else {
			state = sl.initToken(c)
		}
		break
	case dfa_state.GT:
		if c == '=' {
			state = dfa_state.GE
			sl.token.Type = token_type.GE
			sl.token.Text += s
		} else {
			state = sl.initToken(c)
		}
		break
	case dfa_state.GE, dfa_state.Equal, dfa_state.Plus,
		dfa_state.Minus, dfa_state.Star, dfa_state.Slash:
		state = sl.initToken(c)
		break
	case dfa_state.IntLiteral:
		if utils.IsDigit(c) {
			sl.token.Text += s
		} else {
			state = sl.initToken(c)
		}
		break
	case dfa_state.Int1:
		if c == 'n' {
			state = dfa_state.Int2
			sl.token.Text += s
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				sl.token.Type = token_type.Identifier
				sl.token.Text += s
			} else {
				state = sl.initToken(c)
			}
		}
		break
	case dfa_state.Int2:
		if c == 't' {
			state = dfa_state.Int3
			sl.token.Text += s
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				sl.token.Type = token_type.Identifier
				sl.token.Text += s
			} else {
				state = sl.initToken(c)
			}
		}
		break
	case dfa_state.Int3:
		if c == ' ' {
			state = sl.initToken(c)
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				sl.token.Type = token_type.Identifier
				sl.token.Text += s
			} else {
				state = sl.initToken(c)
			}
		}
		break
	}
	return state
}
