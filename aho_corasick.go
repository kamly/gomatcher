package gomatcher

import "container/list"

type AhoCorasickNode struct {
	count int
	fail  *AhoCorasickNode
	child map[rune]*AhoCorasickNode
}

func newAhoCorasickNode() *AhoCorasickNode {
	return &AhoCorasickNode{
		count: 0,
		fail:  nil,
		child: make(map[rune]*AhoCorasickNode),
	}
}

func (ac *AhoCorasickNode) Insert(word string) {
	curNode := ac
	for _, v := range word {
		if curNode.child[v] == nil {
			curNode.child[v] = newAhoCorasickNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
}

func (ac *AhoCorasickNode) InsertFailed() {
	ll := list.New()
	ll.PushBack(ac)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*AhoCorasickNode)
		var p *AhoCorasickNode = nil
		for i, v := range temp.child {
			if temp == ac {
				v.fail = ac
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = ac
				}
			}
			ll.PushBack(v)
		}
	}
}
