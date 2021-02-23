package selo

import (
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type Selo struct {
	cpfDoAluno *sharedaluno.CPF
	nome       string
}

func NewSelo(cpf *sharedaluno.CPF, nome string) *Selo {
	return &Selo{
		cpfDoAluno: cpf,
		nome:       nome,
	}
}

func (s *Selo) GetCpf() *sharedaluno.CPF {
	return s.cpfDoAluno
}
