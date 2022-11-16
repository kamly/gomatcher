package main

import (
	"fmt"

	"github.com/kamly/gomatcher"
)

func main() {
	words := []string{"golang语言", "十分好用"}

	// dfa
	{
		m := gomatcher.NewDFAMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言")) // true
	}

	// ac
	{
		m := gomatcher.NewAhoCorasickMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言"))
	}

	// 正则
	{
		m := gomatcher.NewRegMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言"))
	}
}
