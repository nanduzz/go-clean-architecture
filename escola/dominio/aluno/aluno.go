package aluno

import "errors"

var ErrMaxTelefones = errors.New("Número maximo de telefones já atingido!")

//AGGREGATE ROOT
type Aluno struct {
	cpf       *CPF
	nome      string
	email     *Email
	telefones []*Telefone
}

func (a *Aluno) AdicionarTelefone(ddd, numero string) error {
	if len(a.telefones) == 2 {
		return ErrMaxTelefones
	}

	telefone, err := NewTelefone(ddd, numero)
	if err != nil {
		return err
	}

	a.telefones = append(a.telefones, telefone)
	return nil
}

func NewAluno(nome string, cpf *CPF, email *Email) *Aluno {
	return &Aluno{
		nome:  nome,
		email: email,
		cpf:   cpf,
	}
}

func (a *Aluno) GetCpf() *CPF {
	return a.cpf
}
