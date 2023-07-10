package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(stage string) error {
	if err := u.format(stage); err != nil {
		return err
	}
	if err := u.validate(stage); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(stage string) error {
	if u.Name == "" {
		return errors.New("O campo nome é obrigatório")
	}

	if u.Nick == "" {
		return errors.New("O campo Nick é obrigatório")
	}

	if u.Email == "" {
		return errors.New("O campo E-mail é obrigatório")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("E-mail informado é inválido")
	}

	if stage == "register" && u.Password == "" {
		return errors.New("O campo Senha é obrigatório")
	}

	return nil
}

func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if stage == "register" {
		passwordHash, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = passwordHash
	}
	return nil
}

