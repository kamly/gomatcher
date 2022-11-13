package matcher

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())

}

func TestDFAMatcher(t *testing.T) {
	m := NewDFAMatcher()
	_ = m.Build(words)

	var str string

	str = "golang语言"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "golang语言好"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "赞啊，golang语言"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "赞啊，go语言"
	if m.Has(str) == true {
		t.Fatal(str)
	}

	str = "赞啊，golang语言好，十分好用"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "赞啊，golang，十分好用"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "赞啊，golang语言好，十分好"
	if m.Has(str) == false {
		t.Fatal(str)
	}

	str = "赞啊，go语言，十分好"
	if m.Has(str) == true {
		t.Fatal(str)
	}
}

func BenchmarkDFAMatcher_Has(b *testing.B) {
	m := NewDFAMatcher()
	_ = m.Build(words)

	for idx := 1; idx < b.N; idx++ {
		m.Has(hasSensStr)
	}
}

func BenchmarkDFAMatcher_No(b *testing.B) {
	m := NewDFAMatcher()
	_ = m.Build(words)

	for idx := 1; idx < b.N; idx++ {
		m.Has(noSensStr)
	}
}

func BenchmarkDFAMatcher_ManyWords_Has(b *testing.B) {
	m := NewDFAMatcher()
	_ = m.Build(wordsMany)

	//fmt.Println(m.Has(wordsManyHasSensStr))
	for idx := 1; idx < b.N; idx++ {
		m.Has(wordsManyHasSensStr)
	}
}

func BenchmarkDFAMatcher_ManyWords_No(b *testing.B) {
	m := NewDFAMatcher()
	_ = m.Build(wordsMany)

	//fmt.Println(m.Has(wordsManyNoSensStr))
	for idx := 1; idx < b.N; idx++ {
		m.Has(wordsManyNoSensStr)
	}
}
