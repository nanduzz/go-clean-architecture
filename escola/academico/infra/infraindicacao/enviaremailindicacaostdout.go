package infraindicacao

import (
	"log"

	"github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"
)

type EnviarEmailIndicacaoStdOut struct{}

func (e *EnviarEmailIndicacaoStdOut) enviarPara(aluno aluno.Aluno) error {
	log.Printf("Email sendo enviado para %s \n", aluno.GetCpf().GetNumero())
	return nil
}
