package aluno

import "github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"

type FabricaDeAluno struct {
	aluno *Aluno
}

func NewFabricaDeAluno() *FabricaDeAluno {
	return &FabricaDeAluno{
		aluno: &Aluno{},
	}
}

func (f *FabricaDeAluno) ComNomeCPFEMAIL(nome, numeroCPF, enderecoEmail string) (*FabricaDeAluno, error) {
	cpf, err := sharedaluno.NewCPF(numeroCPF)
	if err != nil {
		return nil, err
	}

	email, err := NewEmail(enderecoEmail)
	if err != nil {
		return nil, err
	}

	f.aluno = NewAluno(nome, cpf, email)
	return f, nil
}

func (f *FabricaDeAluno) ComTelefone(ddd, numero string) (*FabricaDeAluno, error) {
	err := f.aluno.AdicionarTelefone(ddd, numero)
	return f, err
}

func (f *FabricaDeAluno) Criar() *Aluno {
	return f.aluno
}
