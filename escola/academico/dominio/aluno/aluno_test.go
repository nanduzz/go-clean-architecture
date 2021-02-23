package aluno_test

import (
	"testing"

	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

func TestMaxTelefone(t *testing.T) {
	t.Run("Feve permitir dois telefones", func(t *testing.T) {

		fabrica, _ := iniciaFabrica(t)

		fabrica, err := fabrica.ComTelefone("44", "44444444")
		if err != nil {
			t.Errorf("Didn't expected error, but got %q", err)
		}

		fabrica, err = fabrica.ComTelefone("44", "44444445")

		if err != nil {
			t.Errorf("Didn't expected error, but got %q", err)
		}

	})

	t.Run("Não deve permitir mais do que dois telefones", func(t *testing.T) {

		fabrica, _ := iniciaFabrica(t)

		fabrica, _ = fabrica.ComTelefone("44", "44444444")
		fabrica, _ = fabrica.ComTelefone("44", "44444445")
		fabrica, err := fabrica.ComTelefone("44", "44444446")

		if err != aluno.ErrMaxTelefones {
			t.Errorf("Expected %q, but got %q", aluno.ErrMaxTelefones, err)
		}

	})
}

func iniciaFabrica(t *testing.T) (*aluno.FabricaDeAluno, error) {
	t.Helper()
	return aluno.NewFabricaDeAluno().ComNomeCPFEMAIL("teste", "111.111.111-11", "teste@teste.com")
}
