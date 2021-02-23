package aluno_test

import (
	"testing"

	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

func TestEmeil(t *testing.T) {

	t.Run("Deveria permitir criar email com endereco Validos", func(t *testing.T) {
		endereco := "fernando@fernando.com.br"
		email, err := aluno.NewEmail(endereco)
		if err != nil {
			t.Fatalf("Erro inesperado, %v", err)
		}

		if endereco != email.GetEndereco() {
			t.Errorf("Esperava %s, mas obteve %s", endereco, email.GetEndereco())
		}
	})

	t.Run("Não deveria criar emails com enderecos invalidos", func(t *testing.T) {
		_, err := aluno.NewEmail("")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
		_, err = aluno.NewEmail("email invalido")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}

	})
}
