package conf

import (
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/rs/xid"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	seed       = time.Now().UTC().UnixNano()
	NG         = namegenerator.NewNameGenerator(seed)
	App        = kingpin.New("bund", "[ theBund ] is an stack-based functional programming language")
	Debug      = App.Flag("debug", "Enable debug mode.").Bool()
	Production = App.Flag("production", "Switch to a production mode.").Bool()
	Color      = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	ID         = App.Flag("id", "Unique application ID").Default(xid.New().String()).String()
	Name       = App.Flag("name", "Application name").Default(NG.Generate()).String()

	Version = App.Command("version", "Display information about [ theBund ]")
	VBanner = App.Flag("banner", "Display [ theBund ] banner .").Default("true").Bool()
	VTable  = App.Flag("table", "Display [ theBund ] inner information .").Default("true").Bool()

	Interactive = App.Command("interactive", "Run [ theBund ] interactive shell")

	Shell  = App.Command("shell", "Run [ theBund ] batch execution")
	Script = Shell.Arg("script", "File containing [ theBund ] code.").Required().ExistingFile()

	Eval = App.Command("eval", "Evaluate [ theBund ] expression")
	Expr = Eval.Arg("expression", "[ theBund ] expression.").Required().String()
)
