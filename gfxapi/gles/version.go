package gles

import (
	"fmt"
	"regexp"
	"strconv"

	"android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"
)

// Version represents the GL version major and minor numbers,
// and whether its flavour is ES, as opposed to Desktop GL.
type Version struct {
	IsES  bool
	Major int
	Minor int
}

var versionRe = regexp.MustCompile(`^(OpenGL ES.*? )?(\d+)\.(\d+).*`)

// ParseVersion parses the GL version major, minor and flavour from the output of glGetString(GL_VERSION).
func ParseVersion(str string) (*Version, error) {
	if match := versionRe.FindStringSubmatch(str); match != nil {
		isES := len(match[1]) > 0 // Desktop GL doesn't have a flavour prefix.
		major, _ := strconv.Atoi(match[2])
		minor, _ := strconv.Atoi(match[3])
		return &Version{IsES: isES, Major: major, Minor: minor}, nil
	}
	return nil, fmt.Errorf("Unknown GL_VERSION format: %s", str)
}

// GLSLVersion returns the highest supported GLSL version for the given GL version.
func GLSLVersion(glVersion string) (glsl.Version, error) {
	v, err := ParseVersion(glVersion)
	if err != nil {
		return glsl.Version{}, err
	}
	major, minor, isES := v.Major, v.Minor, v.IsES
	switch {
	case major == 2 && isES:
		return glsl.Version{Major: 1, Minor: 0}, nil
	case major == 3 && isES:
		return glsl.Version{Major: 3, Minor: 0}, nil

	case major == 2 && minor == 0 && !isES:
		return glsl.Version{Major: 1, Minor: 1}, nil
	case major == 2 && minor == 1 && !isES:
		return glsl.Version{Major: 1, Minor: 2}, nil
	case major == 3 && minor == 0 && !isES:
		return glsl.Version{Major: 1, Minor: 3}, nil
	case major == 3 && minor == 1 && !isES:
		return glsl.Version{Major: 1, Minor: 4}, nil
	case major == 3 && minor == 2 && !isES:
		return glsl.Version{Major: 1, Minor: 5}, nil

	default:
		return glsl.Version{Major: major, Minor: minor}, nil
	}
}
