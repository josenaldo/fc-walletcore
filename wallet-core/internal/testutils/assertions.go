package testutils

import (
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func IsULID(t *testing.T, id string) {
	t.Helper()
	_, err := ulid.Parse(id)
	assert.NoError(t, err, "ID should be a valid ULID")
}
