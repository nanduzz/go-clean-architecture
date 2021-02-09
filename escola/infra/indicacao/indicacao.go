package indicacao

import (
	"errors"
	"time"

	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

type Indicacao struct {
	indicado      *aluno.Aluno
	indicante     *aluno.Aluno
	dataIndicacao time.Time
}

var ErrIndicacaoInvalida = errors.New("indicação inválida!")

func NewIndicacao(indicado, indicante *aluno.Aluno) (*Indicacao, error) {
	if indicado == nil || indicante == nil {
		return nil, ErrIndicacaoInvalida
	}
	return &Indicacao{
		indicado:      indicado,
		indicante:     indicante,
		dataIndicacao: time.Now(),
	}, nil
}
