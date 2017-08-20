package jsonfeed

import (
	"encoding/json"
)

type FeedParser struct{}

// Parse returns a parsed JSON Feed object.
func (p *FeedParser) Parse(instance []byte) (*Root, error) {
	var root Root
	if err := json.Unmarshal(instance, &root); err != nil {
		return nil, err
	}
	return &root, nil
}
