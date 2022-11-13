package matcher

type RegMatcher struct {
	root *RegNode
}

func NewRegMatcher() Matcher {
	return &RegMatcher{
		root: newRegNode(),
	}
}

func (r *RegMatcher) Build(words []string) error {
	for _, item := range words {
		r.root.Insert(item)
	}

	r.root.Compile()
	return nil
}

func (r *RegMatcher) Has(text string) bool {
	return r.root.compiler.MatchString(text)
}
