package models

import (
	"errors"
	"strings"
	"time"
)

type Publish struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (p *Publish) Prepare() error {
	p.format()
	if err := p.validate(); err != nil {
		return err
	}
	return nil
}

func (p *Publish) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}

func (p *Publish) validate() error {
	if p.Title == "" {
		return errors.New("Campo título é obrigatório")
	}
	if p.Content == "" {
		return errors.New("Campo conteúdo é obrigatório")
	}
	return nil
}
