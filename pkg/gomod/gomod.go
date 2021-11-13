package gomod

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// GoMod : Collects version information from the go.mod file.
type GoMod struct {
}

// GetInfo : Returns the go version from a go.mod file.
func (s GoMod) GetInfo(path string) *goModInfo {
	goModPath := filepath.Join(path, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return nil
	}
	buf, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return nil
	}

	info := &goModInfo{}

	// Get go version
	expr := `go ([\d]*.[\d]*)`
	r := regexp.MustCompile(expr)
	result := r.FindString(string(buf))
	info.version = result[3:]

	// Get go module path
	expr = `module ([\S]*)`
	r = regexp.MustCompile(expr)
	result = r.FindString(string(buf))
	info.modulePath = result[7:]

	return info
}