package main

import (
	"fmt"

	"github.com/kamly/gomatcher"
)

func main() {
	words := []string{"golang语言", "十分好用"}

	// dfa
	{
		m := matcher.NewDFAMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言")) // true
	}

	// ac
	{
		m := matcher.NewAhoCorasickMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言"))
	}

	// 正则
	{
		m := matcher.NewRegMatcher()
		err := m.Build(words)
		if err != nil {
			return
		}
		fmt.Println(m.Has("golang语言"))
	}
}
