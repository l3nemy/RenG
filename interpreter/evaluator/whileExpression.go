package evaluator

import (
	"RenG/interpreter/ast"
	"RenG/interpreter/object"
)

func evalWhileExpression(node *ast.WhileExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	for isTruthy(condition) {
		result := Eval(node.Body, env)
		if _, ok := result.(*object.ReturnValue); ok {
			return result
		}
		condition = Eval(node.Condition, env)
	}
	return nil
}
