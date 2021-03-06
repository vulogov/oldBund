package channel

import (
	"bufio"
	"io"
	"os"

	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
	crunch "github.com/superwhiskers/crunch/v3"

	"github.com/vulogov/Bund/internal/conf"
)

type CHANNEL struct {
	Name   string
	Input  io.Reader
	Output io.Writer
	Buffer *crunch.Buffer
}

var Channels cmap.Cmap

func (ch *CHANNEL) Read() bool {
	_in := bufio.NewReader(ch.Input)
	text, err := _in.ReadString('\n')
	if err != nil {
		log.Errorf("Error reading from %v: %v", ch.Name, err)
		return false
	}
	if len(text) > *conf.BufSize {
		log.Errorf("Buffer overflow for %v", ch.Name)
		return false
	}
	ch.Buffer = crunch.NewBuffer()
	ch.Buffer.Grow(int64(len(text)))
	ch.Buffer.WriteBytes(0, []byte(text))
	return true
}

func (ch *CHANNEL) Write(s string) bool {
	out := bufio.NewWriter(ch.Output)
	out.WriteString(s)
	out.Flush()
	return true
}

func (ch *CHANNEL) B() string {
	return string(ch.Buffer.Bytes())
}

func C(ch string) (CHANNEL, bool) {
	var nilCh CHANNEL
	if ch, ok := Channels.Load(ch); ok {
		return ch.(CHANNEL), true
	}
	return nilCh, false
}

func InitStdio() {
	var c CHANNEL

	log.Debug("Configuring stdin/stdout channel")
	c.Buffer = crunch.NewBuffer()
	c.Buffer.Grow(int64(*conf.BufSize))
	c.Input = bufio.NewReader(os.Stdin)
	c.Output = os.Stdout
	c.Name = "stdio"
	Channels.Store("stdio", c)
}

func InitChannels() {
	log.Debug("Configuring channels")
	InitStdio()
}
