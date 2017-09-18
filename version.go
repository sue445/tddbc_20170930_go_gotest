package tddbc

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(major int, minor int, patch int) *Version {
	return &Version{Major: major, Minor: minor, Patch: patch}
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
