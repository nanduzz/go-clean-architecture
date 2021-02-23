package aluno_test

import (
	"testing"

	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

func TestTelefone(t *testing.T) {
	t.Run("não deveria criar telefones com ddd invalido", func(t *testing.T) {
		_, err := aluno.NewTelefone("", "12345678")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
		_, err = aluno.NewTelefone("432", "12345678")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
		_, err = aluno.NewTelefone("1", "12345678")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
	})
	t.Run("não deveria criar telefones com numerps invalido", func(t *testing.T) {
		_, err := aluno.NewTelefone("11", "")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
		_, err = aluno.NewTelefone("11", "12345")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
	})
	t.Run("deveria criar telefones valido com 8 digitos", func(t *testing.T) {
		ddd := "44"
		numero := "12345678"
		telefone, err := aluno.NewTelefone(ddd, numero)
		if err != nil {
			t.Errorf("Didn't expected an error but got one, %v", err)
		}

		if telefone.GetDDD() != ddd {
			t.Errorf("Expected %s, but got %s", ddd, telefone.GetDDD())
		}

		if telefone.GetNumero() != numero {
			t.Errorf("Expected %s, but got %s", ddd, telefone.GetNumero())
		}

	})
}
