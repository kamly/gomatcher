package gomatcher

// Matcher 匹配
type Matcher interface {
	// Build 构建
	Build(words []string) error

	// Has 是否存在
	Has(text string) bool
}
