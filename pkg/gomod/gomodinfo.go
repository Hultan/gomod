package gomod

// goModInfo contains information from the go.mod file
type goModInfo struct {
	version    string
	modulePath string
}

// Version returns the go version
func (s *goModInfo) Version() string {
	return s.version
}

// ModulePath returns the go module path
func (s *goModInfo) ModulePath() string {
	return s.modulePath
}
