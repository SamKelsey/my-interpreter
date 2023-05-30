package parser

import "interpreter/token"

/*
Language Grammar

expr 		--> NUM exprOpts
exprOpts 	--> '+' NUM

AddExpr     --> ExprOpts '+' integer

EquateExpr  --> Expr '==' Expr
Expr        --> Identifier | string | integer | boolean
*/

/*
AST Examples

1.
var x = 2

AST = {
	type: "Program",
	start: 0,
	end: ...,
	body: [
		VariableDeclarator: {
			type: VariableDeclarator,
			start: 0,
			end: 9,
			id: Identifier{
				type: Identifier,
				start: 4,
				end: 5,
				name: "x"
			}
			value: Value{
				type: Number,
				start: 8,
				end: 9
				raw: "2"
			}
		}
	]
}
*/

func Parse(tokens []*token.Token) aSTNode {
	// TODO: implement me
	return aSTNode{}
}
