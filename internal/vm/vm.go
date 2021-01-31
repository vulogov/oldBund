package vm

import (
	"sync"

	"github.com/gammazero/deque"

	bcmd "github.com/vulogov/Bund/internal/cmd"
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
