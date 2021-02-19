package limqa

type optType int

//Option types
const (
	durable optType = iota + 1
	autoDelete
	internal
	noWait
	declareExchange
	exclusive
	autoAck
	noLocal
)

var _ Option = (*option)(nil)

// Option :
type Option interface {
	otype() optType
	value() interface{}
}

type option struct {
	_value interface{}
	_otype optType
}

func newOption(t optType, v interface{}) *option {
	return &option{
		_otype: t,
		_value: v,
	}
}

func (o *option) value() interface{} { return o._value }

func (o *option) otype() optType { return o._otype }

// DeclareExchange :
func DeclareExchange(f bool) Option {
	return newOption(declareExchange, f)
}

// Durable :
func Durable(f bool) Option {
	return newOption(durable, f)
}

// NoLocal :
func NoLocal(f bool) Option {
	return newOption(noLocal, f)
}

// AutoAck :
func AutoAck(f bool) Option {
	return newOption(autoAck, f)
}

// AutoDelete :
func AutoDelete(f bool) Option {
	return newOption(autoDelete, f)
}

// Internal :
func Internal(f bool) Option {
	return newOption(internal, f)
}

// NoWait :
func NoWait(f bool) Option {
	return newOption(noWait, f)
}

// Exclusive :
func Exclusive(f bool) Option {
	return newOption(exclusive, f)
}

func parseOptions(defaults []Option, opt []Option) map[optType]Option {
	m := make(map[optType]Option)

	for _, o := range defaults {
		m[o.otype()] = o
	}

	for _, o := range opt {
		m[o.otype()] = o
	}

	return m
}

func getFlags(opts map[optType]Option) map[optType]bool {
	r := make(map[optType]bool)
	for k, v := range opts {
		r[k] = v.value().(bool)
	}
	return r
}