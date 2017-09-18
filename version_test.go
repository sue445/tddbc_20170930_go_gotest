package tddbc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVersion(t *testing.T) {
	version := NewVersion(1, 4, 2)

	assert.Equal(t, 1, version.Major)
	assert.Equal(t, 4, version.Minor)
	assert.Equal(t, 2, version.Patch)
}

func TestVersion_String(t *testing.T) {
	version := NewVersion(1, 4, 2)

	assert.Equal(t, "1.4.2", version.String())
	assert.Equal(t, "1.4.2", fmt.Sprintf("%s", version))
}
