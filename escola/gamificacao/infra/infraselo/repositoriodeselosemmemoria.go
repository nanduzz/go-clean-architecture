package infraselo

import (
	"github.com/nanduzz/go-clean-architecture/escola/gamificacao/dominio/selo"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type RepositorioDeSelosEmMemoria struct {
	selos map[sharedaluno.CPF][]*selo.Selo
}

func NewRepositorioDeSelosEmMemoria() *RepositorioDeSelosEmMemoria {
	return &RepositorioDeSelosEmMemoria{
		selos: make(map[sharedaluno.CPF][]*selo.Selo),
	}
}

func (r *RepositorioDeSelosEmMemoria) Adicionar(selo *selo.Selo) {
	r.selos[*selo.GetCpf()] = append(r.selos[*selo.GetCpf()], selo)
}

func (r *RepositorioDeSelosEmMemoria) SelosDoAlunoDeCPF(cpf sharedaluno.CPF) []*selo.Selo {
	if selos, ok := r.selos[cpf]; ok {
		return selos
	}
	return make([]*selo.Selo, 0)
}
