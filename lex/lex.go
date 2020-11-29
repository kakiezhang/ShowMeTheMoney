package lex

import (
	"calculater_interpreter/dfa_state"
	"calculater_interpreter/token_type"
	"calculater_interpreter/utils"
	"fmt"
)

var tokens []*Token

type Token struct {
	Type string
	Text string
}

func LexiRecog(script string) Token {
	var token Token
	var state int = dfa_state.Inital

	for k, v := range script {
		fmt.Printf("loop start k: %d, v: %c, state: %d ||| token: %+v\n",
			k, v, state, token)
		state = token.tokenize(v, state)
		fmt.Printf("loop end k: %d, v: %c, state: %d ||| token: %+v\n\n",
			k, v, state, token)
	}

	return token
}

func (token *Token) reset() {
	token.Type = ""
	token.Text = ""
}

func (token *Token) init(c rune) int {
	if token.Text != "" {
		fmt.Printf("init append: token[ %+v ]\n", token)
		var aToken *Token
		aToken = new(Token)
		aToken.Type = token.Type
		aToken.Text = token.Text

		tokens = append(tokens, aToken)
		// token = &Token{"", ""}
		token.reset()
	}
	s := string(c)
	newState := dfa_state.Inital

	if utils.IsAlpha(c) {
		if c == 'i' {
			newState = dfa_state.Int1
			token.Type = token_type.Int
		} else {
			newState = dfa_state.Id
			token.Type = token_type.Identifier
		}
		token.Text += s
	} else if utils.IsDigit(c) {
		newState = dfa_state.IntLiteral
		token.Type = token_type.IntLiteral
		token.Text += s
	} else if c == '>' {
		newState = dfa_state.GT
		token.Type = token_type.GT
		token.Text += s
	} else if c == '=' {
		newState = dfa_state.Equal
		token.Type = token_type.Assignment
		token.Text += s
	}

	fmt.Printf("after init: tokens[ %+v ]\n", tokens)
	return newState
}

func (token *Token) tokenize(c rune, state int) int {
	fmt.Printf("tokenize start: token[ %+v ]\n", token)
	s := string(c)
	// fmt.Printf("state: %+v\n", state)
	switch state {
	case dfa_state.Inital:
		state = token.init(c)
		break
	case dfa_state.Id:
		// fmt.Printf("Id token1: %+v\n", token)
		if utils.IsAlpha(c) || utils.IsDigit(c) {
			token.Text += s
			// fmt.Printf("Id token2: %+v\n", token)
		} else {
			state = token.init(c)
		}
		break
	case dfa_state.GT:
		if c == '=' {
			state = dfa_state.GE
			token.Type = token_type.GE
			token.Text += s
		} else {
			state = token.init(c)
		}
		break
	case dfa_state.GE:
	case dfa_state.Equal:
		state = token.init(c)
		break
	case dfa_state.IntLiteral:
		if utils.IsDigit(c) {
			token.Text += s
		} else {
			state = token.init(c)
		}
		break
	case dfa_state.Int1:
		if c == 'n' {
			state = dfa_state.Int2
			token.Text += s
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				token.Type = token_type.Identifier
				token.Text += s
			} else {
				state = token.init(c)
			}
		}
		break
	case dfa_state.Int2:
		if c == 't' {
			state = dfa_state.Int3
			token.Text += s
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				token.Type = token_type.Identifier
				token.Text += s
			} else {
				state = token.init(c)
			}
		}
		break
	case dfa_state.Int3:
		if c == ' ' {
			token.Text += s
			state = token.init(c)
		} else {
			if utils.IsAlpha(c) || utils.IsDigit(c) {
				state = dfa_state.Id
				token.Type = token_type.Identifier
				token.Text += s
			} else {
				state = token.init(c)
			}
		}
		break
	}
	return state
}
