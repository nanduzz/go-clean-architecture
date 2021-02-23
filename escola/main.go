package main

import (
	"fmt"
	"log"

	"github.com/nanduzz/go-clean-architecture/escola/academico/aplicacao/aluno/matricular"
	"github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"
	"github.com/nanduzz/go-clean-architecture/escola/academico/infra/infraaluno"
	"github.com/nanduzz/go-clean-architecture/escola/gamificacao/aplicacao/geraselo"
	"github.com/nanduzz/go-clean-architecture/escola/gamificacao/infra/infraselo"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/evento"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

func main() {
	nome := "fernando"
	numeroCPF := "123.456.789-00"
	email := "fernando@fernando.com"

	repositorio := infraaluno.NewRepositorioDeAlunosEmMemoria()
	publicador := evento.NewPublicadorEventos()

	repositorioSelos := infraselo.NewRepositorioDeSelosEmMemoria()

	publicador.Adicionar(aluno.NewLogDeAlunoMatriculado().AbstractOuvinte)
	publicador.Adicionar(geraselo.NewGeraSeloAlunoNovato(repositorioSelos).AbstractOuvinte)

	matricularaluno := matricular.NewMatricularAluno(repositorio, publicador)

	err := matricularaluno.Executa(matricular.NewMatricularAlunoDTO(nome, numeroCPF, email))
	if err != nil {
		log.Fatal(err)
	}

	cpf, err := sharedaluno.NewCPF(numeroCPF)
	if err != nil {
		log.Fatal(err)
	}
	aluno, err := repositorio.BuscaPorCPF(*cpf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(aluno)
}
