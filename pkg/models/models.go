package models

import (
	"errors"
	"time"
)

// ErrNoRecord is for when an record isn't found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is the model of our data
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
