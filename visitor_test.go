package jsonfeed

import (
	"testing"
)

func TestRootValidation(t *testing.T) {
	feed := `
		{}
	`
	// Missing required properties `version` and `title`.
	var parser FeedParser
	root, err := parser.Parse([]byte(feed))
	if err != nil {
		t.Error(err)
	}
	visitor := NewValidationVisitor()
	root.Accept(visitor)
	if !(len(visitor.Errors()) == 2) {
		t.Error("visitor SHOULD have exactly 2 errors")
	}
}
