package sharedaluno

import (
	"errors"
	"regexp"
)

type CPF struct {
	numero string
}

var ErrCPFInvalido = errors.New("CPF INVALIDO")

func NewCPF(numero string) (*CPF, error) {
	validCPF := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}\-\d{2}$`)

	if numero == "" || !validCPF.MatchString(numero) {
		return nil, ErrCPFInvalido
	}
	return &CPF{
		numero: numero,
	}, nil
}

func (c *CPF) GetNumero() string {
	return c.numero
}
