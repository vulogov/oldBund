package bund

import (
	"github.com/pieterclaerhout/go-log"

	blog "github.com/vulogov/Bund/internal/log"
	_ "github.com/vulogov/Bund/packages"
)

func Init() {
	blog.Init()
	log.Debug("[ theBund ] has been initiated")
}
