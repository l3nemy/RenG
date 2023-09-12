package compiler

import (
	"fmt"
	"internal/compiler/core/code"
	"internal/compiler/core/object"
)

func (c *Compiler) loadSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.emit(code.OpGetGlobal, s.Index)
	case LocalScope:
		c.emit(code.OpGetLocal, s.Index)
	case BuiltinScope:
		c.emit(code.OpGetBuiltin, s.Index)
	}
}

func (c *Compiler) ReplaceSymbol() error {
	for _, inform := range c.reservationSymbol {

		// TODO: index out of range
		fn := c.constants[inform.ReplaceFuncIndex]
		fnObj, ok := fn.(*object.CompiledFunction)
		if !ok {
			continue
		}

		//TODO: See this part
		s, ok := c.symbolTable.Resolve(inform.symbol)
		if !ok {
			return fmt.Errorf("the symbol %s is not found", inform.symbol)
		}

		op := code.Make(code.OpGetGlobal, s.Index)

		for n := 0; n < len(op); n++ {
			fnObj.Instructions[inform.pos+n] = op[n]
		}
		c.constants[inform.ReplaceFuncIndex] = fnObj
	}

	return nil
}
