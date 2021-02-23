package selo

import "github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"

type RepositorioDeSelos interface {
	Adicionar(selo *Selo)
	SelosDoAlunoDeCPF(cpf sharedaluno.CPF) []*Selo
}
