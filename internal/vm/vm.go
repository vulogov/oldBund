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

func (* v) Put(v *value.VALUE) int {
  f.lock.Lock()
  v.stack.PushBack(v)
  n := v.stack.Len()
  f.lock.Unlock()
  return n
}

func (* v) Get() *value.VALUE {
  var res *value.VALUE
  f.lock.Lock()
  if v.stack.Len() > 0 {
    res = v.stack.PopBack()
  } else {
    res = nil
  }
  f.lock.Unlock()
  return res
}
