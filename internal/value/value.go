package value

import (
	"strconv"

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

func create(value interface{}, _type int8) *VALUE {
	ret := new(VALUE)
	ret.value = value
	ret.vtype = _type
	return ret
}

func Pair(key string, value interface{}) *VALUE {
	p := new(P)
	p.key = key
	p.pval = value
	return create(value, PAIR)
}

func Block(value interface{}) *VALUE {
	return create(value, BLOCK)
}

func MakeBlock() *VALUE {
	return Block(new(deque.Deque))
}

func Name(value interface{}) *VALUE {
	return create(value, NAME)
}

func Op(value interface{}) *VALUE {
	return create(value, OP)
}

func Cmd(value interface{}) *VALUE {
	return create(value, CMD)
}
