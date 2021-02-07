package packages

import (
	"fmt"

	"github.com/vulogov/Bund/internal/vm"
)

func operationAdd(v *vm.VM) {
	fmt.Println("+")
}

func init() {
	Ops.Register("+", operationAdd, "Operation which will take two elements from stack and perform Add")
}
