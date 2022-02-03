package semver

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/Masterminds/semver/v3"
)

// Semver is a parsed semantic version.
type Semver struct {
	Major      uint64
	Minor      uint64
	Patch      uint64
	Prerelease string
}

func (s Semver) Print() string {
	version := fmt.Sprintf("%v.%v.%v", s.Major, s.Minor, s.Patch)

	if s.Prerelease != "" {
		version = fmt.Sprintf("%s-%s", version, s.Prerelease)
	}

	return version
}

// Bump a version accordingly the fragment
func Bump(sv Semver, fragment string) Semver {
	switch fragment {
	case "major":
		return Semver{
			Major:      sv.Major + 1,
			Minor:      0,
			Patch:      0,
			Prerelease: "",
		}
	case "feature":
		return Semver{
			Major:      sv.Major,
			Minor:      sv.Minor + 1,
			Patch:      0,
			Prerelease: "",
		}
	case "bug":
		return Semver{
			Major:      sv.Major,
			Minor:      sv.Minor,
			Patch:      sv.Patch + 1,
			Prerelease: "",
		}
	case "alpha":
		prefix, preversion := findPreversion(sv.Prerelease)
		if prefix != "alpha" {
			preversion = 0
		}

		return Semver{
			Major:      sv.Major,
			Minor:      sv.Minor,
			Patch:      sv.Patch,
			Prerelease: fmt.Sprintf("alpha%v", preversion+1),
		}
	case "beta":
		prefix, preversion := findPreversion(sv.Prerelease)
		if prefix != "beta" {
			preversion = 0
		}

		return Semver{
			Major:      sv.Major,
			Minor:      sv.Minor,
			Patch:      sv.Patch,
			Prerelease: fmt.Sprintf("beta%v", preversion+1),
		}
	case "rc":
		prefix, preversion := findPreversion(sv.Prerelease)
		if prefix != "rc" {
			preversion = 0
		}

		return Semver{
			Major:      sv.Major,
			Minor:      sv.Minor,
			Patch:      sv.Patch,
			Prerelease: fmt.Sprintf("rc%v", preversion+1),
		}
	default:
		return sv
	}
}

func findPreversion(pre string) (string, int) {
	var prefix string
	preversion := 0

	p := regexp.MustCompile("([a-z]+)").FindAllString(pre, 1)
	if len(p) != 0 {
		prefix = p[0]
	}

	pv := regexp.MustCompile("([0-9]+)").FindAllString(pre, 1)
	if len(pv) != 0 {
		i, err := strconv.Atoi(pv[0])
		if err != nil {
			log.Fatal(err)
		}

		preversion = i
	}

	return prefix, preversion
}

// Parse semantic version.
func Parse(version string) (Semver, error) {
	sv, err := semver.NewVersion(version)
	if err != nil {
		return Semver{}, fmt.Errorf("failed to parse version '%s' as semver: %w", version, err)
	}

	return Semver{
		Major:      sv.Major(),
		Minor:      sv.Minor(),
		Patch:      sv.Patch(),
		Prerelease: sv.Prerelease(),
	}, nil
}
