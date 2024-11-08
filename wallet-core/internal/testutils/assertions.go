package testutils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// IsUUID verifica se a string fornecida é um UUID válido.
func IsUUID(t *testing.T, id string) {
	t.Helper() // Marca esta função como auxiliar para relatórios de teste
	_, err := uuid.Parse(id)
	assert.NoError(t, err, "ID should be a valid UUID")
}
