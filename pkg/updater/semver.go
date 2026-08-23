package updater

import (
	"strings"

	"golang.org/x/mod/semver"
)

// NormalizeVersion trims leading 'v' or 'V' and whitespace.
func NormalizeVersion(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimPrefix(v, "V")
	return v
}

// toCanonicalSemver converts a version string to a valid format for x/mod/semver (prefixed with 'v').
// It expands partial versions (e.g. "0.1" -> "v0.1.0") to adhere to full SemVer 2.0.0.
func toCanonicalSemver(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return ""
	}
	if strings.HasPrefix(v, "v") || strings.HasPrefix(v, "V") {
		v = "v" + v[1:]
	} else {
		v = "v" + v
	}

	if semver.IsValid(v) {
		return v
	}

	// Expand partial versions like "v0.1" or "v1"
	parts := strings.SplitN(v, "-", 2)
	dots := strings.Split(parts[0], ".")
	if len(dots) == 2 {
		expanded := dots[0] + "." + dots[1] + ".0"
		if len(parts) > 1 {
			expanded += "-" + parts[1]
		}
		if semver.IsValid(expanded) {
			return expanded
		}
	} else if len(dots) == 1 {
		expanded := dots[0] + ".0.0"
		if len(parts) > 1 {
			expanded += "-" + parts[1]
		}
		if semver.IsValid(expanded) {
			return expanded
		}
	}

	return v
}

// CompareVersions compares two semver strings (e.g. "1.2.0" and "1.3.0").
// It adheres to Semantic Versioning 2.0.0 via golang.org/x/mod/semver, while
// supporting development tags ("dev") and partial/unprefixed versions.
// Returns -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2.
func CompareVersions(v1, v2 string) int {
	norm1 := NormalizeVersion(v1)
	norm2 := NormalizeVersion(v2)

	if norm1 == norm2 {
		return 0
	}

	// Handle dev / non-semver
	if norm1 == "dev" {
		return -1
	}
	if norm2 == "dev" {
		return 1
	}

	sv1 := toCanonicalSemver(v1)
	sv2 := toCanonicalSemver(v2)

	valid1 := semver.IsValid(sv1)
	valid2 := semver.IsValid(sv2)

	if valid1 && valid2 {
		return semver.Compare(sv1, sv2)
	}

	if valid1 && !valid2 {
		return 1
	}
	if !valid1 && valid2 {
		return -1
	}

	if norm1 < norm2 {
		return -1
	}
	return 1
}
