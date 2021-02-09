package aluno

import (
	"errors"
	"regexp"
)

type Email struct {
	endereco string
}

func NewEmail(email string) (*Email, error) {
	if !regexp.MustCompile(`^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return nil, errors.New("Email invalido")
	}
	return &Email{endereco: email}, nil
}

func (e *Email) GetEndereco() string {
	return e.endereco
}
