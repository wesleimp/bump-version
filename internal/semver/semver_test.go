package semver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidSemver(t *testing.T) {
	require := require.New(t)

	version, err := Parse("v1.2.3-rc1")
	require.NoError(err)
	require.Equal(Semver{
		Major:      1,
		Minor:      2,
		Patch:      3,
		Prerelease: "rc1",
	}, version)
}

func TestInvalidSemver(t *testing.T) {
	_, err := Parse("abcv1.2.3-rc1")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to parse version 'abcv1.2.3-rc1' as semver")
}

func TestFindPreversion(t *testing.T) {
	require := require.New(t)

	prefix, preversion := findPreversion("alpha2")
	require.Equal(prefix, "alpha")
	require.Equal(preversion, 2)

	prefix, preversion = findPreversion("beta1")
	require.Equal(prefix, "beta")
	require.Equal(preversion, 1)

	prefix, preversion = findPreversion("rc1")
	require.Equal(prefix, "rc")
	require.Equal(preversion, 1)

	prefix, preversion = findPreversion("alpha")
	require.Equal(prefix, "alpha")
	require.Zero(preversion)

	prefix, preversion = findPreversion("")
	require.Empty(prefix)
	require.Zero(preversion)
}

func TestMajor(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "major"), Semver{Major: 3, Minor: 0, Patch: 0, Prerelease: ""})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "major"), Semver{Major: 3, Minor: 0, Patch: 0, Prerelease: ""})
}

func TestFeature(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "feature"), Semver{Major: 2, Minor: 12, Patch: 0, Prerelease: ""})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "feature"), Semver{Major: 2, Minor: 12, Patch: 0, Prerelease: ""})
}

func TestBug(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "bug"), Semver{Major: 2, Minor: 11, Patch: 8, Prerelease: ""})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "bug"), Semver{Major: 2, Minor: 11, Patch: 8, Prerelease: ""})
}

func TestAlpha(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "alpha"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha1"})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "alpha"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha4"})
}

func TestBeta(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "beta"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "beta1"})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "beta"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "beta1"})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "beta3"}
	require.Equal(Bump(version, "beta"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "beta4"})
}

func TestRc(t *testing.T) {
	require := require.New(t)

	version := Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: ""}
	require.Equal(Bump(version, "rc"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "rc1"})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "alpha3"}
	require.Equal(Bump(version, "rc"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "rc1"})

	version = Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "rc3"}
	require.Equal(Bump(version, "rc"), Semver{Major: 2, Minor: 11, Patch: 7, Prerelease: "rc4"})
}
