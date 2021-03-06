package packages

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/vm"
)

func operationAdd(v *vm.VM) {
	fmt.Println("+")
}

func InitOperations() {
	log.Debugf("Configuring an embedded operations")
	Ops.Register("+", operationAdd, "Operation which will take two elements from stack and perform Add")
}
