package main

type Component interface {
	Accept(visitor Visitor, arguments ...interface{})
}

type Root struct {
	Version     string        `json:"version"`
	Title       string        `json:"title"`
	HomePageURL string        `json:"home_page_url"`
	FeedURL     string        `json:"feed_url"`
	Description string        `json:"description"`
	UserComment string        `json:"user_comment"`
	NextURL     string        `json:"next_url"`
	Icon        string        `json:"icon"`
	Favicon     string        `json:"favicon"`
	Author      *Author       `json:"author"`
	Expired     bool          `json:"expired"`
	Hubs        []interface{} `json:"hubs"`
	Items       Items         `json:"items"`
}

func (r *Root) Accept(visitor Visitor, arguments ...interface{}) {
	visitor.VisitRoot(r)
}

// Author specifies the feed author.
type Author struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Avatar string `json:"avatar"`
}

// Items describe each object in the list.
type Items []*Item

func (items Items) Accept(visitor Visitor, arguments ...interface{}) {
	visitor.VisitItems(items)
}

type Item struct {
	ID            string      `json:"id"`
	URL           string      `json:"url"`
	ExternalURL   string      `json:"external_url"`
	Title         string      `json:"title"`
	ContentHTML   string      `json:"content_html"`
	ContentText   string      `json:"content_text"`
	Summary       string      `json:"summary"`
	Image         string      `json:"image"`
	BannerImage   string      `json:"banner_image"`
	DatePublished string      `json:"date_published"`
	DateModified  string      `json:"date_modified"`
	Author        *Author     `json:"author"`
	Tags          []string    `json:"tags"`
	Attachments   Attachments `json:"attachments"`
}

func (item *Item) Accept(visitor Visitor, arguments ...interface{}) {
	visitor.VisitItem(item, arguments)
}

type Attachments []*Attachment

func (attachments Attachments) Accept(visitor Visitor, arguments ...interface{}) {
	visitor.VisitAttachments(attachments, arguments)
}

type Attachment struct {
	URL               string  `json:"url"`
	MimeType          string  `json:"mime_type"`
	Title             string  `json:"title"`
	SizeInBytes       float64 `json:"size_in_bytes"`
	DurationInSeconds float64 `json:"duration_in_seconds"`
}

func (attachment *Attachment) Accept(visitor Visitor, arguments ...interface{}) {
	visitor.VisitAttachment(attachment, arguments)
}

type Extension struct{}
