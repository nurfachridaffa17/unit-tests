package tests

import (
	"testing"
	"unit-tests/internal/pertambahan"

	"github.com/stretchr/testify/assert"
)

func TestPertambahan(t *testing.T) {
	actualValue := pertambahan.Pertambahan(1, 2)
	expectedValue := 3
	assert.Equal(t, expectedValue, actualValue)
}
