package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty`
	AuthorNick string    `json:"author_nick,omitempty`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:created_at,omitempty`
}

func (publication *Publication) Prepare() error {
	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.formate()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("title is required")
	}

	if publication.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (publication *Publication) formate() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
