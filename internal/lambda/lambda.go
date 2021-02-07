package lambda

import (
	"github.com/lrita/cmap"

	"github.com/vulogov/Bund/internal/vm"
)

type LAMBDA struct {
	interp      cmap.Cmap
	compiled    cmap.Cmap
	description cmap.Cmap
}

func (l *LAMBDA) Register(name string, f func(v *vm.VM), desc string) bool {
	l.compiled.Store(name, f)
	l.description.Store(name, desc)
	return true
}
