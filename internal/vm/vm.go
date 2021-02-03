package vm

import (
	"errors"
	"sync"

	"github.com/gammazero/deque"

	bcmd "github.com/vulogov/Bund/internal/cmd"
	"github.com/vulogov/Bund/internal/value"
)

type VM struct {
	name    string
	lock    sync.RWMutex
	cmd     bcmd.CMD
	stack   deque.Deque
	sstack  deque.Deque
	block   deque.Deque
	sdirect bool
	isBlock bool
}

func New(name string) *VM {
	res := new(VM)
	res.name = name
	res.sdirect = true
	res.isBlock = false
	return res
}

func (v *VM) Put(val value.VALUE) int {
	v.lock.Lock()
	if v.sdirect {
		v.stack.PushBack(val)
	} else {
		v.stack.PushFront(val)
	}
	n := v.stack.Len()
	v.lock.Unlock()
	return n
}

func (v *VM) Get() (res value.VALUE, err error) {
	v.lock.Lock()
	if v.stack.Len() > 0 {
		if v.sdirect {
			res = v.stack.PopBack().(value.VALUE)
		} else {
			res = v.stack.PopFront().(value.VALUE)
		}
	} else {
		err = errors.New("Primary stack is empty")
	}
	v.lock.Unlock()
	return
}

func (v *VM) Block() {
	if !v.isBlock {
		v.isBlock = true
	} else {
		var block deque.Deque
		for v.block.Len() > 0 {
			block.PushBack(v.block.PopFront())
		}
		v.Put(*value.Block(block))
		v.isBlock = false
	}
}

func (v *VM) PutSecondary(val value.VALUE) int {
	v.lock.Lock()
	v.sstack.PushBack(val)
	n := v.sstack.Len()
	v.lock.Unlock()
	return n
}

func (v *VM) GetSecondary() value.VALUE {
	var res interface{}
	v.lock.Lock()
	if v.sstack.Len() > 0 {
		res = v.sstack.PopBack()
	}
	v.lock.Unlock()
	return res.(value.VALUE)
}
