package main

import (
	"fmt"
	"log"

	"github.com/nanduzz/go-clean-architecture/escola/aplicacao/aluno/matricular"
	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
	"github.com/nanduzz/go-clean-architecture/escola/infra/infraaluno"
)

func main() {
	nome := "fernando"
	numeroCPF := "123.456.789-00"
	email := "fernando@fernando.com"

	repositorio := infraaluno.NewRepositorioDeAlunosEmMemoria()
	matricularaluno := matricular.NewMatricularAluno(repositorio)

	err := matricularaluno.Executa(matricular.NewMatricularAlunoDTO(nome, numeroCPF, email))
	if err != nil {
		log.Fatal(err)
	}

	cpf, err := aluno.NewCPF(numeroCPF)
	if err != nil {
		log.Fatal(err)
	}
	aluno, err := repositorio.BuscaPorCPF(*cpf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(aluno)
}
