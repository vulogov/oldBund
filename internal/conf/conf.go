package conf

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	App        = kingpin.New("bund", "[ theBund ] is an stack-based functional programming language")
	Debug      = App.Flag("debug", "Enable debug mode.").Bool()
	Production = App.Flag("production", "Switch to a production mode.").Bool()
	Color      = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()

	Version = App.Command("version", "Display information about [ theBund ]")

	Interactive = App.Command("interactive", "Run [ theBund ] interactive shell")

	Shell  = App.Command("shell", "Run [ theBund ] batch execution")
	Script = Shell.Arg("script", "File containing [ theBund ] code.").Required().ExistingFile()
)
