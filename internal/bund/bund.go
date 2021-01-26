package bund

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/vulogov/Bund/internal/conf"
)

func Main() {
	switch kingpin.MustParse(conf.App.Parse(os.Args[1:])) {
	case conf.Version.FullCommand():
		Version()
	case conf.Interactive.FullCommand():
		Interactive()
	case conf.Shell.FullCommand():
		Shell()
	}
	fmt.Println(*conf.ID)
}
