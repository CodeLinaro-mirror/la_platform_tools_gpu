package gles

import (
	"fmt"
	"regexp"
	"strconv"
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
