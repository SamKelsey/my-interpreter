package parser

type NodeType int

const (
	NODE_TYPE_VARIABLE_DECLARATOR NodeType = iota
)

type aSTNode struct {
	Start    int
	End      int
	nodeType NodeType
}

func (n *aSTNode) GetNodeType() NodeType {
	return n.nodeType
}

type VariableDeclarationNode struct {
	aSTNode
	Identifier IdentifierNode
}

func NewVariableDeclarationNode() VariableDeclarationNode {
	return VariableDeclarationNode{
		aSTNode: aSTNode{
			nodeType: NODE_TYPE_VARIABLE_DECLARATOR,
		},
	}
}

type IdentifierNode struct {
	aSTNode
	name string
}

type LiteralNode struct {
	aSTNode
	value int
	raw   string
}
