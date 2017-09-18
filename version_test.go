package tddbc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVersion(t *testing.T) {
	version, err := NewVersion(1, 4, 2)

	assert.NoError(t, err)

	assert.Equal(t, 1, version.Major)
	assert.Equal(t, 4, version.Minor)
	assert.Equal(t, 2, version.Patch)
}

type newVersionTest struct {
	major int
	minor int
	patch int
}

var newVersionTests = []newVersionTest{
	{major: -1, minor: 1, patch: 1},
	{major: 1, minor: -1, patch: 1},
	{major: 1, minor: 1, patch: -1},
}

func TestNewVersion_error(t *testing.T) {
	for _, test := range newVersionTests {
		_, err := NewVersion(test.major, test.minor, test.patch)

		assert.Error(t, err, fmt.Sprintf("[major=%d, minor=%d, patch=%d] Expected error, actual no error", test.major, test.minor, test.patch))
	}
}

func TestVersion_String(t *testing.T) {
	version, err := NewVersion(1, 4, 2)

	assert.NoError(t, err)

	assert.Equal(t, "1.4.2", version.String())
	assert.Equal(t, "1.4.2", fmt.Sprintf("%s", version))
}

func TestVersion_Equal_NotEquals(t *testing.T) {
	version1, err := NewVersion(1, 4, 2)

	assert.NoError(t, err)

	version2, err := NewVersion(2, 1, 0)

	assert.NoError(t, err)

	assert.False(t, version1.Equal(version2))
}

func TestVersion_Equal_Equals(t *testing.T) {
	version1, err := NewVersion(1, 4, 2)

	assert.NoError(t, err)

	version2, err := NewVersion(1, 4, 2)

	assert.NoError(t, err)

	assert.True(t, version1.Equal(version2))
}
