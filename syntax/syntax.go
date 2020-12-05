package syntax

type SimpleASTNode struct {
	Type       string
	Text       string
	ChildNodes []*SimpleASTNode
}

func NewSimpleASTNode(typ, text string) *SimpleASTNode {
	return SimpleASTNode{
		Type: typ,
		Text: text,
	}
}

func additive(tokens TokenReader) SimpleASTNode {
}

func multiplicative(tokens TokenReader) SimpleASTNode {
}

func (sat *SimpleASTNode) addChild(node *SimpleASTNode) {
	sat.ChildNodes = append(sat.ChildNodes, node)
}
