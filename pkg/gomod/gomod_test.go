package gomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	m := GoMod{}
	assert.NotNil(t, m, "gomod should not be nil")
}

func TestGoMod_NonGoDirectory(t *testing.T) {
	m := GoMod{}
	info := m.GetInfo("/home/per/")
	assert.Nil(t, info, "info should not nil")
}

func TestGoMod_GoDirectory(t *testing.T) {
	m := GoMod{}
	info := m.GetInfo("/home/per/code/gomod")
	assert.NotNil(t, info, "info should not be nil")
	assert.Equal(t, "1.17", info.Version())
	assert.Equal(t, "github.com/hultan/gomod", info.ModulePath())
}
