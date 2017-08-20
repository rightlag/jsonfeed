package jsonfeed

import (
	"errors"
	"fmt"
)

type Errors []error

func (errs *Errors) Append(err error) {
	*errs = append(*errs, err)
}

type Visitor interface {
	VisitRoot(r *Root, arguments ...interface{})
	VisitItems(items Items, arguments ...interface{})
	VisitItem(item *Item, arguments ...interface{})
	VisitAttachments(attachments Attachments, arguments ...interface{})
	VisitAttachment(attachment *Attachment, arguments ...interface{})
}

type ValidationVisitor struct {
	errors Errors
}

func (v *ValidationVisitor) Errors() Errors {
	return v.errors
}

func (v *ValidationVisitor) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *ValidationVisitor) VisitRoot(root *Root, arguments ...interface{}) {
	if root.Version == "" {
		v.errors.Append(errors.New("missing required property `version`"))
	}
	if root.Title == "" {
		v.errors.Append(errors.New("missing required property `title`"))
	}
	root.Items.Accept(v)
}

func (v *ValidationVisitor) VisitItems(items Items, arguments ...interface{}) {
	for i, item := range items {
		item.Accept(v, (i + 1))
	}
}

func (v *ValidationVisitor) VisitItem(item *Item, arguments ...interface{}) {
	index := arguments[0].([]interface{})[0].(int)
	if item.ID == "" {
		v.errors.Append(fmt.Errorf("item %d missing required property `id`", index))
	}
	if item.ContentHTML == "" && item.ContentText == "" {
		v.errors.Append(fmt.Errorf("item %d is missing one of required properties `content_html` or `content_text`", index))
	}
	item.Attachments.Accept(v, index)
}

func (v *ValidationVisitor) VisitAttachments(attachments Attachments, arguments ...interface{}) {
	index := arguments[0].([]interface{})[0].(int)
	for i, attachment := range attachments {
		attachment.Accept(v, index, (i + 1))
	}
}

func (v *ValidationVisitor) VisitAttachment(attachment *Attachment, arguments ...interface{}) {
	subArguments := arguments[0].([]interface{})
	itemIndex := subArguments[0].(int)
	attachmentIndex := subArguments[1].(int)
	if attachment.URL == "" {
		v.errors.Append(fmt.Errorf("item %d attachment %d missing required property `url`", itemIndex, attachmentIndex))
	}
	if attachment.MimeType == "" {
		v.errors.Append(fmt.Errorf("item %d attachment %d missing required property `mime_type`", itemIndex, attachmentIndex))
	}
}

func NewValidationVisitor() *ValidationVisitor {
	return &ValidationVisitor{make([]error, 0)}
}
