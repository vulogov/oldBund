package vm

import (
	"sync"

	"github.com/gammazero/deque"

	bcmd "github.com/vulogov/Bund/internal/cmd"
	"github.com/vulogov/Bund/internal/value"
)

type VM struct {
	name  string
	lock  sync.RWMutex
	cmd   bcmd.CMD
	stack deque.Deque
}

func New(name string) *VM {
	res := new(VM)
	res.name = name
	return res
}

func (v *VM) Put(val value.VALUE) int {
	v.lock.Lock()
	v.stack.PushBack(val)
	n := v.stack.Len()
	v.lock.Unlock()
	return n
}

func (v *VM) Get() value.VALUE {
	var res interface{}
	v.lock.Lock()
	if v.stack.Len() > 0 {
		res = v.stack.PopBack()
	}
	v.lock.Unlock()
	return res.(value.VALUE)
}
