// @Author: Jiekun
// @Date: 2022/6/20

package parser

import (
	"github.com/jiekun/interpreter/ast"
	"github.com/jiekun/interpreter/lexer"
	"github.com/jiekun/interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
