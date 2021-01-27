package bund

import (
	"github.com/pieterclaerhout/go-log"

	blog "github.com/vulogov/Bund/internal/log"
)

func Init() {
	blog.Init()
	log.Info("[ theBund ] has been initiated")
}
