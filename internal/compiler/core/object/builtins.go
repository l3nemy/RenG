package object

import "fmt"

var FunctionBuiltins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return runtimeError("wrong number of arguments. got=%d", len(args))
				}

				switch arg := args[0].(type) {
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				default:
					return runtimeError("arguments to len not supported, got=%s", args[0].Type())
				}
			},
		},
	},
	{
		"print",
		&Builtin{
			Fn: func(args ...Object) Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}

				return nil
			},
		},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range FunctionBuiltins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}
