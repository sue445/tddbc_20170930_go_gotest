package tddbc

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(major int, minor int, patch int) *Version {
	return &Version{Major: major, Minor: minor, Patch: patch}
}
