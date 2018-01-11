//
// Bet you'll be sick of the word path after reading this documentation.
//

package fileloc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PathSet represents a set of paths to check for a file
type PathSet struct {
	pathSet []string
}

// New creates an empty PathSet
func New() *PathSet {
	p := &PathSet{}
	return p
}

// AppendFromPathString appends to the pathset with paths from a path string (colon-separated string)
func (p *PathSet) AppendFromPathString(pathString string) error {
	paths := strings.Split(pathString, ":")
	for _, path := range paths {
		p.AddPath(path)
	}
	return nil
}

// SetFromPathString sets the pathset with paths from a path string (colon-separated string)
func (p *PathSet) SetFromPathString(pathString string) error {
	p.pathSet = []string{}
	return p.AppendFromPathString(pathString)
}

// AppendFromEnvironment appends paths from the $PATH environment to the set of paths to search.
func (p *PathSet) AppendFromEnvironment() error {
	return p.AppendFromPathString(os.Getenv("PATH"))
}

// SetFromEnvironment sets the pathset to the same as the $PATH environment
func (p *PathSet) SetFromEnvironment() error {
	return p.SetFromPathString(os.Getenv("PATH"))
}

// AddPath adds a path string to the set of paths to look in.
func (p *PathSet) AddPath(path string) {
	p.pathSet = append(p.pathSet, path)
}

// Find searches for a specified file and returns the absolute path or an error.
// If the file is already an absolute, it is returned unchanged.
func (p *PathSet) Find(file string) (string, error) {
	if strings.HasPrefix(file, "/") {
		return file, nil
	}
	for _, path := range p.pathSet {
		fileName := fmt.Sprintf("%s/%s", path, file)
		if _, err := os.Stat(fileName); err == nil {
			absName, err := filepath.Abs(fileName)
			if err != nil {
				return fileName, nil
			}
			return absName, nil
		}
	}
	return "", fmt.Errorf("File Not Found in Path: %s", file)
}
