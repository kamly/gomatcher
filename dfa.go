package gomatcher

type DFANode struct {
	end   bool
	child map[rune]*DFANode
}

func newDFANode() *DFANode {
	return &DFANode{
		end:   false,
		child: make(map[rune]*DFANode),
	}
}

func (dfa *DFANode) Insert(word string) {
	curNode := dfa
	for _, v := range word {
		if curNode.child[v] == nil {
			curNode.child[v] = newDFANode()
		}
		curNode = curNode.child[v]
	}
	curNode.end = true
}
