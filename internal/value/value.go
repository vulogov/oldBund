package value

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

const (
	UNKNOWN = 0
	INT     = 1
	FLOAT   = 2
	STRING  = 3
	BOOL    = 4
	PAIR    = 5
	BLOCK   = 6
	NAME    = 7
	OP      = 8
	CMD     = 9
	NTH     = 10
)

type VALUE struct {
	vtype int8
	value interface{}
}

type P struct {
	key  string
	pval interface{}
}

func (v *VALUE) Value() interface{} {
	return v.value
}

func (v *VALUE) Type() int8 {
	return v.vtype
}

func New(value interface{}) *VALUE {
	ret := new(VALUE)
	ret.value = value
	switch value.(type) {
	case int:
		ret.vtype = INT
	case float64:
		ret.vtype = INT
	case string:
		ret.vtype = STRING
	case bool:
		ret.vtype = BOOL
	default:
		ret.vtype = UNKNOWN
	}
	return ret
}

func MakeInt(value string) (*VALUE, error) {
	val, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return nil, err
	}
	ret := new(VALUE)
	ret.value = val
	ret.vtype = INT
	return ret, nil
}

func MakeFloat(value string) (*VALUE, error) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	ret := new(VALUE)
	ret.value = val
	ret.vtype = FLOAT
	return ret, nil
}

func MakeBool(value string) (*VALUE, error) {
	val, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}
	ret := new(VALUE)
	ret.value = val
	ret.vtype = BOOL
	return ret, nil
}

func create(value interface{}, _type int8) (ret *VALUE, err error) {
	ret = new(VALUE)
	ret.value = value
	ret.vtype = _type
	return
}

func Pair(key string, value interface{}) (res *VALUE, err error) {
	p := new(P)
	p.key = key
	p.pval = New(value)
	res, err = create(p, PAIR)
	return
}

func Block(value interface{}) (*VALUE, error) {
	return create(value, BLOCK)
}

func MakeBlock() (*VALUE, error) {
	return Block(new(deque.Deque))
}

func Name(value interface{}) (*VALUE, error) {
	return create(value, NAME)
}

func Op(value interface{}) (*VALUE, error) {
	return create(value, OP)
}

func Cmd(value interface{}) (*VALUE, error) {
	return create(value, CMD)
}

func MakeNth(value string) (res *VALUE, err error) {
	if !strings.HasSuffix(value, "#") {
		err = errors.New("NTH value have no proper suffix")
		return
	}
	num := strings.TrimSuffix(value, "#")
	val, err := strconv.ParseInt(num, 0, 64)
	if err != nil {
		return
	}
	res, err = create(val, NTH)
	return
}
