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

func TestItemValidation(t *testing.T) {
	feed := `
		{
			"version": "https://jsonfeed.org/version/1",
			"title": "The Record",
			"items": [
				{
					"id": ""
				}
			]
		}
	`
	// Missing required properties `id` and `content_html` OR `content_text`.
	var parser FeedParser
	root, err := parser.Parse([]byte(feed))
	if err != nil {
		t.Error(err)
	}
	visitor := NewValidationVisitor()
	root.Accept(visitor)
	if !(len(visitor.Errors()) == 2) {
		t.Error("visitor SHOULD have exactly 1 error")
	}
}

func TestItemAttachments(t *testing.T) {
	feed := `
		{
			"version": "https://jsonfeed.org/version/1",
			"title": "The Record",
			"items": [
				{
					"id": "http://therecord.co/chris-parrish",
					"content_text": "Chris has worked at Adobe and as a founder of Rogue Sheep, which won an Apple Design Award for Postage. Chris’s new company is Aged & Distilled with Guy English — which shipped Napkin, a Mac app for visual collaboration. Chris is also the co-host of The Record. He lives on Bainbridge Island, a quick ferry ride from Seattle.",
					"attachments": [
						{
							"url": "",
							"mime_type": ""
						}
					]
				}
			]
		}
	`
	// Missing required properties `url` and `mime_type`.
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

// Test HasErrors method for ValidationVisitor.
func TestHasErrors(t *testing.T) {
	feed := `
		{
			"version": "https://jsonfeed.org/version/1"
		}
	`
	// Missing required property `title`.
	var parser FeedParser
	root, err := parser.Parse([]byte(feed))
	if err != nil {
		t.Error(err)
	}
	visitor := NewValidationVisitor()
	root.Accept(visitor)
	if !visitor.HasErrors() {
		t.Error("visitor SHOULD have exactly 1 error")
	}
}

func TestParserError(t *testing.T) {
	feed := `
		"version": "https://jsonfeed.org/version/1"
	`
	// Malformed JSON.
	var parser FeedParser
	if _, err := parser.Parse([]byte(feed)); err == nil {
		t.Error("error should NOT be nil, malformed JSON")
	}
}
