package tests

import (
	"testing"
	"unit-tests/internal/pengurangan"

	"github.com/stretchr/testify/assert"
)

func TestPengurangan(t *testing.T) {
	actualValue := pengurangan.Pengurangan(1, 2)
	expectedValue := -1
	assert.Equal(t, expectedValue, actualValue)
}
