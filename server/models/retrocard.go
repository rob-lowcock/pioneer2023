package models

import (
	"errors"
	"strings"

	"gopkg.in/guregu/null.v4"
)

type Retrocard struct {
	ID          string    `jsonapi:"primary,retrocard"`
	Title       string    `jsonapi:"attr,title"`
	Column      int       `jsonapi:"attr,column"`
	Active      bool      `jsonapi:"attr,active"`
	Focus       bool      `jsonapi:"attr,focus"`
	DiscussedAt null.Time `jsonapi:"attr,discussed_at"`
}

func (r *Retrocard) Tidy() {
	r.Title = strings.TrimSpace(r.Title)
}

func (r *Retrocard) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
