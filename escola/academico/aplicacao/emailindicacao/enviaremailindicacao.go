package indicacao

import "github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"

type EnviarEmailIndicacao interface {
	enviarPara(aluno aluno.Aluno) error
}
