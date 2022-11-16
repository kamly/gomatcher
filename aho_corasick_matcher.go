package gomatcher

// AhoCorasickMatcher
type AhoCorasickMatcher struct {
	root *AhoCorasickNode
}

func NewAhoCorasickMatcher() Matcher {
	return &AhoCorasickMatcher{
		root: newAhoCorasickNode(),
	}
}

func (m *AhoCorasickMatcher) Build(words []string) error {
	// add words
	for _, item := range words {
		m.root.Insert(item)
	}

	// add failed
	m.root.InsertFailed()
	return nil
}

func (m *AhoCorasickMatcher) Has(text string) bool {
	curNode := m.root
	var p *AhoCorasickNode = nil
	for _, v := range text {
		for curNode.child[v] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = m.root
		}
		p = curNode
		for p != m.root && p.count > 0 {
			return true
		}
	}
	return false
}
