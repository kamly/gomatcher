package matcher

type DFAMatcher struct {
	root *DFANode
}

func NewDFAMatcher() Matcher {
	return &DFAMatcher{
		root: newDFANode(),
	}
}

func (m *DFAMatcher) Build(words []string) error {
	// add words
	for _, item := range words {
		m.root.Insert(item)
	}
	return nil
}

func (m *DFAMatcher) Has(text string) bool {
	node := m.root
	for _, code := range text {
		if node.end == true {
			return true // 到达最后一个
		}
		value, ok := node.child[code] // 获取对应子节点
		if !ok {
			node = m.root // 重置
			continue
		}
		// 否则继续往后遍历
		node = value
	}
	if node.end == true {
		// 到达最后一个
		return true
	}
	return false
}
