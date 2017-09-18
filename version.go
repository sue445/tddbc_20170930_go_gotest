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
	return &Version{Major: major, Minor: minor, Patch: patch}, nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Equal(other *Version) bool {
	return v.Major == other.Major && v.Minor == other.Minor && v.Patch == other.Patch
}
