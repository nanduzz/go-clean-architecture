package aluno

import (
	"errors"
	"fmt"
	"regexp"
)

type Telefone struct {
	ddd    string
	numero string
}

func NewTelefone(ddd, numero string) (*Telefone, error) {

	if ddd == "" || numero == "" {
		return nil, errors.New("DDD e Número são obrigatórios")
	}

	if !regexp.MustCompile(`^\d{2}$`).MatchString(ddd) {
		return nil, fmt.Errorf("DDD inválido %s", ddd)
	}

	if !regexp.MustCompile(`^\d{8}|\d{9}$`).MatchString(numero) {
		return nil, errors.New("DDD inválido")
	}

	return &Telefone{
		ddd:    ddd,
		numero: numero,
	}, nil
}

func (t *Telefone) GetDDD() string {
	return t.ddd
}

func (t *Telefone) GetNumero() string {
	return t.numero
}
