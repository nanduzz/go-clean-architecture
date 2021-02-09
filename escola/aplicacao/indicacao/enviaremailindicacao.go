package indicacao

import "github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"

type EnviarEmailIndicacao interface {
	enviarPara(aluno aluno.Aluno) error
}
