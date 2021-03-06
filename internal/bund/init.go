package bund

import (
	"github.com/pieterclaerhout/go-log"

	bch "github.com/vulogov/Bund/internal/channel"
	blog "github.com/vulogov/Bund/internal/log"
	"github.com/vulogov/Bund/packages"
)

func Init() {
	blog.Init()
	log.Debug("[ theBund ] has been initiated")
	packages.InitOperations()
	packages.InitCommands()
	packages.InitPartial()
	bch.InitChannels()
}
