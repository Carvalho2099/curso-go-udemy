package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Teste", "Tes", 0},
		{"Teste", "es", 1},
		{"Teste", "st", 2},
		{"Teste", "te", -1},
	}

	for _, teste := range testes {
		t.Logf("Teste: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
