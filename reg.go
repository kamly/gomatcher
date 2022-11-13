package matcher

import (
	"fmt"
	"regexp"
)

type RegNode struct {
	compiler *regexp.Regexp
	str      string
}

func newRegNode() *RegNode {
	return &RegNode{
		compiler: nil,
		str:      "",
	}
}

func (reg *RegNode) Insert(word string) {
	if reg.str == "" {
		reg.str = word
	} else {
		reg.str = reg.str + "|" + word
	}
}

func (reg *RegNode) Compile() {
	//解析正则表达式，如果成功返回解释器
	regCompile := regexp.MustCompile(reg.str)
	if regCompile == nil {
		fmt.Println("regexp err")
		return
	}
	reg.compiler = regCompile
}
