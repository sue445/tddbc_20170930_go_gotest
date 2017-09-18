package tddbc

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(major int, minor int, patch int) (*Version, error) {
	if major < 0 {
		return nil, fmt.Errorf("Expected majar >= 0, but major is %d", major)
	}
	if minor < 0 {
		return nil, fmt.Errorf("Expected minor >= 0, but minor is %d", minor)
	}
	if patch < 0 {
		return nil, fmt.Errorf("Expected patch >= 0, but patch is %d", patch)
	}

	return &Version{Major: major, Minor: minor, Patch: patch}, nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Equal(other *Version) bool {
	return v.Major == other.Major && v.Minor == other.Minor && v.Patch == other.Patch
}

func (v *Version) BumpPatchVersion() {
	v.Patch++
}

func (v *Version) BumpMinorVersion() {
	v.Patch = 0
	v.Minor++
}

func (v *Version) BumpMajorVersion() {
}
