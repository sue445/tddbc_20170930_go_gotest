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

func TestVersion_Equal_NotEquals(t *testing.T) {
	version1 := NewVersion(1, 4, 2)
	version2 := NewVersion(2, 1, 0)

	assert.False(t, version1.Equal(version2))
}

func TestVersion_Equal_Equals(t *testing.T) {
	version1 := NewVersion(1, 4, 2)
	version2 := NewVersion(1, 4, 2)

	assert.True(t, version1.Equal(version2))
}
