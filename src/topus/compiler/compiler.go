package compile

import (
	"topus/src/topus/lexer"
	"topus/src/topus/token"
)

package compile

import (
	"fmt"
)

type Compiler struct {
	l *lexer.Lexer			// our lexer 
	curToken token.Token    // current token
	peekToken token.Token   // next token
	bytecode []byte			// generated bytecode
	labels map[string]int	// holder for labels
	fixups map[int]string	// holder for fixups
}

func New(l *lexer.Lexer) *Compiler {
	p := &Compiler{l: l}
	p.labels = make(map[string]int)
	p.fixups = make(map[int]string)

	// prime the pump
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Compiler) Compile() {
	
}