package lambda

import (
	"github.com/lrita/cmap"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/vm"
)

type LAMBDA struct {
	name        string
	interp      cmap.Cmap
	compiled    cmap.Cmap
	description cmap.Cmap
}

func New(name string) *LAMBDA {
	var l *LAMBDA
	l = new(LAMBDA)
	l.name = name
	return l
}

func (l *LAMBDA) Register(name string, f func(v *vm.VM), desc string) bool {
	log.Debugf("Registering in %v: %v ", l.name, name)
	l.compiled.Store(name, f)
	l.description.Store(name, desc)
	return true
}

func (l *LAMBDA) Get(name string) func(v *vm.VM) {
	return nil
}
