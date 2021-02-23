package sharedaluno_test

import (
	"testing"

	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

func TestCPF(t *testing.T) {
	t.Run("nao deveria criar cpf com numeros invalidos", func(t *testing.T) {
		_, err := sharedaluno.NewCPF("")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
		_, err = sharedaluno.NewCPF("12345678900")
		if err == nil {
			t.Error("Expected an error, but didn't got one")
		}
	})

	t.Run("deve aceitar cpf com formato ###.###.###-##", func(t *testing.T) {
		cpf, err := sharedaluno.NewCPF("999.999.999-99")

		if err != nil {
			t.Fatalf("Expected no error, but got one %v", err)
		}

		if cpf == nil {
			t.Errorf("Expected to get a cpf, but didn't got one")
		}
	})
}
